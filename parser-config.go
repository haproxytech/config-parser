package parser

type ConfiguredParsers struct {
	_      [0]int
	State  string
	Active ParserTypes

	Comments *ParserTypes
	Defaults *ParserTypes
	Global   *ParserTypes
	Frontend *ParserTypes
	Backend  *ParserTypes
	Listen   *ParserTypes
	Resolver *ParserTypes
	Userlist *ParserTypes
	Peers    *ParserTypes
	Mailers  *ParserTypes
	Cache    *ParserTypes
}
