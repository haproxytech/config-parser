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

func (u *UnProcessed) Parse(line, wholeLine, previousLine string) (changeState string, err error) {
	u.unProcessed = append(u.unProcessed, "*"+wholeLine)
	return "", nil
}

func (u *UnProcessed) Valid() bool {
	if len(u.unProcessed) > 0 {
		return true
	}
	return false
}

func (u *UnProcessed) String() []string {
	return u.unProcessed
}
