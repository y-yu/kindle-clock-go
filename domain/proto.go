package domain

type ProtoMarshalUnmarshal interface {
	ProtoMarshal() ([]byte, error)
	ProtoUnmarshal(data []byte) error
}
