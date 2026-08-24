package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func readJSON(r *http.Request, target any) error {
	return json.NewDecoder(r.Body).Decode(target)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ======================================================
// PRODUCTOS
// GET/POST /api/products
// ======================================================

func productsHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:

			products := store.GetProducts()

			category := strings.TrimSpace(
				r.URL.Query().Get("category"),
			)

			if category != "" {

				products = filterProducts(
					products,
					func(p Product) bool {
						return strings.EqualFold(
							p.Category,
							category,
						)
					},
				)

			}

			writeJSON(
				w,
				http.StatusOK,
				products,
			)

		case http.MethodPost:

			var product Product

			if err := readJSON(r, &product); err != nil {
				writeJSON(
					w,
					http.StatusBadRequest,
					map[string]string{
						"error": "JSON inválido",
					},
				)
				return
			}

			if strings.TrimSpace(product.Name) == "" ||
				product.Price <= 0 ||
				product.Stock < 0 {

				writeJSON(
					w,
					http.StatusBadRequest,
					map[string]string{
						"error": "nombre, precio y stock deben ser válidos",
					},
				)

				return
			}

			created := store.AddProduct(product)

			writeJSON(
				w,
				http.StatusCreated,
				created,
			)

		default:

			writeJSON(
				w,
				http.StatusMethodNotAllowed,
				map[string]string{
					"error": "método no permitido",
				},
			)
		}
	}
}

// ======================================================
// PRODUCTO POR ID
// GET /api/products/{id}
// ======================================================

func productByIDHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(
			strings.TrimPrefix(
				r.URL.Path,
				"/api/products/",
			),
		)

		if err != nil {
			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "ID inválido",
				},
			)
			return
		}

		product, ok := store.GetProduct(id)

		if !ok {
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]string{
					"error": "producto no encontrado",
				},
			)
			return
		}

		writeJSON(
			w,
			http.StatusOK,
			product,
		)
	}
}

// ======================================================
// CATEGORÍAS
// GET /api/categories
// ======================================================

func categoriesHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		seen := map[string]bool{}

		categories := make([]string, 0)

		for _, p := range store.GetProducts() {

			if !seen[p.Category] {

				seen[p.Category] = true

				categories = append(
					categories,
					p.Category,
				)
			}
		}

		writeJSON(
			w,
			http.StatusOK,
			categories,
		)
	}
}

// ======================================================
// CLIENTES
// POST /api/customers
// ======================================================

func customersHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {

			writeJSON(
				w,
				http.StatusMethodNotAllowed,
				map[string]string{
					"error": "método no permitido",
				},
			)

			return
		}

		var customer Customer

		if err := readJSON(r, &customer); err != nil {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "JSON inválido",
				},
			)

			return
		}

		customer.Email =
			normalizeEmail(customer.Email)

		if !validCustomer(customer) {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "nombre o correo inválido",
				},
			)

			return
		}

		created :=
			store.AddCustomer(customer)

		writeJSON(
			w,
			http.StatusCreated,
			created,
		)
	}
}

// ======================================================
// ÓRDENES
// POST /api/orders
// ======================================================

func ordersHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {

			writeJSON(
				w,
				http.StatusMethodNotAllowed,
				map[string]string{
					"error": "método no permitido",
				},
			)

			return
		}

		var order Order

		if err := readJSON(r, &order); err != nil {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "JSON inválido",
				},
			)

			return
		}

		// Verificar cliente.

		if _, ok :=
			store.GetCustomer(order.CustomerID); !ok {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "cliente no encontrado",
				},
			)

			return
		}

		// Verificar que existan productos.

		if len(order.Items) == 0 {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "la orden debe contener productos",
				},
			)

			return
		}

		// Calcular total y comprobar stock.

		total, err :=
			calculateTotal(
				order.Items,
				store,
			)

		if err != nil {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": err.Error(),
				},
			)

			return
		}

		order.Total = total

		order.Status = "PENDIENTE"

		// Guardar orden.

		created :=
			store.AddOrder(order)

		writeJSON(
			w,
			http.StatusCreated,
			created,
		)
	}
}

// ======================================================
// ORDEN POR ID
// GET /api/orders/{id}
// ======================================================

func orderByIDHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(
			strings.TrimPrefix(
				r.URL.Path,
				"/api/orders/",
			),
		)

		if err != nil {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "ID inválido",
				},
			)

			return
		}

		order, ok :=
			store.GetOrder(id)

		if !ok {

			writeJSON(
				w,
				http.StatusNotFound,
				map[string]string{
					"error": "orden no encontrada",
				},
			)

			return
		}

		writeJSON(
			w,
			http.StatusOK,
			order,
		)
	}
}

