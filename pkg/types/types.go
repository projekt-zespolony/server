package types

import "github.com/jinzhu/gorm"

type Status struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type Sensors struct {
	gorm.Model  `json:"-"`
	Timestamp   int64   `json:"timestamp"   gorm:"timestamp"`
	Temperature float32 `json:"temperature" gorm:"temperature"`
	Pressure    float32 `json:"pressure"    gorm:"pressure"`
	Humidity    float32 `json:"humidity"    gorm:"humidity"`
	Gas         float32 `json:"gas"         gorm:"gas"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Topic string `json:"topic"`
}
