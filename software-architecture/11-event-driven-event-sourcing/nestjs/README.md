# 11-event-driven-event-sourcing · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── src
│   ├── notifications
│   │   ├── listeners
│   │   │   └── policy-issued.listener.ts
│   │   └── notifications.module.ts
│   ├── policies
│   │   ├── domain
│   │   │   └── policy.aggregate.ts
│   │   ├── events
│   │   │   └── policy.events.ts
│   │   ├── infrastructure
│   │   │   ├── event-store.ts
│   │   │   └── policy.repository.ts
│   │   ├── issue-policy.dto.ts
│   │   ├── issue-policy.use-case.ts
│   │   ├── policies.module.ts
│   │   ├── policy.controller.ts
│   │   ├── renew-policy.dto.ts
│   │   └── renew-policy.use-case.ts
│   ├── app.module.ts
│   └── main.ts
├── package.json
└── tsconfig.json
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`):

```bash
pnpm install
pnpm --filter arq-11-event-driven-event-sourcing build
pnpm --filter arq-11-event-driven-event-sourcing start        # PORT=3000 por defecto
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3000/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3000/policies/<id>

# Renovar: añade un evento al stream
curl -i -X POST localhost:3000/policies/<id>/renew -H "content-type: application/json" -d '{"months":12}'

# Historial de eventos del agregado
curl -i localhost:3000/policies/<id>/events
```
