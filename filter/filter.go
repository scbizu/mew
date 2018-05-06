package filter

import "strings"

// Filter defines the pkg filter
type Filter struct {
	PkgNames []string
}

// NewFilter init the filter
func NewFilter(pkgNames []string) *Filter {
	return &Filter{
		PkgNames: pkgNames,
	}
}

// Grep greps the pkg list
func (f *Filter) Grep(name string) (pkgNames []string) {
	for _, pn := range f.PkgNames {
		if strings.Contains(pn, name) {
			pkgNames = append(pkgNames, pn)
		}
	}
	return
}
