package pillow

func mergeOptions(options ...Options) (merged Options) {
	if len(options) == 0 {
		return
	}

	merged = make(Options)

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
