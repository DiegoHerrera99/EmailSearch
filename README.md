## Email Search

Este repositorio cuenta de tres partes:

#### **Indexer**

Aplicación de consola que indexa bases de datos de correos con sintaxis [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) y [RFC 6532](https://rfc-editor.org/rfc/rfc6532.html) en el servidor de busquedas Zincsearch. 

#### **Webserver**

 Aplicación de lado del servidor que brinda API Rest para realizar busquedas de la base de datos de correos y sirve contenido estatico.

#### **UI - Email reader**

Aplicación de lado cliente que brinda una UI sencilla y amigable para consumir el servicio API Rest y visualizar los correos.

---

## Indexer

Para indexar una base de datos es sencillo como pasar la ruta como parametro en el main.go o desde un build.

**\* Correr los comandos desde la siguiente ruta:  “REPOSITORIO/indexer/”**

> \<dir> es la ruta donde se encuentra el directorio con los correos
> 
> \<nameExec> es el nombre que le quieras dar al ejecutable

#### Desde el main.go

```plaintext
>> go run main.go <dir>
```

#### Creando un ejecutable

```plaintext
>> go build -o <nameExec> main.go

>> ./<nameExec> <dir>
```

#### Utilizando el ejecutable del repositorio

*   **Primera vez**

El script ./indexer realiza la configuración inicial del servicio. Luego de indexar la base de datos \<dir> el servicio zincsearch queda funcionando en segundo plano.

> Se puede consultar la instancia de zincsearch desde la dirección **http://localhost:4080**

```plaintext
>> ./indexer <dir>
```

*   **Futuras ocasiones**

Para futuras ocasiones ya no es necesario indexar de nuevo \<dir>.

 Se puede seguir consumiendo la base de datos, utilizando el siguiente comando de consola.

> Se puede consultar la instancia de zincsearch desde la dirección **http://localhost:4080**

```plaintext
>> zinc
```

 \* Si se cambian las credenciales de inicio de sesión de zincearch es necesario cambiarlas en el archivo **REPOSITORIO/indexer/globals/globals.go**

---

## Webserver

En el repositorio ya viene cargado un build del UI en la carpeta public del webserver, para levantar el webserver solo es necesario pasar como flag desde la terminal el puerto. Esto puede hacerse desde un build de Go o desde el main.go

*   **El contenido estatico se sirve desde la carpeta “REPOSITORIO/webserver/public/”**

> \<port> es el número de puerto donde se desea levantar el servidor
> 
> \<nameExec> es el nombre que le quieras dar al ejecutable

#### Desde el main.go

```plaintext
>> go run main.go <port>
```

#### Creando un ejecutable

```plaintext
>> go build -o <nameExec> main.go

>> ./<nameExec> <port>
```

#### Utilizando el ejecutable del repositorio

> La ruta de acceso de la interfaz gráfica de consultas de la DB es accesible desde la siguiente dirección **http://localhost:port**

```plaintext
>> ./webserver <port>
```

---

## UI - Email reader

Interfaz sencilla para consumir servicios api desde el webserver realizada en Vue3 y Tailwind CSS. 

> **Ya un build viene cargado en la carpeta public del webserver**

---

## Tecnologías utilizadas

<table><tbody><tr><td>Golang</td><td>Server side</td></tr><tr><td>Chi</td><td>Router</td></tr><tr><td>Vue3</td><td>Web app</td></tr><tr><td>Tailwind CSS</td><td>Styles</td></tr></tbody></table>

Para mayor información sobre la herramienta zincsearch ir a [https://docs.zinc.dev](https://docs.zinc.dev)