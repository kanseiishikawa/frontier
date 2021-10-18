package analyze

import (
	"frontier/config"
)

type AnalyzeData struct {
	Name string
	Data []float32
}

func BitSearch( N int ) ( [][]int ) {
	var result [][]int
	
	for bits := 1; bits <(1<<uint(N)); bits++ {
		instance := []int{}

		for i := 0; i < N; i++ {
			if (bits>>uint(i))&1 == 1 {
				instance = append( instance, i )
			}
		}

		result = append( result, instance )
		
	}

	return result
}

func TargetDataCreate( target *Target ) ( []AnalyzeData ) {
	var result []AnalyzeData

	for i := 0; i < len( target.Info ); i++ {		
		var instance AnalyzeData
		instance.Name = target.Info[i].Name

		for _, v := range target.Value {
			instance.Data = append( instance.Data, v[i] )
		}

		result = append( result, instance )
	}

	return result
}

func ValueDataCreate( feValue *FeValue, bit []int, confData *config.ConfData, re chan []AnalyzeData ) {
	check := map[string][]float32{}
	n := len( bit )
	name := " "

	for i, b := range bit {
		name += feValue.Info[b].Name

		if i != len( bit ) - 1 {
			name += ","
		}
	}

	var valueName string
	
	if n == 1 {
		valueName = "None" + name
		check[valueName] = []float32{}
	}
	
	for i := 0; i < len( feValue.Value ); i++ {
		var instance []float32
		//fmt.Println( feValue.Value[i], bit )
		for _, b := range bit {
			instance = append( instance, feValue.Value[i][b] )
		}

		if n == 1 {
			check[valueName] = append( check[valueName], instance[0] )
		} else {
			paramList := confData.Three
			fu := ThreeCluc
			
			if n == 2 {
				fu = TwoCluc
				paramList = confData.Two
			}
			
			for _, t := range paramList {
				if t.Use {
					valueName = t.Name + name
					v, _ := fu( instance, t.Name )
					check[valueName] = append( check[valueName], v )
				}
			}		
		}
	}
	
	var result []AnalyzeData

	for key := range check {
		var analyzeData AnalyzeData
		analyzeData.Name = key
		analyzeData.Data = check[key]
		result = append( result, analyzeData )
	}
	
	re <- result
}
