# Arquitecturas de frontend

De la más simple a la más compleja, con ejemplos en React y Next.js. Las ideas aplican igual a Vue, Angular o Svelte.

## 1. Por tipo técnico

**Complejidad: 1/5**

```
app/            -> rutas
components/
hooks/
stores/
lib/
types/
```

- **Lo mejor:** obvio, cero decisiones.
- **Lo peor:** igual que en backend, una feature queda repartida; `components/` y `hooks/` llegan a cientos de archivos sin dueño.
- **Úsala:** hasta ~15 pantallas.

## 2. Colocation con App Router

**Complejidad: 1.5/5**

Lo que solo usa una ruta vive junto a esa ruta. Next ignora las carpetas con prefijo `_`.

```
app/
  policy/
    page.tsx
    [id]/page.tsx
    _components/policy-table.tsx
    _hooks/use-policy-filters.ts
    actions.ts
components/ui/          -> solo lo realmente compartido
lib/
```

- **Lo mejor:** borrar una pantalla borra todo lo suyo; casi gratis de adoptar.
- **Lo peor:** lo que comparten dos rutas no tiene lugar claro.
- **Úsala:** como primer paso desde la organización por tipo.

## 3. Por feature (la recomendada para la mayoría)

**Complejidad: 2/5**

Espejo del backend [modular por feature](../03-feature-modules/README.md). `app/` queda delgado: solo enruta y compone.

```
app/
  policy/page.tsx              -> importa de features/policy
  policy/[id]/page.tsx
features/
  policy/
    api/                       -> llamadas al backend y hooks de datos
      get-policies.ts
      issue-policy.ts
    components/
      policy-table.tsx
      issue-policy-form.tsx
    hooks/
    schemas.ts                 -> validación (zod)
    types.ts
    index.ts                   -> API pública de la feature
  claim/
  surety-bond/
components/ui/                 -> design system (shadcn)
lib/                           -> cliente http, auth, utilidades
```

Reglas:

1. Una feature no importa el interior de otra; solo su `index.ts`.
2. Dirección única: `app -> features -> components/ui, lib`. Nunca al revés.
3. Si dos features comparten algo, sube a `components/` o `lib/`, o se compone en `app/`.

- **Lo mejor:** escala a cientos de pantallas, calza 1:1 con los módulos del backend, permite trabajar en paralelo.
- **Lo peor:** requiere disciplina con los imports (se hace cumplir con ESLint).

## 4. Feature-Sliced Design (FSD)

**Complejidad: 3.5/5** · Metodología moderna y formal.

Capas estrictas, de arriba a abajo; cada una solo importa de las de abajo:

```
src/
  app/        -> inicialización, providers, estilos globales
  pages/      -> pantallas completas
  widgets/    -> bloques grandes autónomos (header, tabla de pólizas con filtros)
  features/   -> acciones del usuario con valor (emitir póliza, reportar siniestro)
  entities/   -> conceptos de negocio (policy, claim, user): modelo + UI básica
  shared/     -> ui kit, api client, utilidades, config
```

Cada slice se divide en segmentos: `ui/`, `model/`, `api/`, `lib/`.

- **Lo mejor:** reglas claras de dependencia, vocabulario común, herramientas de lint propias (Steiger). Es el "Clean Architecture" del frontend.
- **Lo peor:** curva de aprendizaje; la diferencia entre `entities`, `features` y `widgets` genera dudas; con Next App Router hay que adaptar, porque `app/` y `pages/` chocan con nombres del framework.
- **Úsala:** equipos frontend de 5 o más, productos grandes de vida larga.

## 5. Clean / Hexagonal en frontend

**Complejidad: 4/5**

Dominio, casos de uso y repositorios en TypeScript puro; React queda como adaptador.

- **Lo mejor:** lógica de negocio reutilizable entre web y móvil, y probada sin navegador.
- **Lo peor:** casi siempre sobreingeniería. En la mayoría de las apps el frontend muestra y captura datos, y las reglas viven en el backend.
- **Úsala:** solo con lógica compleja en cliente: apps offline-first, editores, cálculos pesados en el navegador.

## 6. Micro-frontends

**Complejidad: 5/5**

Varias apps frontend independientes compuestas en una (Module Federation, single-spa). Es el equivalente a [microservices](../12-microservices/README.md).

- **Lo mejor:** equipos que despliegan su parte por separado.
- **Lo peor:** peso duplicado, estado compartido difícil, inconsistencias visuales, operación compleja.
- **Úsala:** solo con varios equipos grandes y autónomos. Un monorepo con paquetes compartidos resuelve la mayoría de los casos con mucho menos costo.

## Decisiones transversales que importan más que la carpeta

**Separa el estado del servidor del estado del cliente.** Es la decisión de mayor impacto.

| Tipo de estado | Herramienta |
|---|---|
| Datos del servidor (listas, detalle) | Server Components, o TanStack Query / SWR en cliente |
| Estado de URL (filtros, página, orden) | `searchParams` (o `nuqs`) |
| Formularios | React Hook Form + Zod |
| Estado global de UI (tema, sidebar) | zustand o Context |

Un hook por recurso con `useEffect` + `useState` (`useAssets`, `useUsers`, ...) reimplementa caché, reintentos y estados de carga en cada uno. TanStack Query lo resuelve una vez.

**Server Components primero.** En Next 14+, la página consulta en el servidor y pasa datos a componentes de cliente solo donde hay interacción. Menos JavaScript y menos estado.

**BFF (Backend for Frontend).** Los route handlers o server actions de Next como capa fina frente a NestJS: guardan el token en el servidor, agregan llamadas y adaptan respuestas a la pantalla.

**Contrato tipado con el backend.** Genera los tipos del cliente desde el Swagger/OpenAPI del backend (`openapi-typescript`, `orval`) o comparte esquemas Zod en un paquete del monorepo. Elimina toda una clase de bugs.

**Atomic Design** (átomos, moléculas, organismos) es una forma de organizar el **design system**, no la aplicación. Úsalo, si quieres, dentro de `components/ui/`.

## Combinaciones recomendadas front + back

| Tamaño | Backend | Frontend |
|---|---|---|
| Muy simple | Transaction script o 2 capas | Por tipo o colocation |
| Mediano | Modular por feature | Por feature |
| Grande | Modular monolith + vertical slices + hexagonal en el core | Por feature estricto o FSD |
| Muy grande, varios equipos | Microservices | FSD o micro-frontends, con BFF |

Mantener los mismos nombres de feature en ambos lados (`policy` aquí, `policy` allá) vale más que cualquier metodología.
