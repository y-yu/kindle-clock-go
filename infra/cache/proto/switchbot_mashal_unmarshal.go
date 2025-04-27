package proto

import "google.golang.org/protobuf/proto"

func (a *SwitchBotDevice) ProtoMarshal() ([]byte, error) {
	data, err := proto.Marshal(a)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *SwitchBotDevice) ProtoUnmarshal(data []byte) error {
	err := proto.Unmarshal(data, a)
	if err != nil {
		return err
	}
	return nil
}

func (a *SwitchBotDevicesDataModel) ProtoMarshal() ([]byte, error) {
	data, err := proto.Marshal(a)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *SwitchBotDevicesDataModel) ProtoUnmarshal(data []byte) error {
	err := proto.Unmarshal(data, a)
	if err != nil {
		return err
	}
	return nil
}
