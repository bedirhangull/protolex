package parser

import (
	"fmt"
	"regexp"
	"strings"
)

func (p *Parser) GetAllReservedFieldsByMessageName(messageName string) ([]string, error) {
	messageBlock := p.GetMessageByName(messageName)
	if messageBlock == "" {
		return nil, fmt.Errorf("message %s not found", messageName)
	}

	// Regular expression to match different types of reserved field declarations:
	// 1. Single numbers: reserved 15;
	// 2. Number ranges: reserved 2, 15, 9 to 11;
	// 3. Field names: reserved "foo", "bar";
	reservedRegex := regexp.MustCompile(`reserved\s+(?:(?:"[^"]+"|[0-9]+(?:\s+to\s+[0-9]+)?)+(?:\s*,\s*(?:"[^"]+"|[0-9]+(?:\s+to\s+[0-9]+)?))*)?\s*;`)

	matches := reservedRegex.FindAllString(messageBlock, -1)
	if matches == nil {
		return []string{}, nil
	}

	var reservedFields []string
	for _, match := range matches {
		cleaned := strings.TrimSpace(strings.TrimPrefix(match, "reserved"))
		cleaned = strings.TrimSuffix(cleaned, ";")

		// Handle single number case
		if strings.Count(cleaned, ",") == 0 && !strings.Contains(cleaned, "\"") && !strings.Contains(cleaned, "to") {
			reservedFields = append(reservedFields, cleaned)
			continue
		}

		parts := strings.Split(cleaned, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" {
				reservedFields = append(reservedFields, part)
			}
		}
	}

	return reservedFields, nil
}
