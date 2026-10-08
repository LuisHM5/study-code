# 11. Event-Driven y Event Sourcing

**Complejidad: 3/5 (eventos en proceso) a 5/5 (event sourcing)**

Son dos ideas distintas que se confunden seguido.

## Parte A: Event-Driven Architecture (EDA)

### Idea

Los componentes no se llaman entre sí: **publican hechos** ("se emitió la póliza") y otros **reaccionan**. Quien publica no sabe quién escucha.

```
                                  ┌─> Notificaciones: envía correo
Pólizas ── PolicyIssued ──> bus ──┼─> Pagos: genera calendario de primas
                                  └─> Auditoría: registra movimiento
```

Niveles:

| Nivel | Transporte | Complejidad |
|---|---|---|
| 1. Eventos en proceso | `EventEmitter` en memoria, mismo proceso | Baja |
| 2. Eventos durables | Tabla outbox + worker, o cola (BullMQ/Redis) | Media |
| 3. Broker distribuido | Kafka, RabbitMQ, NATS entre servicios | Alta |

### Lo mejor

- Desacopla: añadir una reacción nueva no toca al emisor.
- El caso de uso principal queda corto; los efectos secundarios salen de él.
- Trabajo lento (correos, PDFs, sincronizaciones) fuera de la petición HTTP.
- Es el pegamento natural entre módulos de un [modular monolith](../09-modular-monolith/README.md) y entre [microservices](../12-microservices/README.md).

### Lo peor

- El flujo deja de leerse de arriba a abajo: "¿qué pasa cuando se emite una póliza?" obliga a buscar todos los listeners.
- Depurar es más difícil: fallos asíncronos, lejos de la causa.
- Entrega **al menos una vez**: los handlers deben ser idempotentes porque un evento puede llegar repetido.
- Orden no garantizado, consistencia eventual, reintentos, colas de mensajes muertos.
- Si publicas el evento y la transacción falla (o al revés), quedas inconsistente. La solución es el patrón **outbox**.

### Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir una reacción a algo existente | Muy bajo | Un listener nuevo |
| Añadir un flujo nuevo | Medio | Diseñar eventos y contratos |
| Modificar el contenido de un evento | Alto | Es un contrato con todos los consumidores; hay que versionar |
| Fix de bug | Medio a alto | Requiere trazas y correlation IDs |
| Onboarding | Medio | |

## Parte B: Event Sourcing

### Idea

En vez de guardar el **estado actual** de una entidad, guardas **la secuencia de eventos** que le ocurrieron. El estado se obtiene reproduciéndolos.

```
Tradicional:   policy { id: 1, status: 'renewed', endDate: 2027-12-31 }

Event sourcing:
  1. PolicyIssued     { number: 'POL-1', start: 2025-01-01, end: 2025-12-31 }
  2. PolicyEndorsed   { change: 'suma asegurada', from: 500000, to: 750000 }
  3. PolicyRenewed    { newEnd: 2026-12-31 }
  4. PolicyRenewed    { newEnd: 2027-12-31 }
  -> estado actual = aplicar 1..4
```

Casi siempre va junto con [CQRS](../10-cqrs/README.md): el log de eventos es el modelo de escritura, y las **proyecciones** construyen tablas de lectura.

### Lo mejor

- Auditoría completa y gratuita: sabes qué pasó, cuándo y por qué, sin tabla de bitácora aparte.
- Consultas temporales: "¿cómo estaba esta póliza el 15 de marzo?".
- Puedes crear proyecciones nuevas sobre datos históricos (un reporte que nadie pidió hace dos años).
- Depurar reproduciendo los eventos exactos de producción.

### Lo peor

- Es de las arquitecturas más difíciles de operar bien.
- Los eventos son inmutables y para siempre: cambiar su estructura exige versionado y *upcasting*.
- Toda lectura depende de proyecciones, con consistencia eventual.
- Corregir un dato no es un `UPDATE`: es un evento compensatorio.
- Borrar datos personales (derecho al olvido) es complicado.
- Pocos devs tienen experiencia real con ella.

### Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir una proyección o reporte | Bajo a medio | Se reconstruye desde el historial |
| Añadir un evento | Medio | |
| Cambiar un evento existente | Muy alto | Versionado de por vida |
| Fix de datos incorrectos | Alto | Eventos compensatorios |
| Onboarding | Muy alto | |

## Cuándo sí y cuándo no

- **EDA nivel 1:** casi siempre que tengas efectos secundarios (notificar, auditar). Es barato.
- **EDA nivel 2:** cuando perder un evento es inaceptable (cobros, sincronización con ERP).
- **EDA nivel 3:** solo con varios servicios desplegados.
- **Event Sourcing:** solo en el agregado donde el historial **es** el negocio (movimientos contables, endosos con valor legal, trazabilidad regulatoria). Nunca en todo el sistema. Para auditoría normal, una tabla de bitácora resuelve con una fracción del costo.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** El estado de la póliza no se guarda: se reconstruye aplicando sus eventos. Prueba renovar una póliza y luego consulta su historial.

---

[← CQRS (Command Query Responsibility Segregation)](../10-cqrs/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Microservices →](../12-microservices/README.md)
