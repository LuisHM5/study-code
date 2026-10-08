<div align="center">

# 📚 study-code

**Guías de estudio de ingeniería de software, con código que se ejecuta.**

Cada tema se explica de lo simple a lo complejo y se acompaña de proyectos reales<br>
en varios lenguajes para que compares, ejecutes y modifiques.

[Temas](#-temas) · [Cómo estudiar](#-cómo-estudiar-con-este-repositorio) · [Requisitos](#-requisitos) · [Convenciones](#-convenciones)

</div>

---

## 🗂 Temas

| Tema | Qué aprendes | Nivel | Lenguajes |
|---|---|---|---|
| 🏛 [**Software Architecture**](./software-architecture/README.md) | 12 arquitecturas de backend, de Transaction Script a microservicios, más arquitecturas de frontend. Cuándo conviene cada una y cuánto cuesta mantenerla. | Intermedio | NestJS · Go · Rust |

El repositorio crece por temas. Cada uno vive en su propia carpeta y se puede estudiar por separado.

## 🧭 Cómo estudiar con este repositorio

1. **Elige un tema** de la tabla y abre su `README.md`. Ahí está la ruta de estudio completa.
2. **Lee la teoría** de cada sección: es corta y va al punto.
3. **Ejecuta el proyecto** en el lenguaje que conozcas mejor.
4. **Compara** la misma solución en otro lenguaje. Ver qué cambia y qué no es la forma más rápida de separar la idea del detalle.
5. **Modifica algo**: añade un endpoint, cambia una regla. Ahí se nota de verdad cuánto cuesta cada enfoque.

## 🧰 Requisitos

Solo necesitas el lenguaje que vayas a ejecutar.

| Lenguaje | Versión | Instalación |
|---|---|---|
| Node.js + pnpm | Node 22 o superior, pnpm 10 | [nodejs.org](https://nodejs.org) · [pnpm.io](https://pnpm.io/installation) |
| Go | 1.24 o superior | [go.dev/dl](https://go.dev/dl/) |
| Rust | 1.92 o superior | [rustup.rs](https://rustup.rs) |

Ningún proyecto necesita base de datos ni servicios externos: guardan los datos en memoria para que arranquen con un solo comando.

## 📁 Estructura

```
study-code/
├── README.md                  ← estás aquí
└── software-architecture/     ← un tema por carpeta
    ├── README.md              ← ruta de estudio del tema
    ├── 01-transaction-script/
    │   ├── README.md          ← teoría
    │   ├── nestjs/            ← proyecto ejecutable
    │   ├── go/                ← proyecto ejecutable
    │   └── rust/              ← proyecto ejecutable
    └── ...
```

## ✍️ Convenciones

- **Idioma:** las explicaciones están en español. Los términos técnicos se dejan en inglés, que es como aparecen en la documentación, en las ofertas de trabajo y en las búsquedas. La primera vez que aparece uno se da su traducción si ayuda a entenderlo.
- **Un mismo problema por tema:** todos los proyectos de un tema resuelven el mismo caso, así la única variable es lo que se está estudiando.
- **Código mínimo:** los proyectos tienen lo justo para mostrar la idea. No son plantillas de producción.
- **Opiniones marcadas como tales:** donde hay una recomendación, se explica el porqué para que puedas no estar de acuerdo.

## 🤝 Contribuir

Las correcciones y mejoras son bienvenidas: abre un issue o un pull request. Si propones un tema nuevo, sigue la misma estructura (teoría corta más proyectos ejecutables).
