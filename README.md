**Laboratorio 2 - Sistemas distribuidos**

Por ahora se realiza contienentes como un contenedor, donde le manda los nombres a la OMS. Para la entrega, el código de un contienente debe estar presente en cada actor, estos siendo: ONU, OMS, DataNode1, DataNode2; ya que se especifica en las reglas que los contienentes deben estar en cada máquina.

**Compilar:**

Se utiliza *docker-compose* para, por ahora, realizar la conexión entre los distintos contenedores a través de una network y luego ejecutarlos. 

1) Ubicar la terminal en la carpeta donde se encuentre el archivo *docker-compose*
2) Ejecutar en la terminal el comando: **_docker-compose up --build_**