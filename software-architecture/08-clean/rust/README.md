# 08. Clean Architecture en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs                              -> círculo 4: framework y cableado
  domain/                              -> círculo 1: entidades
    mod.rs
    policy.rs
  application/                         -> círculo 2: casos de uso
    mod.rs                             -> AppError
    gateways.rs                        -> traits PolicyGateway, NotificationGateway
    issue_policy.rs                    -> Input, Output, IssuePolicy<P, N>
    get_policy.rs
  adapters/                            -> círculo 3: adaptadores de interfaz
    mod.rs
    http.rs                            -> controllers y DTO de HTTP
    presenter.rs                       -> vistas y errores HTTP
    memory_policy_gateway.rs
    console_notification_gateway.rs
```

Es un solo crate con un módulo por círculo. En un proyecto grande cada círculo sería un crate (`domain`, `application`, `adapters`, `api`).

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-08-clean            # escucha en el puerto 3000; cámbialo con PORT=3005
```

## Probar

```bash
curl -i -X POST localhost:3000/policies \
  -H 'Content-Type: application/json' \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'
# 201 {"id":"..."}   · 400 si endDate <= startDate · 409 si el número ya existe

curl -i localhost:3000/policies/<id>
# 200 {"id","number","startDate","endDate"} · 404 si no existe
```

