<h1 align="center">QuoteMaker</h1>

<p align="center">
  <img src="assets/images/QuoteMaker_logo.png" alt="QuoteMaker Logo" width="150">
</p>

QuoteMaker is a quotation management system built with Go. It allows users to create, update, and manage quotations for products.

| Método | Endpoint                                         | Descripción                                 | Notas                                                            |
|--------|--------------------------------------------------|---------------------------------------------|------------------------------------------------------------------|
| GET    | `/category?limit=&offset=`                       | Listar todas las categorías                 | Soporta paginación con parámetros `limit` y `offset`.            |
| GET    | `/category/:id`                                  | Obtener una categoría por su ID             |                                                                  |
| GET    | `/category/search?name=&limit=&offset=`          | Buscar categorías por nombre                | Soporta paginación con parámetros `limit` y `offset`.            |
| POST   | `/category`                                      | Crear una nueva categoría                   |                                                                  |
| PUT    | `/category`                                      | Actualizar una categoría existente          |                                                                  |
| DELETE | `/category/:id`                                  | Eliminar una categoría por su ID            |                                                                  |
| GET    | `/products?name=&limit=&offset=`                 | Listar productos por nombre                 | Soporta paginación y filtrado por nombre.                        |
| GET    | `/products/:id`                                  | Obtener detalles de un producto por su ID   |                                                                  |
| GET    | `/products/category/:categoryName?limit=&offset=` | Listar productos por categoría             | Soporta paginación con parámetros `limit` y `offset`.            |
| POST   | `/products`                                      | Crear un nuevo producto                     |                                                                  |
| PUT    | `/products/:id`                                  | Actualizar un producto existente            |                                                                  |
| DELETE | `/products/:id`                                  | Eliminar un producto por su ID              |                                                                  |
| GET    | `/quotations?limit=&offset=`                     | Listar todas las cotizaciones               | Soporta paginación con parámetros `limit` y `offset`.            |
| POST   | `/quotations`                                    | Crear una nueva cotización                  | Valida UUIDs de productos.                                       |
