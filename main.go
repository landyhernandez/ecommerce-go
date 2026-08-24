package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewStore()
	seedData(store)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "./frontend/index.html")
	})

	// Servicios Web REST
	mux.HandleFunc("/api/products", productsHandler(store))
	mux.HandleFunc("/api/products/", productByIDHandler(store))
	mux.HandleFunc("/api/categories", categoriesHandler(store))
	mux.HandleFunc("/api/customers", customersHandler(store))
	mux.HandleFunc("/api/orders", ordersHandler(store))
	mux.HandleFunc("/api/orders/", orderByIDHandler(store))
	mux.HandleFunc("/api/payments", paymentsHandler(store))
	mux.HandleFunc("/api/login", loginHandler(store))

	handler := cors(mux)

	log.Println("Servidor e-commerce ejecutándose en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
