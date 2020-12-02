package main

import (
	"log"
	"os"

	"github.com/rwynn/monstache/monstachemap"
)

var infoLog = log.New(os.Stdout, "INFO ", log.Flags())

func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	doc := input.Document
	for k, v := range doc {
		infoLog.Printf("Starting http server k=%s v = %v", k, v)
	}
	output = &monstachemap.MapperPluginOutput{Document: doc}
	return
}

func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error) {
	return
}

func Pipeline(ns string, changeStream bool) (stages []interface{}, err error) {
	return
}
func Process(input *monstachemap.ProcessPluginInput) (err error) {
	return
}
