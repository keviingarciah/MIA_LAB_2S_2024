# Clase 6

Explicación del _frontend_ y _backend_ para la aplicación web. Explicación de la utilizazión de React con ViteJS para el _frontend_ y Go con Fiber para el _backend_. También la explicación de la comunicación entre ambos por medio de una API REST.

## Frontend

Para este proyecto se utilizó React con ViteJS. ViteJS es un _bundler_ que permite trabajar con React de manera más rápida y eficiente. Para instalar ViteJS se hizo lo siguiente:

```bash
# En mi caso utilicé bun, pero se puede utilizar cualquier otro empaquetador
bun create vite

# Entrar a la carpeta del proyecto
cd frontend

# Instalar las dependencias
bun install

# Para correr el servidor de desarrollo
bun dev
```

Para más información sobre ViteJS pueden visitar su [página oficial](https://vitejs.dev/guide/).

### Tailwindcss

También en mi caso utilicé tailwindcss para los estilos. Para instalar tailwindcss se hizo lo siguiente:

```bash
# Instalar tailwindcss
bun install -D tailwindcss postcss autoprefixer
bunx tailwindcss init -p
```

Se va a crear un archivo `tailwind.config.js` en la raíz del proyecto. En este archivo se pueden configurar los estilos de tailwindcss. Se debe agregar el siguiente código:

```js
// Archivo tailwind.config.js
content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
```

Por último se debe importar tailwindcss en el archivo `index.css`:

```css
/* Archivo index.css */
@tailwind base;
@tailwind components;
@tailwind utilities;
```

Para más información sobre tailwindcss pueden visitar su [página oficial](https://tailwindcss.com/docs/guides/vite).

## Backend

Para este proyecto se utilizó Go con Fiber. Fiber es un _framework_ de Go que permite trabajar con APIs de manera más rápida y eficiente. Primero de debe inicializar el _módulo_ de Go en el proyecto:

```bash
go mod init backend
```

Luego se debe instalar el _framework_ de Go llamado Fiber:

```bash
go get -u github.com/gofiber/fiber/v2
```

Una vez instalado Fiber se puede empezar a trabajar con el _backend_. Para correr el servidor de desarrollo se debe correr el siguiente comando:

```bash
go run main.go
```

Para más información sobre Fiber pueden visitar su [página oficial](https://docs.gofiber.io/).

> **Nota:** Si ya tenían algo en excalidraw este link se los va a reemplazar.
> [Excalidraw Link](https://excalidraw.com/#json=WDkZ8Lvfh2LJHYVkzT6e-,ml6ZkYxu8wjdO576pHj4Sg)
