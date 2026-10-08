# 10-cqrs · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── src
│   ├── policies
│   │   ├── commands
│   │   │   └── issue-policy
│   │   │       ├── issue-policy.command.ts
│   │   │       └── issue-policy.handler.ts
│   │   ├── domain
│   │   │   ├── policy-issued.event.ts
│   │   │   └── policy.ts
│   │   ├── infrastructure
│   │   │   ├── policy-issued.notifier.ts
│   │   │   ├── policy-read-model.projector.ts
│   │   │   ├── policy-read-model.ts
│   │   │   └── policy.repository.ts
│   │   ├── queries
│   │   │   ├── get-policy
│   │   │   │   ├── get-policy.handler.ts
│   │   │   │   └── get-policy.query.ts
│   │   │   └── list-expiring-policies
│   │   │       ├── expiring-policy.view.ts
│   │   │       ├── list-expiring-policies.handler.ts
│   │   │       └── list-expiring-policies.query.ts
│   │   ├── issue-policy.dto.ts
│   │   ├── policies.module.ts
│   │   └── policy.controller.ts
│   ├── app.module.ts
│   └── main.ts
├── package.json
└── tsconfig.json
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`):

```bash
pnpm install
pnpm --filter arq-10-cqrs build
pnpm --filter arq-10-cqrs start        # PORT=3000 por defecto
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3000/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3000/policies/<id>

# Consulta servida por el modelo de lectura
curl -i "localhost:3000/policies/expiring?days=30"
```
