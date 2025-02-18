/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// flaskCmd represents the flask command
var flaskCmd = &cobra.Command{
	Use:   "flask",
	Short: "Create a flask template project",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			fmt.Println("Invalid arguments")
			return
		}

		projectPath := ""
		homePath := ""

		switch runtime.GOOS {
		case "windows":
			homePath = os.Getenv("USERPROFILE") + "\\flaskProject"
		case "linux":
			homePath = os.Getenv("HOME") + "/flaskProject"
		default:
			fmt.Println("Unknown OS:", runtime.GOOS)
		}

		suggests := []string{
			"./flaskProject",
			homePath,
		}

		prompt := &survey.Input{
			Message: "Enter the project path to create the flask project: ",
			Suggest: func(toComplete string) []string {
				return suggests
			},
		}

		survey.AskOne(prompt, &projectPath)
		absPath, err := filepath.Abs(projectPath)

		if err != nil {
			fmt.Println("Error getting absolute path:", err)
			return
		}

		err = os.MkdirAll(absPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		git.PlainClone(absPath, false, &git.CloneOptions{
			URL: "https://github.com/yassine20011/flaskTemplate.git",
		})

		confirmPrompt := &survey.Confirm{
			Message: "Do you want to create a virtual environment?",
			Default: true,
		}

		createVenv := false
		survey.AskOne(confirmPrompt, &createVenv)

		if createVenv {
			venvPrompt := &survey.Input{
				Message: "Enter the virtual environment name: ",
				Default: "venv",
			}

			venvName := ""
			survey.AskOne(venvPrompt, &venvName)
			command := exec.Command("python", "-m", "venv", venvName)
			fmt.Println("Creating virtual environment at:", absPath)
			command.Dir = absPath
			err := command.Run()
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("Project created at:", absPath)
	},
}

func init() {
	rootCmd.AddCommand(flaskCmd)
}
