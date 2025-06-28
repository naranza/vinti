// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package log

import (
  "log/syslog"
  "log"
  "fmt"
)

const (
  DEBUG   = int(syslog.LOG_DEBUG)
  INFO    = int(syslog.LOG_INFO)
  WARNING = int(syslog.LOG_WARNING)
  ERR   = int(syslog.LOG_ERR)
)

var (
  minLevel int
)

func Init(level int) error {
  minLevel = level
  return nil
}

func Log(level int, format string, args ...any) {
  if level > minLevel {
    return
  }

  var prefix string
  switch level {
  case DEBUG:
    prefix = "DEBUG"
  case INFO:
    prefix = "INFO"
  case WARNING:
    prefix = "WARNING"
  case ERR:
    prefix = "ERROR"
  default:
    prefix = "INFO"
  }

  log.Printf("[%s] %s", prefix, fmt.Sprintf(format, args...))
}
