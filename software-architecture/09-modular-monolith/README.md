# 09. Modular Monolith con DDD

**Complejidad: 4/5** · La opción moderna que más se recomienda antes de pensar en microservices.

## Idea

Un solo desplegable, pero dividido por dentro en **módulos con fronteras tan estrictas como las de un microservicio**. Cada módulo corresponde a un *bounded context* de DDD: una parte del negocio con su propio modelo y su propio vocabulario.

```
┌───────────────────── un solo proceso, un solo deploy ─────────────────────┐
│                                                                           │
│  ┌──────────┐      ┌──────────┐      ┌──────────┐      ┌──────────────┐   │
│  │ Pólizas  │      │Siniestros│      │  Pagos   │      │Notificaciones│   │
│  │          │      │          │      │          │      │              │   │
│  │ API pub. │<─────│          │      │          │      │              │   │
│  │ interno  │      │ interno  │      │ interno  │      │   interno    │   │
│  │ tablas   │      │ tablas   │      │ tablas   │      │   tablas     │   │
│  └────┬─────┘      └──────────┘      └────▲─────┘      └──────▲───────┘   │
│       └──────── evento: policy.issued ────┴───────────────────┘           │
└───────────────────────────────────────────────────────────────────────────┘
```

Las reglas que lo distinguen de [modular por feature](../03-feature-modules/README.md):

1. **Cada módulo es dueño de sus tablas.** Nadie más las lee ni las escribe. Sin joins entre módulos.
2. **Cada módulo expone una API pública pequeña** (una fachada) y todo lo demás es privado.
3. **La comunicación entre módulos es por esa API o por eventos**, nunca importando servicios internos o entidades.
4. **Las fronteras se verifican con herramientas**, no con buena voluntad.
5. Por dentro, cada módulo usa la arquitectura que le convenga: CRUD simple para catálogos, hexagonal para el núcleo.

## Lo mejor

- Casi todos los beneficios organizativos de microservices (equipos autónomos, fronteras claras) sin su costo operativo.
- Transacciones locales, un solo deploy, un solo log, depuración normal.
- Refactorizar una frontera equivocada es mover código, no migrar servicios.
- Si algún día un módulo necesita ser servicio aparte, ya está aislado.
- Cada módulo elige su nivel de complejidad interna.

## Lo peor

- Las fronteras dependen de disciplina y de herramientas; un atajo ("solo este join") las rompe.
- Sin joins entre módulos, algunas pantallas y reportes necesitan composición o modelos de lectura.
- Definir bien los bounded contexts requiere entender el negocio. Fronteras mal trazadas duelen mucho.
- Sigue escalando como una unidad y un fallo grave tumba todo.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature dentro de un módulo | Bajo a medio | Según la arquitectura interna del módulo |
| Añadir feature que cruza módulos | Medio | Hay que pasar por APIs públicas o eventos |
| Modificar el interior de un módulo | Bajo | Nada externo depende de él |
| Cambiar la API pública de un módulo | Medio a alto | Es un contrato |
| Fix de bug | Bajo a medio | El módulo acota la búsqueda |
| Onboarding | Medio | Un dev puede aprender un solo módulo |

## Cuándo sí y cuándo no

- **Sí:** sistemas de más de ~30 módulos o más de un equipo. Es el destino natural de un proyecto modular por feature que creció.
- **No:** proyectos chicos, donde las fachadas y eventos son peso muerto.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** El módulo de claims solo puede usar la API pública de policies, y notifications se entera por un evento sin que policies sepa que existe.

## Cómo hacer cumplir las fronteras

- **ESLint** con `no-restricted-imports` o `eslint-plugin-boundaries`: prohibir `modules/*/internal/**` desde otro módulo.
- **dependency-cruiser** en CI para detectar ciclos y dependencias prohibidas.
- **Nx** con `enforce-module-boundaries` si usas monorepo.
- Un **schema de Postgres por módulo** (`policies.policy`, `claims.claim`) hace visibles los joins indebidos.
- Tests de arquitectura que fallen si un módulo registra entidades ajenas.

## DDD estratégico en dos párrafos

Antes de escribir código, identifica los bounded contexts: zonas donde una misma palabra significa algo distinto. "Póliza" para emisión es un contrato con vigencia; para siniestros es una cobertura que consultar; para pagos es un calendario de primas. Tres modelos pequeños son mejores que un `PolicyEntity` con 80 columnas.

Clasifica los módulos: **core** (lo que diferencia al negocio; aquí inviertes en dominio rico y tests), **supporting** (necesario pero simple) y **generic** (auth, correo, archivos; compra o usa librerías). La complejidad arquitectónica se gasta en el core.

---

[← Clean Architecture](../08-clean/README.md) · [Índice](../README.md#-ruta-de-estudio) · [CQRS (Command Query Responsibility Segregation) →](../10-cqrs/README.md)
