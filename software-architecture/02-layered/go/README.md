# Capas por tipo técnico en Go

```
cmd/api/main.go
go.mod
internal/handlers/policy_handler.go
internal/models/policy.go
internal/repositories/policy_repository.go
internal/services/policy_service.go
```

Flujo: `handlers` -> `services` -> `repositories`, con `models` compartido.

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
