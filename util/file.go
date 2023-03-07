package util

import (
	"path"
	"runtime"
)

type Path struct {
	File string
	Dir  string
	Line int
}

func GetPath() *Path {
	return GetPathWithCallerSkip(1)
}
func GetPathWithCallerSkip(skip int) *Path {
	_, file, line, _ := runtime.Caller(skip)
	return &Path{
		File: file,
		Dir:  path.Dir(file),
		Line: line,
	}
}
