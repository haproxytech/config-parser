package tcp

type TCPAction interface {
	Parse(parts []string, comment string) error
	String() string
	GetComment() string
}
