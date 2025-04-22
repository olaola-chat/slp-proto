package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	slp_proto "github.com/olaola-chat/slp-proto/protoc-gen-slp-rpc/proto"
)

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func AddImport(name2path map[string]string, path2name map[string]string, fullName string) {
	_, shortName := filepath.Split(fullName)
	if _, ok := path2name[fullName]; ok {
		// log.Println("Has")
		// import过这个package
		return
	}

	if _, ok := name2path[shortName]; !ok {
		// log.Println("None")
		// 名字未占用
		name2path[shortName] = fullName
		path2name[fullName] = shortName
		return
	}
	// log.Println("Dup")

	for i := 2; i < 10000; i++ {
		shortName2 := fmt.Sprintf("%s%d", shortName, i)
		if _, ok := name2path[shortName2]; ok {
			continue
		}
		// 名字未占用
		name2path[shortName2] = fullName
		name2path[fullName] = shortName2
		return
	}
}

func genServices(gen *protogen.Plugin, file *protogen.File, service *protogen.Service) {

	myOptions, ok := proto.GetExtension(service.Desc.Options(), slp_proto.E_RbpService).(*slp_proto.RbpServiceOption)
	if !ok || myOptions == nil {
		log.Println("No Option : ", service.Desc.FullName())
		return
	}
	// log.Println("Has Option : ", myOptions)
	log.Println("Service Name  : ", service.GoName, ", ", myOptions.GetName())

	path := filepath.Dir(file.GeneratedFilenamePrefix)
	g := gen.NewGeneratedFile(path+"/"+myOptions.GetName()+".go", file.GoImportPath)
	g.P("package " + file.GoPackageName)
	g.P()

	name2path := make(map[string]string)
	path2name := make(map[string]string)
	for _, m := range service.Methods {
		fullName := string(m.Input.GoIdent.GoImportPath)
		AddImport(name2path, path2name, fullName)

		fullName = string(m.Output.GoIdent.GoImportPath)
		AddImport(name2path, path2name, fullName)
	}

	var names []string
	for k := range path2name {
		// log.Println(k)
		names = append(names, k)
	}
	sort.Strings(names)
	g.P("import (")
	g.P("  \"context\"")
	g.P()
	g.P("  \"github.com/olaola-chat/slp-proto/rpcclient/base\"")
	g.P()
	for _, n := range names {
		a := path2name[n]
		_, s := filepath.Split(n)
		if a == s {
			g.P("  \"" + n + "\"")
		} else {
			g.P("  " + a + " \"" + n + "\"")
		}
	}
	g.P(")")

	varName := FirstUpper(service.GoName)
	typeName := FirstLower(service.GoName)
	g.P("// " + varName + " *" + typeName)
	g.P("var " + varName + " = &" + typeName + "{")
	g.P("  &base.Base{")
	g.P("    Name: \"" + myOptions.GetName() + "\",")
	g.P("  },")
	g.P("}")
	g.P()
	g.P("type " + typeName + " struct {")
	g.P("  *base.Base")
	g.P("}")
	g.P()

	for _, m := range service.Methods {
		g.P()
		inputPackName := path2name[string(m.Input.GoIdent.GoImportPath)] + "." + string(m.Input.GoIdent.GoName)
		outputPackName := path2name[string(m.Output.GoIdent.GoImportPath)] + "." + string(m.Output.GoIdent.GoName)
		g.P("func (s *" + typeName + ") " + m.GoName + "(ctx context.Context, req *" + inputPackName + ") (*" + outputPackName + ", error) {")
		g.P("  reply := &" + outputPackName + "{}")
		g.P("  err := s.Call(ctx, \"" + m.GoName + "\", req, reply)")
		g.P("  return reply, err")
		g.P("}")
		g.P()
		// log.Println(m.Input.GoIdent.GoImportPath)
	}
}

func main() {

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			for _, s := range f.Services {
				genServices(gen, f, s)
			}
			// generateFile(gen, f)
		}
		return nil
	})
}
