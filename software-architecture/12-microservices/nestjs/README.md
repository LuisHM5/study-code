# 12-microservices · NestJS

Proyecto mínimo y ejecutable. Persistencia en memoria, sin base de datos. La teoría está en [`../README.md`](../README.md).

## Estructura

```
├── apps
│   ├── notifications
│   │   └── src
│   │       ├── main.ts
│   │       ├── notifications.controller.ts
│   │       └── notifications.module.ts
│   └── policies
│       └── src
│           ├── issue-policy.dto.ts
│           ├── main.ts
│           ├── notifications.client.ts
│           ├── policies.controller.ts
│           ├── policies.module.ts
│           └── policies.service.ts
├── libs
│   └── contracts
│       └── src
│           └── policy-issued.event.ts
├── package.json
└── tsconfig.json
```

## Ejecutar

Desde la carpeta del tema (`software-architecture/`):

```bash
pnpm install
pnpm --filter arq-12-microservices build

# Dos procesos, cada uno en su terminal
pnpm --filter arq-12-microservices start:notifications   # PORT=3002 por defecto
pnpm --filter arq-12-microservices start:policies        # PORT=3001, NOTIFICATIONS_URL=http://localhost:3002
```

## Probar

```bash
# Emitir una póliza -> 201 {"id":"..."}  (400 si la vigencia es inválida, 409 si el número ya existe)
curl -i -X POST localhost:3001/policies -H "content-type: application/json" \
  -d '{"number":"POL-001","startDate":"2026-01-01","endDate":"2026-12-31"}'

# Consultar -> 200 {id, number, startDate, endDate}  (404 si no existe)
curl -i localhost:3001/policies/<id>

# Detén notifications y vuelve a emitir: policies responde 201 y solo registra el fallo en el log
```
