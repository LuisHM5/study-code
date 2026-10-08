# 05. Screaming Architecture

**Complejidad: 2.5/5** · Robert C. Martin, 2011.

## Idea

No es un sistema de capas, es un **principio sobre cómo nombrar y organizar las carpetas**: al abrir el proyecto, la estructura debe "gritar" de qué trata el negocio, no qué framework usa.

```
No grita nada:            Grita "sistema de seguros":
src/                      src/
  controllers/              policies/
  services/                 claims/
  entities/                 endorsements/
  dtos/                     surety-bonds/
                            payments/
```

Un paso más: dentro de cada carpeta, los archivos se llaman por la **acción de negocio** (`issue-policy`, `renew-policy`, `report-claim`), no por el patrón (`policy.service`, `policy.helper`).

Se combina con otras: screaming + vertical slice, screaming + hexagonal, screaming + clean. De hecho Clean Architecture la da por supuesta.

## Lo mejor

- Onboarding por negocio: un dev nuevo entiende qué hace el sistema leyendo el árbol.
- Negocio y desarrollo hablan el mismo idioma (lenguaje ubicuo).
- El framework queda como detalle, lo que facilita reemplazarlo.
- Casi gratis de adoptar si ya trabajas por feature.

## Lo peor

- Por sí sola no dice cómo organizar el interior de cada carpeta ni hacia dónde van las dependencias.
- Exige nombrar bien, y nombrar es difícil: requiere entender el dominio.
- El código transversal (auth, logging, config) no encaja y termina en `shared/` o `common/`, que puede volverse basurero.

## Mantenimiento

| Tarea | Costo | Nota |
|---|---|---|
| Añadir feature | Bajo | La carpeta donde va es obvia |
| Modificar feature | Bajo | El nombre del archivo te lleva al lugar |
| Fix de bug | Bajo | Del reporte de negocio al archivo sin escalas |
| Cambiar DB o proveedor | Depende | Lo define la arquitectura con la que la combines |
| Onboarding | Muy bajo | Es su mayor beneficio |

## Cuándo sí y cuándo no

- **Sí:** siempre que el proyecto tenga más de un puñado de features. Es barata.
- **Cuidado:** en librerías técnicas o herramientas sin dominio de negocio, no aporta.

## Código

El mismo caso de uso, emitir una póliza, resuelto con esta arquitectura en tres lenguajes. Cada proyecto es independiente y trae su propio README con el árbol de carpetas y cómo ejecutarlo.

| Lenguaje | Proyecto |
|---|---|
| NestJS (TypeScript) | [`nestjs/`](./nestjs/README.md) |
| Go | [`go/`](./go/README.md) |
| Rust | [`rust/`](./rust/README.md) |

**En qué fijarte:** Lee solo el árbol de carpetas y los nombres de archivo antes de abrir nada: deberían contarte el negocio.

## Prueba rápida

Enseña solo el árbol de `src/` a alguien que no conoce el proyecto. Si responde "es una API en NestJS", no grita. Si responde "es un sistema para administrar pólizas, siniestros y fianzas", sí.

---

[← Vertical Slice Architecture](../04-vertical-slice/README.md) · [Índice](../README.md#-ruta-de-estudio) · [Hexagonal (Ports & Adapters) →](../06-hexagonal/README.md)
