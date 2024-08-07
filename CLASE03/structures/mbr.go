package structures

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type MBR struct {
	Mbr_size           int32   // Tamaño del MBR en bytes
	Mbr_creation_date  float32 // Fecha y hora de creación del MBR
	Mbr_disk_signature int32   // Firma del disco
	Mbr_disk_fit       [1]byte // Tipo de ajuste
	///mbr_partitions     [4]Partition // Particiones del MBR
}

// SerializeMBR escribe la estructura MBR al inicio de un archivo binario
func (mbr *MBR) SerializeMBR(path string) error {
	fmt.Println(mbr)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Serializar la estructura MBR directamente en el archivo
	err = binary.Write(file, binary.LittleEndian, mbr)
	if err != nil {
		return err
	}

	return nil
}

// DeserializeMBR lee la estructura MBR desde el inicio de un archivo binario
func (mbr *MBR) DeserializeMBR(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Obtener el tamaño de la estructura MBR
	mbrSize := binary.Size(mbr)
	if mbrSize <= 0 {
		return fmt.Errorf("invalid MBR size: %d", mbrSize)
	}

	// Leer solo la cantidad de bytes que corresponden al tamaño de la estructura MBR
	buffer := make([]byte, mbrSize)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	// Deserializar los bytes leídos en la estructura MBR
	reader := bytes.NewReader(buffer)
	err = binary.Read(reader, binary.LittleEndian, mbr)
	if err != nil {
		return err
	}

	return nil
}

func (mbr *MBR) Print() {
	// Convertir Mbr_creation_date a time.Time
	creationTime := time.Unix(int64(mbr.Mbr_creation_date), 0)

	// Convertir Mbr_disk_fit a char
	diskFit := rune(mbr.Mbr_disk_fit[0])

	fmt.Printf("MBR Size: %d\n", mbr.Mbr_size)
	fmt.Printf("Creation Date: %s\n", creationTime.Format(time.RFC3339))
	fmt.Printf("Disk Signature: %d\n", mbr.Mbr_disk_signature)
	fmt.Printf("Disk Fit: %c\n", diskFit)
}
