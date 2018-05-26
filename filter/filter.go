package filter

import "strings"

// Filter defines the pkg filter
type Filter struct {
	PkgNames []string
}

// MapFilter defines the pkg map filter
type MapFilter struct {
	PKGMap map[string][]string
}

// NewFilter init the filter
func NewFilter(pkgNames []string) *Filter {
	return &Filter{
		PkgNames: pkgNames,
	}
}

// NewMapFilter init the map filter
func NewMapFilter(pkgmap map[string][]string) *MapFilter {
	return &MapFilter{
		PKGMap: pkgmap,
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

// Grep greps the pkg list
func (f *MapFilter) Grep(name string) map[string][]string {
	res := make(map[string][]string)
	for repo, pkgs := range f.PKGMap {
		f := NewFilter(pkgs)
		pkgs = f.Grep(name)
		res[repo] = pkgs
	}
	return res
}
