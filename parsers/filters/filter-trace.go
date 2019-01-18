package filters

import (
	"fmt"
	"strings"
)

type Trace struct { //filter trace [name <name>] [random-parsing] [random-forwarding] [hexdump]
	Name             string
	RandomParsing    bool
	RandomForwarding bool
	Hexdump          bool
	Comment          string
}

func (f *Trace) Parse(parts []string, comment string) error {
	//we have filter trace [name <name>] [random-parsing] [random-forwarding] [hexdump]
	if comment != "" {
		f.Comment = comment
	}
	index := 2
	for index < len(parts) {
		switch parts[index] {
		case "name":
			index++
			f.Name = parts[index]
		case "random-parsing":
			f.RandomParsing = true
		case "random-forwarding":
			f.RandomForwarding = true
		case "hexdump":
			f.Hexdump = true
		}
		index++
	}
	return nil
}

func (f *Trace) String() string {
	var result strings.Builder
	result.WriteString("  filter trace")
	if f.Name != "" {
		result.WriteString(" ")
		result.WriteString(f.Name)
	}
	if f.RandomParsing {
		result.WriteString(" random-parsing")
	}
	if f.RandomForwarding {
		result.WriteString(" random-forwarding")
	}
	if f.Hexdump {
		result.WriteString(" hexdump")
	}
	if f.Comment != "" {
		result.WriteString(fmt.Sprintf(" # %s", f.Comment))
	}
	return result.String()
}
