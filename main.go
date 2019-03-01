package main

import (
	"github.com/Pingze-github/go-package-plantuml/codeanalysis"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"regexp"
	"strings"
	"path"
)

func main() {

	log.SetLevel(log.InfoLevel)

	var opts struct {
		CodeDir string `long:"codedir" description:"Code dir path" required:"true"`
		GopathDir string `long:"gopath" description:"GOPATH dir path"`
		OutputFile string `long:"outputfile" description:"output file path" required:"true"`
		IgnoreDirs []string `long:"ignoredir" description:"ignore dir paths"`
	}

	if len(os.Args) == 1 {
		fmt.Println("Example:\n" +
			os.Args[0] + " --codedir /appdev/gopath/src/github.com/contiv/netplugin --gopath /appdev/gopath --outputfile  /tmp/result")
		os.Exit(1)
	}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	if opts.CodeDir == "" {
		panic("Codedir cannot be null")
		os.Exit(1)
	}

	if opts.GopathDir == "" {
		gopathEnv := os.Getenv("GOPATH")
		if gopathEnv != "" {
			opts.GopathDir = gopathEnv
			fmt.Println("use default $GOPATH: " + gopathEnv)
		} else {
			panic("GOPATH cannot be null")
			os.Exit(1)
		}
	}

	if ! strings.HasPrefix(opts.CodeDir, opts.GopathDir) {
		panic(fmt.Sprintf("code dir %s must be subdir of GOPATH dir", opts.CodeDir, opts.GopathDir))
		os.Exit(1)
	}

	for _, dir := range opts.IgnoreDirs {
		if ! strings.HasPrefix(dir, opts.CodeDir){
			panic(fmt.Sprintf("ignore dirs %s must be subdirs of code dir", dir, opts.CodeDir))
			os.Exit(1)
		}
	}

	config := codeanalysis.Config{
		CodeDir: FormatSlash(opts.CodeDir),
		GopathDir : FormatSlash(opts.GopathDir),
		VendorDir : path.Join(opts.CodeDir, "vendor"),
		OutputFile : FormatSlash(opts.OutputFile),
		IgnoreDirs: opts.IgnoreDirs,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile(config.OutputFile)

}

func FormatSlash(pathstr string) string {
	reg, _ := regexp.Compile("\\\\")
	return reg.ReplaceAllString(pathstr, "/")
}
