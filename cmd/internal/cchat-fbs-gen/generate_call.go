package main

import (
	"fmt"

	"github.com/diamondburned/cchat/repository"
)

func (g *Generator) writeServiceMethod(tail *Generator, iface string, method repository.Method) {
	paramsName := tail.writeMethodParameterTable(iface, method)
	returnName := tail.writeMethodReturnTable(iface, method)

	name := method.UnderlyingName()

	fieldComment(&g.body, method.UnderlyingComment())
	g.body.WriteString(Indent)
	g.body.WriteString(name)
	g.body.WriteByte('(')
	g.body.WriteString(paramsName)
	g.body.WriteByte(')')

	if returnName != "" {
		g.body.WriteByte(':')
		g.body.WriteString(returnName)
	}

	g.body.WriteString(";\n")
}

var fieldRequired = fieldOpts{
	required: true,
}

func (g *Generator) writeMethodParameterTable(iface string, method repository.Method) string {
	var (
		parameters []repository.NamedType
		hasContext bool
	)

	switch method := method.(type) {
	case repository.GetterMethod:
		parameters = method.Parameters
	case repository.SetterMethod:
		parameters = method.Parameters
	case repository.IOMethod:
		parameters = method.Parameters
	case repository.ContainerMethod:
		parameters = []repository.NamedType{{
			Name: PascalToSnake(method.ContainerType),
			Type: method.ContainerType,
		}}
		hasContext = method.HasContext
	}

	methodName := method.UnderlyingName()
	name := methodName + "Parameters"

	g.WriteComment(repository.Comment{
		Raw: fmt.Sprintf(
			"%s is the parameter table for method %s of interface %s.",
			name, methodName, iface,
		),
	})

	g.body.WriteString("table ")
	g.body.WriteString(name)
	g.body.WriteString(" {\n")

	g.writeFieldDirect("call", namespaceType("call", "ID"), fieldRequired)

	if hasContext {
		g.writeFieldDirect("ctx", namespaceType("call", "Context"), fieldOpts{})
	}

	if len(parameters) > 0 {
		g.WriteLine()

		for _, param := range parameters {
			g.writeField(param, fieldOpts{})
		}
	}

	g.body.WriteByte('}')
	g.WriteLine()
	g.WriteLine()

	return name
}

func (g *Generator) writeMethodReturnTable(iface string, method repository.Method) string {
	var (
		returns    []repository.NamedType
		returnsErr bool
		hasStopFn  bool
	)

	switch method := method.(type) {
	case repository.GetterMethod:
		returnsErr = method.ReturnError()

	case repository.SetterMethod:
		return ""

	case repository.IOMethod:
		// we should always signal that an IO method is done regardless of its
		// types.
		returns = append(returns, method.ReturnValue)
		returnsErr = method.ReturnError()
		// TODO: custom error name attribute

	case repository.ContainerMethod:
		returnsErr = true
		hasStopFn = method.HasStopFn
	}

	methodName := method.UnderlyingName()
	name := methodName + "Returns"

	g.WriteComment(repository.Comment{
		Raw: fmt.Sprintf(
			"%s is the return table for method %s of interface %s.",
			name, methodName, iface,
		),
	})

	g.body.WriteString("table ")
	g.body.WriteString(name)
	g.body.WriteString(" {\n")

	g.writeFieldDirect("call", namespaceType("call", "ID"), fieldRequired)

	if len(returns) > 0 {
		g.WriteLine()

		for _, ret := range returns {
			if ret.Type != "" {
				g.writeField(ret, fieldOpts{})
			}
		}
	}

	if returnsErr {
		typ := namespaceType("core", "Error")
		g.writeFieldDirect("error", typ, fieldOpts{})
	}
	if hasStopFn {
		typ := namespaceType("call", "StopHandle")
		g.writeFieldDirect("stop_handle", typ, fieldRequired)
	}

	g.body.WriteByte('}')
	g.WriteLine()
	g.WriteLine()

	return name
}
