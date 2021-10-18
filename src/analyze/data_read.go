package analyze

import (
	"errors"
	"io/ioutil"
	"encoding/json"
)

type TargetInfo struct {
	Name string
	Up   bool
}

type FeatureInfo struct {
	Name string
}

type Target struct {
	Info  []TargetInfo
	Value [][]float32
}

type FeValue struct {
	Info  []FeatureInfo
	Value [][]float32
}

func TargetRead( file_name string ) ( Target, error ){
	var target Target

	raw, err := ioutil.ReadFile( file_name )

	if err != nil {
		return target, err
	}

	err = json.Unmarshal( raw, &target )

	if err != nil {
		return target, err
	}

	if len( target.Info ) != len( target.Value[0] ) {
		err = errors.New( "target not match value info" )
		return target, err
	}
	
	return target, nil
}

func FeValueRead( file_name string ) ( FeValue, error ){
	var fevalue FeValue

	raw, err := ioutil.ReadFile( file_name )

	if err != nil {
		return fevalue, err
	}

	err = json.Unmarshal( raw, &fevalue )

	if err != nil {
		return fevalue, err
	}

	if len( fevalue.Info ) != len( fevalue.Value[0] ) {
		err = errors.New( "fevalue not match value info" )
		return fevalue, err
	}
	
	return fevalue, nil
}

