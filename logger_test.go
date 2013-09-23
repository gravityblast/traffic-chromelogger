package chromelogger

import (
  "testing"
  "path/filepath"
  "strings"
  assert "github.com/pilu/miniassert"
)

func TestLogger_Add(t *testing.T) {
  logger := NewLogger()

  assert.Equal(t, 0, len(logger.data.Rows))

  logger.add("foo", "log")
  assert.Equal(t, 1, len(logger.data.Rows))
}

func TestLogger_Log(t *testing.T) {
  logger := NewLogger()

  assert.Equal(t, 0, len(logger.data.Rows))

  logger.Log("foo")
  assert.Equal(t, 1, len(logger.data.Rows))

  row := logger.data.Rows[0]
  backtracePath := (*row)[1].(string)
  backtraceFilenameAndLine := filepath.Base(backtracePath)
  backtraceChunks := strings.Split(backtraceFilenameAndLine, ":")

  assert.Equal(t, "logger_test.go", backtraceChunks[0])
  assert.Equal(t, "log", (*row)[2])
}

func TestLogger_Warn(t *testing.T) {
  logger := NewLogger()

  logger.Warn("foo")
  row := logger.data.Rows[0]
  assert.Equal(t, "warn", (*row)[2])
}

func TestLogger_Error(t *testing.T) {
  logger := NewLogger()

  logger.Error("foo")
  row := logger.data.Rows[0]
  assert.Equal(t, "error", (*row)[2])
}

func TestLogger_Info(t *testing.T) {
  logger := NewLogger()

  logger.Info("foo")
  row := logger.data.Rows[0]
  assert.Equal(t, "info", (*row)[2])
}
