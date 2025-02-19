/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
	"github.com/go-git/go-git/v5"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// flaskCmd represents the flask command
var flaskCmd = &cobra.Command{
	Use:   "flask",
	Short: "Create a flask template project",

	Run: func(cmd *cobra.Command, args []string) {
		projectPath := ""
		homePath := ""

		switch runtime.GOOS {
		case "windows":
			homePath = os.Getenv("USERPROFILE") + "\\flaskProject"
		case "linux":
			homePath = os.Getenv("HOME") + "/flaskProject"
		default:
			pterm.Error.Println("Unknown OS:", runtime.GOOS)
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
			pterm.Error.Println("Error getting absolute path:", err)
			return
		}

		err = os.MkdirAll(absPath, 0755)
		if err != nil {
			pterm.Error.Println("Error creating directory:", err)
			return
		}

		confirmPrompt := &survey.Confirm{
			Message: "Do you want to create a virtual environment?",
			Default: true,
		}

		createVenv := false
		survey.AskOne(confirmPrompt, &createVenv)

		venvPrompt := &survey.Input{
			Message: "Enter the virtual environment name: ",
			Default: "venv",
		}

		venvName := ""
		survey.AskOne(venvPrompt, &venvName)

		confirmPrompt = &survey.Confirm{
			Message: "Do you want to install dependencies?",
			Default: true,
		}

		installDeps := false
		survey.AskOne(confirmPrompt, &installDeps)

		pterm.Println()
		pterm.Println("🚀 Setting up your Flask Project...")
		pterm.Println()

		pterm.Info.Println("Cloning flask template project...")
		git.PlainClone(absPath, false, &git.CloneOptions{
			URL: "https://github.com/yassine20011/flaskTemplate.git",
		})

		if createVenv {
			command := exec.Command("python", "-m", "venv", venvName)
			pterm.Info.Printf("Creating virtual environment at %s\n", absPath)
			command.Dir = absPath
			err := command.Run()
			if err != nil {
				pterm.Error.Println(err)
				return
			}

			pterm.Success.Println("Virtual environment created successfully")

			if installDeps {
				pterm.Info.Println("Installing dependencies...")
				command = exec.Command("pip", "install", "-r", "requirements.txt")
				command.Dir = absPath
				err = command.Run()
				spinner, _ := pterm.DefaultSpinner.Start("Installing dependencies...")
				if err != nil {
					pterm.Error.Println(err)
					return
				}
				spinner.Success("Dependencies installed successfully")
			}
		}

		pterm.Println()
		pterm.Bold.Println("📌 Next Steps:")
		pterm.Println()
		pterm.Info.Println("1- Navigate to the project: cd " + absPath)
		switch runtime.GOOS {
		case "windows":
			pterm.Info.Println("2- Activate the virtual environment: .\\venv\\Scripts\\activate")
		case "linux":
			pterm.Info.Println("2- Activate the virtual environment: source venv/bin/activate")
		default:
			pterm.Error.Println("Unknown OS:", runtime.GOOS)
		}
		pterm.Println()
		pterm.Println("🎯 Happy Coding! 🚀")
	},
}

func init() {
	rootCmd.AddCommand(flaskCmd)
}
