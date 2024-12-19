package parser

import (
	"fmt"
	"strings"
)

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

func (p *Parser) GetEnumByName(name string) string {
	var currentEnum strings.Builder
	var inEnum bool

	for _, line := range p.Prescriptive.Lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "enum "+name) {
			inEnum = true
			currentEnum.WriteString(line + "\n")
			continue
		}

		if inEnum {
			currentEnum.WriteString(line + "\n")
			if trimmedLine == "}" {
				return currentEnum.String()
			}
		}
	}

	return ""
}

func (p *Parser) GetAllEnumFieldsByEnumName(enumName string) ([]string, error) {
	enumBlock := p.GetEnumByName(enumName)
	if enumBlock == "" {
		return nil, nil
	}

	var fields []string
	for _, line := range strings.Split(enumBlock, "\n") {
		if isEnumField(line) {
			fields = append(fields, getEnumFieldName(line))
		}
	}

	return fields, nil
}

func (p *Parser) GetAllEnumsByMessageName(messageName string) ([]string, error) {
	message := p.GetMessageByName(messageName)
	if message == "" {
		return nil, fmt.Errorf("message %s not found", messageName)
	}

	var enums []string
	seenEnums := make(map[string]bool)
	var currentEnum strings.Builder
	var inEnum bool

	lines := strings.Split(message, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "enum") {
			inEnum = true
			currentEnum.WriteString(line + "\n")
			continue
		}
		if inEnum {
			currentEnum.WriteString(line + "\n")
			if trimmedLine == "}" {
				enumStr := currentEnum.String()
				enumName := strings.Fields(strings.Split(strings.TrimSpace(enumStr), "{")[0])[1]
				if !seenEnums[enumName] {
					enums = append(enums, enumStr)
					seenEnums[enumName] = true
				}
				currentEnum.Reset()
				inEnum = false
			}
		}
	}

	// Find references to external enums in field declarations
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.Contains(trimmedLine, "=") && !strings.HasPrefix(trimmedLine, "enum") {
			fields := strings.Fields(trimmedLine)
			if len(fields) >= 2 {
				fieldType := fields[0]
				if fieldType == "repeated" && len(fields) >= 3 {
					fieldType = fields[1]
				}

				enumDef := p.GetEnumByName(fieldType)
				if enumDef != "" && !seenEnums[fieldType] {
					enums = append(enums, enumDef)
					seenEnums[fieldType] = true
				}
			}
		}
	}

	return enums, nil
}

func isEnumField(line string) bool {
	return strings.Contains(line, "=")
}

func getEnumFieldName(line string) string {
	return strings.Split(strings.TrimSpace(line), "=")[0]
}
