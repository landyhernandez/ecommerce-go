# 🛒 E-commerce Go

## Proyecto Final – Servicios Web REST

**Estudiante:** Landy Hernandez  
**Lenguaje:** Go  
**Tipo de proyecto:** Aplicación web e-commerce  
**Fecha:** 24 de agosto de 2026

---

## 📌 Descripción del proyecto

E-commerce Go es una aplicación web de comercio electrónico desarrollada con el lenguaje de programación Go.

El proyecto permite administrar productos, categorías, clientes, órdenes de compra, pagos e inicio de sesión mediante servicios web REST.

La aplicación cuenta con un frontend desarrollado en HTML, CSS y JavaScript que se comunica con el backend mediante solicitudes HTTP y datos en formato JSON.

---

## 🎯 Objetivo del proyecto

El objetivo principal es desarrollar una aplicación de comercio electrónico que permita demostrar los conocimientos adquiridos durante las diferentes unidades de la asignatura.

El proyecto integra conceptos de:

- Programación en Go.
- Servicios Web REST.
- APIs HTTP.
- Serialización y deserialización mediante JSON.
- Estructuras y métodos.
- Interfaces.
- Manejo de errores.
- Funciones.
- Encapsulación.
- Concurrencia y protección de datos.
- Desarrollo de una interfaz web.
- Integración entre frontend y backend.

---

## ⚙️ Funcionalidades principales

La aplicación permite:

- 📦 Consultar productos.
- 🔎 Consultar un producto por su ID.
- 🏷️ Consultar categorías.
- 👤 Registrar clientes.
- 🧾 Crear órdenes de compra.
- 🔍 Consultar órdenes.
- 💳 Procesar pagos.
- 🔐 Iniciar sesión.
- 🛒 Agregar productos al carrito.
- 💰 Calcular automáticamente el total de una orden.
- 📊 Controlar el stock disponible.

---

## 🌐 Servicios Web REST

El proyecto implementa 8 servicios web principales:

### 1. Obtener productos

```text
GET /api/products
