package analyze


type ScoreData struct {
	Score float32
	TargetName string
	FeValueName string
}

func AnalyzeScore( targets []AnalyzeData, feValue AnalyzeData, res chan []ScoreData ) {
	var result []ScoreData
	
	for _, target := range targets {
		var instance ScoreData
		instance.TargetName = target.Name
		instance.FeValueName = feValue.Name
		
		tfConv := cov( target.Data, feValue.Data )
		targetStd := std( target.Data )
		feValueStd := std( feValue.Data )
		instance.Score = tfConv / ( targetStd * feValueStd )

		if instance.Score < 0 {
			instance.Score *= -1
		}

		result = append( result, instance )
	}

	res <- result
}
