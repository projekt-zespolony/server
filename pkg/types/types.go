package types

type Status struct {
	Version string
	Commit  string
}

type Sensors struct {
	ID          uint `json:"-" gorm:"primary_key,auto_increment"`
	Timestamp   int64
	Temperature float32
	Pressure    float32
	Humidity    float32
	Gas         float32
}

type Notification struct {
	Title string
	Body  string
	Topic string
}
