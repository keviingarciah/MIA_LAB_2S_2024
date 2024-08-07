package structures

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

// Student representa la estructura de un estudiante con sus atributos
type Student struct {
	Carnet [10]byte // Carnet del estudiante
	CUI    [13]byte // CUI del estudiante
	Name   [25]byte // Nombre del estudiante
	Age    [4]byte  // Edad del estudiante
	// Total Size: 52 bytes
}

const fullPath = "students.dat" // Ruta del archivo binario

// WriteToFile escribe la estructura Student en el primer bloque libre de un archivo binario
func (s *Student) WriteToFile() error {
	// Abre el archivo binario en modo lectura/escritura, lo crea si no existe
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close() // Asegura que el archivo se cierre al finalizar

	// Calcula el tamaño de la estructura Student en bytes
	studentSize := binary.Size(*s)
	// Busca el primer bloque libre en el archivo
	offset, err := findFreeBlock(file, studentSize)
	if err != nil {
		return fmt.Errorf("error al buscar bloque libre: %v", err)
	}

	// Mueve el puntero del archivo a la posición del bloque libre encontrado
	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("error al mover el puntero del archivo: %v", err)
	}

	// Escribe la estructura Student en el archivo usando codificación Little Endian
	err = binary.Write(file, binary.LittleEndian, s)
	if err != nil {
		return fmt.Errorf("error al escribir la estructura Student: %v", err)
	}

	return nil // Retorna nil si no hubo errores
}

// ReadFromFile lee la estructura Student desde un archivo binario en una posición específica
func (s *Student) ReadFromFile(offset int64) error {
	// Abre el archivo binario en modo lectura
	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close() // Asegura que el archivo se cierre al finalizar

	// Mueve el puntero del archivo a la posición especificada por offset
	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("error al mover el puntero del archivo: %v", err)
	}

	// Lee la estructura Student desde el archivo usando codificación Little Endian
	err = binary.Read(file, binary.LittleEndian, s)
	if err != nil {
		return fmt.Errorf("error al leer la estructura Student: %v", err)
	}

	return nil // Retorna nil si no hubo errores
}

// String devuelve una representación en cadena de la estructura Student
func (s Student) String() string {
	return fmt.Sprintf("Carnet: %s, CUI: %s, Name: %s, Age: %s",
		strings.TrimSpace(string(s.Carnet[:])),
		strings.TrimSpace(string(s.CUI[:])),
		strings.TrimSpace(string(s.Name[:])),
		strings.TrimSpace(string(s.Age[:])),
	)
}

// findFreeBlock busca el primer bloque libre en el archivo
func findFreeBlock(file *os.File, blockSize int) (int64, error) {
	buffer := make([]byte, blockSize) // Crea un buffer del tamaño del bloque
	var offset int64

	for {
		// Lee un bloque del archivo en la posición actual de offset
		_, err := file.ReadAt(buffer, offset)
		if err != nil {
			break // Si hay un error (EOF), sale del bucle
		}

		isFree := true
		// Verifica si el bloque está libre (todos los bytes son 0)
		for _, b := range buffer {
			if b != 0 {
				isFree = false
				break
			}
		}

		if isFree {
			return offset, nil // Retorna el offset del bloque libre encontrado
		}

		offset += int64(blockSize) // Incrementa el offset para leer el siguiente bloque
	}

	return offset, nil // Retorna el offset del final del archivo si no se encontró un bloque libre
}
