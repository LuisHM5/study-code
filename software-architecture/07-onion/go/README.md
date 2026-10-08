# Onion en Go

```
cmd/api/main.go
go.mod
internal/policy/application/get_policy.go
internal/policy/application/issue_policy.go
internal/policy/domain/errors.go
internal/policy/domain/number.go
internal/policy/domain/policy.go
internal/policy/domain/repository.go
internal/policy/domain/term.go
internal/policy/domain/uniqueness.go
internal/policy/infrastructure/console/notifier.go
internal/policy/infrastructure/memory/repository.go
internal/policy/presentation/httpapi/handler.go
```

Anillos de adentro hacia afuera: `domain` (entidad, value objects `Term` y `Number`, servicio de dominio `Uniqueness`, interfaz `Repository`), `application`, y por fuera `infrastructure` y `presentation`.

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
