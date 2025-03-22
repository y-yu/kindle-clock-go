package domain

type MapUnmarshal interface {
	MapUnmarshal() ([]byte, error)
}
