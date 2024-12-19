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

// GetServiceByName returns the service block by the given name
func (p *Proto) GetServiceByName(name string) string {
	return p.parser.GetServiceByName(name)
}

// GetRPCByServiceName returns the RPCs of the service by the given name
func (p *Proto) GetRPCByServiceName(name string, RPCname string) string {
	return p.parser.GetRPCByServiceName(name, RPCname)
}

// GetSyntax returns the syntax declaration
func (p *Proto) GetSyntax() string {
	return p.parser.GetSyntax()
}

// GetPackage returns the package declaration
func (p *Proto) GetPackageName() string {
	return p.parser.GetPackageName()
}

// GetMessageByName returns the message block by the given name
func (p *Proto) GetMessageByName(name string) string {
	return p.parser.GetMessageByName(name)
}

// GetAllMessageBlocksByRPCName returns all message blocks by the given RPC name
func (p *Proto) GetAllMessageBlocksByRPCName(service string, RPCname string) ([]string, error) {
	return p.parser.GetAllMessageBlocksByRPCName(service, RPCname)
}

// GetAllFieldByMessageName returns all fields of the message by the given name
func (p *Proto) GetAllFieldsByMessageName(messageName string) ([]string, error) {
	return p.parser.GetAllFieldsByMessageName(messageName)
}

// GetAllTypesByMessageName returns all types of the message by the given name
func (p *Proto) GetAllTypesByMessageName(messageName string) ([]string, error) {
	return p.parser.GetAllTypesByMessageName(messageName)
}

// GetAllReservedFieldsByMessageName returns all reserved fields of the message by the given name
func (p *Proto) GetAllReservedFieldsByMessageName(messageName string) ([]string, error) {
	return p.parser.GetAllReservedFieldsByMessageName(messageName)
}

// GetEnumByName returns the enum block by the given name
func (p *Proto) GetEnumByName(name string) string {
	return p.parser.GetEnumByName(name)
}

// GetAllEnumFieldsByEnumName returns all fields of the enum by the given name
func (p *Proto) GetAllEnumFieldsByEnumName(enumName string) ([]string, error) {
	return p.parser.GetAllEnumFieldsByEnumName(enumName)
}

// GetAllEnumsByMessageName returns all enums of the message by the given name
func (p *Proto) GetAllEnumsByMessageName(messageName string) ([]string, error) {
	return p.parser.GetAllEnumsByMessageName(messageName)
}

// GetAllPackages returns all packages in the proto file
func (p *Proto) GetAllPackages() ([]*parser.Package, error) {
	return p.parser.GetAllPackages()
}

// GetPackageByPath returns the package by the given path
func (p *Proto) GetPackageByPath(path string) (*parser.Package, error) {
	return p.parser.GetPackageByPath(path)
}

// GetChildPackagesByPath returns the child packages by the given path
func (p *Proto) GetChildPackagesByPath(parentPath string) ([]*parser.Package, error) {
	return p.parser.GetChildPackages(parentPath)
}

// GetParentPackageByPath returns the parent package by the given path
func (p *Proto) GetParentPackageByPath(childPath string) (*parser.Package, error) {
	return p.parser.GetParentPackage(childPath)
}
