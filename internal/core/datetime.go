// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, AGPLv3
package core

import (
  "fmt"
  "time"
)

// Datetime returns a sortable timestamp with microsecond precision.
func Datetime(t time.Time) string {
  if t.IsZero() {
    t = time.Now()
  }
  return fmt.Sprintf("%s%09d", t.Format("20060102150405"), t.Nanosecond())
}