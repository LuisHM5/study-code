# 03-feature-modules · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── src
│   ├── notification
│   │   ├── services
│   │   │   └── notification.service.ts
│   │   └── notification.module.ts
│   ├── policy
│   │   ├── controllers
│   │   │   └── policy.controller.ts
│   │   ├── dto
│   │   │   └── issue-policy.dto.ts
│   │   ├── entities
│   │   │   └── policy.entity.ts
│   │   ├── repositories
│   │   │   └── policy.repository.ts
│   │   ├── services
│   │   │   └── policy.service.ts
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
pnpm --filter arq-03-feature-modules build
pnpm --filter arq-03-feature-modules start        # PORT=3000 por defecto
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3000/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3000/policies/<id>

```
