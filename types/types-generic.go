/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

//Enabled is used by parsers Daemon, MasterWorker, ExternalCheck
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
//gen:ExternalCheck
//name:external-check
//create-type:bool
//test:ok:external-check
//test:ok:external-check # comment
type Enabled struct {
	Comment string
}

//Int64 is used by parsers MaxConn, NbProc, NbThread
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
//gen:NbThread
//name:nbthread
//test:ok:nbthread 4
//test:ok:nbthread 4 # comment
//test:fail:nbthread
type Int64C struct {
	Value   int64
	Comment string
}

//String is used by parsers Mode, DefaultBackend, SimpleTimeTwoWords, StatsTimeout
//gen:Mode
//name:mode
//test:ok:mode tcp
//test:ok:mode http
//test:ok:mode health
//test:ok:mode tcp # comment
//test:fail:mode
//gen:DefaultBackend
//name:default_backend
//test:ok:default_backend http
//test:fail:default_backend
//gen:StatsTimeout
//name:stats timeout
//test:ok:stats timeout 4
//test:ok:stats timeout 4 # comment
//test:fail:stats timeout
//test:fail:stats
//test:fail:timeout
type StringC struct {
	Value   string
	Comment string
}

//StringSliceC is used by simple-string-multiple
type StringSliceC struct {
	Value   []string
	Comment string
}

//Filters are not here, see parsers/filters
//==============================================================================
