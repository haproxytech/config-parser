package spoe

import (
	parser "github.com/haproxytech/config-parser/v3"
	configparser "github.com/haproxytech/config-parser/v3/parsers"
	"github.com/haproxytech/config-parser/v3/parsers/extra"
	"github.com/haproxytech/config-parser/v3/parsers/simple"
	"github.com/haproxytech/config-parser/v3/spoe/parsers"
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
	return createParsers([]parser.ParserInterface{
		&simple.String{Name: "groups"},
		&configparser.Log{},
		&simple.Number{Name: "maxconnrate"},
		&simple.Number{Name: "maxerrrate"},
		&simple.Number{Name: "max-frame-size"},
		&simple.Number{Name: "max-waiting-frames"},
		&simple.String{Name: "messages"},
		&simple.Option{Name: "async"},
		&simple.Option{Name: "continue-on-error"},
		&simple.Option{Name: "dontlog-normal"},
		&simple.Option{Name: "force-set-var"},
		&simple.Option{Name: "pipelining"},
		&simple.Option{Name: "send-frag-payload"},
		&simple.TimeTwoWords{Keywords: []string{"option", "set-on-error"}},
		&simple.TimeTwoWords{Keywords: []string{"option", "set-process-time"}},
		&simple.TimeTwoWords{Keywords: []string{"option", "set-total-time"}},
		&simple.TimeTwoWords{Keywords: []string{"option", "var-prefix"}},
		&simple.String{Name: "register-var-names"},
		&simple.TimeTwoWords{Keywords: []string{"timeout", "hello"}},
		&simple.TimeTwoWords{Keywords: []string{"timeout", "idle"}},
		&simple.TimeTwoWords{Keywords: []string{"timeout", "processing"}},
		&simple.Word{Name: "use-backend"},
	})
}
func getSPOEGroupParser() *parser.Parsers {
	return createParsers([]parser.ParserInterface{
		&simple.String{Name: "messages"},
	})
}
func getSPOEMessageParser() *parser.Parsers {
	return createParsers([]parser.ParserInterface{
		&configparser.ACL{},
		&simple.String{Name: "args"},
		&parsers.Event{},
	})
}
