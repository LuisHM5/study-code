# Clean Architecture en Go

```
cmd/api/main.go
go.mod
internal/policies/adapter/controller/http.go
internal/policies/adapter/gateway/console_notifier.go
internal/policies/adapter/gateway/memory.go
internal/policies/adapter/presenter/policy.go
internal/policies/entity/policy.go
internal/policies/infra/router.go
internal/policies/usecase/gateways.go
internal/policies/usecase/get_policy.go
internal/policies/usecase/issue_policy.go
```

Círculos: `entity` (1), `usecase` (2, con Input/Output y gateways), `adapter` (3: controller, presenter, gateway) e `infra` (4, arma todo).

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
