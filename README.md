##  Email Search

Este repositorio cuenta de tres partes:

#### **Indexer** 

Aplicación de consola que indexa bases de datos de correos con sintaxis [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) y [RFC 6532](https://rfc-editor.org/rfc/rfc6532.html) en el servidor de busquedas Zincsearch. 

#### **Webserver**

 Aplicación de lado del servidor que brinda API Rest para realizar busquedas de la base de datos de correos y sirve contenido estatico.

#### **UI** 

Aplicación de lado cliente que brinda una UI sencilla y amigable para consumir el servicio API Rest y visualizar los correos.

---

### Tecnologías utilizadas

<table><tbody><tr><td>Go</td><td>Server side</td></tr><tr><td>Chi</td><td>Router</td></tr><tr><td>Vue3</td><td>Web app</td></tr><tr><td>Tailwind</td><td>Styles</td></tr></tbody></table>

---

## Indexer

Para indexar una base de datos es sencillo como pasar la ruta como parametro en el main.go o desde un build.

**\* Correr los comandos desde la siguiente ruta:  “REPOSITORIO/indexer/”** 

> \<dir> es la ruta donde se encuentra el directorio con los correos
> 
> \<nameExec> es el nombre que le quieras dar al ejecutable

#### Desde el main.go 

```plaintext
go run main.go <dir>
```

#### Creando un ejecutable

```plaintext
go build -o <nameExec> main.go

./<nameExec> <dir>
```

#### Utilizando el ejecutable del repositorio

```plaintext
./indexer <dir>
```

## Formatting blocks with Markdown

Markdown can be used to create various block-level features, such as: