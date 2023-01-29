package modules_separator

import (
	"bufio"
	"fmt"
	"github.com/spf13/afero"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var ( // default <nil> values
	DefaultTime       TimeInterface
	DefaultFS         afero.Fs
	DefaultOutWriter  afero.File
	DefaultErrWriter  afero.File
	DefaultTargetPath string
	DefaultLogger     *log.Logger
)

type TimeInterface interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

// RealTime TimeInterface real (default) implementation
type RealTime struct{}

func (RealTime) Now() time.Time                  { return time.Now() }
func (RealTime) Since(s time.Time) time.Duration { return time.Since(s) }

type ModulesSeparator struct {
	FS        afero.Fs
	Time      TimeInterface
	outWriter afero.File
	errWriter   afero.File
	errorLogger *log.Logger
}

func NewModulesSeparator(FS afero.Fs, Time TimeInterface, outWriter afero.File,
	errWriter afero.File, errorLogger *log.Logger) *ModulesSeparator {
	separator := ModulesSeparator{FS, Time, outWriter, errWriter, errorLogger}

	if FS == DefaultFS {
		separator.FS = afero.OsFs{}
	}
	if Time == DefaultTime {
		separator.Time = RealTime{}
	}
	if outWriter == DefaultOutWriter {
		separator.outWriter = os.Stdout
	}
	if errWriter == DefaultErrWriter {
		separator.errWriter = os.Stderr
	}
	if errorLogger == DefaultLogger {
		separator.errorLogger = DefaultLogger
	}

	return &separator
}

func (separator *ModulesSeparator) logError(msg string) {
	if separator.errorLogger != nil {
		separator.errorLogger.Println(msg)
	} else {
		_, err := fmt.Fprintf(separator.errWriter, msg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, msg)
		}
	}
}

func (separator *ModulesSeparator) ProceedInputFile(filename string, targetPath string) (
	lineNumbers int, numberOfModules int, duration time.Duration) {

	var err error
	var file afero.File
	if targetPath == DefaultTargetPath {
		targetPath, err = filepath.Abs("./") // current opened dir (NOT runner dir)
		if err != nil {
			separator.logError(fmt.Sprintf("ProceedInputFile: While getting cur dir: %v", err))
			os.Exit(1)
		}
	}
	file, err = separator.FS.Open(filename)
	if err != nil {
		separator.logError(fmt.Sprintf("ProceedInputFile: While opening input file: %v", err))
		os.Exit(1)
	}
	defer func() {
		if err = file.Close(); err != nil {
			separator.logError(fmt.Sprintf("ProceedInputFile: While closing input file:  %v", err))
		}
	}()
	scanner := bufio.NewScanner(file)

	return separator.splitModules(*scanner, targetPath)
}

func (separator *ModulesSeparator) splitModules(scanner bufio.Scanner, targetPath string) (
	lineNumbers int, numberOfModules int, duration time.Duration) {

	start := separator.Time.Now()
	outputFilename := ""
	var outputFileCtx []byte
	saveModuleResult := make(chan error)
	numberOfModules = 0
	skipping := false
	lineNumbers = 0
	for scanner.Scan() {
		line := scanner.Bytes()
		lineStr := string(line)
		lineNumbers++
		if strings.HasPrefix(lineStr, "//") && strings.HasSuffix(lineStr, ">") {
			skipping = true
		} else if strings.HasPrefix(lineStr, "//") && !strings.HasSuffix(lineStr, ">") {
			outputFileCtx = append(outputFileCtx, line...)
			skipping = false
			outputFilename = strings.TrimPrefix(lineStr, "// ")
		} else if outputFilename != "" && lineStr == "}" {
			outputFileCtx = append(append(outputFileCtx, append([]byte("\n"), line...)...), []byte("\n")...)
			go separator.saveModule(outputFilename, outputFileCtx, targetPath, saveModuleResult)
			numberOfModules++
			outputFileCtx = []byte{}
			outputFilename = ""
			skipping = true
		} else if !strings.HasPrefix(lineStr, "//") && !skipping {
			outputFileCtx = append(outputFileCtx, append([]byte("\n"), line...)...)
		} else if !skipping {
			separator.logError(fmt.Sprintf("splitModules: unparsed line number: %d line: %s", lineNumbers, lineStr))
		}
	}
	for i := numberOfModules; i > 0; i-- {
		if err := <-saveModuleResult; err != nil {
			separator.logError(fmt.Sprintf("splitModules: %v", err))
		}
	}
	close(saveModuleResult)
	duration = separator.Time.Since(start)

	return lineNumbers, numberOfModules, duration
}

func (separator *ModulesSeparator) saveModule(
	outputFilename string, outputFileCtx []byte, targetPath string, errChan chan<- error) {
	outputDirPath := filepath.Join(targetPath, "cs-modules")
	var err error
	if _, err = separator.FS.Stat(outputDirPath); os.IsNotExist(err) {
		var dirMod uint64
		if dirMod, err = strconv.ParseUint("0775", 8, 32); err == nil {
			err = separator.FS.Mkdir(outputDirPath, os.FileMode(dirMod))
		}
	}
	if err != nil && !os.IsExist(err) {
		separator.logError(fmt.Sprintf("saveModule: %v", err))
		os.Exit(1)
	}
	outputFilepath := filepath.Join(outputDirPath, outputFilename+".cs")
	errChan <- afero.WriteFile(separator.FS, outputFilepath, outputFileCtx, 0644)
}
