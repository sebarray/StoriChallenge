# strori-challenge


# Proyecto AWS Lambda con Golang y MySQL

Este proyecto es una aplicación basada en AWS Lambda escrita en Golang, que interactúa con una base de datos MySQL y utiliza un bucket de Amazon S3 para almacenar y recuperar datos.

## Requisitos previos

Antes de comenzar con este proyecto, asegúrate de tener los siguientes requisitos previos configurados:

- Cuenta de Amazon Web Services (AWS) con acceso a los servicios necesarios.
- AWS CLI (Command Line Interface) configurado en tu máquina local.
- Golang instalado en tu sistema.
- Base de datos MySQL en funcionamiento.

## Variables de entorno

Antes de desplegar la aplicación, deberás configurar las siguientes variables de entorno:

- `pswmail`: Contraseña para la cuenta de correo electrónico que enviará los correos.
- `sendMail`: Dirección de correo electrónico desde la cual se enviarán los correos.
- `STRING_CONNECTION`: Cadena de conexión a la base de datos MySQL.
- `BUCKET_NAME`: Nombre del bucket de Amazon S3 donde se almacenarán los archivos.
- `PUBLIC_KEY`: Clave pública de la cuenta de AWS para acceder a los servicios.
- `SECRET`: Clave secreta correspondiente a la clave pública.
- `REGION`: Región de AWS donde se desplegará la función Lambda y se accederá a los recursos.



## Despliegue del proyecto con Docker

Para utilizar el proyecto en AWS Lambda, es necesario empaquetarlo en una imagen de Docker y subir esta imagen a Amazon Elastic Container Registry (ECR). A continuación, se describen los pasos para hacerlo:

1. **Construir la imagen de Docker:**

   Para construir la imagen de Docker, asegúrate de tener Docker instalado en tu sistema. Luego, en la terminal o línea de comandos, ve al directorio raíz de tu proyecto y ejecuta el siguiente comando:

`docker build -t servicemail:latest .`


Esto creará una imagen de Docker con el nombre `servicemail` y la etiqueta `latest`, basada en el contenido del directorio actual.

2. **Etiquetar la imagen de Docker:**

Una vez que la imagen de Docker se haya construido correctamente, es necesario etiquetarla con la dirección del registro de contenedores de Amazon ECR (Elastic Container Registry) al que queremos subir la imagen. Suponiendo que tienes acceso a la cuenta de AWS y el cliente de AWS CLI configurado con las credenciales adecuadas, puedes etiquetar la imagen con el siguiente comando:

`docker tag servicemail:latest 211868793080.dkr.ecr.us-east-1.amazonaws.com/send_mail:v1`


En este comando, `211868793080.dkr.ecr.us-east-1.amazonaws.com` es el URI del registro de contenedores de ECR donde almacenaremos la imagen. Asegúrate de reemplazarlo con el URI correspondiente a tu cuenta y región.

3. **Subir la imagen a ECR:**

Con la imagen debidamente etiquetada, ahora podemos subirla al registro de contenedores de ECR con el siguiente comando:

`docker push 211868793080.dkr.ecr.us-east-1.amazonaws.com/send_mail:v1`


Esto subirá la imagen al repositorio `send_mail` con la etiqueta `v1` en tu cuenta de AWS.

4. **Utilizar la imagen en Lambda:**

Una vez que la imagen está en ECR, puedes utilizarla para crear la función Lambda. Ve a la consola de AWS Lambda, crea una nueva función Lambda y selecciona "Usar una imagen de contenedor de ECR" como el origen. Luego, busca y selecciona la imagen `send_mail:v1` que acabas de subir.

Configura la función Lambda según tus necesidades y asegúrate de que tenga los permisos adecuados para interactuar con otros servicios de AWS, como S3 y MySQL.

Con estos pasos, habrás empaquetado y desplegado el proyecto de AWS Lambda con Golang y MySQL utilizando una imagen de Docker. Ahora la función Lambda puede interactuar con la base de datos MySQL y usar el bucket de Amazon S3 para almacenar y recuperar datos.