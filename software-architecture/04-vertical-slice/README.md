# 04. Vertical Slice Architecture

**Complejidad: 2.5/5** · Moderna, popularizada por Jimmy Bogard.

## Idea

En vez de cortar por capas horizontales, cortas por **caso de uso**. Cada operación (`emitir póliza`, `renovar póliza`, `cancelar póliza`) es una rebanada vertical con todo lo que necesita: su endpoint, su DTO, su validación, su lógica y su consulta.

```
policy/
  issue-policy/      -> controller + dto + handler
  renew-policy/      -> controller + dto + handler
  get-policy/        -> controller + query
```

La regla: **minimiza el acoplamiento entre slices y maximiza la cohesión dentro del slice.** Se acepta duplicar un poco de código entre slices antes que crear una abstracción compartida prematura.

## Lo mejor

- Resuelve el problema del servicio gigante: `policy.service.ts` de 2,000 líneas se vuelve 20 archivos de 100.
- Añadir una feature es añadir una carpeta. No tocas código existente, así que casi no hay regresiones.
- Cada slice decide su complejidad: un GET simple hace una consulta directa, un caso complejo usa dominio rico.
- Pull requests pequeños y sin conflictos entre devs.
- Fácil de borrar: una feature muerta es una carpeta.

## Lo peor

- Lógica duplicada entre slices si nadie extrae las reglas comunes a tiempo.
- Muchos archivos y carpetas.
- Requiere criterio para decidir qué se comparte y qué no. Sin ese criterio se vuelve transaction script repetido.
- Menos conocida, hay que explicarla al equipo.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Muy bajo | Carpeta nueva, cero cambios en lo existente |
| Modificar feature | Bajo | El impacto queda dentro del slice |
| Fix de bug | Bajo | El bug está en un archivo corto |
| Cambiar regla transversal | Medio a alto | Hay que revisar los slices que la copiaron |
| Cambiar DB o proveedor | Medio a alto | Depende de si los slices usan el ORM directo |
| Onboarding | Bajo a medio | |

## Cuándo sí y cuándo no

- **Sí:** módulos con muchos casos de uso distintos, equipos que entregan features constantemente. Es la mejora de mayor retorno para un proyecto [modular por feature](../03-feature-modules/README.md) con servicios gigantes.
- **No:** CRUD puro, donde un slice por operación es burocracia.

## Cómo adoptarla de forma gradual

No hace falta mover rutas ni carpetas de golpe. Conserva un solo controller por módulo y parte únicamente el service en un handler por caso de uso. Con eso obtienes la mayor parte del beneficio.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Cada caso de uso tiene su propio archivo o carpeta con su request, su lógica y su respuesta. No hay un service con veinte métodos.

---

[← Modular por feature](../03-feature-modules/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Screaming Architecture →](../05-screaming/README.md)
