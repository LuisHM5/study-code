# 09. Monolito modular en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs                -> arma los módulos y el bus de eventos
  platform.rs            -> utilidades técnicas compartidas
  policies/
    mod.rs               -> API PÚBLICA: PoliciesApi, PolicySummary, PolicyIssued, build()
    domain.rs            -> privado
    app.rs               -> privado
    store.rs             -> privado
    http.rs              -> privado
  claims/
    mod.rs               -> build(policies: Arc<dyn PoliciesApi>)
    app.rs
    http.rs
  notifications/
    mod.rs               -> suscrito a PolicyIssued (tokio broadcast)
```

Es un solo crate. Los submódulos de `policies` se declaran sin `pub`, así que `claims` no puede importarlos: solo ve lo que expone `policies/mod.rs`. En un proyecto grande cada módulo sería un crate del workspace.

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-09-modular-monolith            # escucha en el puerto 3000; cámbialo con PORT=3005
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
curl -i -X POST localhost:3000/claims \
  -H 'Content-Type: application/json' \
  -d '{"policyId":"<id de la póliza>","description":"choque"}'
# 201 {"id":"..."} · 422 si la póliza no existe

curl -i localhost:3000/claims/<id>
```
