# 12. Microservices

**Complejidad: 5/5**

## Idea

El sistema se parte en **servicios pequeños que se despliegan por separado**. Cada uno es dueño de su base de datos, tiene su propio ciclo de entrega y se comunica con los demás por red (HTTP/gRPC síncrono, o eventos asíncronos).

```
                      ┌────────────────┐
  Cliente ──────────> │  API Gateway   │
                      └───────┬────────┘
            ┌─────────────────┼──────────────────┐
            v                 v                  v
     ┌────────────┐    ┌────────────┐     ┌────────────┐
     │  Pólizas   │    │ Siniestros │     │   Pagos    │
     │  NestJS    │    │    Go      │     │   Rust     │
     └─────┬──────┘    └─────┬──────┘     └─────┬──────┘
           v                 v                  v
        Postgres          Postgres            Postgres
           └───────── Kafka / RabbitMQ / NATS ──┘
```

Es, sobre todo, una solución a un **problema de organización**: muchos equipos que no pueden esperarse entre sí para desplegar. No es una solución a "mi código está desordenado".

## Lo mejor

- Despliegue y escalado independientes por servicio.
- Equipos autónomos, cada uno dueño de su servicio de principio a fin.
- Aislamiento de fallos: si se cae reportes, emisión sigue funcionando (si se diseñó para ello).
- Libertad tecnológica por servicio (NestJS aquí, Go allá, Rust donde importe el rendimiento).
- Fronteras imposibles de romper por accidente: no hay cómo hacer un join.

## Lo peor

- **Todo lo que era una llamada a función ahora es una llamada de red** que puede fallar, tardar o duplicarse. Necesitas timeouts, reintentos, circuit breakers, idempotencia.
- Sin transacciones entre servicios: aparecen sagas, compensaciones y consistencia eventual.
- Carga operativa enorme: orquestación (Kubernetes), service discovery, trazas distribuidas, logs centralizados, un pipeline por servicio, gestión de secretos.
- Depurar un flujo que cruza 5 servicios es de otro orden de dificultad.
- Las pruebas de integración y el entorno local se complican.
- Un cambio que cruza servicios requiere versionar contratos y coordinar despliegues.
- Si las fronteras están mal trazadas obtienes un **monolito distribuido**: todos los costos y ningún beneficio. Es el resultado más común cuando se adoptan demasiado pronto.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature dentro de un servicio | Bajo | El servicio es pequeño |
| Añadir feature que cruza servicios | Muy alto | Contratos, versiones, despliegues coordinados |
| Modificar un contrato entre servicios | Muy alto | Compatibilidad hacia atrás obligatoria |
| Fix de bug dentro de un servicio | Bajo | |
| Fix de bug entre servicios | Muy alto | Trazas distribuidas, reproducción difícil |
| Operación diaria | Muy alto | Requiere equipo de plataforma |
| Onboarding | Bajo por servicio, alto para el sistema | |

## Cuándo sí y cuándo no

**Sí**, cuando se cumplen varias:

- Más de 3 o 4 equipos estorbándose en el mismo desplegable.
- Partes con necesidades de escalado muy distintas, medidas, no supuestas.
- Requisitos de aislamiento (cumplimiento, disponibilidad).
- Ya existe capacidad de plataforma: CI/CD maduro, observabilidad, orquestación.
- Las fronteras del dominio ya están probadas, idealmente en un [modular monolith](../09-modular-monolith/README.md).

**No:**

- Un equipo de menos de ~10 personas.
- Producto nuevo cuyo dominio todavía se está descubriendo.
- Como remedio para un monolito desordenado. El desorden viaja contigo, ahora con latencia.

El camino recomendado: modular monolith primero, y extraer un servicio solo cuando un módulo concreto demuestre que lo necesita.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Lo que era una llamada a función ahora es una llamada de red con timeout. Apaga el servicio de notifications y observa qué pasa al emitir.

## Patrones que vas a necesitar

| Problema | Patrón |
|---|---|
| Punto de entrada único, auth, rate limit | API Gateway / BFF |
| Transacción que cruza servicios | Saga (orquestada o coreografiada) |
| Publicar evento y guardar de forma atómica | Outbox |
| Mensajes duplicados | Consumidores idempotentes |
| Servicio caído o lento | Timeout, retry con backoff, circuit breaker |
| Seguir una petición entre servicios | Trazas distribuidas (OpenTelemetry) |
| Evolución de contratos | Versionado, contract testing |
| Datos de otro servicio para leer | Réplica local por eventos, o composición en el gateway |

## Variantes modernas

- **Serverless / FaaS:** funciones por evento (Lambda, Cloud Functions). Operación mínima, pero arranque en frío, límites de ejecución y dependencia del proveedor.
- **Macroservices o "servicios del tamaño correcto":** pocos servicios grandes alineados a equipos, en vez de decenas pequeños. Es hacia donde ha vuelto buena parte de la industria.
- **Arquitectura basada en celdas:** copias completas del sistema aisladas por grupo de clientes, para limitar el radio de un fallo. Solo a escala muy grande.

---

[← Event-Driven y Event Sourcing](../11-event-driven-event-sourcing/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Frontend →](../frontend/README.md)
