package file

import (
	"errors"
	"io/ioutil"
	"encoding/json"

	"frontier/analyze"
)

func TargetRead( file_name string ) ( analyze.Target, error ){
	var target analyze.Target

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

func FeValueRead( file_name string ) ( analyze.FeValue, error ){
	var fevalue analyze.FeValue

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

