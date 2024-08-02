package commands

import (
	structures "CLASE02/structures"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// ParseRep lee todos los registros de estudiantes desde el archivo binario y los devuelve
func ParseRep() ([]structures.Student, error) {
	const fullPath = "/home/keviin/University/PRACTICAS/MIA_LAB_S2_2024/CLASE02/disco.mia" // Aquí iría el path completo de su computadora

	// Abre el archivo binario en modo lectura
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close() // Asegura que el archivo se cierre al finalizar

	// Calcula el tamaño de la estructura Student en bytes
	studentSize := binary.Size(structures.Student{})
	offset := int64(0) // Inicializa el offset en 0
	var students []structures.Student

	for {
		var student structures.Student
		// Lee la estructura Student desde el archivo en la posición actual de offset
		err := student.ReadFromFile(offset)
		if err != nil {
			if err == io.EOF {
				break // Si se alcanza el final del archivo, sale del bucle
			}
			return nil, nil // Retorna nil si hay otro tipo de error
		}

		// Verifica si el bloque es libre (todos los bytes son 0)
		if !isBlockFree(student.Carnet[:]) {
			fmt.Printf("Estudiante: %s\n", student.String()) // Imprime la información del estudiante
			students = append(students, student)             // Agrega el estudiante a la lista
		}

		offset += int64(studentSize) // Incrementa el offset para leer el siguiente bloque
	}

	return students, nil // Retorna la lista de estudiantes leídos
}

// isBlockFree verifica si un bloque está libre (todos los bytes son 0)
func isBlockFree(block []byte) bool {
	for _, b := range block {
		if b != 0 {
			return false // Retorna false si encuentra un byte diferente de 0
		}
	}
	return true // Retorna true si todos los bytes son 0
}
