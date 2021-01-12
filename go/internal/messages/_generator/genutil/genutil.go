package genutil

import (
	"strings"
	"unicode"

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
