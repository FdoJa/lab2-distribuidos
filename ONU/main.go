package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	pb "github.com/FdoJa/ONU/proto"
	"google.golang.org/grpc"
)

type onuServer struct {
	mu  sync.Mutex
	ids map[string]string
	pb.UnimplementedONUServer
	omsClient pb.ONUClient
}

func main() {
	// Iniciar ONU
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Fallo en escuchar: %v", err)
	}

	s := grpc.NewServer()
	onuServer := &onuServer{}

	// Conectar con la OMS
	nameNodeAddr := "oms-container:50052"
	conn, err := grpc.Dial(nameNodeAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el NameNode: %v", err)
	}
	defer conn.Close()

	onuServer.omsClient = pb.NewONUClient(conn)

	pb.RegisterONUServer(s, onuServer)

	fmt.Println("Servidor ONU escuchando en :50055")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo en serve: %v", err)
	}

	for {
		var peticion string
		log.Println("Ingresa el estado de las personas que requieres (Infectados/Muertos):")
		_, err := fmt.Scanln(&peticion)
		if err != nil {
			fmt.Printf("Error al leer la entrada: %v\n", err)
			return
		}

		if strings.ToUpper(peticion) == "INFECTADOS" || strings.ToUpper(peticion) == "MUERTOS" {
			ctx := context.Background()
			respuesta, err := onuServer.omsClient.ConsultarNombres(ctx, &pb.Consulta{
				Estado: peticion,
			})
			if err != nil {
				log.Printf("Error al consultar a la OMS: %v\n", err)
			} else {
				for _, nombre := range respuesta.Nombres {
					fmt.Printf("Nombre: %s, Apellido: %s\n", nombre.Nombre, nombre.Apellido)
				}
			}
		} else {
			log.Printf("Error: Elige una opción válida")
		}
	}
}
