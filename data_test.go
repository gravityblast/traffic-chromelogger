package chromelogger

import(
  "testing"
  assert "github.com/pilu/miniassert"
)

func TestLogData_Add(t *testing.T) {
  logData := make(LogData, 0)

  assert.Equal(t, 0, len(logData))

  logData.Add(1)
  assert.Equal(t, 1, len(logData))
}

func TestNewLogRow(t *testing.T) {
  logData := make(LogData, 0)
  row := NewLogRow(&logData, "foo.go:100", "log")

  assert.Equal(t, &logData, (*row)[0])
  assert.Equal(t, "foo.go:100", (*row)[1])
  assert.Equal(t, "log", (*row)[2])
}

func TestNewData(t *testing.T) {
  data := NewData()

  assert.Equal(t, VERSION, data.Version)
  assert.Equal(t, []string{"log", "backtrace", "type"}, data.Columns)
  assert.Equal(t, 0, len(data.Rows))
}
