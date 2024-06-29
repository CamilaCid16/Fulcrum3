package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/yojeje/lab6"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBrokerServer
	pb.UnimplementedIngenierosServer
	pb.UnimplementedKaisServer
}

func EscribirFichero(operacion string, base string, sector string, valor string) {
	file, err := os.Create(sector + ".txt")
	if err != nil {
		fmt.Println("Error al crear eñ archivo", err)
		return
	}

	defer file.Close()

	data := operacion + " " + sector + " " + base

	if valor != "" {
		data += " " + valor
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)

	if err != nil {
		fmt.Println("Error al escribir", err)
		return
	}

	writer.Flush()
}

func AgregarBase(sector string, base string, valor string) {
	EscribirFichero("AgregarBase", base, sector, valor)
}

func RenombrarBase(sector string, base string, nombre string) {
	EscribirFichero("RenombrarBase", base, sector, nombre)
}

func ActualizarValor(sector string, base string, valor string) {
	EscribirFichero("ActualizarValor", sector, base, valor)
}

func BorrarBase(sector string, base string) {
	EscribirFichero("BorrarBase", sector, base, "")
}

func (s *server) EnviarServidor(ctx context.Context, in *pb.Comando) (*pb.Reloj, error) {
	if in.Tipo == "AgregarBase" {
		AgregarBase("SectorAlpha", "Campamento2", "25")
	} else if in.Tipo == "RenombrarBase" {
		RenombrarBase("Campamento2", "SectorAlpha", "Campamento3")
	} else if in.Tipo == "ActualizarValor" {
		ActualizarValor("SectorAlpha", "Campamento2", "30")
	} else if in.Tipo == "BorrarBase" {
		BorrarBase("SectorAlpha", "Campamento2")
	}

	return &pb.Reloj{X: 1, Y: 0, Z: 0}, nil
}

func (s *server) GetEnemigosServidor(ctx context.Context, in *pb.Direccion) (*pb.Enemigos, error) {
	fmt.Println("Consultando enemigos desde el comandante.....")
	return &pb.Enemigos{Cantidad: 0}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50058))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBrokerServer(s, &server{})
	pb.RegisterIngenierosServer(s, &server{})
	pb.RegisterKaisServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falla de servidor: %v", err)
	}
}
