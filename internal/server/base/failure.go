package base

import (
	"fmt"
)

const (
	NORMAL_EXIT = iota
	NETWORK_FAILURE
)

type SystemFailure struct {
	ExitCode int
	Reason   string
}

func SystemExitWithFailure(exitCode int, r interface{}, args ...interface{}) {
	if exitCode == NORMAL_EXIT {
		_SYSTEM_FAIL <- SystemFailure{NORMAL_EXIT, ""}
	} else if reason, ok := r.(string); ok {
		_SYSTEM_FAIL <- SystemFailure{exitCode, fmt.Sprintf(reason, args...)}
	} else if err, ok := r.(error); ok {
		_SYSTEM_FAIL <- SystemFailure{exitCode, err.Error()}
	} else if code, ok := r.(int); ok {
		_SYSTEM_FAIL <- SystemFailure{exitCode, fmt.Sprintf("Code: %d", code)}
	} else {
		_SYSTEM_FAIL <- SystemFailure{exitCode, "Unknown reason"}
	}
}
