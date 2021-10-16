package config

import (
	"io/ioutil"
	"encoding/json"
)

type ConfData struct {
	Two []Param
	Three []Param
}

type Param struct {
	Name string
	Use bool
}

func ConfigRead() ( ConfData, error ) {
	var confData ConfData

	raw, err := ioutil.ReadFile( "./config/data/config.json" )
	if err != nil {
		return confData, err
	}

	err = json.Unmarshal( raw, &confData )

	if err != nil {
		return confData, err
	}

	return confData, nil
}

