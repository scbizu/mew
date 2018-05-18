package linker

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

// Linker defines the link pkg config
type Linker struct {
	GoPath   string
	RepoPath string
}

func init() {
	initLogLevel()
}

// NewLinker inits the linker
func NewLinker(gopath string, repoRootPath string) *Linker {
	return &Linker{
		GoPath:   gopath,
		RepoPath: repoRootPath,
	}
}

func initLogLevel() {
	logrus.SetLevel(logrus.InfoLevel)
}

// SetLinkerLogLevel sets log lv.
func SetLinkerLogLevel(lv logrus.Level) {
	logrus.SetLevel(lv)
}

// GetAllPkgNames get repo related pkgs
func (l *Linker) GetAllPkgNames(allowDup bool, excludeDirs []string) (pkgNames []string, err error) {
	fset := token.NewFileSet()
	fpath := fmt.Sprintf("%s/src/%s/", l.GoPath, l.RepoPath)
	allDirs := []string{fpath}
	var dirs []string
	dirs, err = walkAround(fpath, excludeDirs)
	if err != nil {
		return
	}
	allDirs = append(allDirs, dirs...)
	for _, p := range allDirs {
		// Parse src but stop after processing the imports.
		fs, fsErr := parser.ParseDir(fset, p, nil, parser.ImportsOnly)
		if err != nil {
			return nil, fsErr
		}
		for _, f := range fs {
			for fk, file := range f.Files {
				for _, s := range file.Imports {
					logrus.WithField("file", fk).Debugf("Get PKG:%s", s.Path.Value)
					pkgNames = append(pkgNames, s.Path.Value)
				}
			}
		}
	}
	if !allowDup {
		// remove duplicate elem
		pkgNames = removeDupPkgNames(pkgNames)
	}
	return
}

// GetInvokeSrcMap get pkg names(as value) with his import file(as key)
func (l *Linker) GetInvokeSrcMap() (map[string][]string, error) {
	return nil, nil
}

func removeDupPkgNames(pkgNames []string) []string {
	resMap := make(map[string]struct{}, len(pkgNames))
	var j int
	for _, pkgName := range pkgNames {
		if _, ok := resMap[pkgName]; ok {
			continue
		}
		resMap[pkgName] = struct{}{}
		pkgNames[j] = pkgName
		j++
	}
	return pkgNames[:j]
}

func walkAround(gpath string, excludeDirs []string) (dirs []string, err error) {
	var files []os.FileInfo
	files, err = ioutil.ReadDir(gpath)
	if err != nil {
		return
	}
	for _, file := range files {
		if isIgnore(file, excludeDirs) {
			continue
		}
		absPath := fmt.Sprintf("%s%s", gpath, appendSlash(file.Name()))
		dirs = append(dirs, absPath)
		var subDirs []string
		subDirs, err = walkAround(absPath, excludeDirs)
		if err != nil {
			return
		}
		dirs = append(dirs, subDirs...)
	}
	return
}

func isIgnore(f os.FileInfo, excludeDirs []string) bool {
	if f.IsDir() {
		for _, ed := range excludeDirs {
			if f.Name() == ed {
				return true
			}
		}

		return false
	}
	return true
}

func appendSlash(fileName string) string {
	return fmt.Sprintf("%s/", fileName)
}
