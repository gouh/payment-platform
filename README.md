# Payment Platform

## Características Principales

## Tecnologías Utilizadas

- **Backend**: Go
- **Persistencia de Datos**: Postgresql
- **Contenerización**: Docker

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

Una vez que los contenedores estén arriba y corriendo, podrás acceder a la aplicación web a través de:

- **API Backend**: http://localhost:8080/api/v1

## Despliegue

Este proyecto está configurado para facilitar su despliegue con Docker, asegurando una instalación y ejecución consistentes en cualquier entorno.

## Contribuciones

Las contribuciones son bienvenidas. Si tienes alguna sugerencia para mejorar la aplicación, por favor, considera enviar un pull request o abrir un issue en el repositorio.

## Licencia

Este proyecto está licenciado bajo la MIT License - vea el archivo [LICENSE.md](LICENSE.md) para más detalles.

---

Desarrollado con ❤ por Hugo.
