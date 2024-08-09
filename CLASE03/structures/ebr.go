package structures

type EBR struct {
	Ebr_mount [1]byte
	Ebr_fit   [1]byte
	Ebr_start int32
	Ebr_size  int32
	Ebr_next  int32
	Ebr_name  [16]byte
}
