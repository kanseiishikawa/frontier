package analyze


import (
	"os"
	"sort"
	"strconv"
)

func DataWrite( scores []ScoreData ) ( error ){
	fileName := scores[0].TargetName + "-score.txt"
	sort.Slice( scores, func( i, j int ) bool {
		return scores[i].Score > scores[j].Score
	})

	f, err := os.Create( "score_results/" + fileName )

	if err != nil {
		return err
	}

	for _, score := range scores {
		f.WriteString( score.FeValueName + " " )
		f.WriteString(  strconv.FormatFloat( float64( score.Score ), 'f', -1, 32 ) + "\n" )
	}

	f.Close()
	return nil
}
