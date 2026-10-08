# Hexagonal en Go

```
cmd/api/main.go
go.mod
internal/policy/adapters/console/notifier.go
internal/policy/adapters/httpapi/handler.go
internal/policy/adapters/memory/repository.go
internal/policy/app/get_policy.go
internal/policy/app/issue_policy.go
internal/policy/app/ports.go
internal/policy/domain/policy.go
```

`domain` y `app` son el núcleo; `app/ports.go` define los puertos y `adapters/` los implementa (`httpapi` de entrada; `memory` y `console` de salida). El cableado está en `cmd/api/main.go`.

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
```
