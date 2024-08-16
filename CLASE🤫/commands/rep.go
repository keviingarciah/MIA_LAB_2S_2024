package commands

import (
	structures "CLASEEXTRA/structures" // Importa el paquete "structures" desde el directorio "EDD2021/structures"
	utils "CLASEEXTRA/utils"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// MKDISK estructura que representa el comando mkdisk con sus parámetros
type REP struct {
	path string // Ruta del archivo del disco
}

// CommandRep parsea el comando rep y devuelve una instancia de REP
func ParserRep(tokens []string) (*REP, error) {
	cmd := &REP{} // Crea una nueva instancia de REP

	// Unir tokens en una sola cadena y luego dividir por espacios, respetando las comillas
	args := strings.Join(tokens, " ")
	// Expresión regular para encontrar los parámetros del comando fdisk
	re := regexp.MustCompile(`-size=\d+|-unit=[kKmM]|-fit=[bBfF]{2}|-path="[^"]+"|-path=[^\s]+|-type=[pPeElL]|-name="[^"]+"|-name=[^\s]+`)
	// Encuentra todas las coincidencias de la expresión regular en la cadena de argumentos
	matches := re.FindAllString(args, -1)

	// Itera sobre cada coincidencia encontrada
	for _, match := range matches {
		// Divide cada parte en clave y valor usando "=" como delimitador
		kv := strings.SplitN(match, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("formato de parámetro inválido: %s", match)
		}
		key, value := strings.ToLower(kv[0]), kv[1]

		// Remove quotes from value if present
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = strings.Trim(value, "\"")
		}

		// Switch para manejar diferentes parámetros
		switch key {
		case "-path":
			// Verifica que el path no esté vacío
			if value == "" {
				return nil, errors.New("el path no puede estar vacío")
			}
			cmd.path = value
		default:
			// Si el parámetro no es reconocido, devuelve un error
			return nil, fmt.Errorf("parámetro desconocido: %s", key)
		}
	}

	// Verifica que el parámetro -path haya sido proporcionado
	if cmd.path == "" {
		return nil, errors.New("faltan parámetros requeridos: -path")
	}

	// Crear el disco con los parámetros proporcionados
	err := commandRep(cmd)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return cmd, nil // Devuelve el comando FDISK creado
}

func commandRep(cmd *REP) error {
	// Crear una nueva estructura MBR
	mbr := &structures.MBR{}

	// Deserializar la estructura MBR desde el archivo binario
	err := mbr.Deserialize(cmd.path)
	if err != nil {
		return err
	}

	// Imprimir la información del MBR
	fmt.Println("\nMBR:")
	mbr.Print()

	// Imprimir la información de cada partición
	fmt.Println("\nParticiones:")
	mbr.PrintPartitions()

	// Imprimir partidas montadas
	fmt.Println("\nParticiones montadas:")
	for id, path := range utils.GlobalMounts {
		fmt.Printf("ID: %s, PATH: %s\n", id, path)
	}

	return nil
}
