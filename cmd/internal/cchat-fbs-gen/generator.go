package main

import (
	"bytes"
	"log"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/diamondburned/cchat/repository"
)

const (
	Indent = "  "
)

func init() {
	repository.TabWidth = 2 // per FlatBuffers style guide
}

type Generator struct {
	usedNames map[string]struct{}
	header    bytes.Buffer
	body      bytes.Buffer

	typeAliases map[string]string
	namespaces  map[string]string
}

func NewGenerator() *Generator {
	return &Generator{
		usedNames:   map[string]struct{}{},
		typeAliases: map[string]string{},
		namespaces:  map[string]string{},
	}
}

// clone returns a sub-generator sharing everything but the body.
func (g *Generator) clone() *Generator {
	cpy := *g
	cpy.body = bytes.Buffer{}
	return &cpy
}

func (g *Generator) takeFrom(parent *Generator) {
	g.body.Write(parent.body.Bytes())
}

func (g *Generator) reset() {
	g.usedNames = map[string]struct{}{}
	g.header.Reset()
	g.body.Reset()
}

type customAttribute struct {
	name    string
	comment string
}

var customAttributes = []customAttribute{
	{"service", "service is name of the RPC service name that the table implements, if any."},
}

// Finish renders everything generated and returns the results after resetting
// the generator.
func (g *Generator) Finish() []byte {
	var paths = make([]string, 0, len(g.usedNames))
	for name := range g.usedNames {
		paths = append(paths, g.namespaces[name])
	}

	sort.Strings(paths)

	for _, path := range paths {
		g.header.WriteString("include ")
		g.header.WriteString(strconv.Quote(path))
		g.header.WriteByte(';')
		g.header.WriteByte('\n')
	}

	g.header.WriteByte('\n')

	for _, attr := range customAttributes {
		g.WriteHeaderComment(repository.Comment{Raw: attr.comment})
		g.header.WriteString("attribute ")
		g.header.WriteString(strconv.Quote(attr.name))
		g.header.WriteByte(';')
		g.header.WriteByte('\n')
	}

	g.header.WriteByte('\n')

	output := bytes.Buffer{}
	output.Grow(g.header.Len() + g.body.Len())
	output.Write(g.header.Bytes())
	output.Write(bytes.TrimSuffix(g.body.Bytes(), []byte{'\n'}))

	g.reset()
	return output.Bytes()
}

// RegisterInclude registers an include, connecting from the namespace to the
// path.
func (g *Generator) RegisterInclude(namespace, path string) {
	g.namespaces[namespace] = path
}

// RegisterTypeAlias registers the given actual type to an alias. It pertains
// the registry even after a Reset call.
func (g *Generator) RegisterTypeAlias(actual, alias string) {
	g.typeAliases[actual] = alias
}

func (g *Generator) typeAlias(typeName string) string {
	// Some types we already have in core.
	switch typeName {
	case "map[string]string":
		return namespaceType("core", "StringPair")
	}

	// Find the alias.
	if actual, ok := g.typeAliases[typeName]; ok {
		typeName = actual
	}

	// Record used namespace.
	var namespace string

	pkgPath, realType := repository.TypeQual(typeName)
	if pkgPath != "" {
		if strings.Contains(typeName, "/") {
			base := path.Base(pkgPath)
			namespace = joinNamespace(base)
		} else {
			namespace = dirNamespace(typeName)
		}
	}

	if namespace != "" {
		if _, ok := g.namespaces[namespace]; !ok {
			log.Panicln("illegal namespace", namespace)
		}

		g.usedNames[namespace] = struct{}{}
		realType = namespace + "." + realType
	}

	return realType
}

// WriteNamespace writes the namespace for the current file.
func (g *Generator) WriteNamespace(namespace string) {
	g.header.WriteString("namespace ")
	g.header.WriteString(namespace)
	g.header.WriteByte(';')
	g.header.WriteByte('\n')
	g.header.WriteByte('\n')
}

// WriteHeaderComment writes a comment on the header.
func (g *Generator) WriteHeaderComment(comment repository.Comment) {
	g.header.WriteString(comment.GoString(1))
	g.header.WriteByte('\n')
}

// WriteFileID writes a file identifier on the header.
func (g *Generator) WriteFileID(fileID string) {
	if len(fileID) > 4 {
		fileID = fileID[:4]
	}
	fileID = strings.ToUpper(fileID)

	g.header.WriteString("file_identifier ")
	g.header.WriteString(strconv.Quote(fileID))
	g.header.WriteByte(';')
	g.header.WriteByte('\n')
	g.header.WriteByte('\n')
}

func (g *Generator) WriteLine() { g.body.WriteByte('\n') }

// WriteComment writes a comment.
func (g *Generator) WriteComment(comment repository.Comment) {
	g.body.WriteString(comment.GoString(1))
	g.WriteLine()
}

// WriteStruct writes a struct.
func (g *Generator) WriteStruct(rstruct repository.Struct) {
	g.WriteComment(rstruct.Comment)
	g.writeStruct(rstruct.Name, rstruct.Fields)
}

func (g *Generator) writeStruct(name string, fields []repository.StructField) {
	g.body.WriteString("struct ")
	g.body.WriteString(name)
	g.body.WriteString(" {\n")

	for _, field := range fields {
		fieldComment(&g.body, field.Comment)
		g.writeField(field.NamedType, fieldOpts{})
	}

	g.body.WriteByte('}')
	g.WriteLine()
	g.WriteLine()
}

// WriteEnum writes an enum block.
func (g *Generator) WriteEnum(enum repository.Enumeration) {
	g.WriteComment(enum.Comment)
	g.body.WriteString("enum ")
	g.body.WriteString(enum.Name)

	if enum.Bitwise {
		g.body.WriteString(" bit_flags")
	}

	g.body.WriteString(": ")
	g.body.WriteString(enum.GoType())
	g.body.WriteString(" {\n")

	for _, v := range enum.Values {
		fieldComment(&g.body, v.Comment)
		g.body.WriteString(Indent)
		g.body.WriteString(v.Name)
		g.body.WriteString(",\n")
	}

	g.body.WriteByte('}')
	g.WriteLine()
	g.WriteLine()
}
