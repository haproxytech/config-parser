package extra

type UnProcessed struct {
	unProcessed []string
}

func (u *UnProcessed) Init() {
	u.unProcessed = []string{}
}

func (u *UnProcessed) GetParserName() string {
	return ""
}

func (u *UnProcessed) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	u.unProcessed = append(u.unProcessed, line)
	return "", nil
}

func (u *UnProcessed) Valid() bool {
	if len(u.unProcessed) > 0 {
		return true
	}
	return false
}

func (u *UnProcessed) Result(AddComments bool) []string {
	return u.unProcessed
}
