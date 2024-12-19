package parser

import (
	"fmt"
	"strings"
)

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

func (p *Parser) GetMessageByName(messageName string) string {
	messages := p.GetAllMessageBlocks()
	for _, m := range messages {
		if strings.Contains(m, messageName) {
			return m
		}
	}
	return ""
}

func (p *Parser) GetAllMessageBlocksByRPCName(service string, RPCname string) ([]string, error) {
	var messages []string
	rpc := p.GetRPCByServiceName(service, RPCname)
	if rpc == "" {
		return nil, fmt.Errorf("RPC not found: %s", RPCname)
	}

	requestMessageName := strings.TrimSpace(strings.Split(strings.Split(rpc, "(")[1], ")")[0])
	responseMessageName := strings.TrimSpace(strings.Split(strings.Split(rpc, "returns (")[1], ")")[0])

	requestMessageName = strings.TrimPrefix(requestMessageName, "stream ")
	responseMessageName = strings.TrimPrefix(responseMessageName, "stream ")

	requestMessage := p.GetMessageByName(requestMessageName)
	responseMessage := p.GetMessageByName(responseMessageName)

	if requestMessage != "" {
		messages = append(messages, requestMessage)
	}
	if responseMessage != "" {
		messages = append(messages, responseMessage)
	}

	return messages, nil
}
