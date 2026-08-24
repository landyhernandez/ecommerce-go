package main

import (
	"errors"
	"strings"
)

// calculateTotal es una función pura: con los mismos datos produce el mismo resultado.
func calculateTotal(items []OrderItem, finder ProductFinder) (float64, error) {
	var total float64

	for _, item := range items {
		if item.Quantity <= 0 {
			return 0, errors.New("la cantidad debe ser mayor que cero")
		}

		product, ok := finder.GetProduct(item.ProductID)
		if !ok {
			return 0, errors.New("producto no encontrado")
		}

		if item.Quantity > product.Stock {
			return 0, errors.New("stock insuficiente para el producto: " + product.Name)
		}

		total += product.Price * float64(item.Quantity)
	}

	return total, nil
}

// filterProducts es una función de estilo funcional: recibe una función predicado.
func filterProducts(products []Product, predicate func(Product) bool) []Product {
	result := make([]Product, 0)
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func validCustomer(c Customer) bool {
	return strings.TrimSpace(c.Name) != "" &&
		strings.Contains(c.Email, "@")
}
