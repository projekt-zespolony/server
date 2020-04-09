package types

type Status struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type Sensors struct {
	ID          uint    `json:"-"           gorm:"column:id" gorm:"primary_key,auto_increment"`
	Timestamp   int64   `json:"timestamp"   gorm:"column:timestamp"`
	Temperature float32 `json:"temperature" gorm:"column:temperature"`
	Pressure    float32 `json:"pressure"    gorm:"column:pressure"`
	Humidity    float32 `json:"humidity"    gorm:"column:humidity"`
	Gas         float32 `json:"gas"         gorm:"column:gas"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Topic string `json:"topic"`
}
