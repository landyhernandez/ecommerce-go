package main

// Product representa un producto del catálogo.
type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	Description string  `json:"description"`
}

// Customer representa un cliente.
type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// OrderItem representa un producto dentro de una orden.
type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// Order representa una compra.
type Order struct {
	ID         int         `json:"id"`
	CustomerID int         `json:"customer_id"`
	Items      []OrderItem `json:"items"`
	Total      float64     `json:"total"`
	Status     string      `json:"status"`
}

// Payment representa el pago de una orden.
type Payment struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"`
	Status  string  `json:"status"`
}

// LoginRequest representa los datos de autenticación.
type LoginRequest struct {
	Email string `json:"email"`
}

// LoginResponse es la respuesta de autenticación.
type LoginResponse struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}
