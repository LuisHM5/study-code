# 11. Event-Driven y Event Sourcing en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs
  notifications.rs       -> suscriptor del bus (tokio broadcast)
  policies/
    mod.rs
    events.rs            -> enum PolicyEvent
    aggregate.rs         -> Policy: decidir (issue, renew) y aplicar (fold)
    event_store.rs       -> streams en memoria, concurrencia optimista
    service.rs           -> carga, decide, guarda eventos y publica
    http.rs
```

El estado de una póliza no se guarda: `GET /policies/{id}` lo reconstruye reproduciendo sus eventos.

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-11-event-driven-event-sourcing            # escucha en el puerto 3000; cámbialo con PORT=3005
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

Endpoints adicionales:

```bash
curl -i -X POST localhost:3000/policies/<id>/renew \
  -H 'Content-Type: application/json' -d '{"months":12}'
# 200 con la póliza y su nueva endDate

curl -i localhost:3000/policies/<id>/events
# 200 [{"type":"PolicyIssued",...},{"type":"PolicyRenewed","newEndDate":"..."}]
```
