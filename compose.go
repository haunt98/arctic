package arctic

func ComposeKey(prefix, key string) string {
	return prefix + "/" + key
}
