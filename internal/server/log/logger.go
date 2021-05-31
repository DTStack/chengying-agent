package log

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

var (
	LOGDIR          string
	_LOGGERS        = map[int]*log.Logger{}
	_OUTPUT_LOGGERS = map[string]*log.Logger{}
	LOGGER_MAX_SIZE int
	LOGGER_MAX_BKS  int
	LOGGER_MAX_AGE  int

	makeLogger func(prefix string, tag string, flag int) *log.Logger

	debug bool
)

func SetDebug(b bool) {
	debug = b
}

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

func Output2Path(path string, format string, args ...interface{}) error {
	_, exsits := _OUTPUT_LOGGERS[path]
	if !exsits {
		_OUTPUT_LOGGERS[path] = makeLogger(path, path, log.LstdFlags|log.Lshortfile)
	}
	return _OUTPUT_LOGGERS[path].Output(4, fmt.Sprintf(format, args...))
}

func ConfigureLogger(dir string, maxSize, maxBackups, maxAge int) error {
	LOGDIR = dir
	LOGGER_MAX_SIZE = maxSize
	LOGGER_MAX_BKS = maxBackups
	LOGGER_MAX_AGE = maxAge

	makeLogger = func(prefix string, tag string, flag int) *log.Logger {
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
		_LOGGERS[DEFAULT_LOG] = makeLogger("easyagent-server", "EA-SERVER-DEBUG:", log.LstdFlags|log.Lshortfile)
	} else {
		_LOGGERS[DEFAULT_LOG] = makeLogger("easyagent-server", "EA-SERVER:", log.LstdFlags)
	}
	_LOGGERS[DEFAULT_ERR_LOG] = makeLogger("easyagent-server-error", "EA-SERVER-ERROR:", log.LstdFlags|log.Lshortfile)

	return nil
}
