package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/projekt-zespolony/server/pkg/types"
)

type Database struct {
	db      *gorm.DB
	options *Options
}

type Options struct {
	User string
	Pass string
	Name string
	Addr string
}

func New(options *Options) (*Database, error) {
	args := "charset=utf8&parseTime=True&loc=Local"
	url := fmt.Sprintf("%s:%s@(%s)/%s?%s", options.User, options.Pass, options.Addr, options.Name, args)

	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&types.Sensors{}, &types.OptimizationData{})

	return &Database{
		db:      db,
		options: options,
	}, nil
}

func (database *Database) CreateSensors(sensors *types.Sensors) error {
	return database.db.Create(sensors).Error
}

func (database *Database) LatestSensors() (*types.Sensors, error) {
	sensors := &types.Sensors{}

	err := database.db.Last(sensors).Error
	if err != nil {
		return nil, err
	}

	return sensors, nil
}

func (database *Database) SinceSensors(timestamp int64) ([]*types.Sensors, error) {
	sensors := []*types.Sensors{}

	err := database.db.Where("timestamp > ?", timestamp).Find(&sensors).Error
	if err != nil {
		return nil, err
	}

	return sensors, nil
}

func (database *Database) CreateOptimizationData(opt *types.OptimizationData) error {
	return database.db.Create(opt).Error
}

func (database *Database) LatestOptimizationData() (*types.OptimizationData, error) {
	opt := &types.OptimizationData{}

	err := database.db.Last(opt).Error
	if err != nil {
		return nil, err
	}

	return opt, nil
}
