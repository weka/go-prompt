//go:build !unix && !windows
// +build !unix,!windows

package prompt

func NewStdoutWriter() Writer { return nil }
func NewStderr() Writer       { return nil }
