package genutil

import "github.com/diamondburned/cchat/repository"

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
