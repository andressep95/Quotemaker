<!-- PROJECT LOGO -->
<div align="center">
  <a>
    <img src="/assets/images/QuoteMaker_Logo.png" alt="Logo" width="80" height="80">
  </a>
  <h3 align="center">QuoteMaker</h3>
  <p align="center">
    Una herramienta increíble para obtener cotizaciones rápidas y precisas
    <br />
    <a href="https://github.com/your_username/QuoteMaker"><strong>Explora la documentación »</strong></a>
    <br />
    <br />
    <a href="https://github.com/your_username/QuoteMaker">Ver Demo</a>
    ·
    <a href="https://github.com/your_username/QuoteMaker/issues">Reportar Bug</a>
    ·
    <a href="https://github.com/your_username/QuoteMaker/issues">Solicitar Funcionalidad</a>
  </p>
</div>

## Tabla de Contenidos
- [Tabla de Contenidos](#tabla-de-contenidos)
- [Introducción](#introducción)
- [Arquitectura Hexagonal](#arquitectura-hexagonal)
- [Funcionalidades](#funcionalidades)
- [Empezando](#empezando)
- [Instalación](#instalación)
- [Uso](#uso)
- [Contribuyendo](#contribuyendo)
- [Licencia](#licencia)
- [Contacto](#contacto)
- [Agradecimientos](#agradecimientos)

## Introducción
Facilita la creación de cotizaciones minimizando el tipeo manual y los posibles errores que podrían ocurrir al hacerlo de la forma convencional.

## Arquitectura Hexagonal
Descripción de la estructura del proyecto con la arquitectura hexagonal, incluyendo el flujo entre API, casos de uso, servicios y base de datos.

Estructura del Proyecto:
```
 ┣ .github
 ┣ assets
 ┃ ┗ images
 ┃ ┃ ┗ QuoteMaker_Logo.png
 ┣ cmd
 ┃ ┗ main.go
 ┣ internal
 ┃ ┣ adapters
 ┃ ┃ ┣ api
 ┃ ┃ ┃ ┗ models
 ┃ ┃ ┃ ┃ ┣ customer
 ┃ ┃ ┃ ┃ ┗ product
 ┃ ┃ ┃ ┃ ┃ ┣ productRequest.go
 ┃ ┃ ┃ ┃ ┃ ┗ productResponse.go
 ┃ ┃ ┗ repository
 ┃ ┃ ┃ ┣ sqlCategoryRepository.go
 ┃ ┃ ┃ ┣ sqlCategoryRepository_test.go
 ┃ ┃ ┃ ┣ sqlCustomerRepository.go
 ┃ ┃ ┃ ┣ sqlCustomerRepository_test.go
 ┃ ┃ ┃ ┣ sqlProductRepository.go
 ┃ ┃ ┃ ┣ sqlProductRepository_test.go
 ┃ ┃ ┃ ┣ sqlQuotationRepository.go
 ┃ ┃ ┃ ┣ sqlQuotationRepository_test.go
 ┃ ┃ ┃ ┣ sqlSellerRepository.go
 ┃ ┃ ┃ ┗ sqlSellerRepository_test.go
 ┃ ┣ application
 ┃ ┃ ┣ mappers
 ┃ ┃ ┗ usecases
 ┃ ┃ ┃ ┗ productUseCases.go
 ┃ ┣ config
 ┃ ┃ ┣ config.go
 ┃ ┃ ┗ config.yaml
 ┃ ┣ core
 ┃ ┃ ┣ domain
 ┃ ┃ ┃ ┣ entity
 ┃ ┃ ┃ ┃ ┣ category.go
 ┃ ┃ ┃ ┃ ┣ customer.go
 ┃ ┃ ┃ ┃ ┣ product.go
 ┃ ┃ ┃ ┃ ┣ quotation.go
 ┃ ┃ ┃ ┃ ┣ quoteProduct.go
 ┃ ┃ ┃ ┃ ┗ seller.go
 ┃ ┃ ┃ ┗ service
 ┃ ┃ ┃ ┃ ┗ productService.go
 ┃ ┃ ┗ ports
 ┃ ┃ ┃ ┣ categoryRepository.go
 ┃ ┃ ┃ ┣ customerRepository.go
 ┃ ┃ ┃ ┣ productRepository.go
 ┃ ┃ ┃ ┣ quotationRepository.go
 ┃ ┃ ┃ ┗ sellerRepository.go
 ┃ ┣ infrastructure
 ┃ ┃ ┣ db
 ┃ ┃ ┃ ┗ db.go
 ┃ ┃ ┣ grpc
 ┃ ┃ ┗ http
 ┃ ┗ util
 ┃ ┃ ┣ random.go
 ┃ ┃ ┗ testUtils.go
 ┣ migrations
 ┃ ┣ 000001_init_schema.down.sql
 ┃ ┗ 000001_init_schema.up.sql
 ┣ .env
 ┣ LICENSE
 ┣ Makefile
 ┣ README.md
 ┣ go.mod
 ┣ go.sum
 ┣ go.work
 ┗ go.work.sum
```

## Funcionalidades
- Backend en Go estándar.
- Interfaz de usuario adaptable a cualquier tecnología frontend.
- Creación rápida de interfaces de usuario para programas en Go.
- Multiplataforma con motores de renderizado nativos.

### Built With
- ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## Empezando
Instrucciones para configurar el proyecto localmente.

### Prerrequisitos
- Docker y Postgres.
- Herramientas de migración de base de datos.

### Instalación
Pasos para clonar el repositorio y configurar el entorno de desarrollo.

## Uso
Ejemplos de cómo se puede utilizar el proyecto. Enlace a la documentación para más ejemplos.

## Contribuyendo
Instrucciones para contribuir al proyecto. Incentivar a la comunidad a participar.

## Licencia
Distribuido bajo la Licencia MIT. Ver `LICENSE.txt` para más información.

## Contacto
Tus detalles de contacto y enlace al proyecto en GitHub.

## Agradecimientos
Espacio para agradecer y listar recursos útiles.

