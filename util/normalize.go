package util

func NormalizeNoteInput(input map[string]interface{}) map[string]interface{} {
	if val, ok := input["note_body"]; ok {
		input["body"] = val

		delete(input, "note_body")
	}
	if val, ok := input["note_text"]; ok {
		input["body"] = val

		delete(input, "note_text")
	}

	return input
}
