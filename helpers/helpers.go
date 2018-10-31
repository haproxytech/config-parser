package helpers

func AddToBoolMap(data map[string]bool, items ...string) {
	for _, item := range items {
		data[item] = true
	}
}
