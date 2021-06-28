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

package parser

import (
	"sync"
)

type Section string

const (
	Comments   Section = "#"
	Defaults   Section = "defaults"
	Global     Section = "global"
	Resolvers  Section = "resolvers"
	UserList   Section = "userlist"
	Peers      Section = "peers"
	Mailers    Section = "mailers"
	Frontends  Section = "frontend"
	Backends   Section = "backend"
	Listen     Section = "listen"
	Cache      Section = "cache"
	Program    Section = "program"
	HTTPErrors Section = "http-errors"
	Ring       Section = "ring"
	// spoe sections
	SPOEAgent   Section = "spoe-agent"
	SPOEGroup   Section = "spoe-group"
	SPOEMessage Section = "spoe-message"
)

const (
	CommentsSectionName = "data"
	GlobalSectionName   = "data"
	DefaultSectionName  = "data"
)

type UnlockError struct{}

func (e UnlockError) Error() string {
	return "failed funclock"
}

type Options struct {
	UseV2HTTPCheck bool
	UseMd5Hash     bool
}

// Parser reads and writes configuration on given file
type Parser struct {
	Parsers map[Section]map[string]*Parsers
	Options Options
	mutex   *sync.Mutex
}

func NewParserWithOptions(options Options) *Parser {
	return initParserMaps(&Parser{Options: options})
}

func NewParser() *Parser {
	return initParserMaps(&Parser{})
}
