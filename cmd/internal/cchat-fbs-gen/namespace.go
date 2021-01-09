package main

import (
	"path"
	"strings"

	"github.com/diamondburned/cchat/repository"
)

var rootNamespace = pathToNamespace(path.Base(repository.RootPath))

func namespaceType(namespace, typeName string) string {
	return joinNamespace(namespace) + "." + typeName
}

func joinNamespace(namespace string) string {
	if namespace == rootNamespace {
		return rootNamespace
	}
	return strings.Join([]string{rootNamespace, namespace}, ".")
}

func pathToNamespace(path string) string {
	parts := strings.Split(path, "/")

	// Reverse domain.
	domainParts := strings.Split(parts[0], ".")
	// SliceTricks: Reversing.
	for i := len(domainParts)/2 - 1; i >= 0; i-- {
		opp := len(domainParts) - 1 - i
		domainParts[i], domainParts[opp] = domainParts[opp], domainParts[i]
	}
	parts[0] = strings.Join(domainParts, ".")

	return strings.Join(parts, ".")
}

func dirNamespace(namespace string) string {
	if strings.Contains(namespace, ".") {
		parts := strings.Split(namespace, ".")
		namespace = strings.Join(parts[:len(parts)-1], ".")
	}
	return namespace
}

func trimPrefix(rootPrefix, path string) string {
	return strings.Trim(strings.TrimPrefix(path, rootPrefix), "/")
}
