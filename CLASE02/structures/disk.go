package structures

import (
	"errors" // Importa el paquete "errors" para manejar errores
	"fmt"
	"os" // Importa el paquete "os" para operaciones del sistema operativo
)

// Ruta completa donde se creará el archivo binario
const fullPath = "/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE02/disco.mia" // Aquí iría el path completo de su computadora

// CreateBinaryFile crea un archivo binario con el tamaño y la unidad especificados
func CreateBinaryFile(size int, unit string) error {
	// Convierte el tamaño a bytes
	sizeInBytes, err := convertToBytes(size, unit)
	if err != nil {
		return err
	}

	// Crea el archivo en la ruta especificada
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close() // Asegura que el archivo se cierre al finalizar

	// Escribe los bytes en el archivo
	return writeToFile(file, sizeInBytes)
}

// convertToBytes convierte el tamaño y la unidad a bytes
func convertToBytes(size int, unit string) (int, error) {
	switch unit {
	case "K":
		return size * 1024, nil // Convierte kilobytes a bytes
	case "M":
		return size * 1024 * 1024, nil // Convierte megabytes a bytes
	default:
		return 0, errors.New("unidad inválida") // Devuelve un error si la unidad es inválida
	}
}

// writeToFile escribe los bytes en el archivo
func writeToFile(file *os.File, sizeInBytes int) error {
	buffer := make([]byte, 1024*1024) // Crea un buffer de 1 MB
	for sizeInBytes > 0 {
		writeSize := len(buffer)
		if sizeInBytes < writeSize {
			writeSize = sizeInBytes // Ajusta el tamaño de escritura si es menor que el buffer
		}
		if _, err := file.Write(buffer[:writeSize]); err != nil {
			return err // Devuelve un error si la escritura falla
		}
		sizeInBytes -= writeSize // Resta el tamaño escrito del tamaño total
	}
	fmt.Println("Archivo creado con éxito!")
	return nil
}
