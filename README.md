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

![](https://33333.cdn.cke-cs.com/kSW7V9NHUXugvhoQeFaf/images/87d46e932aeea20915160b074351f54be2670eaae49a0e56.png) ![](https://33333.cdn.cke-cs.com/kSW7V9NHUXugvhoQeFaf/images/bee0d17dd3ae54d28f53b917f3adf426efef243dd3c0a758.png) ![](https://33333.cdn.cke-cs.com/kSW7V9NHUXugvhoQeFaf/images/77408158af96ac6836ff8e0c013118d8bbc3f1df02d7eb23.png) ![](https://33333.cdn.cke-cs.com/kSW7V9NHUXugvhoQeFaf/images/7eea1ac49260bfe7376f82c57e6c9b75a10190570af4e6f9.png) ![](https://33333.cdn.cke-cs.com/kSW7V9NHUXugvhoQeFaf/images/58c67ee5f3191810e650aef88341b39281bd16cdc500780e.png)

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