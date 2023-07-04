package main

import (
	"fmt"
	"strings"
)

//преобразует строку в число
func StringToInt(str string) int {
	allNum := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	result := 0
	for _, value := range str {
		if valueMap, inMap := allNum[string(value)]; inMap {
			result *= 10
			result += valueMap
		} else {
			return result
		}
	}

	return result
}

// генерирует sql для обновления только некоторых полей таблицы todo
func GenerateSql(t TodoUpdate) string {
	querySlice := []string{}

	if t.Title != nil{
		querySlice = append(querySlice, fmt.Sprintf(`title='%s'`, *t.Title))
	}

	if t.Description != nil{
		querySlice = append(querySlice, fmt.Sprintf(`description='%s'`, *t.Description))
	}

	if t.IsReady != nil{
		querySlice = append(querySlice, fmt.Sprintf(`is_ready='%t'`, *t.IsReady))
	}

	if t.ListId != nil{
		querySlice = append(querySlice, fmt.Sprintf(`list_id='%d'`, *t.ListId))
	}

	query := `UPDATE public.todo SET `

	res := query + strings.Join(querySlice, ", ") +  ` WHERE id=$1`

	return res
}