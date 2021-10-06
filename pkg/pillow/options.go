package pillow

func mergeOptions(options ...map[string]interface{}) (merged map[string]interface{}) {
	if len(options) == 0 {
		return
	}

	merged = make(map[string]interface{})

	for _, options := range options {
		for k, v := range options {
			merged[k] = v
		}
	}

	if len(merged) == 0 {
		return
	}

	return
}
