package prescriptive

import (
	"bufio"
	"strings"
)

type Prescriptive struct {
	ProtoContent string
	Lines        []string
}

func NewPrescriptive(protoContent string) *Prescriptive {
	removedSpaces := strings.Split(protoContent, "\n")

	return &Prescriptive{
		ProtoContent: protoContent,
		Lines:        removedSpaces,
	}
}

func (p *Prescriptive) CleanContent() {
	var cleanedLines []string
	var inMultilineComment bool

	scanner := bufio.NewScanner(strings.NewReader(p.ProtoContent))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "/*") {
			inMultilineComment = true
			continue
		}
		if strings.Contains(line, "*/") {
			inMultilineComment = false
			continue
		}
		if inMultilineComment {
			continue
		}

		if idx := strings.Index(line, "//"); idx != -1 {
			line = line[:idx]
		}

		line = strings.TrimRightFunc(line, func(r rune) bool {
			return r == ' ' || r == '\t' || r == '\r'
		})

		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		cleanedLines = append(cleanedLines, line)
	}

	p.ProtoContent = strings.Join(cleanedLines, "\n")
	p.Lines = cleanedLines
}
