package portal

import "runtime"

func ErrorFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

func ErrorLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}
