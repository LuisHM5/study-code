# 02. Layered Architecture

**Complejidad: 1.5/5** · La más extendida: casi todos empezamos aquí.

## Idea

El código se divide en capas horizontales y las carpetas se nombran por **tipo técnico**: rutas, controladores, servicios, entidades. Cada capa solo llama a la de abajo.

```
Ruta -> Controlador -> Servicio -> Repositorio/ORM -> Base de datos
```

Otros nombres: Layered Architecture, N-tier, Controller-Service-Repository, package-by-layer.

## Lo mejor

- Todo el mundo la conoce. Onboarding casi nulo.
- Separa HTTP de lógica: el servicio se puede reusar desde un cron o un CLI.
- Predecible: sabes dónde va cada archivo.

## Lo peor

- **Una feature queda repartida en 4 o 5 carpetas.** Para tocar "licencias" abres `routes/`, `controllers/`, `services/`, `entities/`, `dtos/`.
- No hay fronteras: cualquier servicio importa cualquier otro y con el tiempo todo depende de todo.
- La capa de servicio se vuelve cajón de sastre y aparecen archivos de más de 1,000 líneas.
- La lógica de negocio depende del ORM, así que los tests unitarios necesitan simular el repositorio completo.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Bajo | 3 o 4 archivos en carpetas distintas |
| Modificar feature | Medio | Hay que saltar entre carpetas; sube con el tamaño |
| Fix de bug | Bajo a medio | Fácil de ubicar la capa, difícil medir el impacto |
| Cambiar DB o proveedor | Alto | El ORM se usa directo en los servicios |
| Borrar una feature | Medio | Hay que cazar archivos en cada carpeta |
| Onboarding | Muy bajo | |

## Cuándo sí y cuándo no

- **Sí:** hasta ~30 endpoints o ~15 entidades, equipo de 1 a 3, lógica mayormente CRUD.
- **No:** cuando ya tienes 40 archivos por carpeta. Ahí conviene agrupar por [feature](../03-feature-modules/README.md).

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Sigue una petición de arriba a abajo: handler/controller, service, repository. Cada capa solo conoce a la siguiente.

## Reglas que la mantienen sana

1. El controlador no tiene lógica: solo traduce HTTP a una llamada de servicio.
2. Un servicio de más de ~400 líneas se parte por caso de uso (`issue-policy.service.ts`, `renew-policy.service.ts`).
3. Las reglas puras (cálculos, validaciones) van en funciones sin dependencias, fuera del servicio. Se prueban sin mocks.
4. Un servicio no importa controladores ni objetos `req`/`res`.

---

[← Transaction Script](../01-transaction-script/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Modular por feature →](../03-feature-modules/README.md)
