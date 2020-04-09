package types

type Status struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type Sensors struct {
	ID          uint    `json:"-"           sql:"column:id;primary_key;auto_increment"`
	Timestamp   int64   `json:"timestamp"   sql:"column:timestamp;type:decimal(15)"`
	Temperature float32 `json:"temperature" sql:"column:temperature;type:decimal(3,2)"`
	Pressure    float32 `json:"pressure"    sql:"column:pressure;type:decimal(4,2)"`
	Humidity    float32 `json:"humidity"    sql:"column:humidity;type:decimal(2,2)"`
	Gas         float32 `json:"gas"         sql:"column:gas;type:decimal(5,2)"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Topic string `json:"topic"`
}
