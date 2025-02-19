# EzSetup 🚀

A command-line tool to quickly set up development environments and project templates.

# What is ezSetup?

ezSetup is a lightweight yet robust tool that:

- Creates the structure for various languages.
- Installs required dependencies automatically.
- Initializes Git using the appropriate .gitignore.
- Sets up environment variables effortlessly.
- set up a project using a costum template using YAML.

# Why Go and Cobra?

I built this project with the goal of learning Golang while solving a real problem for beginners or even just to speed up the process for some people.

# Why did I create this?

As a student, I noticed that many beginners struggle when setting up their projects. Some get stuck on dependencies, while others aren’t sure how to structure their code. ezSetup helps them dive straight into coding without the setup hassle!

## Features

- 🌟 Interactive CLI with user-friendly prompts
- 🐍 Flask project setup with:
  - Virtual environment creation
  - Automatic dependency installation
  - Project structure from template
- 💻 Cross-platform support (Windows & Linux)

## Contributing

### Setting Up Development Environment

```bash
# Fork the repository and clone it
git clone https://github.com/yourusername/EzSetup.git

# Navigate to the project directory
cd EzSetup

# Install dependencies
go mod tidy
```

### Example: Adding a New Command

Here's an example of how to add a new command called "django" that sets up a Django project:

1. Create a new file `cmd/django.go`:

```go
package cmd

import (
    "github.com/spf13/cobra"
    "github.com/AlecAivazis/survey/v2"
    "github.com/pterm/pterm"
)

var djangoCmd = &cobra.Command{
    Use:   "django",
    Short: "Create a Django project template",
    Run: func(cmd *cobra.Command, args []string) {
        // 1. Ask for project name
        projectName := ""
        prompt := &survey.Input{
            Message: "Enter the project name: ",
            Default: "djangoProject",
        }
        survey.AskOne(prompt, &projectName)

        // 2. Your command implementation
        pterm.Info.Println("Setting up Django project:", projectName)

        // Add your Django setup logic here
    },
}

func init() {
    rootCmd.AddCommand(djangoCmd)
}
```

2. Test your command:

```bash
go run main.go django
```

### Best Practices

- Follow the existing code structure
- Use the survey package for interactive prompts
- Use pterm for consistent output styling
- Add proper error handling
- Update this README when adding new features

### Pull Request Process

1. Create a new branch for your feature
2. Write clear commit messages
3. Test your changes
4. Update documentation
5. Submit a pull request with a description of your changes

## Usage

### Flask Project Setup

Create a new Flask project with a single command:

```bash
ezsetup flask
```

This will:

1. Prompt for project location
2. Create virtual environment (optional)
3. Install dependencies (optional)
4. Set up project structure from template

## Project Structure

```
EzSetup/
├── cmd/
│   ├── root.go
│   └── flask.go
└── main.go
```

## Dependencies

- github.com/spf13/cobra
- github.com/AlecAivazis/survey/v2
- github.com/go-git/go-git/v5
- github.com/pterm/pterm

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
