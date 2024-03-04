package maps

func Filter[Key comparable, Val comparable](data map[Key]Val, filters ...Val) {
	for _, filter := range filters {
		if len(data) == 0 {
			return
		}
		for key := range data {
			if data[key] == filter {
				delete(data, key)
			}
		}
	}
}
