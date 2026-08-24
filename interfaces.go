package main

// ProductFinder define el comportamiento necesario para buscar productos.
// La interfaz permite desacoplar la lógica de negocio del almacenamiento.
type ProductFinder interface {
	GetProduct(id int) (Product, bool)
}

// PaymentProcessor representa una abstracción para procesar pagos.
type PaymentProcessor interface {
	Process(amount float64, method string) (string, error)
}

// SimplePaymentProcessor es una implementación sencilla para el proyecto académico.
type SimplePaymentProcessor struct{}

func (SimplePaymentProcessor) Process(amount float64, method string) (string, error) {
	if amount <= 0 {
		return "", errorsNew("el monto debe ser mayor que cero")
	}
	if method == "" {
		return "", errorsNew("el método de pago es obligatorio")
	}
	return "APROBADO", nil
}

// errorsNew mantiene el ejemplo simple y permite demostrar manejo de errores
// sin agregar dependencias externas.
func errorsNew(message string) error {
	return simpleError(message)
}

type simpleError string

func (e simpleError) Error() string { return string(e) }
