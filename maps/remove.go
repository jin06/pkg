package maps

// Remove remove data's elements while key is in keys
func Remove[Key comparable](data map[Key]any, keys ...Key) {
	for _, k := range keys {
		delete(data, k)
	}
}
