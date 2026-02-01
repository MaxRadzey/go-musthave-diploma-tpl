package handler

// RegisterRequest — тело запроса регистрации.
type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// LoginRequest — тело запроса логина.
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// OrderItem — элемент списка заказов.
type OrderItem struct {
	Number     string   `json:"number"`
	Status     string   `json:"status"`
	Accrual    *float64 `json:"accrual,omitempty"`
	UploadedAt string   `json:"uploaded_at"`
}
