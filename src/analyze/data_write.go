package analyze


import (
	"os"
	"sort"
	"strconv"
)

func RDataWrite( scores []ScoreData ) ( error ){
	fileName := scores[0].TargetName + "-Rscore.txt"
	sort.Slice( scores, func( i, j int ) bool {
		return scores[i].RScore > scores[j].RScore
	})

	f, err := os.Create( "score_results/" + fileName )

	if err != nil {
		return err
	}

	for _, score := range scores {
		f.WriteString( score.FeValueName + " " )
		f.WriteString( strconv.FormatFloat( float64( score.RScore ), 'f', -1, 32 ) + "\n" )
	}

	f.Close()
	return nil
}

func DiffDataWrite( scores []ScoreData ) ( error ){
	fileName := scores[0].TargetName + "-Diffscore.txt"
	sort.Slice( scores, func( i, j int ) bool {
		return scores[i].DiffScore < scores[j].DiffScore
	})

	f, err := os.Create( "score_results/" + fileName )

	if err != nil {
		return err
	}

	for _, score := range scores {
		f.WriteString( score.FeValueName + " " )
		f.WriteString( " a:" + strconv.FormatFloat( float64( score.DiffA ), 'f', -1, 32 ) +
			" b:" + strconv.FormatFloat( float64( score.DiffB ), 'f', -1, 32 ) +
			" score:" +  strconv.FormatFloat( float64( score.DiffScore ), 'f', -1, 32 ) + "\n" )
	}

	f.Close()
	return nil
}

