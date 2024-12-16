package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"

	data "github.com/einsy-dev/GoPM/internal"
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
				if c.Args().Get(0) == "" {
					log.Fatal("package name is required")
				}
				goGet(c.Args())
				return nil
			},
		},
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "remove unused packages",
			Action: func(c *cli.Context) error {
				goDel()
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list package",
			Action: func(c *cli.Context) error {
				for i, v := range data.Data {
					fmt.Println(i, " ", v)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func goGet(name []string) *exec.Cmd {
	for _, v := range name {
		if goList(v) {
			fmt.Println(data.Data[v] + " package is already installed")
			continue
		}
		if _, ok := data.Data[v]; ok {
			cmd := exec.Command("go", "get", "-u", "-v", data.Data[v])
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(output))
		}
	}
	return nil
}

func goDel() *exec.Cmd {

	cmd := exec.Command("go", "mod", "tidy")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
	fmt.Println("Unused dependencies is removed")

	return nil
}

func goList(name string) bool {
	cmd := exec.Command("go", "list", data.Data[name])
	err := cmd.Run()
	return err == nil
}
