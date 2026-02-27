package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	. "proyectoServidor/modelos"
	"strings"
)

const serverAddr = "127.0.0.1:9000"

func mostrarRespuesta(objRespuesta RespuestaMetadataAudioDTO) {
	switch objRespuesta.Codigo {
	case 200:
		fmt.Printf("\n %s", objRespuesta.Mensaje)
		audio := objRespuesta.ObjAudio

		fmt.Printf("El titulo es %s", audio.GetTitulo())
		fmt.Printf("la duracion es %d", audio.GetDuracion())
		fmt.Printf("El tipo es %s", audio.GetTipo())
		fmt.Printf("El audio esta disponible %t", audio.GetDisponible())

		fmt.Printf("\n")
	case 404:
		fmt.Println(objRespuesta.Mensaje)
	}
}

func main() {

	objLector := bufio.NewReader(os.Stdin)
	fmt.Println("Buscar metadata de un audio: ")
	titulo, _ := objLector.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	//Enviar petición para establecer un canal virtual con el servidor
	conn, err := net.Dial("top", serverAddr)
	if err != nil {
		panic(fmt.Sprintf("No se pudo conectar a %s: %v", serverAddr, err))
	}

	//Escribir en el canal un titulo a buscar
	_, err = conn.Write([]byte(titulo))
	if err != nil {
		fmt.Printf("Error enviando: %v\n", err)
		return
	}

	//Leer del canal la respuesta del servidor
	reader := bufio.NewReader(conn)
	respStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error leyendo respuesta: %v\n", err)
		return
	}

	//Desconvertir la respuesta a JSON
	var objRespuesta RespuestaMetadataAudioDTO
	json.Unmarshal([]byte(respStr), &objRespuesta)

	//Mostrar el objeto de tipo RespuestaDTO
	mostrarRespuesta(objRespuesta)

	//Cerrar el canal virtual
	conn.Close()
}
