package main

import (
	"time"

	"github.com/rwynn/monstache/monstachemap"
)

// var infoLog = log.New(os.Stdout, "INFO ", log.Flags())

//Map function plugins must implement a function named "Map" with the following signature
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {

	doc := input.Document

	createdAt := doc["created_at"]
	if createdAt != nil {
		createdAtStr, ok := createdAt.(string)
		if ok {
			createdAtt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
			if err == nil {
				doc["created_at"] = createdAtt
			}

		}
	}

	uid := doc["uid"]
	if uid != nil {
		uidStr, ok := uid.(string)
		if ok {
			rs := []rune(uidStr)
			length := len(rs)

			doc["uid_last_6char"] = string(rs[length-6:])

		}
	}

	output = &monstachemap.MapperPluginOutput{Document: doc}
	// infoLog.Printf("plugin Map input=%v output = %v", input, output)
	return output, nil
}

//Filter function
// func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error) {
// 	infoLog.Printf("plugin Filter input=%v ", input)
// 	return true, nil
// }

//Pipeline function
// func Pipeline(ns string, changeStream bool) (stages []interface{}, err error) {
// 	infoLog.Printf("plugin Pipeline ns=%v changeStream=%v stages = %v", ns, changeStream, stages)
// 	return
// }

//Process function
// func Process(input *monstachemap.ProcessPluginInput) (err error) {
// 	infoLog.Printf("plugin  Process input= %v", input)
// 	return
// }
