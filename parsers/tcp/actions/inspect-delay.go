package actions

import (
	"fmt"
	"strings"
)

type InspectDelay struct {
	Timeout string
	Comment string
}

func (f *InspectDelay) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}
	if len(parts) >= 3 {
		if len(parts) > 1 {
			f.Timeout = parts[2]
		} else {
			return fmt.Errorf("not enough params")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *InspectDelay) String() string {
	var result strings.Builder
	result.WriteString("inspect-delay ")

	result.WriteString(f.Timeout)

	if f.Comment != "" {
		result.WriteString(" # ")
		result.WriteString(f.Comment)
	}
	return result.String()
}

func (f *InspectDelay) GetComment() string {
	return f.Comment
}
