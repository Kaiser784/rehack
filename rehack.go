package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli/v2"
)

// This is in testing Period
func installScripts() {
	r := bufio.NewReader(os.Stdin)
	var s string
	for {
		fmt.Fprintf(os.Stderr, "Install full or base (Full/base/exit): ")
		s, _ = r.ReadString('\n')
		s = strings.ToLower(strings.TrimSpace(s))
		if s == "" {
			continue
		}

		absPath, _ := filepath.Abs("./setup-scripts")
		switch s {
		case "full":
			fileList, _ := os.ReadDir(absPath)
			for _, v := range fileList {
				out, err := exec.Command("/bin/sh", absPath+"/"+v.Name()).CombinedOutput()
				if err != nil {
					log.Fatal("Err: ", err)
				} else {
					fmt.Println(string(out))
				}

			}
			fmt.Println("Installed files successfully")
			return
		case "base":
			out, err := exec.Command("/bin/sh", absPath+"/base.sh").CombinedOutput()
			if err != nil {
				log.Fatal("Err: ", err)
			} else {
				fmt.Println(string(out))
			}
			return
		case "exit":
			fmt.Println("I am exit")
			return
		}
	}
}

func main() {

	app := &cli.App{
		Name:                 "rehack",
		Usage:                "Tools on go",
		Version:              "0",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Used to install the base code",
				After: func(ctx *cli.Context) error {
					installScripts()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}
