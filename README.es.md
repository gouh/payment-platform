# Payment Platform

Este proyecto es una aplicación de backend desarrollada en Go, utilizando el framework Gin para el manejo de solicitudes HTTP y `pgx` para la interacción con una base de datos PostgreSQL. La aplicación sirve como una plataforma de pagos, permitiendo a los comerciantes procesar transacciones, gestionar pagos y realizar consultas sobre detalles de pagos anteriores.

## Características

- **API RESTful**: La aplicación expone una API RESTful para el procesamiento de pagos, la consulta de detalles de pagos y la gestión de comerciantes.
- **Gestión de Pagos**: Los comerciantes pueden procesar pagos y consultar detalles de pagos previos.
- **Tokenización de Tarjetas**: Implementa la tokenización de información de tarjetas de crédito para mejorar la seguridad.
- **Simulación Bancaria**: Incluye un simulador bancario para simular interacciones con bancos adquirentes. Se ha configurado un porcentaje de aprobación del 40% para las transacciones simuladas, lo que significa que alrededor del 40% de los pagos procesados a través del simulador serán aprobados, mientras que el resto se considerarán fallidos.
- **Paginación y Filtro**: Soporta paginación y filtros en consultas de pagos para una mejor gestión de datos.
- **Middleware Personalizado**: Usa middlewares personalizados para el manejo de errores y logs en formato JSON. Esto incluye un middleware de recuperación que captura y loguea errores no capturados en el formato deseado, y un middleware para personalizar las respuestas de error y asegurar que siempre se devuelvan en formato JSON, incluso cuando ocurren errores de binding de datos.
- **Autorización de Pagos**: La lógica de autorización de pagos simula la comunicación con sistemas de autorización bancaria, utilizando un enfoque probabilístico para determinar la aprobación o rechazo de transacciones de pago basado en el porcentaje configurado.

## Tecnologías Utilizadas

- **Go (Golang)**: Lenguaje de programación para el desarrollo del backend.
- **Gin**: Framework web utilizado para crear la API RESTful.
- **pgx y pgxpool**: Bibliotecas para la conexión y operaciones con PostgreSQL.
- **goqu**: Biblioteca constructora de consultas SQL para Go, utilizada para construir consultas dinámicas.
- **logrus**: Biblioteca de logging para Go, configurada para imprimir logs en formato JSON.
- **PostgreSQL**: Sistema de gestión de bases de datos relacional utilizado para almacenar datos de comerciantes y pagos.

## Modelo de base de datos

