package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" db:"id"`
	Url       string `json:"url" db:"url"`
	Width     int    `json:"width" db:"width"`
	Height    int    `json:"height" db:"height"`
	CloudName string `json:"cloud_name, omitempty" db:"cloud_name"`
	Extension string `json:"extension, omitempty" db:"extension"`
}

func (Image) TableName() string { return "images" }

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
