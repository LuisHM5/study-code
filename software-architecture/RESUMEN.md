# Resumen: las 12 arquitecturas en una página

Hoja de repaso. Para la guía completa empieza por el [README](./README.md).

## De la más simple a la más compleja

Costos de mantenimiento: B = bajo, M = medio, A = alto.

| # | Arquitectura | Compl. | Añadir feature | Modificar | Fix | Lo mejor | Lo peor |
|---|---|---|---|---|---|---|---|
| 01 | Transaction Script | 1 | B | M | B | Rapidez, cero ceremonia | Duplicación, no se puede probar aislado |
| 02 | Capas por tipo (MVC) | 1.5 | B | M | B | Todos la conocen | Feature repartida en 5 carpetas, servicios gigantes |
| 03 | Modular por feature | 2 | B | B-M | B | Lo que cambia junto vive junto | Fronteras débiles, lógica pegada al ORM |
| 04 | Vertical Slice | 2.5 | B | B | B | Un caso de uso por archivo, sin regresiones | Duplicación si falta criterio |
| 05 | Screaming | 2.5 | B | B | B | El árbol explica el negocio | No define dependencias por sí sola |
| 06 | Hexagonal | 3.5 | M | B | B-M | Negocio probado sin infraestructura, proveedores intercambiables | Puertos, adaptadores y mapeo |
| 07 | Onion | 3.5 | M-A | B | B-M | Dominio rico, calza con DDD | Más anillos y más discusiones |
| 08 | Clean | 4 | A | B | B-M | Independencia total, vida larga | La que más código repetitivo genera |
| 09 | Modular monolith + DDD | 4 | B-M | B | B-M | Beneficios de microservices sin su operación | Exige disciplina y herramientas |
| 10 | CQRS | 2.5-4 | M | B | M | Lecturas simples, escrituras enfocadas | Más clases; consistencia eventual en el nivel alto |
| 11 | Event-Driven / Event Sourcing | 3-5 | B-M | A (contratos) | M-A | Desacople, auditoría total | Flujo difícil de seguir y de depurar |
| 12 | Microservices | 5 | B dentro, A entre servicios | A | A | Equipos y despliegues independientes | Sistema distribuido, operación muy cara |

## Cómo se relacionan

No son 12 opciones excluyentes, son 4 ejes que se combinan:

1. **Carpetas:** por tipo (02) o por feature (03, 04, 05).
2. **Dependencias:** hexagonal, onion y clean (06, 07, 08) son una misma familia: el dominio al centro, la infraestructura depende de él. Cambia cuánto prescriben. Hexagonal es la más ligera, Clean la más estricta.
3. **Partición:** monolito, modular monolith (09), microservices (12).
4. **Flujo de datos:** CQRS (10) y eventos (11).

## Qué uso según el tamaño

| Proyecto | Arquitectura |
|---|---|
| **Muy simple:** script, webhook, cron, CRUD de pocas tablas, prototipo | Transaction Script (01) o dos capas controller + service. Conservando validación de entrada y manejo global de errores. |
| **Chico:** hasta ~30 endpoints, 1 a 3 devs | Capas (02), o directamente modular por feature (03) si es NestJS, porque cuesta lo mismo. |
| **Mediano:** API de negocio típica | Modular por feature (03) + un caso de uso por archivo (04) en los módulos grandes. |
| **Grande:** reglas complejas, integraciones, vida larga | Modular monolith (09). Dentro: CRUD simple en catálogos, hexagonal (06) en el núcleo, CQRS ligero (10) para reportes, eventos en proceso (11) para efectos secundarios. |
| **Muy grande:** varios equipos, escalado desigual comprobado | Extraer microservices (12) desde el modular monolith, solo los módulos que lo necesiten. |

## Mis recomendaciones

**La que más recomiendo para enterprise:** modular monolith con complejidad interna variable por módulo. Da la mejor relación entre orden y costo, y deja abierta la puerta a microservices.

**La de mejor retorno inmediato:** Vertical Slice. Ataca justo el problema del servicio gigante y se adopta de forma gradual.

**Las que recomiendo con cuidado:**

- **Clean estricta en todo el proyecto.** En CRUD produce capas que solo copian datos. Úsala en el núcleo, o prefiere hexagonal.
- **Event Sourcing.** Solo donde el historial es el negocio. Para auditoría alcanza una bitácora.
- **Microservices.** Resuelven un problema de equipos, no de código. Sin plataforma y fronteras probadas, terminan en monolito distribuido.

**Regla general:** la complejidad de la arquitectura debe seguir a la complejidad del dominio, módulo por módulo. Un catálogo no merece el mismo tratamiento que el motor de workflows.

## Cómo mejorar un proyecto que ya existe

No hace falta reescribir. En un proyecto por capas o modular que ya duele, estos pasos van de mayor a menor retorno:

1. Partir los servicios gigantes en un handler por caso de uso.
2. Sacar las reglas de negocio a funciones o clases puras, sin dependencias.
3. Usar eventos en proceso para notificaciones, auditoría y logs.
4. Poner puertos y adaptadores en las integraciones externas (correo, ERP, almacenamiento, proveedor de identidad).
5. Cerrar fronteras: cada módulo registra solo sus entidades y expone una fachada.
6. Dejar los catálogos CRUD como están.

## Frontend

| Tamaño | Arquitectura |
|---|---|
| Simple | Por tipo técnico, o colocation con App Router |
| Mediano y grande | **Por feature** (`features/policy/{api,components,hooks}`), con `app/` solo para enrutar |
| Equipos frontend grandes | Feature-Sliced Design |
| Varios equipos autónomos | Micro-frontends, con mucha cautela |

Más importante que la carpeta: separar estado de servidor (Server Components o TanStack Query) de estado de cliente (zustand), formularios con React Hook Form + Zod, tipos generados desde el OpenAPI de NestJS, y los mismos nombres de feature en front y back.

Detalle en [`frontend/`](./frontend/README.md).

## Combinación recomendada front + back

| Tamaño | Backend | Frontend |
|---|---|---|
| Muy simple | Transaction script o 2 capas | Por tipo o colocation |
| Mediano | Modular por feature | Por feature |
| Grande | Modular monolith + vertical slices + hexagonal en el core | Por feature estricto o FSD |
| Muy grande | Microservices | FSD o micro-frontends, con BFF |

## NestJS, Go y Rust

Cada arquitectura trae un proyecto ejecutable en `nestjs/`, `go/` y `rust/` con el mismo caso de uso (emitir una póliza). Diferencias que vale la pena saber:

- **NestJS** da módulos e inyección de dependencias de fábrica, pero las fronteras entre módulos hay que vigilarlas con linters.
- **Go** hace cumplir fronteras con paquetes e `internal/`, y sus interfaces implícitas vuelven muy barata la arquitectura hexagonal. La comunidad prefiere estructuras planas y desconfía de la Clean estricta.
- **Rust** hace cumplir las capas con crates del workspace (el `Cargo.toml` es la regla de dependencias) y su sistema de tipos es el mejor de los tres para modelar dominio: los estados inválidos no compilan.
