package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"kego.io/process"
	"kego.io/uerr"
)

func main() {
	var testFlag = flag.Bool("test", false, "test mode? e.g. don't write the files")
	var pathFlag = flag.String("path", "", "full package path e.g. github.com/foo/bar")
	flag.Parse()
	testMode := *testFlag
	packagePath := *pathFlag

	currentDir, err := filepath.Abs("")
	if err != nil {
		log.Fatalf("Error getting the current working directory:\n%v\n", err.Error())
	}

	if packagePath == "" {
		path, err := GetPackage(currentDir)
		if err != nil {
			log.Fatalf("Error while getting package path from current working directory. You can specify the full package path with the -path=github.com/foo\n%v\n", err.Error())
		}
		packagePath = path
	}

	parts := strings.Split(packagePath, string(os.PathSeparator))
	packageName := parts[len(parts)-1]

	imports := map[string]string{}

	if err := process.Scan(currentDir, packagePath, imports); err != nil {
		panic(err)
	}

	if err := process.Generate(currentDir, packageName, packagePath, imports, testMode); err != nil {
		panic(err)
	}
}

func GetPackage(dir string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if strings.HasPrefix(gopath, "file://") {
		// This is to fix a bug when running the code in IntelliJ. This can be removed when the
		// bug is fixed: https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1600
		gopath = gopath[7:]
	}
	return getPackage(dir, gopath)
}
func getPackage(dir string, gopathEnv string) (string, error) {
	gopaths := filepath.SplitList(gopathEnv)
	var savedError error
	for _, gopath := range gopaths {
		if strings.HasPrefix(dir, gopath) {
			gosrc := fmt.Sprintf("%s/src", gopath)
			relpath, err := filepath.Rel(gosrc, dir)
			if err != nil {
				savedError = err
				continue
			}
			if relpath == "" {
				continue
			}
			return relpath, nil
		}
	}
	if savedError != nil {
		return "", savedError
	}
	return "", uerr.New("CXOETFPTGM", nil, "process/cmd/kego.getPackage", "Package not found for %s", dir)
}
