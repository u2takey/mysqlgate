package log

import (
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

var defaultLogger = log.With(log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)),
	"ts", log.DefaultTimestampUTC, "caller", log.Caller(6))

func SetLogLevel(levelStr string){
	if levelValue, err := level.Parse(levelStr); err == nil{
		defaultLogger = level.NewFilter(defaultLogger, level.Allow(levelValue))
	}else{
		defaultLogger = level.NewFilter(defaultLogger, level.Allow(level.InfoValue()))
		_ = level.Info(defaultLogger).Log("msg", "level not valid, use default level: info")
	}
}

// log format: [reqId=%d] [module=%s] [method=%s] msg=%s [err=%s]
func Debug(keyvals ...interface{}){
	_ = level.Debug(defaultLogger).Log(keyvals...)
}

func Info(keyvals ...interface{}){
	_ = level.Info(defaultLogger).Log(keyvals...)
}

func Warn(keyvals ...interface{}){
	_ = level.Warn(defaultLogger).Log(keyvals...)
}

func Error(keyvals ...interface{}){
	_ = level.Error(defaultLogger).Log(keyvals...)
}

type MLogger struct {
	log.Logger
}

func (m* MLogger)Debug(keyvals ...interface{}){
	_ = level.Debug(m.Logger).Log(keyvals...)
}

func (m* MLogger)Info(keyvals ...interface{}){
	_ = level.Info(m.Logger).Log(keyvals...)
}

func (m* MLogger)Warn(keyvals ...interface{}){
	_ = level.Warn(m.Logger).Log(keyvals...)
}

func (m* MLogger)Error(keyvals ...interface{}){
	_ = level.Error(m.Logger).Log(keyvals...)
}

func ModuleLogger(module string)MLogger{
	return MLogger{log.With(defaultLogger, "module", module)}
}