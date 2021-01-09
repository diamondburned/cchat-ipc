package main

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"

	"github.com/diamondburned/cchat/repository"
)

type fieldOpts struct {
	required bool
	attribs  []string
}

func (g *Generator) writeField(named repository.NamedType, opts fieldOpts) {
	g.writeFieldDirect(named.Name, named.Type, opts)
}

func (g *Generator) writeFieldDirect(name, goType string, opts fieldOpts) {
	var comment string

	var typ = goType
	switch {
	case strings.HasPrefix(goType, "[]"):
		underlying := g.typeAlias(strings.TrimPrefix(typ, "[]"))
		typ = "[" + underlying + "]"

		if name == "" {
			name = underlying + "_array"
		}

	case strings.HasPrefix(goType, "map["):
		k, v := mapTypes(goType)

		// If we have a type in core, then prefer it over the Go type.
		if newTyp := g.typeAlias(typ); newTyp != typ {
			typ = newTyp
		} else {
			typ = "[ubyte]"
			comment = goType
			opts.attribs = append(opts.attribs, "flexbuffer")
		}

		if name == "" {
			name = "map_" + g.typeAlias(k) + "_" + g.typeAlias(v)
		}

	default:
		typ = g.typeAlias(typ)

		if name == "" {
			// use the type name as the name
			_, name = repository.TypeQual(goType)
		}
	}

	name = PascalToSnake(name)

	// Disallow forbidden names.
	switch name {
	case "error": // conflicts with type Error
		name = "err"
	}

	g.body.WriteString(Indent)
	g.body.WriteString(name)
	g.body.WriteString(":")

	g.body.WriteString(typ)

	if opts.required {
		opts.attribs = append(opts.attribs, "required")
	}

	if len(opts.attribs) > 0 {
		g.body.WriteByte(' ')
		g.body.WriteByte('(')
		g.body.WriteString(strings.Join(opts.attribs, ", "))
		g.body.WriteByte(')')
	}

	g.body.WriteByte(';')

	if comment != "" {
		g.body.WriteString(" // ")
		g.body.WriteString(comment)
	}

	g.WriteLine()
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

func fieldComment(buf *bytes.Buffer, comment repository.Comment) {
	if comment.Raw == "" {
		return
	}

	lines := strings.Split(comment.GoString(2), "\n")
	for i, line := range lines {
		lines[i] = Indent + line
	}

	buf.WriteString(strings.Join(lines, "\n"))
	buf.WriteByte('\n')
}

var mapRegex = regexp.MustCompile(`map\[(\w+)\](\w+)`)

func mapTypes(mapType string) (k, v string) {
	matches := mapRegex.FindStringSubmatch(mapType)
	if matches == nil {
		return "", ""
	}
	return matches[1], matches[2]
}
