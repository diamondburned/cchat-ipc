package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/diamondburned/cchat/repository"
)

// WriteInterface writes an interface.
func (g *Generator) WriteInterface(pkg repository.Package, iface repository.Interface) {
	methods := interfaceMethods(pkg, iface)
	service := g.clone()

	// Write services.
	if len(methods.others) > 0 {
		name := iface.Name + "Service"
		service.WriteComment(repository.Comment{Raw: fmt.Sprintf(
			"%s is the service containing call methods for implementations of interface %s.",
			name, iface.Name,
		)})

		service.body.WriteString("rpc_service ")
		service.body.WriteString(iface.Name + "Service")
		service.body.WriteString(" {\n")

		tail := service.clone()
		for _, method := range methods.others {
			service.writeServiceMethod(tail, iface.Name, method)
		}

		service.body.WriteByte('}')
		service.WriteLine()
		service.WriteLine()
		service.takeFrom(tail)
	}

	tail := g.clone()

	g.WriteComment(iface.Comment)
	g.body.WriteString("table ")
	g.body.WriteString(iface.Name)

	if len(methods.services) > 0 {
		g.body.WriteByte(' ')
		g.body.WriteString("(services: ")
		g.body.WriteString(strconv.Quote(strings.Join(methods.services, ", ")))
		g.body.WriteString(")")
	}

	g.body.WriteString(" {")

	if len(methods.fields) > 0 || len(methods.asserters) > 0 {
		g.WriteLine()
	}

	for _, getter := range methods.fields {
		fieldComment(&g.body, getter.Comment)

		field := repository.NamedType{
			Name: getter.Name,
			Type: getter.Returns[0].Type,
		}

		if len(getter.Returns) > 1 {
			var returns = make([]repository.StructField, len(getter.Returns))
			for i, ret := range getter.Returns {
				returns[i].NamedType = ret
			}
			field.Type = iface.Name + field.Name

			tail.WriteComment(repository.Comment{
				Raw: fmt.Sprintf(
					"%s is the return of method %s for interface %s.",
					field.Type, field.Name, iface.Name,
				),
			})
			tail.writeStruct(field.Type, returns)
		}

		g.writeField(field, fieldOpts{})
	}

	if len(methods.asserters) > 0 {
		g.WriteLine()
		fieldComment(&g.body, repository.Comment{Raw: "Asserters"})
	}

	for _, field := range methods.asserters {
		fieldType := repository.NamedType{
			Name: PascalToSnake(field.ChildType),
			Type: field.ChildType,
		}
		g.writeField(fieldType, fieldOpts{})
	}

	g.body.WriteByte('}')
	g.WriteLine()
	g.WriteLine()

	g.takeFrom(service)
	g.takeFrom(tail)
}

type methods struct {
	fields    []repository.GetterMethod // no parameters
	asserters []repository.AsserterMethod
	others    []repository.Method

	pkg      repository.Package
	services []string
}

func newMethods(pkg repository.Package) methods { return methods{pkg: pkg} }

func interfaceMethods(pkg repository.Package, iface repository.Interface) methods {
	methods := newMethods(pkg)
	methods.load(iface)

	return methods
}

func (methods *methods) load(iface repository.Interface) {
	for _, method := range iface.Methods {
		switch method := method.(type) {
		case repository.GetterMethod:
			if len(method.Parameters) == 0 {
				methods.fields = append(methods.fields, method)
				continue
			}
		case repository.AsserterMethod:
			methods.asserters = append(methods.asserters, method)
			continue
		}

		methods.others = append(methods.others, method)
	}

	if len(methods.others) > 0 {
		methods.services = append(methods.services, iface.Name+"Service")
	}

	// Load embedded interfaces recursively into another methods list, and then
	// prepend that.
	var child = newMethods(methods.pkg)

	for _, embedded := range iface.Embeds {
		iface := child.pkg.Interface(embedded.InterfaceName)
		if iface != nil {
			child.load(*iface)
		}
	}

	methods.fields = append(child.fields, methods.fields...)
	methods.asserters = append(child.asserters, methods.asserters...)

	// We have methods that we cannot inherit directly, which means that a
	// Service will be required. However, since that would lead to namespace
	// conflicts, we'll just have to indicate that there's inheritance and not
	// copy over the methods.
	methods.services = append(methods.services, child.services...)

	return
}
