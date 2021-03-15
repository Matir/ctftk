package testdata

import (
	"os"
	"path"
	"runtime"
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
