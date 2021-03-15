package testdata

import (
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

func GetTestdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

func OpenTestdataFile(filename string) (*os.File, error) {
	fname := path.Join(GetTestdataDir(), filename)
	return os.Open(fname)
}

func MustOpenTestdataFile(filename string) *os.File {
	fp, err := OpenTestdataFile(filename)
	if err != nil {
		panic(err)
	}
	return fp
}

func ExpectErrorContains(t *testing.T, e error, s string) bool {
	if e == nil {
		t.Fatalf("Expected error containing %s, got nil.", s)
		return false
	}
	eStr := e.Error()
	if !strings.Contains(eStr, s) {
		t.Fatalf("Expected error containing %s, got %s", s, eStr)
		return false
	}
	return true
}
