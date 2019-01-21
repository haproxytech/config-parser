package http

type HTTPAction interface {
	Parse(parts []string, comment string) error
	String() string
}
