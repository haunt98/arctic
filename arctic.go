package arctic

type Arctic interface {
	Get(key string) []byte
}
