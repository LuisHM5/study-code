# 06. Hexagonal (Ports & Adapters) en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs                -> conecta cada puerto con su adaptador
  policy/
    mod.rs
    domain.rs            -> núcleo: Policy y sus reglas
    ports.rs             -> núcleo: traits PolicyRepository y Notifier
    issue_policy.rs      -> núcleo: caso de uso
    get_policy.rs        -> núcleo: caso de uso
    adapters/
      mod.rs
      http.rs            -> adaptador de entrada
      memory.rs          -> adaptador de salida (persistencia)
      console.rs         -> adaptador de salida (notificación)
```

Los adaptadores `memory.rs` y `console.rs` ocupan el lugar de `postgres.rs` y `smtp.rs`: cambiarlos no toca el núcleo.

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-06-hexagonal            # escucha en el puerto 3000; cámbialo con PORT=3005
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

