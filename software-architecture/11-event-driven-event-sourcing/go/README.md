# Event-Driven y Event Sourcing en Go

```
cmd/api/main.go
go.mod
internal/notifications/subscriber.go
internal/platform/eventbus/bus.go
internal/policies/date.go
internal/policies/events.go
internal/policies/eventstore.go
internal/policies/http.go
internal/policies/policy.go
internal/policies/service.go
```

`policy.go` es un agregado cuyo estado se reconstruye desde `eventstore.go`. Tras guardar, `service.go` publica los eventos en `platform/eventbus` y `notifications` reacciona.

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
curl -i -X POST localhost:3000/policies/<id>/renew -H "Content-Type: application/json" -d '{"months":12}'
curl -i localhost:3000/policies/<id>/events
```
