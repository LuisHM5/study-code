# 09-modular-monolith · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── src
│   ├── modules
│   │   ├── claims
│   │   │   ├── internal
│   │   │   │   ├── claim.controller.ts
│   │   │   │   ├── claim.repository.ts
│   │   │   │   ├── claim.ts
│   │   │   │   ├── report-claim.dto.ts
│   │   │   │   └── report-claim.use-case.ts
│   │   │   ├── claims.module.ts
│   │   │   └── index.ts
│   │   ├── notifications
│   │   │   ├── index.ts
│   │   │   ├── notifications.module.ts
│   │   │   └── policy-issued.listener.ts
│   │   └── policies
│   │       ├── internal
│   │       │   ├── application
│   │       │   │   └── issue-policy.use-case.ts
│   │       │   ├── domain
│   │       │   │   └── policy.ts
│   │       │   └── infrastructure
│   │       │       ├── issue-policy.dto.ts
│   │       │       ├── policy.controller.ts
│   │       │       └── policy.repository.ts
│   │       ├── public
│   │       │   ├── events
│   │       │   │   └── policy-issued.event.ts
│   │       │   ├── policies.facade.ts
│   │       │   └── policy.dto.ts
│   │       ├── index.ts
│   │       └── policies.module.ts
│   ├── app.module.ts
│   └── main.ts
├── package.json
└── tsconfig.json
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`):

```bash
pnpm install
pnpm --filter arq-09-modular-monolith build
pnpm --filter arq-09-modular-monolith start        # PORT=3000 por defecto
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3000/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3000/policies/<id>

# Reportar un siniestro: claims consulta a policies solo por su fachada (422 si la póliza no existe)
curl -i -X POST localhost:3000/claims -H "content-type: application/json" \
  -d '{"policyId":"<id>","description":"choque"}'
```
