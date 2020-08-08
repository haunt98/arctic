package arctic

// Arctic is used to get key value
type Arctic interface {
	Get(key string) []byte
}
