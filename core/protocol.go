package core

//Protocol publish model service protocol
//go:generate stringer -type=Protocol
type Protocol int32

const (
	_ Protocol = iota
	//SeldonCoreV1 Seldon Core V1 protocol
	SeldonCoreV1
	//SeldonCoreV1Alpha1 Seldon Core v1alpha1 protocol
	SeldonCoreV1Alpha1
	//SeldonCoreV1Alpha2 Seldon Core v1aplha2 protocol
	SeldonCoreV1Alpha2
)

var protocols = map[string]int32{
	SeldonCoreV1.String():       int32(SeldonCoreV1),
	SeldonCoreV1Alpha1.String(): int32(SeldonCoreV1Alpha1),
	SeldonCoreV1Alpha2.String(): int32(SeldonCoreV1Alpha2),
}

//GetProtocols ///
func GetProtocols() map[string]int32 {
	return protocols
}
