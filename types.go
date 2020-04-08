package main

type Status struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type Sensors struct {
	Timestamp   int64   `json:"timestamp"`
	Temperature float32 `json:"temperature"`
	Pressure    float32 `json:"pressure"`
	Humidity    float32 `json:"humidity"`
	Gas         float32 `json:"gas"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Topic string `json:"topic"`
}
