# 04. Vertical Slice en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs
  app.rs                 -> AppContext (dependencias) y AppError
  policy/
    mod.rs               -> router: una línea por slice
    domain.rs            -> Policy y validate_term, lo único compartido
    issue_policy.rs      -> slice: Request + Response + handler
    get_policy.rs        -> slice: Response + handler
```

No se incluye el slice `renew_policy` que menciona la guía: añadirlo sería crear `policy/renew_policy.rs` y una línea en `policy/mod.rs`.

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-04-vertical-slice            # escucha en el puerto 3000; cámbialo con PORT=3005
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

