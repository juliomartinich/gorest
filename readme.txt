Este programa implementa un webservice rest del recurso articles 
con los verbos GET y POST en el puerto 10000

----------------------------------------------------------------
RUN

Para hacer correr el programa

go run rest.go

-----------------------------------------------------------------
Prueba

desde una sesion de terminal en el mismo servidor

curl -X GET http://localhost:10000/articles
curl -X POST http://localhost:10000/articles

también 

curl  http://localhost:10000/articles

que es igual a la opción GET porque es el default


o desde un browser 

http://localhost:10000/

