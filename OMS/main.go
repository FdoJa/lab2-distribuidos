package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	pb "github.com/FdoJa/OMS/proto"
	"google.golang.org/grpc"
)

var id int = 0

const (
	DataNode1 = "DataNode1"
	DataNode2 = "DataNode2"
	Address1  = "datanode1-container:50053"
	Address2  = "datanode2-container:50054"
)

type continentServer struct {
	mu sync.Mutex
	pb.UnimplementedContinentsServer
}

type onuServer struct {
	mu sync.Mutex
	pb.UnimplementedONUServer
}

func (o *onuServer) ConsultarNombres(ctx context.Context, consulta *pb.Consulta) (*pb.RespuestaNombres, error) {
	// Lógica para pedir las personas "Infectadas" o "Muertas"
	txt := "DATA.txt"

	archivo, err := os.Open(txt)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer archivo.Close()

	var nombres []*pb.Nombre
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		if consulta.Estado == data[3] {
			id := data[0]
			var address string

			if data[1] == DataNode1 {
				address = Address1
			} else {
				address = Address2
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("No se pudo conectar al DataNode: %v", err)
			}
			defer conn.Close()

			client := pb.NewDataNodeClient(conn)

			respuesta, err := client.PedirNombre(context.Background(), &pb.Peticion{
				Id: id,
			})

			if err != nil {
				log.Fatalf("Error al registrar nombre en el DataNode: %v", err)
			} else {
				nombres = append(nombres, &pb.Nombre{
					Nombre:   respuesta.Nombre,
					Apellido: respuesta.Apellido,
				})
			}
		}
	}

	respuestaNombres := &pb.RespuestaNombres{
		Nombres: nombres,
	}

	return respuestaNombres, nil
}

func (s *continentServer) InformePersona(ctx context.Context, informe *pb.Informe) (*pb.RespuestaVacia, error) {
	// Lógica para registrar la persona en DATA.txt y elegir un DataNode

	// Generar ID
	id++
	personaID := strconv.Itoa(id)

	// Elige un DataNode para almacenar la información y registra la persona en DATA.txt
	var dataNode, address string = elegirDataNode(informe.Apellido)

	// Registrar la persona en DATA.txt
	registrarEnTxt(personaID, dataNode, informe.Estado)

	// Enviar datos al DataNode correspondiente
	enviarDatosAlDataNode(address, personaID, informe)

	// Una vez registrada la persona, enviar una respuesta vacía al continente
	return &pb.RespuestaVacia{}, nil
}

func registrarEnTxt(personaID string, dataNode string, estado string) {
	txt := "DATA.txt"

	archivo, err := os.OpenFile(txt, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error al crear o abrir archivo: %v", err)
	}
	defer archivo.Close()

	datos := personaID + " " + dataNode + " " + estado + "\n"
	_, err = archivo.WriteString(datos)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}
}

func elegirDataNode(apellido string) (string, string) {
	var dataNode string
	var address string
	if apellido != "" {
		initial := apellido[0]
		if 'A' <= initial && initial <= 'M' {
			dataNode = DataNode1
			address = Address1
		} else {
			dataNode = DataNode2
			address = Address2
		}
	}
	return dataNode, address
}

func enviarDatosAlDataNode(dataNodeAddr string, id string, informe *pb.Informe) {
	conn, err := grpc.Dial(dataNodeAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al DataNode: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataNodeClient(conn)

	_, err = client.RegistrarNombre(context.Background(), &pb.Registro{
		Id:       id,
		Nombre:   informe.Nombre,
		Apellido: informe.Apellido,
	})
	if err != nil {
		log.Fatalf("Error al registrar nombre en el DataNode: %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50052") // Puedes elegir el puerto que desees
	if err != nil {
		log.Fatalf("Fallo en escuchar: %v", err)
	}

	s := grpc.NewServer()
	o := grpc.NewServer()
	pb.RegisterContinentsServer(s, &continentServer{})
	pb.RegisterONUServer(o, &onuServer{})

	fmt.Println("Servidor NameNode/OMS escuchando en :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo en serve: %v", err)
	}
}
