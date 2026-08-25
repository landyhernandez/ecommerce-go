# 🛒 E-Commerce Go

## Proyecto Final — Servicios Web REST

**Estudiante:** Landy Hernandez  
**Fecha:** 24 de agosto de 2026  
**Proyecto:** E-Commerce Go

---

## 📌 1. Descripción del proyecto

**E-Commerce Go** es una aplicación web de comercio electrónico desarrollada utilizando el lenguaje de programación **Go**.

El sistema permite administrar productos, categorías, clientes, órdenes de compra y pagos mediante **Servicios Web REST**.

La aplicación cuenta con un frontend desarrollado con **HTML, CSS y JavaScript**, conectado al backend mediante solicitudes HTTP y utilizando **JSON** para el intercambio de información.

Este proyecto integra los conocimientos adquiridos durante las cuatro unidades de la asignatura.

---

## 🎯 2. Objetivo del proyecto

El objetivo principal es desarrollar una aplicación de comercio electrónico funcional que permita aplicar los conocimientos de programación y desarrollo de Servicios Web.

Los objetivos específicos son:

- Crear Servicios Web REST.
- Utilizar HTTP para la comunicación entre aplicaciones.
- Utilizar JSON para el intercambio de información.
- Implementar operaciones para productos y categorías.
- Registrar y consultar clientes.
- Crear y consultar órdenes.
- Procesar pagos.
- Implementar inicio de sesión.
- Integrar frontend y backend.
- Utilizar Git y GitHub para el control de versiones.

---

## ✨ 3. Funcionalidades principales

### 🛍️ Productos

La aplicación permite:

- Consultar todos los productos.
- Consultar un producto mediante su ID.
- Filtrar productos por categoría.
- Consultar categorías.
- Visualizar precios.
- Visualizar stock.
- Visualizar descripciones.

### 👤 Clientes

El sistema permite:

- Registrar clientes.
- Validar información.
- Consultar clientes.
- Realizar inicio de sesión.

### 🛒 Carrito y compras

El sistema permite:

- Seleccionar productos.
- Agregar productos al carrito.
- Calcular cantidades.
- Calcular el total de la compra.
- Validar el stock disponible.
- Crear órdenes.
- Consultar órdenes.

### 💳 Pagos

El sistema permite:

- Procesar pagos.
- Validar el monto.
- Validar el método de pago.
- Asociar el pago con una orden.
- Mostrar el estado del pago.

---

# 🌐 4. Servicios Web REST

El proyecto implementa al menos **8 Servicios Web REST** de diferentes funcionalidades.

| # | Método | Endpoint | Funcionalidad |
|---|---|---|---|
| 1 | GET | `/api/products` | Obtener productos |
| 2 | GET | `/api/products/{id}` | Obtener producto por ID |
| 3 | GET | `/api/categories` | Obtener categorías |
| 4 | POST | `/api/customers` | Registrar cliente |
| 5 | POST | `/api/orders` | Crear una orden |
| 6 | GET | `/api/orders/{id}` | Consultar una orden |
| 7 | POST | `/api/payments` | Procesar un pago |
| 8 | POST | `/api/login` | Inicio de sesión |

Estos servicios permiten que el frontend se comunique con el backend mediante solicitudes HTTP.

---

# 🔄 5. Serialización mediante JSON

La comunicación entre el frontend y el backend se realiza utilizando **JSON**.

Ejemplo de información de un producto:

```json
{
  "id": 1,
  "name": "Laptop Lenovo",
  "category": "Tecnologia",
  "price": 899.99,
  "stock": 10,
  "description": "Laptop para trabajo y estudio"
}