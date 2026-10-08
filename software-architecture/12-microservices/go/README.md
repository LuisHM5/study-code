# Microservicios en Go

```
contracts/notifications.go
go.mod
services/notifications/cmd/server/main.go
services/policies/cmd/server/main.go
services/policies/internal/app/service.go
services/policies/internal/domain/policy.go
services/policies/internal/httpapi/handler.go
services/policies/internal/memory/repository.go
services/policies/internal/notificationsclient/client.go
```

Dos servicios con proceso propio. `policies` llama a `notifications` por **HTTP + JSON** con timeout de 2 s; si falla, solo lo registra en el log. Se usa HTTP y no gRPC para no añadir dependencias; `contracts/` hace el papel del `.proto`.

## Ejecutar

```bash
go run ./services/notifications/cmd/server   # PORT=3002 por defecto
go run ./services/policies/cmd/server        # PORT=3001, NOTIFICATIONS_URL=http://localhost:3002
```

Requiere Go 1.24 o superior. Sin dependencias externas; los datos viven en memoria.

## Probar

```bash
curl -i -X POST localhost:3001/policies -H "Content-Type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'
curl -i localhost:3001/policies/<id>
```
