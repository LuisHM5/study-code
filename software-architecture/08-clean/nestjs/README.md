# 08-clean · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── src
│   ├── policies
│   │   ├── application
│   │   │   ├── gateways
│   │   │   │   ├── id-generator.ts
│   │   │   │   ├── notification.gateway.ts
│   │   │   │   └── policy.gateway.ts
│   │   │   ├── get-policy
│   │   │   │   ├── get-policy.output.ts
│   │   │   │   └── get-policy.use-case.ts
│   │   │   └── issue-policy
│   │   │       ├── issue-policy.input.ts
│   │   │       ├── issue-policy.output.ts
│   │   │       └── issue-policy.use-case.ts
│   │   ├── domain
│   │   │   ├── policy.entity.ts
│   │   │   └── policy.errors.ts
│   │   ├── infrastructure
│   │   │   └── policies.module.ts
│   │   └── interface-adapters
│   │       ├── controllers
│   │       │   ├── domain-error.filter.ts
│   │       │   ├── issue-policy.request.ts
│   │       │   └── policy.controller.ts
│   │       ├── gateways
│   │       │   ├── console-notification.gateway.ts
│   │       │   ├── in-memory-policy.gateway.ts
│   │       │   └── uuid-generator.ts
│   │       └── presenters
│   │           └── policy.presenter.ts
│   ├── app.module.ts
│   └── main.ts
├── package.json
└── tsconfig.json
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`):

```bash
pnpm install
pnpm --filter arq-08-clean build
pnpm --filter arq-08-clean start        # PORT=3000 por defecto
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3000/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3000/policies/<id>

```
