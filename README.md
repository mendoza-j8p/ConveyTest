# ConveyTest

Este es un proyecto de ejemplo en Go que incluye funciones para dividir números y sumar una lista de números. El proyecto utiliza la biblioteca **GoConvey** para realizar pruebas unitarias y mostrar los resultados en una interfaz web interactiva.

## Requisitos

- Go 1.16 o superior instalado en tu sistema.

## Instalación

1. Clona el repositorio de GitHub:

2. Instala las dependencias del proyecto, incluyendo GoConvey:

```bash
go get github.com/smartystreets/goconvey
```

## Uso

El proyecto incluye un archivo `main.go` con un programa de ejemplo que utiliza las funciones `divide()` y `sum`.

Para ejecutar el programa y ver los resultados en la consola, simplemente ejecuta el siguiente comando:

```bash
go run main.go
```

Esto imprimirá el cociente y el resto de la división de 17 entre 5, y también mostrará el resultado de sumar 1, 2, 3, 4 y 5.

## Pruebas

El proyecto utiliza la biblioteca GoConvey para realizar pruebas unitarias y mostrar los resultados en una interfaz web interactiva.

1. Inicia el servidor web de GoConvey ejecutando el siguiente comando:

```bash
goconvey
```

Esto iniciará un servidor web en `http://localhost:8080`, donde se mostrarán los resultados de las pruebas.

1. Abre tu navegador web y visita `http://localhost:8080`. Verás la interfaz web de GoConvey.

2. La interfaz web de GoConvey supervisará los archivos en el proyecto y ejecutará automáticamente las pruebas cada vez que se detecten cambios en los archivos. Los resultados de las pruebas se actualizarán en tiempo real en la interfaz web.

3. Puedes hacer clic en los enlaces para expandir y ver los detalles de las pruebas individuales. También puedes ver la cobertura del código y los resultados de las pruebas en general.

4. Si deseas ejecutar manualmente las pruebas en la línea de comandos en lugar de utilizar la interfaz web, puedes ejecutar el siguiente comando:

```bash
go test -v
```

Esto ejecutará las pruebas y mostrará los resultados en la consola.

## Contribución

Si deseas contribuir a este proyecto, ¡eres bienvenido! Puedes enviar pull requests con tus mejoras o correcciones de errores.
