package parser

import (
	"fmt"
	"strings"
)

type Package struct {
	Name     string
	Path     string
	Parent   *Package
	Children []*Package
}

func (p *Parser) GetAllPackages() ([]*Package, error) {
	var packages []*Package

	packageMap := make(map[string]*Package)

	for _, line := range p.Prescriptive.Lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "//") || strings.HasPrefix(trimmedLine, "/*") {
			continue
		}

		// Check for package declarations
		if strings.HasPrefix(trimmedLine, "package") {
			pkgPath := strings.TrimSpace(strings.TrimPrefix(trimmedLine, "package"))
			pkgPath = strings.TrimSuffix(pkgPath, ";")

			parts := strings.Split(pkgPath, ".")

			var fullPath string
			var parent *Package

			for i, part := range parts {
				if i == 0 {
					fullPath = part
				} else {
					fullPath = fullPath + "." + part
				}

				if pkg, exists := packageMap[fullPath]; exists {
					parent = pkg
					continue
				}

				newPkg := &Package{
					Name:   part,
					Path:   fullPath,
					Parent: parent,
				}

				if parent != nil {
					parent.Children = append(parent.Children, newPkg)
				}

				if parent == nil {
					packages = append(packages, newPkg)
				}

				packageMap[fullPath] = newPkg
				parent = newPkg
			}
		}
	}

	return packages, nil
}

func (p *Parser) GetPackageByPath(path string) (*Package, error) {
	packages, err := p.GetAllPackages()
	if err != nil {
		return nil, err
	}

	var findPackage func([]*Package, string) *Package
	findPackage = func(pkgs []*Package, searchPath string) *Package {
		for _, pkg := range pkgs {
			if pkg.Path == searchPath {
				return pkg
			}
			if found := findPackage(pkg.Children, searchPath); found != nil {
				return found
			}
		}
		return nil
	}

	pkg := findPackage(packages, path)
	if pkg == nil {
		return nil, fmt.Errorf("package %s not found", path)
	}

	return pkg, nil
}

func (p *Parser) GetChildPackages(parentPath string) ([]*Package, error) {
	parent, err := p.GetPackageByPath(parentPath)
	if err != nil {
		return nil, err
	}
	return parent.Children, nil
}

func (p *Parser) GetParentPackage(childPath string) (*Package, error) {
	child, err := p.GetPackageByPath(childPath)
	if err != nil {
		return nil, err
	}
	if child.Parent == nil {
		return nil, fmt.Errorf("package %s has no parent", childPath)
	}
	return child.Parent, nil
}
