package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/diamondburned/cchat/repository"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var destDir string
	if len(os.Args) > 1 {
		destDir = os.Args[1]
	}

	gen := NewGenerator()
	gen.RegisterTypeAlias("io.Reader", namespaceType("core", "Reader"))
	gen.RegisterTypeAlias("time.Time", namespaceType("core", "Time"))
	gen.RegisterTypeAlias("time.Duration", "int64") // nanoseconds
	gen.RegisterTypeAlias("context.Context", namespaceType("call", "Context"))
	gen.RegisterInclude(joinNamespace("core"), "core.fbs")
	gen.RegisterInclude(joinNamespace("call"), "call.fbs")

	for pkgPath := range repository.Main {
		base := path.Base(pkgPath)
		dest := base + ".fbs"

		gen.RegisterInclude(joinNamespace(base), dest)
	}

	for pkgPath, pkg := range repository.Main {
		base := path.Base(pkgPath)
		dest := filepath.Join(destDir, base+".fbs")

		for _, alias := range pkg.TypeAliases {
			gen.RegisterTypeAlias(alias.Name, alias.Type)
		}

		gen.WriteHeaderComment(pkg.Comment)
		gen.WriteNamespace(joinNamespace(base))
		gen.WriteFileID(base)

		for _, enum := range pkg.Enums {
			gen.WriteEnum(enum)
		}
		for _, sstruct := range pkg.Structs {
			gen.WriteStruct(sstruct)
		}
		for _, estruct := range pkg.ErrorStructs {
			gen.WriteStruct(estruct.Struct)
		}
		for _, iface := range pkg.Interfaces {
			gen.WriteInterface(pkg, iface)
		}

		if err := ioutil.WriteFile(dest, gen.Finish(), 0644); err != nil {
			log.Fatalln("Failed to write file:", err)
		}
	}
}
