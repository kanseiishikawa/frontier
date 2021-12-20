package analyze

import (
	"math"
)

type Diff struct {
	Score float32
	Ave float32
	A float32
	B float32
}

type ScoreData struct {
	RScore float32
	Diff Diff
	TargetName string
	FeValueName string
}

func correlationCoefficient( xList []float32, yList []float32 ) ( float32 ) {
	tfConv := cov( xList, yList )
	targetStd := std( xList )
	feValueStd := std( yList )
	score := tfConv / ( targetStd * feValueStd )

	if score < 0 {
		score *= -1
	}

	return score	
}

func leastSquare( xList []float32, yList []float32 ) ( float32, float32 ) {
	xStd := std( xList )
	xMean := mean( xList )
	yMean := mean( yList )
	tfConv := cov( xList, yList )
	a := tfConv / xStd
	b := yMean - a * xMean

	return a, b
}

func diffScore( xList []float32, yList []float32, a float32, b float32 ) ( float32, float32 ) {
	var diff float32
	var ave float32
	diff = 0
	ave = 0
	
	for i := 0; i < len( xList ); i++ {
		predictTarget := a * xList[i] + b
		diff += float32( math.Abs( float64( yList[i] - predictTarget ) ) )
		ave += predictTarget
	}

	diff /= float32( len( xList ) )
	ave /= float32( len( xList ) )
	
	return diff, ave
}

func AnalyzeScore( targets []AnalyzeData, feValue *AnalyzeData, res chan []ScoreData ) {
	var result []ScoreData
	
	for _, target := range targets {
		var instance ScoreData
		instance.TargetName = target.Name
		instance.FeValueName = feValue.Name
		instance.RScore = correlationCoefficient( target.Data, feValue.Data )
		a, b := leastSquare( feValue.Data, target.Data )
		instance.Diff.Score, instance.Diff.Ave = diffScore( feValue.Data, target.Data, a, b )
		instance.Diff.A = a
		instance.Diff.B = b
		result = append( result, instance )
	}

	res <- result
}
