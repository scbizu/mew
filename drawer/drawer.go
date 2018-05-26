// Package drawer gen the graph and etc.
package drawer

import (
	"io/ioutil"
	"os"

	"github.com/awalterschulze/gographviz"
)

const (
	// DefaultFileName defines default file name
	DefaultFileName = "mew.dot"
	// GraphName defines parent graph name
	GraphName = "Mew"
)

// DrawWithSlice returns the DOT lang  of a slice
func DrawWithSlice(baseNode string, pkgNames []string) (string, error) {
	g := gographviz.NewGraph()
	g.SetName(GraphName)
	g.SetDir(true)
	paresedNodeName := string(parseBaseNodeName([]byte(baseNode)))
	g.AddNode(GraphName, paresedNodeName, nil)
	for _, name := range pkgNames {
		g.AddNode(GraphName, name, nil)
		g.AddEdge(paresedNodeName, name, true, nil)
	}
	return g.String(), nil
}

// DrawWithSliceAndSave draws the grpah and save to current path by a slice
func DrawWithSliceAndSave(filename string, baseNode string, pkgNames []string) error {
	dotTree, err := DrawWithSlice(baseNode, pkgNames)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, []byte(dotTree), os.ModePerm); err != nil {
		return err
	}
	return nil
}

// DrawWithMap returns DOT lang of a map
func DrawWithMap(pkgMap map[string][]string) (string, error) {
	g := gographviz.NewGraph()
	g.SetName(GraphName)
	g.SetDir(true)
	for repo, pkgs := range pkgMap {
		paresedNodeName := string(parseBaseNodeName([]byte(repo)))
		for _, pkg := range pkgs {
			g.AddNode(GraphName, pkg, nil)
			g.AddEdge(paresedNodeName, pkg, true, nil)
		}
	}
	return g.String(), nil
}

// DrawWithMapAndSave draws the graph and save to current path by a map
func DrawWithMapAndSave(filename string, pkgMap map[string][]string) error {
	dotTree, err := DrawWithMap(pkgMap)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, []byte(dotTree), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func parseBaseNodeName(baseNode []byte) []byte {
	bs := []byte{'"'}
	for _, b := range baseNode {
		bs = append(bs, b)
	}
	bs = append(bs, '"')
	return bs
}
