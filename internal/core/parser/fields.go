package parser

import (
	"fmt"
	"strings"
)

func (p *Parser) GetAllFieldsByMessageName(messageName string) ([]string, error) {
	var fields []string

	message := p.GetMessageByName(messageName)
	if message == "" {
		return nil, fmt.Errorf("message not found: %s", messageName)
	}

	for _, line := range strings.Split(message, "\n") {
		if isField(line) {
			repeat := isRepeated(line)
			if repeat {
				field := getFieldName(line, 2)
				fields = append(fields, field)
			} else {
				field := getFieldName(line, 1)
				fields = append(fields, field)
			}
		}
	}

	return fields, nil
}

func (p *Parser) GetAllTypesByMessageName(messageName string) ([]string, error) {
	var types []string

	message := p.GetMessageByName(messageName)
	if message == "" {
		return nil, fmt.Errorf("message not found: %s", messageName)
	}

	for _, line := range strings.Split(message, "\n") {
		if isField(line) {
			repeat := isRepeated(line)
			if repeat {
				fieldType := getFieldType(line, 1)
				types = append(types, fieldType)
			} else {
				fieldType := getFieldType(line, 0)
				types = append(types, fieldType)
			}
		}
	}

	return types, nil
}

func isRepeated(line string) bool {
	return strings.Contains(line, "repeated")
}

func isField(line string) bool {
	return strings.Contains(line, "=")
}

func getFieldName(line string, section int16) string {
	return strings.Fields(line)[section]
}

func getFieldType(line string, section int16) string {
	return strings.Fields(line)[section]
}
