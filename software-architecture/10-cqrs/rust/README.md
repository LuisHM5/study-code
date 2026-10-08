# 10. CQRS en Rust

Proyecto real y mínimo con axum. Persistencia en memoria, sin base de datos.

## Estructura

```
Cargo.toml
src/
  main.rs
  policies/
    mod.rs
    domain.rs                  -> modelo de escritura
    commands/
      mod.rs
      issue_policy.rs          -> comando + handler
    queries/
      mod.rs
      get_policy.rs
      list_expiring.rs
    infra/
      mod.rs
      write_repo.rs            -> almacén de escritura
      read_model.rs            -> almacén de lectura (PolicyView)
    http.rs
```

Sin bus de comandos. El handler del comando proyecta la póliza al modelo de lectura de forma síncrona; todas las consultas (incluido `GET /policies/{id}`) leen de ese modelo.

## Ejecutar

Desde la carpeta del tema (`software-architecture/`, donde está el workspace de Cargo):

```bash
cargo run -p arq-10-cqrs            # escucha en el puerto 3000; cámbialo con PORT=3005
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

Endpoint adicional (servido por el modelo de lectura):

```bash
curl -i 'localhost:3000/policies/expiring?days=365'
# 200 [{"id","number","startDate","endDate"}, ...]  (days por defecto: 30)
```
