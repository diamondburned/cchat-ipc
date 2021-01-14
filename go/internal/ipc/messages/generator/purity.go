package main

import (
	"log"

	"github.com/diamondburned/cchat-ipc/go/internal/ipc/messages/generator/genutil"
	"github.com/diamondburned/cchat/repository"
)

type purityChecker struct {
	pkgs  repository.Packages
	cache map[string]bool
	fail  bool
}

// guaranteePurity guarantees that the entire repository does not violate the
// struct purity assumption, that is, all structs must be able to be prefetched
// recursively down to the bottom-most level.
func guaranteePurity(pkgs repository.Packages) {
	purityChecker := purityChecker{
		pkgs:  pkgs,
		cache: map[string]bool{},
	}

	purityChecker.do()
}

func (pc purityChecker) isPure(typePath string) bool {
	for _, typePath := range genutil.UnderlyingTypes(typePath) {
		if pure, ok := pc.cache[typePath]; ok && !pure {
			return false
		}

		pure := pc.isPureDirect(typePath)
		pc.cache[typePath] = pure

		if !pure {
			return false
		}
	}

	return true
}

func (pc purityChecker) isPureDirect(typePath string) bool {
	// Check if the type is handled specifically. If it is, then we
	// automatically treat it as pure.
	if _, ok := typePathAliases[typePath]; ok {
		return true
	}

	// Resolve the types.
	for _, typePath := range genutil.UnderlyingTypes(typePath) {
		path, typ := repository.TypeQual(typePath)
		if path == "" {
			// We can assume that built-in types are always pure; they
			// will always be resolved internally to other values.
			continue
		}

		// Try and see if this type is resolvable.
		refPkg, ok := pc.pkgs[path]
		if !ok {
			log.Panicln("Type", typePath, "is not resolvable: unknown package.")
		}

		pure, _ := pc.typeIsPure(refPkg.FindType(typ))
		if !pure {
			return false
		}
	}

	return true
}

func (pc purityChecker) structIsPure(sstruct repository.Struct) (pure bool, where string) {
	// pure by default
	pure = true

	for _, field := range sstruct.Fields {
		if !pc.isPure(field.Type) {
			pure = false
			where = field.Name
			break
		}
	}
	return
}

func (pc purityChecker) typeIsPure(typ interface{}) (pure bool, where string) {
	// pure by default
	pure = true

	switch v := typ.(type) {
	case *repository.Interface:
		// Check all embedded interfaces recursively.
		for _, ifaceName := range v.Embeds {
			if !pc.isPure(ifaceName.InterfaceName) {
				return false, ifaceName.InterfaceName
			}
		}

		// Check for pure methods.
		for _, method := range v.Methods {
			switch method := method.(type) {
			case repository.AsserterMethod:
				// pass
			case repository.GetterMethod:
				if len(method.Parameters) > 0 {
					return false, method.UnderlyingName()
				}
			default:
				return false, method.UnderlyingName()
			}
		}

	case *repository.Struct:
		pure, where = pc.structIsPure(*v)

	case *repository.ErrorStruct:
		pure, where = pc.structIsPure(v.Struct)
	}

	return
}

func (pc purityChecker) do() {
	// Reset fail state.
	pc.fail = false

	for pkgPath, pkg := range pc.pkgs {
		for _, sstruct := range pkg.Structs {
			pure, field := pc.structIsPure(sstruct)
			if !pure {
				pc.markFail(sstruct.Name, field, pkgPath)
			}
		}

		for _, estruct := range pkg.ErrorStructs {
			pure, field := pc.structIsPure(estruct.Struct)
			if !pure {
				pc.markFail(estruct.Name, field, pkgPath)
			}
		}
	}

	if pc.fail {
		log.Fatalln("Purity check failed.")
	}
}

func (pc *purityChecker) markFail(structName, fieldName, pkgPath string) {
	pc.fail = true
	log.Printf(
		"struct %s of package %s has impure field %s.\n",
		structName, pkgPath, fieldName,
	)
}
