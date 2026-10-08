# CQRS en Go

```
cmd/api/main.go
go.mod
internal/policies/adapters/console_notifier.go
internal/policies/adapters/http.go
internal/policies/adapters/memory_db.go
internal/policies/adapters/memory_read_model.go
internal/policies/adapters/memory_repository.go
internal/policies/app/app.go
internal/policies/app/command/issue_policy.go
internal/policies/app/query/get_policy.go
internal/policies/app/query/list_expiring.go
internal/policies/app/query/read_model.go
internal/policies/domain/policy.go
```

Sin bus: `app/command` cambia estado pasando por `domain`; `app/query` lee vistas desde un read model (`adapters/memory_read_model.go`) distinto del repositorio de escritura. Ambos comparten `MemoryDB` (CQRS nivel 1).

## Ejecutar

```bash
go run ./cmd/api        # PORT=3000 por defecto
```

Requiere Go 1.24 o superior. Sin dependencias externas; los datos viven en memoria.

## Probar

```bash
curl -i -X POST localhost:3000/policies -H "Content-Type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'
curl -i localhost:3000/policies/<id>
curl -i "localhost:3000/policies/expiring?days=365"
```
