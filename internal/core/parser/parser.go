package parser

import (
	"strings"

	"github.com/bedirhangull/protolex/internal/adapter/prescriptive"
)

type Parser struct {
	Prescriptive *prescriptive.Prescriptive
}

func (p *Parser) GetSyntax() string {
	for _, line := range p.Prescriptive.Lines {
		if strings.HasPrefix(strings.TrimSpace(line), "syntax") {
			return strings.TrimSpace(line)
		}
	}
	return ""
}

func (p *Parser) GetPackageName() string {
	for _, line := range p.Prescriptive.Lines {
		if strings.HasPrefix(strings.TrimSpace(line), "package") {
			return strings.TrimSpace(line)
		}
	}
	return ""
}
