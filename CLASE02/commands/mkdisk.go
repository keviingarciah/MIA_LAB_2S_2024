package commands

import (
	structures "CLASE02/structures" // Importa el paquete "structures" desde el directorio "PRUEBA01/structures"
	"errors"                        // Importa el paquete "errors" para manejar errores
	"fmt"                           // Importa el paquete "fmt" para formatear e imprimir texto
	"strconv"                       // Importa el paquete "strconv" para convertir cadenas a otros tipos
	"strings"                       // Importa el paquete "strings" para manipulación de cadenas
)

// MKDISK estructura que representa el comando mkdisk con sus parámetros
type MKDISK struct {
	size int    // Tamaño del disco
	unit string // Unidad de medida del tamaño (K o M)
}

/*
   Ejemplos de uso del comando mkdisk:
   mkdisk -size=1 -unit=K
   mkdisk -size=1
*/

// ParseMkdisk analiza los tokens del comando mkdisk y crea un objeto MKDISK
func ParseMkdisk(tokens []string) (*MKDISK, error) {
	cmd := &MKDISK{} // Crea una nueva instancia de MKDISK
	for _, token := range tokens {
		// Divide cada token en clave y valor usando "=" como delimitador
		parts := strings.SplitN(token, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("formato de parámetro inválido: %s", token)
		}
		key, value := strings.ToLower(parts[0]), parts[1]

		// Switch para manejar diferentes parámetros
		switch key {
		case "-size":
			// Convierte el valor del tamaño a un entero
			size, err := strconv.Atoi(value)
			if err != nil || size <= 0 {
				return nil, errors.New("el tamaño debe ser un número entero positivo")
			}
			cmd.size = size
		case "-unit":
			// Verifica que la unidad sea "K" o "M"
			if value != "K" && value != "M" {
				return nil, errors.New("la unidad debe ser K o M")
			}
			cmd.unit = value
		default:
			// Si el parámetro no es reconocido, devuelve un error
			return nil, fmt.Errorf("parámetro desconocido: %s", key)
		}
	}

	// Verifica que el parámetro -size haya sido proporcionado
	if cmd.size == 0 {
		return nil, errors.New("faltan parámetros requeridos: -size")
	}

	// Si no se proporcionó la unidad, se establece por defecto a "M"
	if cmd.unit == "" {
		cmd.unit = "M"
	}

	// Llama a la función CreateBinaryFile del paquete disk para crear el archivo binario
	err := structures.CreateBinaryFile(cmd.size, cmd.unit)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return cmd, nil // Devuelve el comando MKDISK creado
}
