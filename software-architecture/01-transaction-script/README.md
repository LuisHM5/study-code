# 01. Transaction Script

**Complejidad: 1/5**

## Idea

Cada operación es un procedimiento de arriba a abajo: recibe la petición, valida, consulta la base de datos y responde. No hay capas, o hay una sola. El handler lo hace todo.

Es la arquitectura que uso para proyectos muy simples: un webhook, un cron, un CRUD interno de 5 tablas, un prototipo, una lambda.

## Lo mejor

- Lo más rápido de escribir y de leer: todo lo de un endpoint está en un solo lugar.
- Cero ceremonia. No hay interfaces, mappers ni módulos que configurar.
- Fácil de borrar y rehacer.

## Lo peor

- La lógica se duplica en cuanto dos endpoints necesitan la misma regla.
- Solo se puede probar con base de datos real o con tests e2e.
- No escala en personas: dos devs tocando el mismo handler se pisan.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Muy bajo | Un handler nuevo |
| Modificar feature | Bajo al inicio, alto después | Hay que buscar las copias de la regla |
| Fix de bug | Bajo | Todo está a la vista |
| Cambiar DB o proveedor | Alto | El SQL/ORM está en todos lados |
| Onboarding | Inmediato | |

## Cuándo sí y cuándo no

- **Sí:** menos de ~10 endpoints, un dev, vida corta, o lógica que es solo "guardar y leer".
- **No:** en cuanto aparezcan reglas compartidas o más de un dev. Ahí pasa a [capas](../02-layered/README.md).

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Todo el caso de uso cabe en un solo handler: validación, regla, guardado y respuesta juntos.

---

[← Guía del tema](../README.md) · [Índice](../README.md#-ruta-de-estudio) · [Capas por tipo técnico (MVC / N-capas) →](../02-layered/README.md)
