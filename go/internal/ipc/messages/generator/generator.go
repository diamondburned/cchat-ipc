package main

import (
	"github.com/dave/jennifer/jen"
	"github.com/diamondburned/cchat-ipc/go/internal/ipc/messages/generator/genutil"
	"github.com/diamondburned/cchat/repository"
)

type Generator struct {
	file   *jen.File
	pkg    repository.Package
	path   string
	prefix string
}

func (g *Generator) typeQual(typePath string) (string, string) {
	return repository.TypeQual(g.replTypePath(typePath))
}

var typePathAliases = map[string]string{
	"error":     IPCCore + ".Error",
	"time.Time": IPCCore + ".Time",
	"io.Reader": IPCCore + ".Reader",
}

func (g *Generator) replTypePath(typePath string) string {
	path, typ := repository.TypeQual(typePath)
	if pathIsKnown(path) {
		// If we know of this path, then we'll be generating code for it.
		return pathPrefix(path) + typ
	}

	switch typePath {
	case "error":
		return IPCCore + ".Error"
	case "time.Time":
		return IPCCore + ".Time"
	case "io.Reader":
		return IPCCore + ".Reader"
	}

	return genutil.ProcessType(typePath, func(typePath string) string {
		path, _ := repository.TypeQual(typePath)

		// If the current type belongs to another package whose types require
		// prefixing, then we'll have to prefix its children types as well.
		//
		// The `g.path == ""` check ensures the condition that the current type
		// is from another package, and the `g.prefix != ""` check ensures that
		// this package is not the root package.

		if g.prefix != "" && path == "" && genutil.IsExported(typePath) {
			return g.prefix + typePath
		}

		return typePath
	})
}

func (g *Generator) namedType(namedType repository.NamedType) repository.NamedType {
	var name = namedType.Name
	if name == "" {
		_, name = g.typeQual(namedType.Type)
	}

	return repository.NamedType{
		Name: name,
		Type: g.replTypePath(namedType.Type),
	}
}

// genTypes generates unimportant types.
func (g *Generator) genTypes() {
	for _, sstruct := range g.pkg.Structs {
		g.generateStruct(sstruct)
	}
	for _, estruct := range g.pkg.ErrorStructs {
		g.generateStruct(estruct.Struct)
	}
}

func (g *Generator) generateStruct(sstruct repository.Struct) {
	var fields []jen.Code

	for _, field := range sstruct.Fields {
		namedType := g.namedType(field.NamedType)

		f := jen.Id(namedType.Name).Qual(namedType.Qual())
		f.Tag(map[string]string{
			"msg": genutil.PascalToSnake(namedType.Name),
		})

		fields = append(fields, f)
	}

	g.file.Type().Id(g.prefix + sstruct.Name).Struct(fields...)
	g.file.Line()

	// Generate a converter func body.
	var originalFields = []jen.Code{}

	for _, field := range sstruct.Fields {
		namedType := g.namedType(field.NamedType)

		f := jen.Id(namedType.Name).Op(":")
		f.Id(genutil.ReceiverID(sstruct.Name)).Dot(namedType.Name)

		// TODO: array support
		if namedType.Type != field.Type {
			f.Dot("Convert").Params()
		}

		f.Op(",")

		originalFields = append(originalFields, f)
	}

	converter := g.file.Func()
	converter.Params(genutil.Receiver(g.prefix+sstruct.Name, false)).Id("Convert")
	converter.Params()
	converter.Qual(g.path, sstruct.Name)
	converter.Block(
		jen.Return(jen.Qual(g.path, sstruct.Name).Block(originalFields...)),
	)
}