![payment-platform es](https://github.com/gouh/payment-platform/assets/13145599/1ece23cb-b9ac-44df-9b67-f56cf232fc7d)

## Estructura del Proyecto

```
.
├── cmd
│   └── server
│       └── main.go
├── config
│   ├── auth.go
│   ├── config.go
│   ├── database_config.go
│   └── init.sql
├── docker
│   └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── container
│   │   ├── container.go
│   │   ├── database_dialect.go
│   │   └── database.go
│   ├── dao
│   │   ├── customer_dao.go
│   │   ├── merchant_dao.go
│   │   ├── payment_dao.go
│   │   └── tokenized_cards_dao.go
│   ├── dto
│   ├── http
│   │   ├── handlers
│   │   │   ├── auth_handler.go
│   │   │   ├── customer_handler.go
│   │   │   ├── health_handler.go
│   │   │   ├── merchant_handler.go
│   │   │   ├── payment_handler.go
│   │   │   └── tokenized_card_handler.go
│   │   ├── middleware
│   │   │   ├── auth_middleware.go
│   │   │   ├── cors_middleware.go
│   │   │   ├── error_middleware.go
│   │   │   └── recovery_middleware.go
│   │   ├── routes
│   │   │   ├── auth_routes.go
│   │   │   ├── customer_routes.go
│   │   │   ├── health_routes.go
│   │   │   ├── merchant_routes.go
│   │   │   ├── payment_routes.go
│   │   │   └── tokenized_card_routes.go
│   │   └── routes.go
│   ├── models
│   │   ├── customer.go
│   │   ├── merchant.go
│   │   ├── payment.go
│   │   └── tokenized_card.go
│   ├── requests
│   │   ├── auth_request.go
│   │   ├── customer_request.go
│   │   ├── merchant_request.go
│   │   ├── pagination_request.go
│   │   ├── payment_request.go
│   │   └── tokenized_card_request.go
│   ├── responses
│   │   ├── common_response.go
│   │   ├── customer_response.go
│   │   ├── health_response.go
│   │   ├── merchant_response.go
│   │   ├── payment_response.go
│   │   └── tokenized_card_response.go
│   ├── services
│   └── utils
│       ├── bank_simulator.go
│       ├── body_validator.go
│       └── card_tokenizer.go
├── LICENSE
└── README.md

```

## Instrucciones de Instalación

1. **Clonar el Repositorio**

```bash
git clone https://github.com/gouh/payment-platform.git
cd payment-platform
```

2. **Levantar los Servicios con Docker Compose**

```bash
docker-compose up -d
```

Este comando inicia todos los servicios necesarios (backend, postgresql) en contenedores Docker.

## Acceder al Proyecto

Una vez que los contenedores estén arriba y corriendo, podrás acceder al API a través de:

| Método | Endpoint                            | Descripción                                      | Autenticación Requerida |
|--------|-------------------------------------|--------------------------------------------------|-------------------------|
| POST   | `/auth/login`                       | Inicia sesión y devuelve un token de acceso.     | No                      |
| GET    | `/v1/health`                        | Verifica el estado de salud del servicio.        | Sí                      |
| GET    | `/v1/merchants`                     | Obtiene una lista de comerciantes.               | Sí                      |
| GET    | `/v1/merchants/:id`                 | Obtiene un comerciante por su ID.                | Sí                      |
| POST   | `/v1/merchants`                     | Crea un nuevo comerciante.                       | Sí                      |
| PATCH  | `/v1/merchants/:id`                 | Actualiza los datos de un comerciante por su ID. | Sí                      |
| DELETE | `/v1/merchants/:id`                 | Elimina un comerciante por su ID.                | Sí                      |
| GET    | `/v1/customers`                     | Obtiene una lista de clientes.                   | Sí                      |
| GET    | `/v1/customers/:id`                 | Obtiene un cliente por su ID.                    | Sí                      |
| POST   | `/v1/customers`                     | Crea un nuevo cliente.                           | Sí                      |
| PATCH  | `/v1/customers/:id`                 | Actualiza los datos de un cliente por su ID.     | Sí                      |
| DELETE | `/v1/customers/:id`                 | Elimina un cliente por su ID.                    | Sí                      |
| GET    | `/v1/customers/:id/cards`           | Obtiene las tarjetas tokenizadas de un cliente.  | Sí                      |
| GET    | `/v1/customers/:id/cards/:token`    | Obtiene una tarjeta tokenizada por su token.     | Sí                      |
| POST   | `/v1/customers/:id/cards`           | Crea una nueva tarjeta tokenizada para un cliente.| Sí                    |
| DELETE | `/v1/customers/:id/cards/:token`    | Elimina una tarjeta tokenizada por su token.     | Sí                      |
| GET    | `/v1/payments`                      | Obtiene una lista de pagos.                      | Sí                      |
| GET    | `/v1/payments/:id`                  | Obtiene un pago por su ID.                       | Sí                      |
| POST   | `/v1/payments`                      | Procesa un nuevo pago.                           | Sí                      |

## Despliegue

Este proyecto está configurado para facilitar su despliegue con Docker, asegurando una instalación y ejecución consistentes en cualquier entorno.

## Contribuciones

Las contribuciones son bienvenidas. Si tienes alguna sugerencia para mejorar la aplicación, por favor, considera enviar un pull request o abrir un issue en el repositorio.

## Licencia

Este proyecto está licenciado bajo la MIT License - vea el archivo [LICENSE](LICENSE) para más detalles.

---

Desarrollado con ❤ por Hugo.
