# RETO MELI

Éste proyecto permite desplegar un frontend (Flask) junto con dos APIs (Flask, Golang) en una instancia EC2 de AWS.

## Local

### Golang

<strong>Versión Golang:</strong> go1.18.1 linux/amd64

Para copilar la aplicación Golang, navegue hasta la carpeta /golang del proyecto y ejecute el siguiente comando:
<code>CGO_ENABLED=1 go build -v</code>.

Una vez copilado, ejecute el servidor usando:
<code>./meli-go</code>

La API podrá ser consultada en la ruta http://localhost:8080.

### Flask

Para ejecutar la aplicación Flask, navegue hasta la carpeta /flask del proyecto y ejecute los siguientes comandos:
<code>virtualenv meli_venv</code>
<code>source meli_venv/bin/activate</code>
<code>pip install -r requirements.txt</code>
<code>python run.py</code>

Primero se crea un ambiente virtual donde correrá la aplicación, despues se activa el ambiente virtual, se instalan las dependencias necesarias relacionadas en el archivo "requirements.txt" y finalmente se puede correr la aplicación. Se debe tener en cuenta que el servidor de flask que se ejecuta es un servidor de desarrollo. La aplicación podrá ser accedida en la ruta http://localhost:5000.

## Flask API-Front

La aplicación permite crear y consultar usuarios. El uso de la misma está protegido bajo inicio de sesión. Las rutas disponibles en la aplicación son:

/login: Página de inicio de sesión.
/users: Permite consultar todos los usuarios creados.
/users/new: Permite crear nuevos usuarios.

## Golang API

La api permite consultar los usuarios creados con la aplicación Flask. Las rutas disponibles de la API son:
/: Retorna información relacionada a las capacidades de la API.
/users: Permite consultar todos los usuarios creados.
/users/<id>: Permite consultar un usuario en específico.

## Despliegue

El despliegue de la aplicación se realiza con el apoyo de Serverless Framework, con el cual podemos definir unos templates .yml que contendrán los recursos necesarios en formato CloudFormation. Para desplegar la aplicación se usa una instancia EC2 de AWS, en la cual se corren ambas aplicaciones, y para gestionar el ruteo de las mismas dentro de el servidor se usa nginx.

Al navegar a la ruta "ip_server/" se podrá acceder a la aplicación de flask.
Al navegar a la ruta "ip_server/golang/" se podrá acceder a la API de Golang.

ip_server será la ip pública asignada a la instancia EC2.

Para realizar el despliegue será necesario contar con el instalador de dependencias de node "npm", además, el cli de aws debe estar instalado y las credenciales necesarias deben estar configuradas. Una vez cumplidas éstas condiciones iniciales se debe navegar a la raíz del proyecto, y ahí ejecutar el despliegue con el siguiente comando de serverless:
<code>sls deploy --config serverless-service.yml --stage dev --region us-east-1 --app meli-challenge --verbose</code>
