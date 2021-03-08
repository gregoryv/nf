package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/gregoryv/cmdline"
)

func main() {
	var (
		cli       = cmdline.NewParser(os.Args...)
		help      = cli.Flag("-h, --help")
		_         = cli.Required("FILES...").String("")
		filenames = cli.Args()
	)
	log.SetFlags(0)
	switch {
	case help:
		writeUsage(os.Stdout, cli)
		os.Exit(0)
	case !cli.Ok():
		log.Fatal(cli.Error())
	case len(filenames) == 0:
		log.Fatal("missing files")
	}
	err := writeFiles(filenames)
	if err != nil {
		log.Fatal(err)
	}
}

func writeFiles(filenames []string) error {
	for _, filename := range filenames {
		if fi, _ := os.Stat(filename); fi != nil {
			log.Println(" skip", filename)
			continue
		}
		err := writeFile(os.Stdout, filename)
		if err != nil {
			return err
		}
		switch path.Ext(filename) {
		case ".sh":
			err := os.Chmod(filename, 0744)
			if err != nil {
				return err
			}
		}
		log.Println("write", filename)
	}
	return nil
}

func writeFile(w io.Writer, filename string) error {
	content, err := assets.ReadFile(path.Join("assets", filename))
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(content)
	return nil
}

func writeUsage(w io.Writer, cli *cmdline.Parser) {
	cli.WriteUsageTo(w)
	fmt.Println("Files")
	files, _ := assets.ReadDir("assets")
	for _, f := range files {
		fmt.Fprintln(w, "   ", f.Name())
	}
}

//go:embed "assets/*"
var assets embed.FS
