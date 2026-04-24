package api

func existsInOperations(operation Operation, operations []Operation) bool {
	for _, candidate := range operations {
		if candidate == operation {
			return true
		}
	}

	return false
}
