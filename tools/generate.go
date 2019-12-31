package main

import (
	"fmt"
	"os"

	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
)

func main() {
	// TODO: check for arguments.
	pkg := "github.com/mattermost/mattermost-server/v5/plugin"
	iface := "Hooks"

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pkgAst, err := getPackageAST(pkg, cwd)
	if err != nil {
		panic(err)
	}

	ifaceDecl, err := findInterfaceDeclaration(pkgAst, iface)
	if err != nil {
		panic(err)
	}

	methods, err := getInterfaceMethods(ifaceDecl)
	if err != nil {
		panic(err)
	}

	for _, m := range methods {
		fmt.Println(m.Name)
	}
}

func getPackageAST(pkgId string, dir string) (*ast.File, error) {
	p, err := build.Default.Import(pkgId, dir, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to import package %q: %v", pkgId, err)
	}

	fset := token.NewFileSet()

	// Parse package directory.
	pkgs, err := parser.ParseDir(fset, p.Dir, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to parse package directory %qv", p.Dir, err)
	}

	// Get package.
	pkgAst, ok := pkgs[p.Name]
	if !ok {
		return nil, fmt.Errorf("AST for package %q not loaded", pkgId)
	}

	// Combined AST of all package files.
	return ast.MergePackageFiles(pkgAst, 0), nil
}

func findInterfaceDeclaration(f *ast.File, interfaceName string) (*ast.InterfaceType, error) {
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			continue
		}
		for _, spec := range decl.Specs {
			spec, ok := spec.(*ast.TypeSpec)
			if !ok || spec.Name.Name != interfaceName {
				continue
			}
			specType, ok := spec.Type.(*ast.InterfaceType)
			if !ok {
				return nil, fmt.Errorf("%q is not an interface", interfaceName)
			}
			return specType, nil
		}
	}
	return nil, fmt.Errorf("declaration of interface %q not found", interfaceName)
}

func getInterfaceMethods(i *ast.InterfaceType) ([]method, error) {
	methods := make([]method, 0, len(i.Methods.List))
	for _, fn := range i.Methods.List {
		m := method{
			Name: fn.Names[0].Name,
		}
		fnType := fn.Type.(*ast.FuncType)
		params, err := getMethodParams(methodType)
		if err != nil {
			return
		}
		m.Params = params
		methods = append(methods, m)
	}
	return methods, nil
}

func getMethodParams(m *ast.FuncType) ([]arg, error) {
	args := make([]arg, 0, len(m.Params.List))

	for _, arg := range m.Params.List {
		t := paramType(arg.Type)
	}
	return args, nil
}

func paramType(e ast.Expr) string {
	ast.Inspect(e, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.Ident:
			if n.IsExported() {
				n.Name = p.Package.Name + "." + n.Name
			}
		case *ast.SelectorExpr:
			return false
		}
		return true
	})
	return p.gofmt(e)
}

func parseInterface(pkg string, iface string) error {
	// TODO: refactor
	// getPackageAST -> getInterfaceDeclaration

	// Find interface declaration.
	/*
		for _, decl := range f.Decls {
			decl, ok := decl.(*ast.GenDecl)
			if !ok || decl.Tok != token.TYPE {
				continue
			}
			for _, spec := range decl.Specs {
				spec, ok := spec.(*ast.TypeSpec)
				if !ok || spec.Name.Name != iface {
					continue
				}
				specType, ok := spec.Type.(*ast.InterfaceType)
				if !ok {
					return fmt.Errorf("%q is not an interface", iface)
				}

				for _, fn := range specType.Methods.List {
					fmt.Printf("%s ", fn.Names[0].Name)
					fnType := fn.Type.(*ast.FuncType)
					for _, fnParam := range fnType.Params.List {
						fmt.Printf(" %v ", fnParam.Names)
					}
					fmt.Println("")
				}
			}
		}
	*/
	return nil
}

type method struct {
	Name    string
	Args    []arg
	Results []string
}

type arg struct {
	Name string
	Type string
}
