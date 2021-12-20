package analyze

import (
	"math"
	"errors"
)


func add( a float32, b float32 ) ( float32 ) {
	return a + b
}

func mul( a float32, b float32 ) ( float32 ) {
	return a * b
}

func mean( valueList []float32 ) ( float32 ) {
	var all float32 = 0
	var c float32 = float32( len( valueList ) )

	for _, v := range valueList {
		all += v
	}

	return all / c
}

func diff( a float32, b float32 ) ( float32 ) {
	return float32( math.Abs( float64(  a - b ) ) )
}

func pow( a float32 ) ( float32 ) {
	return a * a
}

func std( valueList []float32 ) ( float32 ) {
	var all float32 = 0
	var m float32 = mean( valueList )
	var c float32 = float32( len( valueList ) )
	
	for _, v := range valueList {
		all += pow( m - v )
	}

	all /= c
	return float32( math.Sqrt( float64( all ) ) )
}

func cov( xList []float32, yList []float32 ) ( float32 ) {
	xMean := mean( xList )
	yMean := mean( yList )

	var c float32 = 0
	
	for i := 0; i < len( xList ); i++ {
		c += ( xList[i] - xMean ) * ( yList[i] - yMean )
	}

	return c / float32( len( xList ) )
}

func TwoCluc( valueList []float32, name string ) ( float32, error ) {
	a := valueList[0]
	b := valueList[1]
	
	if name == "add" {
		return add( a, b ), nil
	} else if name == "mul" {
		return mul( a, b ), nil
	} else if name == "diff" {
		return diff( a, b ), nil
	} else if name == "mean" {
		return mean( valueList ), nil
	}

	return 0, errors.New( "not mutch calc" )
}

func ThreeCluc( valueList []float32, name string ) ( float32, error ) {
	
	if name == "mean" {
		return mean( valueList ), nil
	} else if name == "std" {
		return std( valueList ), nil
	}

	return 0, errors.New( "not mutch calc" )
}
