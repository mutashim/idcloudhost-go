package idcloudhost

type VMParameter struct {
	OS        []VMOS
	Disks     VMRange
	RAM       VMRange
	Password  VMPassword
	PublicKey VMPublicKey
}

type VMOS struct {
	Name    string   `json:"os_name"`
	Version []string `json:"values"`
}

type VMRange struct {
	Min int64
	Max int64
}

type VMRAM struct {
}

type VMPassword struct {
	Regex string
}

type VMPublicKey struct {
	Desc  string
	Regex string
}
