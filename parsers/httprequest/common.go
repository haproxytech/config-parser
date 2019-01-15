package httprequest

//searches for "if" or "unless" and returns result
func splitRequest(parts []string) (command, condition []string) {
	if len(parts) == 0 {
		return []string{}, []string{}
	}
	index := 0
	found := false
	for index < len(parts) {
		switch parts[index] {
		case "if", "unless":
			found = true
			break
		}
		if found {
			break
		}
		index++
	}
	command = parts[:index]
	condition = parts[index:]
	return command, condition
}
