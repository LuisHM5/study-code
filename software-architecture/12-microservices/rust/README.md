# 12. Microservicios en Rust

Dos servicios reales y mínimos con axum, cada uno con su propio binario. Persistencia en memoria.

Para mantenerlo simple se comunican por **HTTP + JSON** (con `reqwest`), no por gRPC/tonic como sugiere la guía.

## Estructura

```
policies/                      -> servicio de pólizas (puerto 3001)
  Cargo.toml
  src/
    main.rs
    http.rs
    policy.rs                  -> Policy y PolicyStore (datos propios del servicio)
    notifications_client.rs    -> cliente HTTP con timeout de 2 s
notifications/                 -> servicio de notificaciones (puerto 3002)
  Cargo.toml
  src/
    main.rs                    -> POST /notifications
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`), en dos terminales:

```bash
cargo run -p arq-12-microservices-notifications   # PORT=3002 por defecto
cargo run -p arq-12-microservices-policies        # PORT=3001 por defecto
```

`policies` busca al otro servicio en `NOTIFICATIONS_URL` (por defecto `http://localhost:3002`).

## Probar

```bash
curl -i -X POST localhost:3001/policies \
  -H 'Content-Type: application/json' \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'
# 201 {"id":"..."}   · 400 si endDate <= startDate · 409 si el número ya existe
# En la terminal de notifications aparece: [notificación] póliza emitida POL-001 (id ...)

curl -i localhost:3001/policies/<id>
```

Apaga `notifications` y vuelve a emitir: `policies` responde 201 igual y solo registra el fallo en su log.
