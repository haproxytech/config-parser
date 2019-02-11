package types

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
