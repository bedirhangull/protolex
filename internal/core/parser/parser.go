package parser

import (
	"strings"

	"github.com/bedirhangull/protolex/internal/adapter/prescriptive"
)

type Parser struct {
	Prescriptive *prescriptive.Prescriptive
}

func (p *Parser) GetAllServiceBlocks() []string {
	var services []string
	var currentService strings.Builder
	var inService bool

	for _, line := range p.Prescriptive.Lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "service") {
			inService = true
			currentService.WriteString(line + "\n")
			continue
		}

		if inService {
			currentService.WriteString(line + "\n")
			if trimmedLine == "}" {
				services = append(services, currentService.String())
				currentService.Reset()
				inService = false
			}
		}
	}

	return services
}

func (p *Parser) GetAllMessageBlocks() []string {
	var messages []string
	var currentMessage strings.Builder
	var inMessage bool

	for _, line := range p.Prescriptive.Lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "message") {
			inMessage = true
			currentMessage.WriteString(line + "\n")
			continue
		}

		if inMessage {
			currentMessage.WriteString(line + "\n")
			if trimmedLine == "}" {
				messages = append(messages, currentMessage.String())
				currentMessage.Reset()
				inMessage = false
			}
		}
	}

	return messages
}

func (p *Parser) GetAllEnumBlocks() []string {
	var enums []string
	var currentEnum strings.Builder
	var inEnum bool

	for _, line := range p.Prescriptive.Lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "enum") {
			inEnum = true
			currentEnum.WriteString(line + "\n")
			continue
		}

		if inEnum {
			currentEnum.WriteString(line + "\n")
			if trimmedLine == "}" {
				enums = append(enums, currentEnum.String())
				currentEnum.Reset()
				inEnum = false
			}
		}
	}

	return enums
}

func (p *Parser) GetSyntax() string {
	for _, line := range p.Prescriptive.Lines {
		if strings.HasPrefix(strings.TrimSpace(line), "syntax") {
			return strings.TrimSpace(line)
		}
	}
	return ""
}

func (p *Parser) GetPackage() string {
	for _, line := range p.Prescriptive.Lines {
		if strings.HasPrefix(strings.TrimSpace(line), "package") {
			return strings.TrimSpace(line)
		}
	}
	return ""
}
