package types

//Enabled is used by parsers Daemon, MasterWorker
//gen:Daemon
//name:daemon
//create-type:bool
//test:ok:daemon
//test:ok:daemon # comment
//gen:MasterWorker
//name:master-worker
//create-type:bool
//test:ok:master-worker
//test:ok:master-worker # comment
type Enabled struct {
	Comment string
}

//Int64 is used by parsers MaxConn, NbProc
//gen:MaxConn
//name:maxconn
//test:ok:maxconn 10000
//test:ok:maxconn 10000 # comment
//test:fail:maxconn
//gen:NbProc
//name:nbproc
//test:ok:nbproc 4
//test:ok:nbproc 4 # comment
//test:fail:nbproc
type Int64C struct {
	Value   int64
	Comment string
}

//String is used by parsers Mode, DefaultBackend, SimpleTimeTwoWords,
//gen:Mode
//name:mode
//test:ok:mode tcp
//test:ok:mode http
//test:ok:mode health
//test:ok:mode tcp # comment
//test:fail:mode
type StringC struct {
	Value   string
	Comment string
}

//StringSliceC is used by parsers stats.Timeout (this may change)
//gen:StatsTimeout
//name:stats timeout
//test:ok:stats timeout 4
//test:ok:stats timeout 4 # comment
//test:fail:stats timeout
//test:fail:stats
//test:fail:timeout
type StringSliceC struct {
	Value   []string
	Comment string
}

//Filters are not here, see parsers/filters
//==============================================================================
