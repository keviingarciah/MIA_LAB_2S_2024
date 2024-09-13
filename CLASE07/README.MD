# Clase 7

Explicación de los comando _mkdir_ y _mfile_ para nuesto sistema de archivos. Explicación de la lógica tras la creación de archivos y directorios, a partir de los inodos y bloques de nuestro sistema de archivos.

## air

Antes que nada, les dejo la recomendación de usar air; air es una herramienta que nos permite recargar el servidor de manera automática cada vez que guardamos un archivo. Para instalarlo de manera global, ejecuten el siguiente comando:

```bash
# Deben tener go 1.22 o superior
go install github.com/air-verse/air@latest
```

Luego en su carpeta donde tengan el backend ejecuten el siguiente comando:

```bash
air init
```

Se creará un archivo `.air.toml` en el cual pueden configurar el puerto y otras opciones de air. Para levantar el servidor, ejecuten el siguiente comando:

```bash
air
```

Y listo, cada vez que guarden un archivo, el servidor se recargará automáticamente.

Para más información sobre air pueden visitar su [página oficial](https://github.com/air-verse/air).

> **Nota:** Si ya tenían algo en excalidraw este link se los va a reemplazar.
> [Excalidraw Link](https://excalidraw.com/#json=ZdEbD-3k1eUIJxRbcXGsB,3qPtsREmRJjSpj-IR6tszw)
