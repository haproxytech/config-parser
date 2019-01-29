package parsers

import (
	"fmt"
	"strconv"

	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

type LogLines struct {
	data              []types.Log
	allowedLevels     map[string]bool
	allowedFacitlites map[string]bool
}

func (l *LogLines) Init() {
	l.data = []types.Log{}
	l.allowedFacitlites = map[string]bool{}
	l.allowedLevels = map[string]bool{}
	common.AddToBoolMap(l.allowedFacitlites,
		"kern", "user", "mail", "daemon", "auth", "syslog", "lpr", "news",
		"uucp", "cron", "auth2", "ftp", "ntp", "audit", "alert", "cron2",
		"local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7")
	common.AddToBoolMap(l.allowedLevels, "emerg", "alert", "crit", "err", "warning", "notice", "info", "debug")
}

func (l *LogLines) GetParserName() string {
	return "log"
}

func (l *LogLines) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(l.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return l.data, nil
}

func (l *LogLines) Set(data common.ParserData) error {
	if data == nil {
		l.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.Log:
		l.data = newValue
	case *types.Log:
		l.data = append(l.data, *newValue)
	case types.Log:
		l.data = append(l.data, newValue)
	default:
		return fmt.Errorf("casting error")
	}
	return nil
}

func (l *LogLines) SetStr(data string) error {
	parts, comment := common.StringSplitWithCommentIgnoreEmpty(data, ' ')
	oldData, _ := l.Get(false)
	l.Init()
	_, err := l.Parse(data, parts, []string{}, comment)
	if err != nil {
		l.Set(oldData)
	}
	return err
}

func (l *LogLines) parseLogLine(line string, parts []string, comment string) (*types.Log, error) {
	if len(parts) > 1 && parts[1] == "global" {
		return &types.Log{Global: true}, nil
	}
	log := &types.Log{
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
	if _, ok := l.allowedFacitlites[facility]; !ok {
		return log, &errors.ParseError{Parser: "Log", Line: line}
	}
	log.Facility = facility
	currIndex++
	//level is optional
	if currIndex >= len(parts) {
		return log, nil
	}
	level := parts[currIndex]
	if _, ok := l.allowedLevels[level]; !ok {
		return log, nil
	}
	log.Level = level
	currIndex++
	//min level is optional
	if currIndex >= len(parts) {
		return log, nil
	}
	level = parts[currIndex]
	if _, ok := l.allowedLevels[level]; !ok {
		return log, nil
	}
	log.MinLevel = level
	return log, nil
}

func (l *LogLines) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if parts[0] == "log" {
		if log, err := l.parseLogLine(line, parts, comment); err == nil {
			l.data = append(l.data, *log)
		}
		return "", nil
	}
	if parts[0] == "no" && parts[1] == "log" {
		l.data = append(l.data, types.Log{
			NoLog:   true,
			Comment: comment,
		})
		return "", nil
	}
	return "", &errors.ParseError{Parser: "LogLines", Line: line}
}

func (l *LogLines) Result(AddComments bool) ([]common.ReturnResultLine, error) {
	if len(l.data) == 0 {
		return nil, errors.FetchError
	}
	result := make([]common.ReturnResultLine, len(l.data))
	for index, log := range l.data {
		if log.Global {
			result[index] = common.ReturnResultLine{
				Data:    "log global",
				Comment: log.Comment,
			}
		} else if log.NoLog {
			result[index] = common.ReturnResultLine{
				Data:    "no log",
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
	return result, nil
}
