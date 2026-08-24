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
Permite obtener todos los productos disponibles en el catálogo.

También permite filtrar productos por categoría:

GET /api/products?category=Tecnologia

Servicio 2 - Obtener producto por ID
Método:

GET /api/products/{id}

Permite consultar un producto específico utilizando su identificador.

Ejemplo:

GET /api/products/1

Servicio 3 - Obtener categorías
Método:

GET /api/categories

Permite obtener las categorías disponibles en el catálogo.

Servicio 4 - Registrar cliente
Método:

POST /api/customers

Permite registrar un nuevo cliente.

Ejemplo de información enviada:

{
  "name": "Landy Hernandez",
  "email": "landy@demo.com"
}

Servicio 5 - Crear orden
Método:

POST /api/orders

Permite crear una orden de compra asociada a un cliente.

El sistema verifica:

Que el cliente exista.
Que los productos existan.
Que la cantidad sea válida.
Que exista suficiente stock.
El total de la compra.
Servicio 6 - Consultar orden
Método:

GET /api/orders/{id}

Permite consultar una orden específica mediante su ID.

Ejemplo:

GET /api/orders/1

Servicio 7 - Procesar pago
Método:

POST /api/payments

Permite procesar el pago de una orden.

El sistema verifica:

Que la orden exista.
Que el monto sea correcto.
Que el método de pago sea válido.
El resultado puede ser:

APROBADO

Servicio 8 - Inicio de sesión
Método:

POST /api/login

Permite validar el correo electrónico de un cliente registrado.

Ejemplo:

{
  "email": "landy@demo.com"
}

🔄 5. Uso de JSON
La comunicación entre el frontend y el backend se realiza mediante JSON.

Ejemplo de un producto:

{
  "id": 1,
  "name": "Laptop Lenovo",
  "category": "Tecnologia",
  "price": 899.99,
  "stock": 10,
  "description": "Laptop para trabajo y estudio"
}

Go utiliza estructuras con etiquetas JSON para realizar la serialización y deserialización de los datos.

