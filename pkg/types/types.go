package types

type Status struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type Sensors struct {
	ID          uint    `json:"-"           sql:"column:id;primary_key;auto_increment"`
	Timestamp   int64   `json:"timestamp"   sql:"column:timestamp;type:bigint unsigned"`
	Temperature float32 `json:"temperature" sql:"column:temperature;type:float(5,2)"`
	Pressure    float32 `json:"pressure"    sql:"column:pressure;type:float(6,2) unsigned"`
	Humidity    float32 `json:"humidity"    sql:"column:humidity;type:float(4,2) unsigned"`
	Gas         float32 `json:"gas"         sql:"column:gas;type:float(8,2) unsigned"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Topic string `json:"topic"`
}
