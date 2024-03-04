package maps

func Remove[Key comparable](data map[Key]any, keys ...Key) {
	for _, k := range keys {
		delete(data, k)
	}
}
