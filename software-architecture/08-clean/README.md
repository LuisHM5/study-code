# 08. Clean Architecture

**Complejidad: 4/5** · Robert C. Martin, 2012.

## Idea

Es la síntesis de hexagonal y onion, con nombres y reglas más estrictas. Cuatro círculos, y **la regla de dependencia**: el código de un círculo solo puede nombrar cosas de círculos más internos.

```
4. Frameworks y drivers      -> NestJS, TypeORM, Postgres, SMTP
3. Adaptadores de interfaz   -> controllers, presenters, gateways (repos)
2. Casos de uso              -> reglas de la aplicación
1. Entidades                 -> reglas de negocio de la empresa
```

Lo que añade sobre las anteriores:

- **Un caso de uso por clase**, con modelos de entrada y salida propios (`Input`/`Output`) que no son ni el DTO HTTP ni la entidad.
- **Input boundary / output boundary:** el caso de uso se expone por una interfaz y devuelve su resultado por otra (presenter).
- Los datos cruzan cada frontera en estructuras simples; una entidad de dominio nunca llega al controlador ni una entidad ORM al caso de uso.
- Asume [Screaming Architecture](../05-screaming/README.md) para las carpetas.

## Hexagonal, Onion y Clean en una tabla

| | Hexagonal | Onion | Clean |
|---|---|---|---|
| Regla base | Afuera depende de adentro | Igual | Igual |
| Estructura interna | Libre | Anillos de dominio | 4 círculos definidos |
| Foco | Puertos y adaptadores | Modelo de dominio | Casos de uso |
| Modelos por frontera | Opcional | Opcional | Obligatorio |
| Ceremonia | Media | Media | Alta |

Son la misma idea con distinto nivel de prescripción. Elegir entre ellas importa menos que respetar la regla de dependencia.

## Lo mejor

- Independencia real de framework, base de datos y UI. Sistemas que viven 10 años cambiando de tecnología.
- La lista de casos de uso es la documentación del sistema.
- Testeabilidad máxima: entidades y casos de uso sin nada externo.
- Fronteras explícitas que resisten equipos grandes y rotación de personal.

## Lo peor

- Es la que más código repetitivo genera: por cada operación, DTO HTTP + Input + Output + entidad + entidad ORM + mappers.
- Añadir un campo puede tocar 6 u 8 archivos.
- En NestJS vas contra la corriente: el framework quiere decoradores en todos lados y Clean quiere el dominio sin ellos.
- Aplicada a CRUD produce capas que solo copian datos. Es la causa más común de proyectos "clean" lentos y odiados.
- Curva de aprendizaje alta; mal entendida degenera en carpetas con nombres bonitos y dependencias cruzadas.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Alto | Muchos archivos por operación |
| Modificar regla de negocio | Bajo | Aislada y cubierta por tests |
| Fix de bug | Bajo a medio | Fácil de localizar, hay más saltos que seguir |
| Cambiar DB, framework o proveedor | Muy bajo | Su razón de ser |
| Añadir un campo | Alto | Atraviesa todos los modelos |
| Onboarding | Alto | |

## Cuándo sí y cuándo no

- **Sí:** núcleo de negocio con reglas complejas, vida esperada larga, varios equipos, requisitos regulatorios o cambio tecnológico probable.
- **No:** CRUD, MVPs, equipos de 1 a 3 con presión de entrega. Ahí cuesta más de lo que devuelve.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** El use case tiene su propio Input y Output, y no importa nada del framework. Cuenta cuántos modelos distintos atraviesa un mismo dato.

## Clean pragmática

La versión estricta rara vez compensa. Concesiones habituales que conservan casi todo el valor:

1. Permitir `@Injectable()` en los casos de uso. Es un decorador de metadatos; evitas las fábricas del módulo.
2. Usar el presenter solo cuando la respuesta difiere del `Output`.
3. Para lecturas, saltarse el dominio y consultar directo (ver [CQRS](../10-cqrs/README.md)).
4. Aplicarla solo en los módulos del núcleo de negocio.

---

[← Onion Architecture](../07-onion/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Modular monolith con DDD →](../09-modular-monolith/README.md)
