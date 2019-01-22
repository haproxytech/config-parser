package parsers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
)

type Log struct {
	Global   bool
	Address  string
	Length   int64
	Facility string
	Level    string
	MinLevel string
	Comment  string
}

type LogLines struct {
	LogLines          []Log
	AllowedLevels     map[string]bool
	AllowedFacitlitys map[string]bool
}

func (l *LogLines) Init() {
	l.LogLines = []Log{}
	l.AllowedFacitlitys = map[string]bool{}
	l.AllowedLevels = map[string]bool{}
	common.AddToBoolMap(l.AllowedFacitlitys,
		"kern", "user", "mail", "daemon", "auth", "syslog", "lpr", "news",
		"uucp", "cron", "auth2", "ftp", "ntp", "audit", "alert", "cron2",
		"local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7")
	common.AddToBoolMap(l.AllowedLevels, "emerg", "alert", "crit", "err", "warning", "notice", "info", "debug")
}

func (l *LogLines) GetParserName() string {
	return "log"
}

func (l *LogLines) parseLogLine(line string, parts []string, comment string) (Log, error) {
	if len(parts) > 1 && parts[1] == "global" {
		return Log{Global: true}, nil
	}
	log := Log{
		Address: parts[1],
		Comment: comment,
	}
	//see if we have length
	currIndex := 2
	if num, err := strconv.ParseInt(parts[currIndex], 10, 64); err == nil {
		log.Length = num
		currIndex++
	}
	//we must have facility
	if currIndex >= len(parts) {
		return log, &errors.ParseError{Parser: "Log", Line: line}
	}
	facility := parts[currIndex]
	if _, ok := l.AllowedFacitlitys[facility]; !ok {
		return log, &errors.ParseError{Parser: "Log", Line: line}
	}
	log.Facility = facility
	currIndex++
	//level is optional
	if currIndex >= len(parts) {
		return log, nil
	}
	level := parts[currIndex]
	if _, ok := l.AllowedLevels[level]; !ok {
		return log, nil
	}
	log.Level = level
	currIndex++
	//min level is optional
	if currIndex >= len(parts) {
		return log, nil
	}
	level = parts[currIndex]
	if _, ok := l.AllowedLevels[level]; !ok {
		return log, nil
	}
	log.MinLevel = level
	return log, nil
}

func (l *LogLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "log" {
		if log, err := l.parseLogLine(line, parts, comment); err == nil {
			l.LogLines = append(l.LogLines, log)
		}
		return "", nil
	}
	return "", &errors.ParseError{Parser: "LogLines", Line: line}
}

func (l *LogLines) Valid() bool {
	if len(l.LogLines) > 0 {
		return true
	}
	return false
}

func (l *LogLines) Result(AddComments bool) []common.ReturnResultLine {
	result := make([]common.ReturnResultLine, len(l.LogLines))
	for index, log := range l.LogLines {
		if log.Global {
			result[index] = common.ReturnResultLine{
				Data:    "log global",
				Comment: log.Comment,
			}
		} else {
			line := fmt.Sprintf("log %s", log.Address)
			if log.Length > 0 {
				line = fmt.Sprintf("%s %d", line, log.Length)
			}
			line = fmt.Sprintf("%s %s", line, log.Facility)
			if log.Level != "" {
				line = fmt.Sprintf("%s %s", line, log.Level)
				if log.MinLevel != "" {
					line = fmt.Sprintf("%s %s", line, log.MinLevel)
				}
			}
			result[index] = common.ReturnResultLine{
				Data:    line,
				Comment: log.Comment,
			}
		}
	}
	return result
}
