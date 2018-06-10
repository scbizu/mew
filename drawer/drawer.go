// Package drawer gen the graph and etc.
package drawer

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/sirupsen/logrus"
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
	paresedNodeName := addQuotation(baseNode)
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
func DrawWithMap(baseNode string, pkgMap map[string][]string) (string, error) {
	g := gographviz.NewGraph()
	g.SetName(GraphName)
	g.SetDir(true)

	drawTree(g, baseNode, pkgMap)

	return g.String(), nil
}

func drawTree(g *gographviz.Graph, base string, pkgMap map[string][]string) {

	ps, ok := pkgMap[base]

	if !ok || len(ps) == 0 {
		return
	}

	g.AddNode(GraphName, addQuotation(base), nil)

	for _, p := range ps {
		g.AddNode(GraphName, addQuotation(p), nil)
		if err := g.AddEdge(addQuotation(base), addQuotation(p), true, nil); err != nil {
			logrus.Error(err)
		}
		drawTree(g, p, pkgMap)
	}

	return

}

// DrawWithMapAndSave draws the graph and save to current path by a map
func DrawWithMapAndSave(baseNode string, filename string, pkgMap map[string][]string) error {
	dotTree, err := DrawWithMap(baseNode, pkgMap)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, []byte(dotTree), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func addQuotation(node string) string {
	bs := []byte{'"'}
	for _, b := range []byte(node) {
		bs = append(bs, b)
	}
	bs = append(bs, '"')
	return addNAfterSlash(string(bs))
}

// addNAfterSlash enter the newline after /
func addNAfterSlash(nodeName string) string {
	return strings.Replace(nodeName, "/", "/\n", -1)
}
