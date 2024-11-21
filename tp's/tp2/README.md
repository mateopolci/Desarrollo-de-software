# Objetivo

Desarrollar íntegramente una API web con persistencia de datos en un sistema de gestión de base de datos (SGBD).

## 1. Web API

RESTful de dos endpoints (como mínimo) que posean algún tipo de relación. DTOs para intercambio de datos hacia el cliente y entre servicios (internos). persistencia en base de datos PostgreSQL con usuario limitado. Metodología Code First.

## `main.go`

Se encarga de iniciar el servidor y cargar las rutas.

## `database.go`

Define la conexión a la base de datos y realiza la migración de los modelos.

## `models.go`

Definir las estructuras de datos (modelos) para la base de datos.

## `cafe_controller.go`

Define las funciones que controlan los endpoints de la API. Estos se encargan de manejar las solicitudes HTTP.