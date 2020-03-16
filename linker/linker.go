package linker

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Linker defines the link pkg config
type Linker struct {
	GoPath   string
	RepoPath string
	FullPath string
}

func init() {
	initLogLevel()
}

var (
	scannedRepo = make(map[string]bool)
	pkgMap      = make(map[string][]string)
)

// NewLinker inits the linker
func NewLinker(gopath string, repoRootPath string) *Linker {
	l := &Linker{
		GoPath:   gopath,
		RepoPath: repoRootPath,
	}
	return l
}

func initLogLevel() {
	logrus.SetLevel(logrus.InfoLevel)
}

// SetLinkerLogLevel sets log lv.
func SetLinkerLogLevel(lv logrus.Level) {
	logrus.SetLevel(lv)
}

// GetAllPKGNames gets the full layers packages names
func (l *Linker) GetAllPKGNames(allowDup bool, excludeDirs []string) (map[string][]string, error) {

	names, err := l.GetLayerPKGNames(allowDup, excludeDirs)
	if err != nil {
		if os.IsNotExist(err) {
			logrus.Warnf("[file not exists]:%v", err)
			pkgMap[l.RepoPath] = []string{}
			return pkgMap, nil
		}
		return nil, err
	}
	logrus.Debugf("pkgnames:[%v]", names)
	for _, repo := range names {
		if !isThirdPartyPackage(repo) {
			continue
		}
		if _, ok := scannedRepo[repo]; ok {
			logrus.Warnf("scanned repo:[%v]", repo)
			continue
		}
		scannedRepo[repo] = true
		logrus.Debugf("repo path:%v,package:%v", l.RepoPath, repo)
		pkgMap[l.RepoPath] = append(pkgMap[l.RepoPath], repo)
		lk := NewLinker(l.GoPath, repo)
		pkgMap, err = lk.GetAllPKGNames(allowDup, excludeDirs)
		if err != nil {
			return nil, err
		}
	}

	return pkgMap, nil
}

// GetLayerPKGNames gets the layer(depends on the repo) package names
// DO NOT SUPPORT GOROOT ENV
func (l *Linker) GetLayerPKGNames(allowDup bool, excludeDirs []string) (pkgNames []string, err error) {
	fset := token.NewFileSet()
	fpath := fmt.Sprintf("%s/src/%s/", l.GoPath, l.RepoPath)
	if l.FullPath != "" {
		fpath = l.FullPath
	}
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
					pkgNames = append(pkgNames, trimQuotation(s.Path.Value))
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

func trimQuotation(pkgName string) string {
	pkgName = strings.Replace(pkgName, `"`, ``, -1)
	return pkgName
}

func isThirdPartyPackage(repo string) bool {
	tpPrefixes := []string{"github.com"}
	for _, p := range tpPrefixes {
		ok := strings.HasPrefix(repo, p)
		if ok {
			return true
		}
	}
	return false
}

// GetInvokeSrcMap get pkg names(as value) with his import file(as key)
// TODO
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
