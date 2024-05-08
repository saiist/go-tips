package client

type Store interface {
	Get(key string) string
}
