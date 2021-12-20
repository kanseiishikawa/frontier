package file


import (
	"os"
	"sort"
	"strconv"

	"frontier/analyze"
)

func RDataWrite( scores []analyze.ScoreData ) ( error ){
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

func DiffDataWrite( scores []analyze.ScoreData ) ( error ){
	fileName := scores[0].TargetName + "-Diffscore.txt"
	sort.Slice( scores, func( i, j int ) bool {
		return scores[i].Diff.Score < scores[j].Diff.Score
	})

	f, err := os.Create( "score_results/" + fileName )

	if err != nil {
		return err
	}

	for _, score := range scores {
		f.WriteString( score.FeValueName + " " )
		f.WriteString( " a:" + strconv.FormatFloat( float64( score.Diff.A ), 'f', -1, 32 ) +
			" b:" + strconv.FormatFloat( float64( score.Diff.B ), 'f', -1, 32 ) +
			" score:" + strconv.FormatFloat( float64( score.Diff.Score ), 'f', -1, 32 ) +
			" ave:" + strconv.FormatFloat( float64( score.Diff.Ave ), 'f', -1, 32 ) + "\n" )
	}

	f.Close()
	return nil
}

