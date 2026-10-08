# Modular por feature en Go

```
cmd/api/main.go
go.mod
internal/notification/service.go
internal/platform/web.go
internal/policy/handler.go
internal/policy/policy.go
internal/policy/repository.go
internal/policy/service.go
```

Un paquete por feature (`policy`, `notification`); `platform` tiene lo técnico compartido. `policy` usa notificaciones a través de la interfaz `Notifier` que él mismo declara.

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
