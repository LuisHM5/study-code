# 06. Hexagonal (Ports & Adapters)

**Complejidad: 3.5/5** · Alistair Cockburn, 2005.

## Idea

La aplicación es un núcleo que no sabe nada del mundo exterior. Todo lo que entra o sale pasa por un **puerto** (una interfaz definida por el núcleo) y lo implementa un **adaptador** (código que habla con HTTP, Postgres, un correo, otra API).

```
   Adaptadores de entrada            NÚCLEO               Adaptadores de salida
   (driving / primarios)                                  (driven / secundarios)

   Controller HTTP  ──┐                              ┌──>  TypeORM / Postgres
   Cron job         ──┼──> puerto de ──> lógica ──> puerto ──>  SMTP / SES
   Consumer de cola ──┘    entrada      de negocio  de salida └──>  API de un ERP
```

Solo hay dos zonas: **adentro** (negocio) y **afuera** (tecnología). La regla es que afuera depende de adentro, nunca al revés. El núcleo define la interfaz `PolicyRepository`; la infraestructura la implementa.

## Lo mejor

- El negocio se prueba sin base de datos ni framework: en los tests usas adaptadores en memoria.
- Cambiar de proveedor (correo, ORM, API externa) es escribir otro adaptador. El núcleo no se toca.
- Un mismo caso de uso se expone por HTTP, cron, cola o CLI sin duplicar lógica.
- Es la más simple de la familia hexagonal/onion/clean: solo exige la frontera adentro/afuera, no dicta capas internas.

## Lo peor

- Más archivos: puerto + adaptador + mapper por cada dependencia externa.
- Mapeo entre el modelo de dominio y la entidad del ORM. Es el costo más molesto en el día a día.
- En CRUD simple es sobreingeniería: mucha indirección para pasar datos.
- Si el equipo no entiende la regla de dependencias, los puertos terminan siendo una copia 1:1 del ORM y no aíslan nada.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Medio | Caso de uso + puerto + adaptador + mapper |
| Modificar regla de negocio | Bajo | Está aislada y cubierta por tests rápidos |
| Fix de bug | Bajo a medio | Fácil saber si es de negocio o de adaptador |
| Cambiar DB o proveedor | Bajo | Para esto se diseñó |
| Añadir un campo a una entidad | Medio | Dominio + entidad ORM + mapper + DTO |
| Onboarding | Medio | |

## Cuándo sí y cuándo no

- **Sí:** módulos con reglas reales y con integraciones externas (ERP, proveedores de correo, OCR, pasarelas, Keycloak). Aplícala **por módulo**, no en todo el proyecto.
- **No:** catálogos, CRUD, prototipos.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Los ports (interfaces) los define el núcleo y los adapters los implementan. Busca dónde se conecta cada port con su adapter: es el único lugar que conoce a ambos.

## Versión pragmática

Si el mapeo dominio/ORM te frena, aplica hexagonal **solo a las salidas externas** (correo, APIs de terceros, almacenamiento, OCR) y deja que los casos de uso usen el repositorio de TypeORM directo. Te llevas buena parte del beneficio (tests sin red, proveedores intercambiables) por una fracción del costo.

---

[← Screaming Architecture](../05-screaming/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Onion Architecture →](../07-onion/README.md)
