package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/FdoJa/DataNode1/proto"
	"google.golang.org/grpc"
)

type dataNodeServer struct {
	pb.UnimplementedDataNodeServer
}

func (s *dataNodeServer) RegistrarNombre(ctx context.Context, registro *pb.Registro) (*pb.RespuestaVacia, error) {
	// Almacenar la información en DATA.txt
	txt := "DATA.txt"

	archivo, err := os.OpenFile(txt, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error al crear o abrir archivo: %v", err)
	}
	defer archivo.Close()

	datos := registro.Id + " " + registro.Nombre + " " + registro.Apellido + "\n"
	_, err = archivo.WriteString(datos)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}

	// Devolver respuesta vacia a OMS
	return &pb.RespuestaVacia{}, nil
}

func (s *dataNodeServer) PedirNombre(ctx context.Context, peticion *pb.Peticion) (*pb.Nombre, error) {
	txt := "DATA.txt"
	archivo, err := os.Open(txt)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
		return nil, err
	}
	defer archivo.Close()

	var nombre, apellido string
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		if data[0] == peticion.Id {
			nombre = data[1]
			apellido = data[2]
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error al leer el archivo: %v", err)
		return nil, err
	}

	return &pb.Nombre{
		Nombre:   nombre,
		Apellido: apellido,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Fallo en escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataNodeServer(s, &dataNodeServer{})

	fmt.Println("Servidor DataNode1 escuchando en :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo en serve: %v", err)
	}
}
