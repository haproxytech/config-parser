package spoe

import (
	parser "github.com/haproxytech/config-parser/v2"
	configparser "github.com/haproxytech/config-parser/v2/parsers"
	"github.com/haproxytech/config-parser/v2/parsers/extra"
	"github.com/haproxytech/config-parser/v2/parsers/simple"
	"github.com/haproxytech/config-parser/v2/spoe/parsers"
)

func getStartParser() *parser.Parsers {
	return createParsers(map[string]parser.ParserInterface{
		"configversion": &extra.ConfigVersion{},
		"comments":      &extra.Comments{},
	})
}

func createParsers(psrs map[string]parser.ParserInterface) *parser.Parsers {
	m := make(map[string]parser.ParserInterface)
	for k, v := range psrs {
		m[k] = v
	}

	m["spoe-agent"] = &parsers.SPOESection{Name: "spoe-agent"}
	m["spoe-group"] = &parsers.SPOESection{Name: "spoe-group"}
	m["spoe-message"] = &parsers.SPOESection{Name: "spoe-message"}

	p := parser.Parsers{
		Parsers: m,
	}

	for _, psr := range p.Parsers {
		psr.Init()
	}
	return &p
}

func getSPOEAgentParser() *parser.Parsers {
	return createParsers(map[string]parser.ParserInterface{
		"groups":                  &simple.String{Name: "groups"},
		"log":                     &configparser.Log{},
		"maxconnrate":             &simple.Number{Name: "maxconnrate"},
		"maxerrrate":              &simple.Number{Name: "maxerrrate"},
		"max-frame-size":          &simple.Number{Name: "max-frame-size"},
		"max-waiting-frames":      &simple.Number{Name: "max-waiting-frames"},
		"messages":                &simple.String{Name: "messages"},
		"async":                   &simple.Option{Name: "async"},
		"continue-on-error":       &simple.Option{Name: "continue-on-error"},
		"dontlog-normal":          &simple.Option{Name: "dontlog-normal"},
		"force-set-var":           &simple.Option{Name: "force-set-var"},
		"pipelining":              &simple.Option{Name: "pipelining"},
		"send-frag-payload":       &simple.Option{Name: "send-frag-payload"},
		"option-set-on-error":     &simple.TimeTwoWords{Keywords: []string{"option", "set-on-error"}},
		"option-set-process-time": &simple.TimeTwoWords{Keywords: []string{"option", "set-process-time"}},
		"option-set-total-time":   &simple.TimeTwoWords{Keywords: []string{"option", "set-total-time"}},
		"option-var-prefix":       &simple.TimeTwoWords{Keywords: []string{"option", "var-prefix"}},
		"register-var-names":      &simple.String{Name: "register-var-names"},
		"timeout-hello":           &simple.TimeTwoWords{Keywords: []string{"timeout", "hello"}},
		"timeout-idle":            &simple.TimeTwoWords{Keywords: []string{"timeout", "idle"}},
		"timeout-processing":      &simple.TimeTwoWords{Keywords: []string{"timeout", "processing"}},
		"use-backend":             &simple.Word{Name: "use-backend"},
	})
}
func getSPOEGroupParser() *parser.Parsers {
	return createParsers(map[string]parser.ParserInterface{
		"messages": &simple.String{Name: "messages"},
	})
}
func getSPOEMessageParser() *parser.Parsers {
	return createParsers(map[string]parser.ParserInterface{
		"acl":   &configparser.ACL{},
		"args":  &simple.String{Name: "args"},
		"event": &parsers.Event{},
	})
}
