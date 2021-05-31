package base

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
)

const (
	DEFAULT_LOG = iota
	DEFAULT_ERR_LOG
)

var debug = true

func SetDebug(b bool) {
	debug = b
}

var _LOGGERS = map[int]*log.Logger{}

func Infof(format string, args ...interface{}) {
	_LOGGERS[DEFAULT_LOG].Output(2, fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...interface{}) {
	if debug {
		_LOGGERS[DEFAULT_LOG].Output(2, fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	_LOGGERS[DEFAULT_ERR_LOG].Output(2, fmt.Sprintf(format, args...))
}

func ConfigureLogger(dir string, maxSize, maxBackups, maxAge int) error {
	makeLogger := func(prefix string, tag string, flag int) *log.Logger {
		return log.New(&lumberjack.Logger{
			Filename:   filepath.Join(dir, prefix+".log"),
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
		}, tag, flag)
	}

	// mkdir -p
	os.MkdirAll(dir, os.FileMode(0755))

	if debug {
		_LOGGERS[DEFAULT_LOG] = makeLogger("agent", "AGENT-DEBUG:", log.LstdFlags|log.Lshortfile)
	} else {
		_LOGGERS[DEFAULT_LOG] = makeLogger("agent", "AGENT:", log.LstdFlags)
	}
	_LOGGERS[DEFAULT_ERR_LOG] = makeLogger("agent-error", "AGENT-ERROR:", log.LstdFlags|log.Lshortfile)

	return nil
}
