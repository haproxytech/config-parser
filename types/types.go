package types

import "github.com/haproxytech/config-parser/params"

//Enabled is used by parsers Daemon, MasterWorker
type Enabled struct {
	Comment string
}

//Int64 is used by parsers MaxConn, NbProc
type Int64C struct {
	Value   int64
	Comment string
}

//String is used by parsers Mode, DefaultBackend, SimpleTimeTwoWords,
//SimpleString
type StringC struct {
	Value   string
	Comment string
}

//StringSliceC is used by parsers stats.Timeout (this may change)
type StringSliceC struct {
	Value   []string
	Comment string
}

//Filters are not here, see parsers/filters
//==============================================================================

type Bind struct {
	Path    string //can be address:port or socket path
	Params  []params.BindOption
	Comment string
}

type Balance struct {
	Algorithm string
	Arguments []string
	Comment   string
}

type CpuMap struct {
	Name    string
	Value   string
	Comment string
}

type DefaultServer struct {
	Params  []params.ServerOption
	Comment string
}

type ErrorFile struct {
	Code    string
	File    string
	Comment string
}

type Group struct {
	Name    string
	Users   []string
	Comment string
}

type Log struct {
	Global   bool
	NoLog    bool
	Address  string
	Length   int64
	Facility string
	Level    string
	MinLevel string
	Comment  string
}

type Mailer struct {
	Name    string
	IP      string
	Port    int64
	Comment string
}

type OptionHttpchk struct {
	Method  string
	Uri     string
	Version string
	Comment string
}

type Peer struct {
	Name    string
	IP      string
	Port    int64
	Comment string
}

type Section struct {
	Name    string
	Comment string
}
type Server struct {
	Name    string
	Address string
	Params  []params.ServerOption
	Comment string
}

type SimpleOption struct {
	NoOption bool
	Comment  string
}

type SimpleTimeout struct {
	Value   string
	Comment string
}

type Socket struct {
	Path    string //can be address:port
	Params  []params.BindOption
	Comment string
}

type Nameserver struct {
	Name    string
	Address string
	Comment string
}

type UseBackend struct {
	Name          string
	Condition     string
	ConditionKind string
	Comment       string
}

type User struct {
	Name       string
	Password   string
	IsInsecure bool
	Groups     []string
	Comment    string
}
type UseServer struct {
	Name          string
	Condition     string
	ConditionType string
	Comment       string
}
