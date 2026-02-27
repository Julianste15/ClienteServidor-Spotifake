package main

import (
	"encoding/json"
	"fmt"
	"net"
	alias "proyectoServidor/modelos"
	"proyectoServidor/servicios"
	"strings"
	"time"
)

const addr = "localhost:9000"

func funcionManejarPeticion(conn net.Conn, vectorMetadataAudios []alias.MetadataAudio) {

	remote := conn.RemoteAddr().String()
	obtenerDatosCliente(conn)

	//Leer del canal el titulo enviado por el cliente
	buffer := make([]byte, 1024) // tamaño máximo esperado
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error leyendo desde %s: %v\n", remote, err)
		return
	}

	titulo := string(buffer[:n])
	titulo = strings.TrimSpace(titulo)
	fmt.Printf("Título recibido de %s: %q\n", remote, titulo)

	objRespuesta := servicios.BuscarAudio(titulo, vectorMetadataAudios)

	//convertimos la respuesta a json para enviarla al cliente
	jsonResp, _ := json.Marshal(objRespuesta)

	fmt.Printf("%s \n", jsonResp)
	//simulación de procesamiento de la petición, se duerme el proceso por 10 segundos
	//para que el cliente pueda ver que el servidor sigue atendiendo otras peticiones
	//mientras atiende esta petición
	time.Sleep(10 * time.Second)

	//Escribir en el canal la respuesta
	_, err = conn.Write(append(jsonResp, '\n'))
	if err != nil {
		fmt.Printf("Error enviando a %s: %v\n", remote, err)
		return
	}

	fmt.Printf("Atendido %s, título: %q\n", remote, titulo)

	//Cerrar el canal virtual con el cliente
	conn.Close()
}

func obtenerDatosCliente(conn net.Conn) {

	remote := conn.RemoteAddr().String()
	fmt.Printf("Conexión entrante desde %s\n", remote)

	// Obtener IP y Puerto del cliente por separado
	host, port, err := net.SplitHostPort(remote)
	if err != nil {
		fmt.Printf("Error al separar host y puerto: %v\n", err)
	} else {
		fmt.Printf("Cliente IP: %s, Puerto Efímero: %s\n", host, port)
	}
}

func main() {

	//Al iniciar el servidor se cargan los metadatos de los audios
	vectorMetadataAudios := make([]alias.MetadataAudio, 5)
	servicios.CargarMetadataAudios(vectorMetadataAudios)

	//Colocar el servidor en estado de escucha
	ln, err := net.Listen("top", addr)
	if err != nil {
		panic(fmt.Sprintf("No se pudo escuchar en %s: %v", addr, err))
	}
	// defer permite cerrar el canal virtual cuando se termine la ejecución del main
	defer ln.Close()

	fmt.Printf("Servidor escuchando en %s\n", addr)

	for {
		objReferenciaCanal, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error al aceptar conexión: %v\n", err)
			continue
		}
		go funcionManejarPeticion(objReferenciaCanal, vectorMetadataAudios)
	}
}
