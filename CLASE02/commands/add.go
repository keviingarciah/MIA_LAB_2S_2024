package commands

import (
	structures "CLASE02/structures" // Importa el paquete "structures" con alias "student"
	"errors"                        // Importa el paquete "errors" para manejar errores
	"fmt"                           // Importa el paquete "fmt" para formateo de entrada/salida
	"strconv"                       // Importa el paquete "strconv" para conversiones de string a int
	"strings"                       // Importa el paquete "strings" para manipulación de strings
)

// Estructura ADD que contiene los parámetros del comando
type ADD struct {
	Carnet string
	CUI    string
	Name   string
	Age    int
}

/*
   Ejemplos de comandos:
   add -carnet=2017000000 -cui=2017000000001 -name="Alejandro del Valle Rodri" -age=21
   add -carnet=2018000000 -cui=2018000000001 -name=Mariano -age=20
*/

// ParseAdd analiza los tokens del comando y devuelve una estructura ADD
func ParseAdd(tokens []string) (*ADD, error) {
	cmd := &ADD{} // Inicializa una nueva estructura ADD
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		key, value, err := parseToken(token) // Divide el token en clave y valor
		if err != nil {
			return nil, err
		}

		// Asigna los valores a la estructura ADD según la clave
		switch key {
		case "-carnet":
			cmd.Carnet = value
		case "-cui":
			cmd.CUI = value
		case "-name":
			value, i, err = handleQuotedName(value, tokens, i) // Maneja nombres entre comillas
			if err != nil {
				return nil, err
			}
			cmd.Name = value
		case "-age":
			age, err := strconv.Atoi(value) // Convierte el valor de edad a entero
			if err != nil || age <= 0 {
				return nil, errors.New("la edad debe ser un número entero positivo")
			}
			cmd.Age = age
		default:
			return nil, fmt.Errorf("parámetro desconocido: %s", key)
		}
	}

	// Valida que todos los parámetros requeridos estén presentes
	if err := validateRequiredParameters(cmd); err != nil {
		return nil, err
	}

	// Crea y escribe el estudiante en el archivo
	if err := createAndWriteStudent(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// parseToken divide un token en clave y valor
func parseToken(token string) (string, string, error) {
	parts := strings.SplitN(token, "=", 2) // Divide el token en dos partes usando "="
	if len(parts) != 2 {
		return "", "", fmt.Errorf("formato de parámetro inválido: %s", token)
	}
	return strings.ToLower(parts[0]), parts[1], nil
}

// handleQuotedName maneja nombres entre comillas
func handleQuotedName(value string, tokens []string, i int) (string, int, error) {
	if strings.HasPrefix(value, "\"") { // Verifica si el valor empieza con comillas
		for !strings.HasSuffix(value, "\"") && i+1 < len(tokens) {
			i++
			value += " " + tokens[i] // Agrega los tokens siguientes hasta encontrar el cierre de comillas
		}
		if strings.HasSuffix(value, "\"") {
			value = value[1 : len(value)-1] // Remueve las comillas del inicio y final
		} else {
			return "", i, errors.New("comillas no emparejadas en el parámetro de nombre")
		}
	}
	return value, i, nil
}

// validateRequiredParameters valida que todos los parámetros requeridos estén presentes
func validateRequiredParameters(cmd *ADD) error {
	if cmd.Carnet == "" || cmd.CUI == "" || cmd.Name == "" || cmd.Age == 0 {
		return errors.New("faltan parámetros requeridos")
	}
	return nil
}

// createAndWriteStudent crea un estudiante y lo escribe en un archivo
func createAndWriteStudent(cmd *ADD) error {
	student := &structures.Student{}
	copy(student.Carnet[:], cmd.Carnet)         // Copia el valor de Carnet a la estructura Student
	copy(student.CUI[:], cmd.CUI)               // Copia el valor de CUI a la estructura Student
	copy(student.Name[:], cmd.Name)             // Copia el valor de Name a la estructura Student
	copy(student.Age[:], strconv.Itoa(cmd.Age)) // Copia el valor de Age a la estructura Student

	// Escribe la estructura Student en el archivo
	if err := student.WriteToFile(); err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return err
	}

	fmt.Println("Estudiante escrito en el archivo exitosamente")
	return nil
}
