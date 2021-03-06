// Package drawer gen the graph and etc.
package drawer

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"

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
func DrawWithSlice(g *gographviz.Graph, baseNode string, pkgNames []string) (string, error) {
	paresedNodeName := addQuotation(baseNode)
	_ = g.AddNode(GraphName, paresedNodeName, nil)
	for _, name := range pkgNames {
		_ = g.AddNode(GraphName, addQuotation(name), nil)
		_ = g.AddEdge(paresedNodeName, addQuotation(name), true, nil)
	}
	return g.String(), nil
}

type DotTree struct {
	g      *gographviz.Graph
	buffer *bytes.Buffer
}

func NewDot() *DotTree {
	g := gographviz.NewGraph()
	g.SetName(GraphName)
	g.SetDir(true)
	return &DotTree{g: g, buffer: bytes.NewBufferString(g.String())}
}

func (d *DotTree) AddDep(baseNode string, pkgNames []string) error {
	paresedNodeName := addQuotation(baseNode)
	_ = d.g.AddNode(GraphName, paresedNodeName, nil)
	for _, name := range pkgNames {
		_ = d.g.AddNode(GraphName, addQuotation(name), nil)
		_ = d.g.AddEdge(paresedNodeName, addQuotation(name), true, nil)
	}
	d.buffer.Reset()
	if _, err := d.buffer.WriteString(d.g.String()); err != nil {
		return err
	}
	return nil
}

func (d *DotTree) WriteFile(filename string) error {
	if err := ioutil.WriteFile(filename, d.buffer.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil
}

// DrawWithSliceAndSave draws the grpah and save to current path by a slice
func DrawWithSliceAndSave(filename string, baseNode string, pkgNames []string) error {
	g := gographviz.NewGraph()
	_ = g.SetName(GraphName)
	_ = g.SetDir(true)
	dotTree, err := DrawWithSlice(g, baseNode, pkgNames)
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
	_ = g.SetName(GraphName)
	_ = g.SetDir(true)

	drawTree(g, baseNode, pkgMap)

	return g.String(), nil
}

func drawTree(g *gographviz.Graph, base string, pkgMap map[string][]string) {

	ps, ok := pkgMap[base]

	if !ok || len(ps) == 0 {
		return
	}

	_ = g.AddNode(GraphName, addQuotation(base), nil)

	for _, p := range ps {
		_ = g.AddNode(GraphName, addQuotation(p), nil)
		if err := g.AddEdge(addQuotation(base), addQuotation(p), true, nil); err != nil {
			logrus.Error(err)
		}
		drawTree(g, p, pkgMap)
	}

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
	return strconv.Quote(node)
}
