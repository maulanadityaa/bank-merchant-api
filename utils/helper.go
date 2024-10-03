package utils

func StringToPointer(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}
