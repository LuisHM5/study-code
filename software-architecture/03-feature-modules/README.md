# 03. Feature Modules

**Complejidad: 2/5** · El default de NestJS y el mejor punto de partida para la mayoría.

## Idea

Las mismas capas de [02](../02-layered/README.md), pero agrupadas primero por **feature de negocio** y después por tipo técnico. Cada feature es un módulo que declara qué expone y qué importa.

```
policy/   -> controllers, services, entities, dto
claim/    -> controllers, services, entities, dto
```

Otros nombres: package-by-feature, feature modules, modular monolith básico.

## Lo mejor

- Lo que cambia junto vive junto: para tocar "pólizas" abres una carpeta.
- Borrar o mover una feature es mover una carpeta.
- NestJS lo soporta de fábrica (`@Module`, `imports`, `exports`), sin librerías extra.
- Varios devs trabajan en paralelo con pocos conflictos.
- Es la base natural para evolucionar a vertical slice, hexagonal o modular monolith.

## Lo peor

- Las fronteras son débiles: nada impide que un módulo importe entidades de otro y haga joins a sus tablas.
- Dependencias circulares entre módulos (`forwardRef`) en cuanto el dominio se enreda.
- La capa `services/` sigue creciendo sin tope dentro de cada módulo.
- La lógica sigue pegada al ORM.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Bajo | `nest g resource` y listo |
| Modificar feature | Bajo a medio | Todo en una carpeta; sube si el servicio es gigante |
| Fix de bug | Bajo | |
| Cambiar DB o proveedor | Alto | El ORM sigue directo en servicios |
| Borrar una feature | Bajo | Si las fronteras se respetaron |
| Onboarding | Bajo | Es lo que enseña la documentación de NestJS |

## Cuándo sí y cuándo no

- **Sí:** es el punto de partida correcto para casi cualquier API de negocio en NestJS, de 10 a 100+ módulos.
- **No alcanza:** cuando hay reglas complejas o muchas integraciones externas. Ahí se refuerzan los módulos críticos con [vertical slice](../04-vertical-slice/README.md) o [hexagonal](../06-hexagonal/README.md), sin cambiar el resto.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Lo que el módulo de pólizas exporta y lo que se guarda para sí. Todo lo de una feature vive en una carpeta.

## Reglas que la mantienen sana

1. **Un módulo solo registra sus propias entidades** en `TypeOrmModule.forFeature`. Si necesita datos de otro, importa ese módulo y usa su servicio exportado.
2. `exports` es la API pública. Lo que no se exporta es privado.
3. Si aparece `forwardRef`, hay un ciclo: extrae lo común a un tercer módulo o usa un evento.
4. Parte los servicios por caso de uso antes de que pasen de ~400 líneas.
5. Los catálogos CRUD no necesitan más que esto. No los compliques.

---

[← Capas por tipo técnico (MVC / N-capas)](../02-layered/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Vertical Slice Architecture →](../04-vertical-slice/README.md)