// ======================================================
// PAGOS
// POST /api/payments
// ======================================================

func paymentsHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {

			writeJSON(
				w,
				http.StatusMethodNotAllowed,
				map[string]string{
					"error": "método no permitido",
				},
			)

			return
		}

		var payment Payment

		if err := readJSON(r, &payment); err != nil {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "JSON inválido",
				},
			)

			return
		}

		order, ok :=
			store.GetOrder(payment.OrderID)

		if !ok {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "orden no encontrada",
				},
			)

			return
		}

		if payment.Amount != order.Total ||
			payment.Amount <= 0 {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "monto de pago incorrecto",
				},
			)

			return
		}

		processor :=
			SimplePaymentProcessor{}

		status, err :=
			processor.Process(
				payment.Amount,
				payment.Method,
			)

		if err != nil {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": err.Error(),
				},
			)

			return
		}

		payment.Status = status

		created :=
			store.AddPayment(payment)

		writeJSON(
			w,
			http.StatusCreated,
			created,
		)
	}
}

// ======================================================
// LOGIN
// POST /api/login
// ======================================================

func loginHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {

			writeJSON(
				w,
				http.StatusMethodNotAllowed,
				map[string]string{
					"error": "método no permitido",
				},
			)

			return
		}

		var request LoginRequest

		if err := readJSON(r, &request); err != nil {

			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]string{
					"error": "JSON inválido",
				},
			)

			return
		}

		email :=
			normalizeEmail(request.Email)

		for _, c := range store.customers {

			if c.Email == email {

				writeJSON(
					w,
					http.StatusOK,
					LoginResponse{
						Message: "Inicio de sesión válido",

						Email: email,
					},
				)

				return
			}
		}

		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]string{
				"error": "cliente no encontrado",
			},
		)
	}
}

// ======================================================
// DATOS INICIALES
// ======================================================
func seedData(store *Store) {

	store.AddProduct(Product{
		Name:        "Laptop Lenovo",
		Category:    "Tecnologia",
		Price:       899.99,
		Stock:       10,
		Description: "Laptop para trabajo y estudio",
	})

	store.AddProduct(Product{
		Name:        "Monitor Samsung 24",
		Category:    "Tecnologia",
		Price:       189.99,
		Stock:       12,
		Description: "Monitor Full HD de 24 pulgadas",
	})

	store.AddProduct(Product{
		Name:        "Mouse inalámbrico",
		Category:    "Accesorios",
		Price:       24.99,
		Stock:       30,
		Description: "Mouse inalámbrico USB",
	})

	store.AddProduct(Product{
		Name:        "Teclado mecánico RGB",
		Category:    "Accesorios",
		Price:       59.99,
		Stock:       15,
		Description: "Teclado mecánico con iluminación RGB",
	})

	store.AddProduct(Product{
		Name:        "Audífonos Bluetooth",
		Category:    "Audio",
		Price:       79.99,
		Stock:       20,
		Description: "Audífonos inalámbricos con conexión Bluetooth",
	})

	store.AddProduct(Product{
		Name:        "Webcam HD",
		Category:    "Accesorios",
		Price:       44.99,
		Stock:       18,
		Description: "Cámara web HD para videollamadas",
	})

	store.AddProduct(Product{
		Name:        "Smartphone Android",
		Category:    "Celulares",
		Price:       499.99,
		Stock:       8,
		Description: "Smartphone Android con pantalla Full HD",
	})

	store.AddProduct(Product{
		Name:        "Parlante Bluetooth",
		Category:    "Audio",
		Price:       39.99,
		Stock:       25,
		Description: "Parlante portátil con conexión Bluetooth",
	})

	store.AddProduct(Product{
		Name:        "Disco SSD 1TB",
		Category:    "Tecnologia",
		Price:       89.99,
		Stock:       14,
		Description: "Unidad SSD de 1TB para almacenamiento rápido",
	})

	store.AddProduct(Product{
		Name:        "Hub USB 4 Puertos",
		Category:    "Accesorios",
		Price:       19.99,
		Stock:       35,
		Description: "Hub USB con cuatro puertos",
	})

	store.AddProduct(Product{
		Name:        "Control inalámbrico",
		Category:    "Gaming",
		Price:       54.99,
		Stock:       16,
		Description: "Control inalámbrico para videojuegos",
	})

	store.AddProduct(Product{
		Name:        "Impresora Multifunción",
		Category:    "Tecnologia",
		Price:       149.99,
		Stock:       7,
		Description: "Impresora multifunción para oficina y hogar",
	})

	store.AddCustomer(Customer{
		Name:  "Cliente Demo",
		Email: "cliente@demo.com",
	})
}
