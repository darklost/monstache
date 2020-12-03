package main

import (
	"log"
	"os"

	"github.com/rwynn/monstache/monstachemap"
)

var infoLog = log.New(os.Stdout, "INFO ", log.Flags())

func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	infoLog.Printf("plugin Map input=%v output = %v", input, output)

	return
}

func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error) {
	infoLog.Printf("plugin Filter input=%v output = %v", input, keep)
	return 
}

func Pipeline(ns string, changeStream bool) (stages []interface{}, err error) {
	infoLog.Printf("plugin Pipeline ns=%v changeStream=%v output = %v", ns, changeStream, stages)
	return
}
func Process(input *monstachemap.ProcessPluginInput) (err error) {
	infoLog.Printf("plugin  Process input= %v", input)
	return
}
