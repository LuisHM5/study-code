# 02. Capas por tipo técnico en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs                  -> rutas y cableado
  error.rs                 -> AppError -> respuesta HTTP
  handlers/policy.rs       -> HTTP
  services/policy.rs       -> lógica de negocio
  repositories/policy.rs   -> datos (en memoria)
  models/policy.rs         -> Policy, NewPolicy
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-02-layered            # escucha en el puerto 3000; cámbialo con PORT=3005
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

