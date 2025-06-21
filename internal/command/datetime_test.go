// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
  "testing"
  "time"
)

func TestDatetime_ZeroTime(t *testing.T) {
  
  var testTime time.Time
  wig := Datetime(testTime)

  if len(wig) == 24 {
    t.Errorf("wig: %s", wig)
  }
}

func TestDatetime_SpecificTime(t *testing.T) {
  
  testTime := time.Date(2023, time.April, 15, 10, 20, 30, 123456789, time.UTC)
  wie := "20230415102030123456789" // YYYYMMDDHHMMSSnnnnnnnnn

  wig := Datetime(testTime)

  if wie != wig {
    t.Errorf("wie:%s - wig: %s", wie, wig)
  }
}
