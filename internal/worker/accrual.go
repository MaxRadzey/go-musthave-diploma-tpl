package worker

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/accrual"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/logger"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/service"
	"go.uber.org/zap"
)

const (
	maxRetries429  = 3
	maxConcurrent  = 10
)

// OrderAccrualService — минимальный интерфейс сервиса заказов для воркера начислений.
type OrderAccrualService interface {
	GetOrderNumbersPendingAccrual(ctx context.Context) ([]string, error)
	ApplyAccrualResult(ctx context.Context, number, status string, accrual *int) error
}

// AccrualWorker — воркер опроса системы начислений и обновления заказов.
type AccrualWorker struct {
	baseURL  string
	orderSvc OrderAccrualService
	interval time.Duration
	client   accrual.Client
}

// NewAccrualWorker создаёт воркер. baseURL — адрес системы начислений (пустой = опрос отключён). interval — пауза между проходами.
func NewAccrualWorker(baseURL string, orderSvc OrderAccrualService, interval time.Duration) *AccrualWorker {
	return &AccrualWorker{baseURL: baseURL, orderSvc: orderSvc, interval: interval}
}

// Run запускает цикл опроса; блокируется до отмены ctx.
// Если baseURL пустой — логирует отключение и сразу выходит.
func (w *AccrualWorker) Run(ctx context.Context) {
	if w.baseURL == "" {
		logger.Log.Info("Accrual worker disabled: no accrual system address")
		return
	}
	w.client = accrual.NewHTTPClient(w.baseURL, nil)
	logger.Log.Info("Accrual worker started", zap.String("accrual_address", w.baseURL))
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	for {
		w.doOnePass(ctx)
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// следующий проход
		}
	}
}

func (w *AccrualWorker) doOnePass(ctx context.Context) {
	numbers, err := w.orderSvc.GetOrderNumbersPendingAccrual(ctx)
	if err != nil || len(numbers) == 0 {
		return
	}
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	for _, number := range numbers {
		wg.Add(1)
		go func(num string) {
			sem <- struct{}{}
			defer func() { <-sem }()
			defer wg.Done()
			w.processOrder(ctx, num)
		}(number)
	}
	wg.Wait()
}

func (w *AccrualWorker) processOrder(ctx context.Context, number string) {
	for retry := 0; retry <= maxRetries429; retry++ {
		resp, err := w.client.GetOrder(ctx, number)
		if err == nil {
			_ = w.orderSvc.ApplyAccrualResult(ctx, resp.Order, resp.Status, resp.Accrual)
			return
		}
		var notReg *accrual.ErrOrderNotRegistered
		if errors.As(err, &notReg) {
			_ = w.orderSvc.ApplyAccrualResult(ctx, number, service.OrderStatusInvalid, nil)
			return
		}
		var rateLimit *accrual.ErrRateLimit
		if errors.As(err, &rateLimit) {
			if retry < maxRetries429 {
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Duration(rateLimit.RetryAfter) * time.Second):
					continue
				}
			}
			return
		}
		// 500 / сеть — пропускаем, в следующем цикле попадёт снова
		return
	}
}
