package bmap

// Index returns the index of the first occurrence of v in s,
// or "" if not present.
func Index[E comparable](s map[string]E, v E) string {
	for key, vs := range s {
		if v == vs {
			return key
		}
	}
	return ""
}

// ContainKey map是否包含目标key
func ContainKey[E comparable](s map[string]E, key string) bool {
	if _, ok := s[key]; ok {
		return true
	}
	return false
}

// ContainValue map是否包含目标value
func ContainValue[E comparable](s map[string]E, v E) bool {
	return len(Index(s, v)) > 0
}
