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
)

// DrawWithSlice returns the DOT tree
func DrawWithSlice(baseNode string, pkgNames []string) (string, error) {
	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	paresedNodeName := string(parseBaseNodeName([]byte(baseNode)))
	g.AddNode("G", paresedNodeName, nil)
	for _, name := range pkgNames {
		g.AddNode("G", name, nil)
		g.AddEdge(paresedNodeName, name, true, nil)
	}
	return g.String(), nil
}

// DrawWithSliceAndSave draws the grpah and save
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

func parseBaseNodeName(baseNode []byte) []byte {
	bs := []byte{'"'}
	for _, b := range baseNode {
		bs = append(bs, b)
	}
	bs = append(bs, '"')
	return bs
}
