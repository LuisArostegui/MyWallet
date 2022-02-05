# Contenedor de pruebas

Para tener una forma de hacer que la aplicación sea portable y esté lista para integrarse con CI/CD, debemos de elegir una imagen base que lo acompañe. Los principios básicos que se debe de seguir esta imagen base son:

* Debe de ser estable, esto implica que siempre debe de funcionar siempre de igual manera, es decir, que dadas las mismas entradas y condiciones se producirán invariablemente las mismas salidas y condiciones. Esto evitará errores y problemas que dependan del entorno. Por tanto siempre tendremos que utilizar bibliotecas compatibles.
* Debe de ser una imagen ligera, siempre que se pueda, es decir, tener las funcionalidades necesarias de Go para cumplir con la correcta construcción y ejecución de nuestro proyecto. Básicamente, esto sirve para acelerar la construcción, la implementación y también reducir costos con el almacenamiento y la salida de la red si está utilizando algún proveedor de la nube.
* Debe de recibir actualizaciones frecuentes, de esta manera se evitarán problemas de seguridad y rendimiento.

* Debe ofrecer un buen rendimiento, esto abarca desde el tamaño de la imagen hasta la duración de la ejecución de los tests pasando por el tiempo construcción del contenedor.


## Imagen de Golang

Teniendo en cuenta los requisitos nombrados podremos lograr un tamaño mínimo de imagen de Docker utilizando imágenes base que se centran en el minimalismo, como Alpine Linux. Dentro de Docker Hub nos vamos a centrar en la imagen oficial de Golang suministrada por Dockerhub, ya que hoy día es la que más actualizaciones recibe y con mayor frecuencia. Tenemos otras opciones de *Verified Publisher* que son entidades comerciales que publican imágenes muy confiables y están mantenidas por ellos, como Circle CI o portainer, en el caso de Circle CI las actualizaciones se reciben cada 3/4 semanas y en el caso de portainer la última actualización se recibió hace 4 años. Por tanto, teniendo en cuenta los requisitos nombrados vamos a centrarnos en las imágenes oficiales de Dockerhub.

Las variantes que nos encontramos son:

