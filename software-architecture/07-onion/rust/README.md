# 07. Onion en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs
  domain/                      -> anillo central
    mod.rs                     -> DomainError y re-exports
    policy.rs                  -> entidad
    term.rs                    -> value object
    number.rs                  -> value object
    uniqueness.rs              -> servicio de dominio
    repository.rs              -> interfaz del repositorio
  application/                 -> servicios de aplicación
    mod.rs                     -> trait PolicyNotifier
    issue_policy.rs
    get_policy.rs
  infrastructure/              -> implementa las interfaces
    mod.rs
    memory_repository.rs
    console_notifier.rs
  presentation/
    mod.rs
    http.rs
```

Es un solo crate con un módulo por anillo. En un proyecto grande cada anillo sería un crate (`policy-domain`, `policy-application`, ...) y el `Cargo.toml` de cada uno haría cumplir la regla de dependencias.

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-07-onion            # escucha en el puerto 3000; cámbialo con PORT=3005
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

