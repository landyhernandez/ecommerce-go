# Sistema de Gestión de E-commerce - Go

Proyecto integrador desarrollado en Go para demostrar programación funcional,
encapsulación, estructuras, interfaces/funciones, manejo de errores y
generación de Servicios Web REST con JSON.

## Requisitos

- Go 1.22 o superior
- Visual Studio Code recomendado
- Git/GitHub

## Ejecutar

Desde la carpeta del proyecto:

```bash
go run .
```

El servidor quedará disponible en:

`http://localhost:8080`

## Servicios Web

1. `GET /api/products` - consultar catálogo.
2. `POST /api/products` - registrar producto.
3. `GET /api/products/{id}` - consultar producto por ID.
4. `GET /api/categories` - consultar categorías.
5. `POST /api/customers` - registrar cliente.
6. `POST /api/orders` - crear orden y calcular total.
7. `GET /api/orders/{id}` - consultar orden.
8. `POST /api/payments` - registrar pago.
9. `POST /api/login` - validar cliente.

Todos los datos de entrada y salida se manejan mediante JSON.

## Ejemplos

### Consultar productos

```bash
curl http://localhost:8080/api/products
```

### Crear producto

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Audifonos\",\"category\":\"Accesorios\",\"price\":39.99,\"stock\":20,\"description\":\"Audifonos Bluetooth\"}"
```

### Crear cliente

```bash
curl -X POST http://localhost:8080/api/customers \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Juan Perez\",\"email\":\"juan@email.com\"}"
```

### Crear orden

```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d "{\"customer_id\":1,\"items\":[{\"product_id\":1,\"quantity\":2}]}"
```

### Registrar pago

Si la orden anterior tiene total 1799.98:

```bash
curl -X POST http://localhost:8080/api/payments \
  -H "Content-Type: application/json" \
  -d "{\"order_id\":1,\"amount\":1799.98,\"method\":\"TARJETA\"}"
```

## Programación funcional

El proyecto usa funciones puras y funciones de orden superior:

- `calculateTotal`: calcula el total sin modificar el estado global.
- `filterProducts`: recibe una función predicado para filtrar productos.
- `normalizeEmail`: transforma datos sin efectos secundarios.
- `validCustomer`: valida un cliente.
- `PaymentProcessor`: interfaz para desacoplar el procesamiento de pagos.

Go no es un lenguaje puramente funcional, por lo que se combinan estos
conceptos con estructuras y métodos de Go.

## Encapsulación

Los mapas internos de `Store` son privados:

- `products`
- `customers`
- `orders`
- `payments`

El acceso se realiza mediante métodos como `AddProduct`, `GetProduct`,
`AddCustomer`, `GetOrder`, etc.

## Paquetes externos

El proyecto no requiere paquetes externos. Se utiliza la biblioteca estándar
de Go, principalmente:

- `net/http`
- `encoding/json`
- `strconv`
- `strings`
- `sync`

## Manejo de errores e interfaces

El servidor valida JSON, IDs, clientes, productos, stock, montos y métodos de pago. Además, se utiliza la interfaz `ProductFinder` para desacoplar la lógica de cálculo de pedidos y `PaymentProcessor` para representar el procesamiento de pagos.
## GitHub

Comandos básicos:

```bash
git init
git add .
git commit -m "Proyecto integrador e-commerce"
git branch -M main
git remote add origin URL_DE_TU_REPOSITORIO
git push -u origin main
```
