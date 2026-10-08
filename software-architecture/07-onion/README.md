# 07. Onion Architecture

**Complejidad: 3.5/5** · Jeffrey Palermo, 2008.

## Idea

Misma regla que hexagonal (las dependencias apuntan hacia adentro), pero organiza el interior en **anillos concéntricos**:

```
┌──────────────────────────────────────────────────┐
│ Infraestructura / UI / Tests                     │
│  ┌────────────────────────────────────────────┐  │
│  │ Servicios de aplicación (casos de uso)     │  │
│  │  ┌──────────────────────────────────────┐  │  │
│  │  │ Servicios de dominio                 │  │  │
│  │  │  ┌────────────────────────────────┐  │  │  │
│  │  │  │ Modelo de dominio              │  │  │  │
│  │  │  │ (entidades, value objects)     │  │  │  │
│  │  │  └────────────────────────────────┘  │  │  │
│  │  └──────────────────────────────────────┘  │  │
│  └────────────────────────────────────────────┘  │
└──────────────────────────────────────────────────┘
        un anillo solo conoce a los de más adentro
```

- **Modelo de dominio:** entidades y value objects con sus reglas.
- **Servicios de dominio:** reglas que involucran varias entidades, más las interfaces de repositorio.
- **Servicios de aplicación:** orquestan un caso de uso, manejan transacciones.
- **Infraestructura:** ORM, HTTP, correo. Implementa las interfaces de adentro.

## Diferencia con hexagonal

Hexagonal solo distingue adentro y afuera, y pone el foco en los puertos. Onion dice además **cómo estratificar el adentro** y pone el foco en el modelo de dominio. Onion es la que más naturalmente acompaña a DDD.

En la práctica, un proyecto onion bien hecho también es hexagonal.

## Lo mejor

- Empuja a un modelo de dominio rico: las reglas viven en las entidades, no en servicios anémicos.
- Deja claro dónde va cada regla: en la entidad, en un servicio de dominio o en la aplicación.
- Negocio probado sin infraestructura, igual que hexagonal.
- Calza muy bien con DDD táctico (agregados, value objects, eventos de dominio).

## Lo peor

- Más anillos, más decisiones: "¿esto es servicio de dominio o de aplicación?" genera discusiones.
- Mismo costo de mapeo dominio/ORM que hexagonal.
- Si el dominio es simple, los anillos internos quedan vacíos y solo estorban.
- La organización por anillos choca con organizar por feature si se aplica de forma global. Hay que aplicarla dentro de cada módulo.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Medio a alto | Toca varios anillos |
| Modificar regla de negocio | Bajo | Vive en la entidad o en un servicio de dominio |
| Fix de bug | Bajo a medio | |
| Cambiar DB o proveedor | Bajo | |
| Añadir un campo | Medio | Dominio + ORM + mapper + DTO |
| Onboarding | Medio a alto | Requiere entender DDD táctico |

## Cuándo sí y cuándo no

- **Sí:** dominios con invariantes fuertes (estados, cálculos, vigencias, flujos de aprobación) y equipo cómodo con DDD.
- **No:** CRUD, o equipos donde nadie ha modelado dominio antes.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Los value objects no pueden existir en estado inválido, y el domain service resuelve la regla que una entidad sola no puede verificar.

---

[← Hexagonal (Ports & Adapters)](../06-hexagonal/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Clean Architecture →](../08-clean/README.md)
