package servicios

import (
	. "proyectoServidor/modelos"
)

func CargarMetadataAudios(vec []MetadataAudio) {
	var objAudio1, objAudio2, objAudio3, objAudio4, objAudio5 MetadataAudio

	objAudio1.SetTitulo("Canción 1")
	objAudio1.SetDuracion(10)
	objAudio1.SetTipo("Música")
	objAudio1.SetDisponible(true)

	objAudio2.SetTitulo("Podcast 2")
	objAudio2.SetDuracion(20)
	objAudio2.SetTipo("Podcast")
	objAudio2.SetDisponible(false)

	objAudio3.SetTitulo("Ruido Blanco 3")
	objAudio3.SetDuracion(30)
	objAudio3.SetTipo("Ruido Blanco")
	objAudio3.SetDisponible(true)

	objAudio4.SetTitulo("Audiolibro 4")
	objAudio4.SetDuracion(40)
	objAudio4.SetTipo("Audiolibros")
	objAudio4.SetDisponible(true)

	objAudio5.SetTitulo("Meditación 5")
	objAudio5.SetDuracion(50)
	objAudio5.SetTipo("Meditaciones guiadas")
	objAudio5.SetDisponible(false)

	(vec)[0] = objAudio1
	(vec)[1] = objAudio2
	(vec)[2] = objAudio3
	(vec)[3] = objAudio4
	(vec)[4] = objAudio5
}

func BuscarAudio(titulo string, vectorMetadataAudios []MetadataAudio) RespuestaMetadataAudioDTO {
	for i := 0; i < len(vectorMetadataAudios); i++ {
		if vectorMetadataAudios[i].GetTitulo() == titulo {
			var resp RespuestaMetadataAudioDTO
			resp.ObjAudio = vectorMetadataAudios[i]
			resp.Codigo = 200
			resp.Mensaje = "Métadata del audio encontrada"
			return resp
		}
	}
	var resp RespuestaMetadataAudioDTO
	resp.Codigo = 404
	resp.Mensaje = "La métadata del audio no se encontró"
	return resp
}