* `golang:<version>`, es la imagen por defecto. Si no estamos seguros de cuáles son nuestras necesidades, probablemente esta es la mejor opción. Además puede incluir etiquetas como pueden ser *bullseye*, *buster* o *stretch*. Estas etiquetas son los nombres de código de la suite para las versiones de Debian e indican en que versión se basa la imagen. 
* `golang:<version>-alpine`, esta imagen se basa en el proyecto Alpine Linux. Las imágenes Alpine Linux son mucho más livianas que la mayoría de imágenes base de distribución (~5 MB). Esta variante es experimental y no es oficialmente compatible con el [proyecto Go](https://github.com/golang/go/issues/19938). La principal advertencia a tener en cuenta es que utiliza **musl libc** en lugar de **glibc**, puede llegar a provocar un comportamiento inesperador en nuestra aplicación. En [este artículo](https://news.ycombinator.com/item?id=10782897) se conversa acerca de los problemas que puede traer este tipo de imágenes. [En esta página](https://wiki.musl-libc.org/functional-differences-from-glibc.html) se comentan las diferencias funcionales entre `glibc` y `musl libc`. Las principales diferencias que hay entre ambas librerías son cuestiones que no van a afectar directamente a nuestro proyecto, por ejemplo, si hiciesemos uso de un gestor de paquetes con una imagen Alpine tendremos que usar `apk` y con imagenes que usen la librería glibc se usaría `apt`, pero esto no nos concierne ya que no tenemos que instalar paquetes para que nuestro proyecto funcione correctamente. [Aquí](https://honnef.co/posts/2015/06/statically_compiled_go_programs__always__even_with_cgo__using_musl/) se comenta que si se usa el paquete cgo para usar funciones de C en Go necesitará este enlazarse con una librería libc, en este caso si podriamos tener problemas al elegir una imagen y otra, pero como digo para nuestro proyecto no existen tales dependecias, solamente existe la de Task y las dependecias reflejadas en go.mod.
* `golang:<version>-windowsservercore`, esta imagen se basa en Windows Server Core.


### Versiones de Go
Dentro de las posibles imágenes tenemos que saber elegir la versión de Go para ejecutar nuestro proyecto, las distintas versiones las podemos encontrar [aquí](https://go.dev/doc/devel/release). Tenemos que usar una versión que permita obtener los resultados esperados en nuestra aplicación, que tenga soporte y actualizaciones frecuentemente. Tanto [aquí](https://endoflife.date/go) como la [página oficial](https://go.dev/doc/devel/release), podemos ver que las versiones 1.16 y 1.17 son actualmente tienen soporte, por tanto las versiones 1.15 y anteriores quedan descartadas, según los requisitos nombrados. Vamos a tener en cuenta, por ahora, ambas versiones y realizaremos pruebas con ambas para obtener una conclusión final. Cabe mencionar que la versión 1.17 recibe actualizaciones más a menudo que la versión 1.16.


Como candidatos a nuestro proyecto:

* [golang:1.17.6-bullseye](https://github.com/docker-library/golang/blob/6b93987c3ec7bb3082dd54a46e9b6b8de95b0eb1/1.17/bullseye/Dockerfile) Debian 11, se selecciona como candidato porque actualmente es la versión estable de Debian, lo podemos comprobar tanto [aquí](https://wiki.debian.org/Status/Stable) como [aquí](https://wiki.debian.org/DebianBullseye). Dado que uno de nuestros requisitos es elegir una versión estable, esta opción se ajusta a nuestros criterios.
* [golang:1.17.6-alpine](https://github.com/docker-library/golang/blob/6b93987c3ec7bb3082dd54a46e9b6b8de95b0eb1/1.17/alpine3.15/Dockerfile) Alpine 3.15, se elige está porque es rápida y ligera, una de las más populares imágenes base para contenedores Docker.
* [golang:1.16-bullseye](https://github.com/docker-library/golang/blob/6b93987c3ec7bb3082dd54a46e9b6b8de95b0eb1/1.16/bullseye/Dockerfile).
* [golang:1.16-alpine](https://github.com/docker-library/golang/blob/6b93987c3ec7bb3082dd54a46e9b6b8de95b0eb1/1.16/alpine3.15/Dockerfile).

Vamos a analizar las opciones.

## Análisis de construcción

### Tiempos de construcción

En esta sección se generan los contenedores usando la herramienta `time`. Se ejecuta la orden `time docker build --no-cache . -f ./Dockerfile -t go_<bullseye|alpine>_1.1X`. Los resultados son los siguientes:

* Alpine, versión de Go 1.17

```shell
Successfully built 8c6c1dda8a53
Successfully tagged go_alpine_1.17:latest

real	0m5,843s
user	0m0,063s
sys	0m0,057s
```

* Alpine, versión de Go 1.16

```shell
Successfully built 7f201f405681
Successfully tagged go_alpine_1.16:latest

real	0m5,742s
user	0m0,044s
sys	0m0,054s
```

* Bullseye, versión de Go 1.17

```shell
Successfully built d2c15c730270
Successfully tagged go_bullseye_1.17:latest

real	0m5,761s
user	0m0,068s
sys	0m0,050s
```
* Bullseye, versión de Go 1.17

```shell
Successfully built dafa5fb683e3
Successfully tagged go_bullseye_1.16:latest

real	0m5,757s
user	0m0,068s
sys	0m0,046s
```

Los resultados son prácticamente iguales, variando a nivel de milésimas.

### Tamaño de la imagen
Se analiza en espacio que ocupa cada imagen base:

```
go_alpine_1.16             latest          7f201f405681       337MB
go_bullseye_1.16           latest          dafa5fb683e3       954MB
go_bullseye_1.17           latest          d2c15c730270       976MB
go_alpine_1.17             latest          8c6c1dda8a53       351MB
```

Encontramos una diferencia bastante notoria relativa al peso de las imágenes, la imagen Alpine es bastante más ligera que el resto. La más ligera es la de la versión 1.16 de Go, pero en tan solo 14MB. Teniendo en cuenta nuestros requisitos, la imagen de Alpine es una clara candidata para nuestro proyecto.

Aun no se puede tomar una decisión final, analizando el rendimiento de las imágenes encontré [esta página](https://nickjanetakis.com/blog/benchmarking-debian-vs-alpine-as-a-base-docker-image) donde se realiza un benchmarking entre imágenes Debian y Alpine. La conclusión es que son dos imágenes con las que se obtienen resultados muy similares y que a no ser que se encuentren errores significativos con Alpine por su tamaño y velocidad es más recomendable.


### Tiempos de ejecución

Se ejecutan los contenedores que hemos generado. Los resultados que obtenemos son los siguientes:

* 1.17 BULLSEYE

ok  	MyWallet/internal/mywallet	0.009s

ok  	MyWallet/internal/mywallet	0.001s

Me parecio extraño el primer tiempo, volví a ejecutar y se obtuvieron varias veces el segundo tiempo.

* 1.17 ALPINE

ok  	MyWallet/internal/mywallet	0.001s

* 1.16 ALPINE

ok  	MyWallet/internal/mywallet	0.001s

* 1.16 BULLSEYE

ok  	MyWallet/internal/mywallet	0.001s


### Conclusión

Los tiempos de ejecución y de construcción son prácticamente iguales, la principal diferencia se encuentra en el tamaño de la imagen, en este apartado Alpine sale claro vencedor. Por tanto, mi decisión final es que a no ser que se encuentre algún error importante con imágenes Alpine se usará `golang:1.17-alpine` como imagen base para nuestro proyecto. La versión de Go que he elegido es la 1.17, si es verdad que la versión 1.16 pesaba algo menos, pero me parece algo poco significativo, más significativo me parece que una versión reciba más actualizaciones que la otra, por esto se elije la versión 1.17.

## Facilitar uso de Docker con nuestro task runner

Se ha automatizado la ejecución de los tests en el task runner. Esto se consigue añadiendo la siguiente directiva a nuestro Taskfile.yml:

```docker run -t -v `pwd`:/app/test luisarostegui/mywallet```

## Buenas prácticas en nuestro Dockerfile

* Vamos a usar una imagen ligera (Alpine) para optimizar el tamaño y poder tener el control sobre los paquetes necesarios.
* Usar variables con ENV para directorios de trabajo.
* Ejecutar tanto las instalaciones de las dependencias como el task runner como usuario y no como superusuario.

## Justificación de directivas en el Dockerfile

1. Con la directiva FROM especificamos la imagen base.
2. La directiva LABEL la utilizamos para especificar nombre y correo de la persona encargada del Docker.
3. Con ENV especificamos la ruta donde queremos que se ejecuten los test, sirve para crear una variable de entorno.
4. RUN lo usamos para ejecutar ordenes de terminal, tales como crear un grupo y un usuario.
5. USER lo usamos para cambiar de usuario.
6. La directiva COPY permite copiar los ficheros de dependecias a la carpeta /app/test.
7. WORKDIR, especificamos la ruta donde queremos trabajar.
8. ENTRYPOINT, indicamos la acción a ejecutar, en este caso `task test`.

## Comentarios en el Dockerfile

```
#Imagen base para docker
FROM golang:1.17-alpine

# Metadatos de información del encargado de mantenimiento
LABEL maintainer="Luis Aróstegui Ruiz <luisarostegui@correo.ugr.es>"

# Creamos variable de entorno para el directorio donde vamos a ejecutar los tests
ENV TEST_DIR=/app/test/

# Añadimos usuario sin privilegios de superusuario y cremos un grupo para dicho usuario
RUN addgroup -S mywallet && adduser -S mywallet -G mywallet

# Cambiamos al nuevo usuario
USER mywallet

#Instalamos modulos necesarios para compilar
COPY go.mod /app/

#Ahora podemos descargar y actualizar las dependecias
RUN go mod download

#Instalamos nuestro task runner
RUN go install github.com/go-task/task/v3/cmd/task@latest

#Establecemos el directorio donde vamos a ejecutar los tests con nuestro nuevo usuario
WORKDIR $TEST_DIR

#Especificamos el ejecutable que usará el contenedor
ENTRYPOINT ["task", "test"]
```

## Uso de Docker Hub

Para realizar la configuración de mi repositorio con Docker Hub he consultado [esta página](https://docs.docker.com/docker-hub/builds/).

Se tiene que crear un Github Action que crea la imagen del contenedor y la publica en Docker Hub de forma automática. Con esto conseguiremos que cada vez que avancemos de objetivo (se haga un push a la rama principal), se automatice la construcción de la imagen. También me ha sido de utilidad la siguiente [documentación](https://docs.docker.com/ci-cd/github-actions/).

### Construcción de nuestro fichero para el workflow

1. Indicamos cuando se debe de publicar la imagen en docker hub.

```yaml
on:
  push:
    paths: # Indicamos los ficheros que tiene que analizar para realizar la publicación de la imagen.
      - Dockerfile #  Si estos ficheros no se han modificado no se realiza la publicación
      - go.mod

```

La imagen depende de las dependencias y del Dockerfile, por tanto:

* Se indica que cuando se haga un push bien a la rama main o bien se modifiquen las dependencias del proyecto, go.mod, o se modifique el Dockerfile se generará una imagen de nuestro proyecto.
* Cuando se realice un pull request hacia la rama main o bien se modique el Dockerfile o go.mod, se generará una nueva imagen.
* Antes se tenía que cuando se hiciese un push a la rama main se actualizase la imagen, pero en este proyecto se avanza mergeando PR de una rama al main por tanto nunca se va a hacer un push a la rama main directamente.

3. Especificamos que queremos que suceda dentro de nuestro flujo de trabajo.
```yaml
jobs:
  build:
    runs-on: ubuntu-latest # Indicamos que se ejecute en las últimas instancias de Ubuntu disponibles.
    
    steps:
      - name: Checkout # Revisa nuestro repo en $GITHUB_WORKSPACE para que nuestro workflow pueda acceder a el.
        uses: actions/checkout@v2
      - name: Login to Docker Hub # Iniciamos sesión en docker hub
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.DOCKER_HUB_USERNAME }}
          password: ${{ secrets.DOCKER_HUB_ACCESS_TOKEN }}
      - name: Build and push # Construimos la imagen y la publicamos
        uses: docker/build-push-action@v2
        with:
          context: .
          file: ./Dockerfile
          push: true
          tags: ${{ secrets.DOCKER_HUB_USERNAME }}/mywallet:latest
```


# Cambios de framework de testing

Tras completar el fichero de CI workflow, en nuestro caso [este](.github/workflows/dockerhub.yml), me han surgido problemas para poder ejecutar el Dockerfile. Aparentemente, el Dockerfile tenía buena pinta no había error sintáticos. Pero al ejecutar ```docker run -t -v `pwd`:/app/test luisarostegui/mywallet``` surgió el siguiente error:

```console
 Objetivo-5 U:5 ?:1  ~/UGR/IV/Mi-repo/MyWallet                                                                                                                      12:54:09  luismsi 
❯ task docker
task: [docker] docker run -t -v `pwd`:/app/test luisarostegui/mywallet
task: [test] go test -v ./...
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
go: updates to go.mod needed; to update it:
        go mod tidy
task: Failed to run task "test": exit status 1
task: Failed to run task "docker": exit status 1
```

Lo primero que se me vino a la mente fue que hay un error al instalar nuestro task runner y no consegue ejecutar correctamente la sentencia `Task test` correctamente. Ya que el error nos indica que hay paquetes que no están actualizados caí en que el error tenía que estar referido a los paquetes que importo en mi proyecto pero esto no debería de ser un gran problema, `go mod tidy` actualiza nuestras dependecias y `go mod download` las descarga... ¿Entonces porque no consigue descargar las dependecias? El único import que tenía que podía suponer problemas era el de **Testify**, probé a quitar este paquete y a ejecutar los tests sin framework de test y correcto, ahi estaba el error, la versión 1.17 de Go no puede incluir esta dependencia. Como el error parecía de la imagen base seleccionada, opté por seleccionar otra imagen, como por ejemplo una imagen que no fuese Alpine (obtenia el mismo error) u otra versión de Go. La opción de otra versión de Go parecía atractiva ya que pude observar en la documentación de Testify que su proyecto funciona de manera estable en versiones de Go desde la 1.13 hasta la 1.15, es decir, realmente no soporta la versión 1.17 que se estaba usando hasta el momento. Cambiando a una versión 1.15, al ejecutar nuestro docker run obtenía el siguiente error:

```console
  Objetivo-5 U:5 ?:1  ~/UGR/IV/Mi-repo/MyWallet                                                                                                                                13:14:07  luismsi 
❯ task docker
task: [docker] docker run -t -v `pwd`:/app/test luisarostegui/mywallet
task: [test] go test -v ./...
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
# runtime/cgo
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in $PATH
FAIL    github.com/LuisArostegui/MyWallet/internal/mywallet [build failed]
FAIL
task: Failed to run task "test": exit status 2
task: Failed to run task "docker": exit status 1
```

Otro error a nuestro catalogo. Al buscar el error parece que tenía que incluir en el Dockerfile una operación como `RUN apk add g++` ya que el error parece que viene de un paquete de Testify que necesita funciones escritas en C. Pero aun incluyendo tenía el mismo error una y otra vez.

Al seguir indagando sobre el error, quizas este estaba en como se montaba el fichero para el workflow de dockerhub, encontré una sección en la documentación oficial para crear este fichero especificamente para este lenguaje, Go, pero al cambiar este fichero no suponía ningún cambio en el error.

Por esto he decidido abandonar Testify y volver a la documentación del anterior objetivo en busca de un framework de test que se ajuste al proyecto. Probé con go-testdeep pero me pareció demasiasdo complejo para los tests que tengo actualmente en mi proyecto. He probado con **Goblin**, del que investigué acerca de este framework en el anterior objetivo y pienso que se adapta perfectamente al proyecto, sencillo crear tests y se informa de manera muy gráfica y agradable si se pasan o no los tests.