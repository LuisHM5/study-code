# 05-screaming · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── src
│   ├── policies
│   │   ├── issuing
│   │   │   ├── issue-policy.controller.ts
│   │   │   ├── issue-policy.dto.ts
│   │   │   └── issue-policy.use-case.ts
│   │   ├── lookup
│   │   │   ├── find-policy.controller.ts
│   │   │   └── find-policy.use-case.ts
│   │   ├── policies.module.ts
│   │   ├── policy.entity.ts
│   │   └── policy.store.ts
│   ├── shared
│   │   └── notifications
│   │       └── notification.service.ts
│   ├── app.module.ts
│   └── main.ts
├── package.json
└── tsconfig.json
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`):

```bash
pnpm install
pnpm --filter arq-05-screaming build
pnpm --filter arq-05-screaming start        # PORT=3000 por defecto
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3000/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3000/policies/<id>

```
