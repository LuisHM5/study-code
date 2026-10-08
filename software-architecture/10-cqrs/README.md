# 10. CQRS

**Complejidad: 4/5** (2.5/5 en su versión ligera) · Greg Young.

## Idea

Separar el modelo que **cambia** datos del modelo que los **lee**.

- **Command:** expresa una intención (`IssuePolicy`, `CancelPolicy`). Pasa por las reglas de negocio. No devuelve datos, o solo un ID.
- **Query:** pregunta algo (`GetPolicyDetail`, `ListExpiringPolicies`). No cambia nada. Devuelve un DTO hecho a la medida de la pantalla.

```
                ┌─> Command -> Handler -> Dominio -> Repositorio ─> DB escritura
Controller ─────┤                                                       │ (misma DB o sincronizada)
                └─> Query   -> Handler -> SQL directo / vista   ─> DB lectura
```

Tiene niveles, y conviene no saltar al último:

| Nivel | Qué separas | Complejidad |
|---|---|---|
| 1. CQRS ligero | Solo el código: handlers de comando y de consulta distintos, misma base de datos | Baja |
| 2. Modelos de lectura | Vistas, vistas materializadas o tablas desnormalizadas para consultar | Media |
| 3. Almacenes separados | Otra base o motor para lectura (réplica, Elasticsearch), sincronizado por eventos | Alta |

## Lo mejor

- Las consultas dejan de pelear con el modelo de dominio: un listado con 6 joins es un SQL directo, sin hidratar agregados.
- El lado de escritura queda pequeño y enfocado en reglas.
- Lectura y escritura se optimizan y escalan por separado.
- Combina muy bien con Clean/Hexagonal: resuelve su mayor queja ("tanto mapeo para un GET").
- Los comandos con nombre de negocio documentan lo que el sistema hace.

## Lo peor

- Más clases: comando, handler, query, handler, DTO de lectura.
- En el nivel 3 aparece **consistencia eventual**: el usuario guarda y por unos instantes no ve su cambio. La UI tiene que contemplarlo.
- El bus de comandos añade indirección: seguir el flujo en el editor cuesta más ("ir a la definición" llega al bus, no al handler).
- Mantener sincronizados los modelos de lectura es trabajo y una fuente de bugs.
- En CRUD simple es puro costo.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir una consulta o reporte | Bajo | Un handler con su SQL, sin tocar el dominio |
| Añadir un comando | Medio | Comando + handler + reglas |
| Modificar regla de negocio | Bajo | Solo el lado de escritura |
| Fix de bug | Medio | Bajo en nivel 1; alto en nivel 3 por la sincronización |
| Añadir un campo | Medio a alto | Modelo de escritura y cada modelo de lectura que lo muestre |
| Onboarding | Medio | |

## Cuándo sí y cuándo no

- **Sí (nivel 1):** casi cualquier módulo con dominio rico. Es barato y ordena.
- **Sí (nivel 2 y 3):** dashboards, analítica, búsquedas, lecturas mucho más frecuentes que las escrituras.
- **No:** CRUD donde lo que guardas es lo que muestras.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** El command pasa por las reglas del dominio; la query lee de un read model hecho a la medida, sin tocar el dominio.

## CQRS sin bus

Si el bus te parece demasiada indirección, obtienes el mismo beneficio inyectando los handlers directamente: una carpeta `commands/` y otra `queries/` con clases `@Injectable()` normales. Es [Vertical Slice](../04-vertical-slice/README.md) con la separación lectura/escritura explícita, y suele ser el mejor punto de entrada.

---

[← Modular monolith con DDD](../09-modular-monolith/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Event-Driven y Event Sourcing →](../11-event-driven-event-sourcing/README.md)
