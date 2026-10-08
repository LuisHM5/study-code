# 04-vertical-slice · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── src
│   ├── notification
│   │   └── notification.service.ts
│   ├── policy
│   │   ├── domain
│   │   │   └── policy-term.ts
│   │   ├── entities
│   │   │   ├── policy.entity.ts
│   │   │   └── policy.store.ts
│   │   ├── get-policy
│   │   │   └── get-policy.controller.ts
│   │   ├── issue-policy
│   │   │   ├── issue-policy.controller.ts
│   │   │   ├── issue-policy.dto.ts
│   │   │   └── issue-policy.handler.ts
│   │   ├── renew-policy
│   │   │   ├── renew-policy.controller.ts
│   │   │   ├── renew-policy.dto.ts
│   │   │   └── renew-policy.handler.ts
│   │   └── policy.module.ts
│   ├── app.module.ts
│   └── main.ts
├── package.json
└── tsconfig.json
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`):

```bash
pnpm install
pnpm --filter arq-04-vertical-slice build
pnpm --filter arq-04-vertical-slice start        # PORT=3000 por defecto
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3000/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3000/policies/<id>

# Renovar (otro slice)
curl -i -X POST localhost:3000/policies/<id>/renew -H "content-type: application/json" -d '{"months":12}'
```