Ejemplo:

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Category    string  `json:"category"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    Description string  `json:"description"`
}

🏗️ 6. Arquitectura del proyecto
La aplicación utiliza una arquitectura sencilla donde el frontend se comunica con los servicios REST del backend.

                 👤 USUARIO
                     │
                     ▼
              🖥️ FRONTEND
          HTML + CSS + JavaScript
                     │
                     │ HTTP / JSON
                     ▼
              ⚙️ BACKEND GO
                     │
                     ▼
               🌐 REST API
                     │
                     ▼
                  STORE
                     │
          ┌──────────┼──────────┐
          ▼          ▼          ▼
      Productos   Clientes    Órdenes
                                  │
                                  ▼
                                Pagos

📁 7. Estructura del proyecto
ecommerce_go/
│
├── frontend/
│   └── index.html
│
├── main.go
├── handlers.go
├── models.go
├── services.go
├── interfaces.go
├── store.go
├── go.mod
├── README.md
└── .gitignore

📄 8. Descripción de los archivos
main.go
Contiene el punto de entrada del programa.

También configura el servidor HTTP y registra las rutas de los servicios REST.

handlers.go
Contiene los manejadores HTTP de los diferentes servicios web.

Aquí se implementan los servicios de:

Productos.
Categorías.
Clientes.
Órdenes.
Pagos.
Login.
models.go
Contiene las estructuras principales del sistema:

Product
Customer
Order
OrderItem
Payment
LoginRequest
LoginResponse
services.go
Contiene funciones relacionadas con la lógica del negocio.

Entre ellas se encuentra el cálculo del total de las órdenes, filtros de productos y validaciones.

interfaces.go
Define interfaces utilizadas para separar la lógica de negocio de las implementaciones concretas.

Se utilizan interfaces para:

Buscar productos.
Procesar pagos.
store.go
Administra los datos almacenados en memoria.

Utiliza mapas para almacenar:

Productos.
Clientes.
Órdenes.
Pagos.
También utiliza sync.RWMutex para proteger el acceso concurrente a los datos.

frontend/index.html
Contiene la interfaz gráfica de la tienda.

Utiliza:

HTML.
CSS.
JavaScript.
El JavaScript realiza solicitudes al backend mediante fetch().

🛡️ 9. Manejo de errores
La aplicación realiza validaciones para evitar operaciones incorrectas.

Algunos ejemplos son:

JSON inválido.
Producto inexistente.
Cliente inexistente.
Orden inexistente.
Cantidad menor o igual a cero.
Stock insuficiente.
Precio inválido.
Correo electrónico inválido.
Monto de pago incorrecto.
Método de pago vacío.
Método HTTP no permitido.
Las respuestas de error se envían mediante códigos HTTP y mensajes en formato JSON.

🔐 10. Encapsulación
El almacenamiento de los datos se encuentra encapsulado dentro de la estructura Store.

Los mapas internos son privados:

products
customers
orders
payments

El acceso se realiza mediante métodos como:

AddProduct()
GetProducts()
GetProduct()
AddCustomer()
GetCustomer()
AddOrder()
GetOrder()
AddPayment()
GetPayment()

Esto permite separar el almacenamiento de la lógica que utiliza los datos.

🔄 11. Concurrencia
El proyecto utiliza:

sync.RWMutex

para proteger los datos almacenados en memoria.

Se utilizan bloqueos de lectura y escritura:

RLock()
RUnlock()

y:

Lock()
Unlock()

Esto permite controlar el acceso concurrente a los mapas utilizados por la aplicación.

💻 12. Tecnologías utilizadas
Backend
Go
HTTP
REST
JSON
Frontend
HTML5
CSS3
JavaScript
Herramientas
Visual Studio Code
Git
GitHub
PowerShell
🚀 13. Cómo ejecutar el proyecto
Primero se debe ingresar a la carpeta principal del proyecto.

Después ejecutar:

go run .

Cuando el servidor se encuentre ejecutándose, abrir en el navegador:

http://localhost:8080

🧪 14. Pruebas realizadas
Los servicios principales fueron probados durante el desarrollo.

Resultados:

✅ Obtener productos.
✅ Obtener producto por ID.
✅ Obtener categorías.
✅ Registrar cliente.
✅ Crear orden.
✅ Consultar orden.
✅ Procesar pago.
✅ Inicio de sesión.
Los servicios respondieron correctamente y permitieron comprobar la comunicación entre el frontend y el backend.

🔮 15. Visualización del futuro
El proyecto puede evolucionar en el futuro hasta convertirse en una plataforma de comercio electrónico más completa.

Las futuras mejoras podrían incluir:

☁️ Despliegue en la nube.
🗄️ Base de datos PostgreSQL o MySQL.
🔐 Autenticación mediante JWT.
💳 Integración con plataformas de pago reales.
📱 Aplicación móvil.
📊 Panel administrativo.
📦 Control de inventario en tiempo real.
📧 Notificaciones por correo electrónico.
🔎 Búsqueda avanzada.
🤖 Inteligencia artificial para recomendaciones de productos.
La visión futura consiste en transformar el prototipo académico en una plataforma escalable y preparada para usuarios reales.

🎓 16. Integración de las cuatro unidades
El proyecto integra conocimientos de las diferentes unidades de la asignatura.

Unidad 1
Se aplicaron fundamentos de programación, variables, estructuras, funciones y organización del código.

Unidad 2
Se aplicaron conceptos de programación en Go, estructuras, métodos, interfaces y manejo de errores.

Unidad 3
Se desarrollaron Servicios Web REST utilizando HTTP y JSON para el intercambio de información.

Unidad 4
Se integró el backend con el frontend y se utilizaron herramientas como Git y GitHub para administrar el proyecto.

También se aplicaron conceptos de concurrencia y organización de una aplicación completa.

📚 17. Aplicaciones prácticas
El proyecto representa una aplicación práctica de los Servicios Web en un escenario de comercio electrónico.

Este tipo de arquitectura puede utilizarse en:

Tiendas en línea.
Sistemas de inventario.
Plataformas de venta.
Aplicaciones móviles.
Sistemas empresariales.
Plataformas de gestión de pedidos.
Los servicios REST permiten que diferentes aplicaciones puedan comunicarse con el backend.

⚠️ 18. Dificultades encontradas
Durante el desarrollo se presentaron diferentes dificultades.

Entre ellas:

Integrar correctamente frontend y backend.
Configurar las rutas REST.
Manejar las solicitudes HTTP.
Trabajar con JSON.
Validar los datos recibidos.
Controlar el stock.
Implementar las órdenes y pagos.
Configurar Git y GitHub.
Organizar correctamente los archivos del proyecto.
Estas dificultades permitieron mejorar la comprensión de las tecnologías utilizadas.

🧠 19. Aprendizajes
El desarrollo del proyecto permitió aprender y reforzar conocimientos relacionados con:

Desarrollo de aplicaciones en Go.
Creación de APIs REST.
Comunicación mediante HTTP.
Uso de JSON.
Diseño de estructuras.
Uso de interfaces.
Manejo de errores.
Concurrencia.
Desarrollo frontend.
Integración frontend-backend.
Uso de Git.
Uso de GitHub.
🎥 20. Demostración
El proyecto cuenta con una demostración de sus principales funcionalidades.

Durante la demostración se presentan:

Página principal.
Productos.
Categorías.
Registro de cliente.
Carrito.
Creación de orden.
Consulta de orden.
Procesamiento de pago.
Inicio de sesión.
Servicios REST.
👨‍💻 21. Autor
Landy Hernandez
Proyecto desarrollado como trabajo final de la asignatura.

📅 22. Fecha
24 de agosto de 2026

📝 23. Conclusión
El desarrollo de E-Commerce Go permitió integrar los conocimientos adquiridos durante las cuatro unidades de la asignatura en una aplicación práctica.

Se logró construir una aplicación funcional utilizando Go como lenguaje principal, implementando ocho Servicios Web REST y utilizando JSON para la comunicación entre el frontend y el backend.

El proyecto permitió comprender la importancia de organizar correctamente el código, utilizar interfaces, manejar errores, controlar el acceso concurrente a los datos y utilizar herramientas de control de versiones como Git y GitHub.

Una de las principales dificultades fue integrar todos los componentes del sistema para que funcionaran correctamente. Sin embargo, mediante pruebas y correcciones se logró obtener una aplicación funcional.

Como trabajo futuro, el proyecto podría incorporar una base de datos, autenticación avanzada, pagos reales, servicios en la nube, una aplicación móvil e inteligencia artificial.

En conclusión, este proyecto permitió aplicar los conocimientos de programación y Servicios Web en un escenario práctico de comercio electrónico y comprender cómo estas tecnologías pueden utilizarse para desarrollar soluciones reales.
