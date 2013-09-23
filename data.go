package chromelogger

const VERSION = "0.1.0"

type LogData []interface{}

func (data *LogData) Add(item interface{}) {
  *data = append(*data, item)
}

type LogRow []interface{}

func NewLogRow(logData *LogData, backtrace, logType string) *LogRow {
  row := LogRow{
    logData,
    backtrace,
    logType,
  }

  return &row
}

type Data struct {
  Version string    `json:"version"`
  Columns []string  `json:"columns"`
  Rows    []*LogRow `json:"rows"`
}

func (data *Data) AddRow(item interface{}, backtrace, logType string) {
  logData := make(LogData, 0)
  logData.Add(item)
  row := NewLogRow(&logData, backtrace, logType)
  data.Rows = append(data.Rows, row)
}

func NewData() *Data {
  data := &Data{
    Version: VERSION,
    Columns: []string{"log", "backtrace", "type"},
  }
  data.Rows = make([]*LogRow, 0)

  return data
}

