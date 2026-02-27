package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"proyecto/modelos"
	"proyecto/servicios"
)

func main() {

	vectorMetadataAudios := make([]modelos.MetadataAudio, 5)
	servicios.CargarMetadataAudios(&vectorMetadataAudios)

	var opcion int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==== Menú ====")
		fmt.Println("1. Buscar metadata de un audio")
		fmt.Println("2. Salir")
		fmt.Print("Opción: ")
		fmt.Scan(&opcion)
		reader.ReadString('\n')

		switch opcion {
		case 1:
			fmt.Println("Digite el título del audio: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			objRespuesta := servicios.BuscarAudio(titulo, vectorMetadataAudios)

			switch objRespuesta.Codigo {
			case 200:
				fmt.Printf("\n%s", objRespuesta.Mensaje)
				audio := objRespuesta.ObjAudio
				fmt.Printf("\nTítulo del audio: %s", audio.GetTitulo())
				fmt.Printf("\nDuración del audio: %d", audio.GetDuracion())
				fmt.Printf("\nTipo de audio: %s", audio.GetTipo())
				fmt.Printf("\nEl audio está disponible: %t\n", audio.GetDisponible())
			case 404:
				fmt.Println(objRespuesta.Mensaje)
			}

		case 2:
			fmt.Println("Programa terminado")
			return

		default:
			fmt.Println("Opción incorrecta")
		}
	}
}
