package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"

	data "app/internal"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go Package Manager"
	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install package",
			Action: func(c *cli.Context) error {
				goGet(c.Args().Get(0))
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove package",
			Action: func(c *cli.Context) error {
				goDel(c.Args().Get(0))
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		println("Greetings")
		return nil
	}
	app.Run(os.Args)
}

func goGet(name string) *exec.Cmd {
	if goList(name) {
		log.Fatal(data.Data[name] + " package is already installed")
	}
	if _, ok := data.Data[name]; ok {
		cmd := exec.Command("go", "get", data.Data[name])
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
	}
	return nil
}

func goDel(name string) *exec.Cmd {
	if !goList(name) {
		log.Fatal(data.Data[name] + " package is not installed")
	}
	if _, ok := data.Data[name]; ok {
		cmd := exec.Command("go", "mod", "tidy")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(output))
		fmt.Println(data.Data[name] + " is removed")
	}
	return nil
}

func goList(name string) bool {
	cmd := exec.Command("go", "list", data.Data[name])
	err := cmd.Run()
	return err == nil
}
