package types

type Status struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type Sensors struct {
	ID          uint    `json:"-"           sql:"column:id;          primary_key;              auto_increment"`
	Timestamp   int64   `json:"timestamp"   sql:"column:timestamp;   type:bigint unsigned;     not null"`
	Temperature float32 `json:"temperature" sql:"column:temperature; type:float(5,2) signed;   not null"`
	Pressure    float32 `json:"pressure"    sql:"column:pressure;    type:float(6,2) unsigned; not null"`
	Humidity    float32 `json:"humidity"    sql:"column:humidity;    type:float(4,2) unsigned; not null"`
	Gas         float32 `json:"gas"         sql:"column:gas;         type:float(8,2) unsigned; not null"`
}

type OptimizationData struct {
	ID				 uint 	`json:"-"				 sql:"column:id;			  primary_key;		 auto_increment"`
	SensorsID   	 uint   `json:"-" 				 sql:"column:sensorsID	 	  type:int unsigned; not null"`
	PeopleInTheRoom  uint 	`json:"peopleInTheRoom"	 sql:"column:peopleInTheRoom  type:int unsigned; not null"`
	WindowsAreOpened bool	`json:"WindowsAreOpened" sql:"column:windowsAreOpened type:boolean;      not null"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Topic string `json:"topic"`
}
