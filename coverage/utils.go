package coverage

import "strings"

func GetHeaderValues(header []string) []string {
	s := make([]string, 0)
	for _, item := range header {
		s = append(s, strings.Split(item, ", ")...)
	}
	return s
}
