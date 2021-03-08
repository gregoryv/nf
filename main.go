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
		cli  = cmdline.NewParser(os.Args...)
		help = cli.Flag("-h, --help")

		filename = cli.Required("FILE").String("")
	)

	log.SetFlags(0)

	switch {
	case help:
		writeUsage(os.Stdout, cli)
		os.Exit(0)

	case !cli.Ok():
		log.Fatal(cli.Error())

	}

	err := writeFile(os.Stdout, filename)
	if err != nil {
		log.Fatal(err)
	}
	switch path.Ext(filename) {
	case ".sh":
		err := os.Chmod(filename, 0744)
		if err != nil {
			log.Fatal(err)
		}
	}
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
