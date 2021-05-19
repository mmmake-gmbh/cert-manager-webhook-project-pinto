package testutils

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

// getCallerDir stacktraceStep is the number calls we travel up to identify the source Go file. This will usually be 2.
func getCallerDir(stacktraceStep int) (string, error) {
	_, filename, _, ok := runtime.Caller(stacktraceStep)
	if ok {
		return path.Dir(filename), nil
	}
	return "", fmt.Errorf("could not retrieve current dir")
}

// ReadRelativeFile read file relative to the calling go file
func ReadRelativeFile(relativePath string) ([]byte, error) {
	sourceDir, callErr := getCallerDir(2)
	if callErr != nil {
		return nil, callErr
	}

	filepath := path.Join(sourceDir, relativePath)

	return os.ReadFile(filepath)
}

// ReadFixture read a file from the fixture directory right below of the calling go file. E.g. if `a/b/main.go` is calling this function, the fixture will be searched in `a/b/fixture. Subdirs are allowed for fileName e.g. `subDir/fileName` will lead to `a/b/fixtures/subDir/FileName`
func ReadFixture(fileName string) ([]byte, error) {
	sourceDir, callErr := getCallerDir(2)
	if callErr != nil {
		return nil, callErr
	}

	filepath := path.Join(sourceDir, "fixtures", fileName)

	return os.ReadFile(filepath)
}
