package utils

import (
	"encoding/binary"
	"errors"
)

func Int32ToBytes(n int32) [4]byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(n))
	return buf
}

func Float64ToBytes(f float64) [4]byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(f))
	return buf
}

func ConvertToBytes(size int, unit string) (int, error) {
	switch unit {
	case "K":
		return size * 1024, nil // Convierte kilobytes a bytes
	case "M":
		return size * 1024 * 1024, nil // Convierte megabytes a bytes
	default:
		return 0, errors.New("invalid unit") // Devuelve un error si la unidad es inválida
	}
}
