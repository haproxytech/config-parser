package spoe

import (
	parser "github.com/haproxytech/config-parser/v2"
	"github.com/haproxytech/config-parser/v2/parsers/extra"
	"github.com/haproxytech/config-parser/v2/spoe/parsers"
)

func getStartParser() *parser.Parsers {
	return createParsers([]parser.ParserInterface{
		&extra.ConfigVersion{},
		&extra.Comments{},
	})
}

func createParsers(psrs []parser.ParserInterface) *parser.Parsers {
	p := parser.Parsers{
		Parsers: append(psrs, []parser.ParserInterface{
			&parsers.SPOESection{Name: "spoe-agent"},
			&parsers.SPOESection{Name: "spoe-group"},
			&parsers.SPOESection{Name: "spoe-message"},
			&extra.UnProcessed{},
		}...),
	}
	for _, psr := range p.Parsers {
		psr.Init()
	}
	return &p
}

func getSPOEAgentParser() *parser.Parsers {
	return createParsers([]parser.ParserInterface{})
}
func getSPOEGroupParser() *parser.Parsers {
	return createParsers([]parser.ParserInterface{})
}
func getSPOEMessageParser() *parser.Parsers {
	return createParsers([]parser.ParserInterface{})
}
