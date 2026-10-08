# Monolito modular en Go

```
cmd/api/main.go
go.mod
modules/claims/internal/app/report_claim.go
modules/claims/internal/httpapi/routes.go
modules/claims/module.go
modules/notifications/module.go
modules/policies/api.go
modules/policies/internal/app/service.go
modules/policies/internal/domain/policy.go
modules/policies/internal/httpapi/routes.go
modules/policies/internal/memory/repository.go
modules/policies/module.go
platform/eventbus/bus.go
platform/web/web.go
```

Cada módulo expone solo `api.go`/`module.go`; su interior está en `internal/`, que el compilador impide importar desde otro módulo. `claims` usa `policies.API` y `notifications` escucha el evento `policy.issued` del bus.

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
curl -i -X POST localhost:3000/claims -H "Content-Type: application/json" \
  -d '{"policyId":"<id>","description":"choque"}'      # 422 si la póliza no existe
```
