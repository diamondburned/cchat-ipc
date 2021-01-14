package genutil

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dave/jennifer/jen"
	"github.com/diamondburned/cchat/repository"
)

// FrontendInterfaces filters from the given package a list of interfaces that
// belong to the frontend.
//
// Frontend interfaces should be treated asynchronously over the wire, that is,
// backends must not rely on the reply of these calls.
func FrontendInterfaces(pkg repository.Package) []repository.Interface {
	var ifaces = make([]repository.Interface, 0, len(pkg.Interfaces))

	for _, iface := range pkg.Interfaces {
		if iface.IsContainer() {
			ifaces = append(ifaces, iface)
		}
	}

	return ifaces
}

// PascalToSnake converts PascalCase to snake_case.
func PascalToSnake(pascal string) string {
	build := strings.Builder{}
	build.Grow(len(pascal))

	runes := []rune(pascal)
	for i, r := range runes {
		if shouldUnderscore(runes, i) {
			build.WriteByte('_')
		}

		build.WriteRune(unicode.ToLower(rune(r)))
	}

	return build.String()
}

func shouldUnderscore(runes []rune, i int) bool {
	if i == 0 || i == len(runes)-1 {
		return false
	}

	// The current rune is upper case.
	return unicode.IsUpper(runes[i]) &&
		// The rune before it must be lower OR the rune after it must be lower.
		// This is an edge case to cover mixed caps.
		(unicode.IsLower(runes[i-1]) || unicode.IsLower(runes[i+1]))
}

func peakLast(str string, i int) byte {
	if i < 0 || i >= len(str) {
		return 0
	}
	return str[i]
}

func firstRune(str string) rune {
	r, _ := utf8.DecodeRuneInString(str)
	return r
}

// IsExported returns true if the type is exported. It does not handle slice or
// map types and will return false.
func IsExported(typ string) bool {
	_, typ = repository.TypeQual(typ)
	return unicode.IsUpper(firstRune(typ))
}

var mapRegex = regexp.MustCompile(`map\[(\w+)\](\w+)`)

// ProcessType breaks a type down into multiple parts when it's a slice or map
// and calls transFn on each of those parts, then reassemble and return the
// newly-processed type. The type is passed as-is into the callback without
// running TypeQual.
func ProcessType(typ string, transFn func(string) string) string {
	switch {
	case strings.HasPrefix(typ, "[]"):
		underlying := transFn(typ[2:]) // trim the "[]"
		return "[]" + underlying

	case strings.HasPrefix(typ, "map["):
		matches := mapRegex.FindStringSubmatch(typ)
		if matches == nil {
			// Invalid map syntax; pretend it's a regular type.
			return transFn(typ)
		}

		matches[1] = transFn(matches[1])
		matches[2] = transFn(matches[2])

		return "map[" + matches[1] + "]" + matches[2]

	default:
		return transFn(typ)
	}
}

func ReceiverID(typ string) string {
	if typ == "error" {
		return "err"
	}

	_, name := repository.TypeQual(typ)
	return string(unicode.ToLower(firstRune(name)))
}

func Receiver(typ string, ptr bool) *jen.Statement {
	recv := jen.Id(ReceiverID(typ))
	if ptr {
		recv.Op("*")
	}

	recv.Qual(repository.TypeQual(typ))
	return recv
}
