package main

import (
	"flag"
	"fmt"
	"os"

	"frontier/config"
	"frontier/analyze"
)

func main() {
	t := flag.String( "target", "target.json", "target json file" )	
	f := flag.String( "fevalue", "fevalue.json", "fevalue json file" )
	flag.Parse()

	//設定ファイルの読み込み
	confData, err := config.ConfigRead()
	fmt.Println( "config read finish" )
	
	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}

	//targetの読み込み
	target, err := analyze.TargetRead( *t )
	fmt.Println( "target read finish" )

	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}

	//特徴量の読み込み
	feValue, err := analyze.FeValueRead( *f )
	fmt.Println( "fevalue read finish" )
	
	if err != nil {
		fmt.Println( err )
		os.Exit( 1 )
	}
	
	feValueN := len( feValue.Info )

	//bit全探索で全てのパターンを網羅
	BitSearchList := analyze.BitSearch( feValueN )
	//並列処理で受け取るchannel変数の定義
	chanAnalyzeData := make( chan []analyze.AnalyzeData )

	//時間がかかるので並列処理
	for i := 0; i < len( BitSearchList ); i++ {
		go analyze.ValueDataCreate( &feValue, BitSearchList[i], &confData, chanAnalyzeData )
	}

	var analyzeDataList []analyze.AnalyzeData
	//結果を一つづつ受け取る
	for i := 0; i < len( BitSearchList ); i++ {
		res := <-chanAnalyzeData
		analyzeDataList = append( analyzeDataList, res... )
		fmt.Println( "BitSearch", len( BitSearchList ) - i )
	}

	targetDataList := analyze.TargetDataCreate( &target )
	chanScoreData := make( chan []analyze.ScoreData )

	scores := make( map[string][]analyze.ScoreData )

	for _, t := range target.Info {
		scores[t.Name] = []analyze.ScoreData{}
	}
	
	for i := 0; i < len( analyzeDataList ); i++ {
		go analyze.AnalyzeScore( targetDataList, &analyzeDataList[i], chanScoreData )
	}
	
	for i := 0; i < len( analyzeDataList ); i++ {
		res := <-chanScoreData
		fmt.Println( "ScoreData", len( analyzeDataList ) - i )
		
		for _, v := range res {
			scores[v.TargetName] = append( scores[v.TargetName], v )
		}
	}

	//結果を作成するディレクトリを作成する
	os.Remove( "score_results" )
	os.Mkdir( "score_results", 0777 )
	
	for _, score := range scores {
		err = analyze.RDataWrite( score )
		
		if err != nil {
			fmt.Println( err )
			os.Exit( 1 )
		}

		err = analyze.DiffDataWrite( score )
		
		if err != nil {
			fmt.Println( err )
			os.Exit( 1 )
		}		
	}


	fmt.Println( "finish!" )
}
