# Clase 10

En esta clase se abordó el flujo de la Web App, al igual que se enseño como desplegar una aplicación en AWS S3.

## Buildear la aplicación

En mi caso utilicé Vite con Bun para el desarrollo de la aplicación, por lo que para buildear la aplicación se debe ejecutar el siguiente comando:

```bash
bunx --bun vite build
```

## Política de acceso a S3

Para que los archivos de la aplicación sean accesibles desde la web, se debe configurar una política de acceso en el bucket de S3. Para ello se debe acceder a la consola de AWS y en el bucket correspondiente se debe ir a la pestaña de permisos y en la sección de "Políticas de bucket" se debe agregar la siguiente política:

```bash
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PublicReadGetObject",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::NOMBRE_DEL_BUCKET/*"
    }
  ]
}
```
