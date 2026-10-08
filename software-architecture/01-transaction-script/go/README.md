# Transaction Script en Go

```
cmd/api/main.go
go.mod
```

Todo vive en `cmd/api/main.go`: cada handler valida, guarda y responde.

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
