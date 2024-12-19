package parser

import "strings"

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

func (p *Parser) GetServiceByName(serviceName string) string {
	service := p.GetAllServiceBlocks()
	for _, s := range service {
		if strings.Contains(s, serviceName) {
			return s
		}
	}
	return ""
}

func (p *Parser) GetRPCByServiceName(serviceName string, RPCname string) string {
	service := p.GetServiceByName(serviceName)
	if service == "" {
		return ""
	}

	var rpcServices = make(map[string]string)

	for _, line := range strings.Split(service, "\n") {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "rpc") {
			rpcName := strings.Split(trimmedLine, "(")[0]
			rpcName = strings.TrimSpace(strings.Split(rpcName, "rpc")[1])
			rpcServices[rpcName] = line
			continue
		}

	}

	for k := range rpcServices {
		if k == RPCname {
			return rpcServices[k]
		}
	}

	return ""
}
