package proto

import (
	"google.golang.org/protobuf/proto"
)

func (a *AwairDataModel) ProtoMarshal() ([]byte, error) {
	data, err := proto.Marshal(a)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AwairDataModel) ProtoUnmarshal(data []byte) error {
	err := proto.Unmarshal(data, a)
	if err != nil {
		return err
	}
	return nil
}
