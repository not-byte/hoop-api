package utils

func DerefString(s *string) string {
	if s != nil {
		return *s
	}
	return "nil"
}
