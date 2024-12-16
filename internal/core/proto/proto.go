package proto

import "github.com/bedirhangull/protolex/internal/core/parser"

type Proto struct {
	parser *parser.Parser
}

func NewProto(p *parser.Parser) *Proto {
	return &Proto{parser: p}
}

// GetMessageBlocks returns all message blocks in the proto file
func (p *Proto) GetAllMessageBlocks() []string {
	return p.parser.GetAllMessageBlocks()
}

// GetEnumBlocks returns all enum blocks in the proto file
func (p *Proto) GetAllEnumBlocks() []string {
	return p.parser.GetAllEnumBlocks()
}

// GetServiceBlocks returns all service blocks in the proto file
func (p *Proto) GetAllServiceBlocks() []string {
	return p.parser.GetAllServiceBlocks()
}

// GetSyntax returns the syntax declaration
func (p *Proto) GetSyntax() string {
	return p.parser.GetSyntax()
}

// GetPackage returns the package declaration
func (p *Proto) GetPackage() string {
	return p.parser.GetPackage()
}
