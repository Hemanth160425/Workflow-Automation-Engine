package transform

func Run(
	config map[string]interface{},
	input map[string]interface{},
) (map[string]interface{}, error) {

	output := make(map[string]interface{})

	fields, ok := config["select"].([]interface{})
	if !ok {
		return output, nil
	}

	for _, f := range fields {
		key := f.(string)
		if val, exists := input[key]; exists {
			output[key] = val
		}
	}

	return output, nil
}
