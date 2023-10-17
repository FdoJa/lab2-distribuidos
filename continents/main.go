package main

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	pb "github.com/FdoJa/distribuidos-lab2/proto"
	"google.golang.org/grpc"
)

const (
	Infectado = "Infectado"
	Muerto    = "Muerto"
)

type Persona struct {
	Nombre   string
	Apellido string
	Estado   string
}

func main() {
	nameNodeAddr := "oms-container:50052" // Reemplazarlo con IP y puerto de la OMS

	conn, err := grpc.Dial(nameNodeAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectar con el NameNode: %v", err)
	}
	defer conn.Close()

	client := pb.NewContinentsClient(conn)

	personas := generarPersonas()

	for i, persona := range personas {
		_, err := client.InformePersona(context.Background(), &pb.Informe{
			Nombre:   persona.Nombre,
			Apellido: persona.Apellido,
			Estado:   persona.Estado,
		})

		if err != nil {
			log.Printf("Error al registrar persona: %v", err)
		} else {
			log.Printf("Estado enviado: %s %s %s", persona.Nombre, persona.Apellido, persona.Estado)
		}

		if i >= 4 {
			time.Sleep(3 * time.Second)
		}
	}
}

// Función para generar lista de personas
func generarPersonas() []Persona {
	rand.Seed(time.Now().UnixNano())

	var personas []Persona

	nombres, err := os.Open("names.txt")
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer nombres.Close()

	scanner := bufio.NewScanner(nombres)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Fields(linea)
		if len(partes) == 2 {
			nombre := partes[0]
			apellido := partes[1]
			estado := Infectado

			if rand.Float64() > 0.45 {
				estado = Muerto
			}

			personas = append(personas, Persona{nombre, apellido, estado})
		}
	}
	return personas
}
