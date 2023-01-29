package main

import (
	"log"
	"os"

	s "github.com/madzohan/ilspy_utils/pkg/modules_separator"
)

func main() {
	var (
		errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	)
	args := os.Args[1:]
	if  len(args) == 0 {
		errorLogger.Fatal("please set input file, example: `./separator -Module-.cs`")
	} else if len(args) > 2 {
		errorLogger.Fatal(
			"maximum 2 arguments allowed, example: `./separator -Module-.cs ./output_path_for_separated_modules`")
	}
	targetPath := s.DefaultTargetPath
	filename := args[0]
	if len(args) == 3 {
		targetPath = args[2]
	}
	separator := s.NewModulesSeparator(
		s.DefaultFS, s.DefaultTime, s.DefaultOutWriter, s.DefaultErrWriter, errorLogger)
	lineNumbers, numberOfModules, duration := separator.ProceedInputFile(filename, targetPath)
	infoLogger.Printf("Processed %d lines: separated %d modules, took %.4f seconds\n",
		lineNumbers, numberOfModules, duration.Seconds())
}
