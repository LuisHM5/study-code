# Vertical Slice en Go

```
cmd/api/main.go
go.mod
internal/policy/get_policy.go
internal/policy/issue_policy.go
internal/policy/policy.go
internal/policy/routes.go
```

Un archivo por caso de uso (`issue_policy.go`, `get_policy.go`). `policy.go` tiene solo lo compartido y `routes.go` conecta cada slice con su ruta.

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
