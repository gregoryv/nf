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
		w := os.Stdout
		cli.WriteUsageTo(w)
		fmt.Println("Files")
		files, _ := assets.ReadDir("assets")
		for _, f := range files {
			fmt.Fprintln(w, "   ", f.Name())
		}
		os.Exit(0)

	case !cli.Ok():
		log.Fatal(cli.Error())

	}

	err := writeFile(os.Stdout, filename)
	if err != nil {
		log.Fatal(err)
	}
}

func writeFile(w io.Writer, filename string) error {
	content, err := assets.ReadFile(path.Join("assets", filename))
	if err != nil {
		return err
	}
	w.Write(content)
	return nil
}

//go:embed "assets/*"
var assets embed.FS
