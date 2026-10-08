# 03. Modular por feature en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs                -> arma los módulos
  policy/
    mod.rs               -> API pública: router(), PolicyService, Notifier
    handler.rs
    service.rs
    repository.rs        -> privado del módulo
    model.rs
  notification/
    mod.rs               -> implementa policy::Notifier
  platform/
    mod.rs
    error.rs             -> AppError compartido
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-03-feature-modules            # escucha en el puerto 3000; cámbialo con PORT=3005
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

