# Contributing to EzSetup

Thank you for your interest in contributing to EzSetup! This document provides guidelines and steps for contributing.

## Getting Started

1. Fork the repository
2. Clone your fork:

```bash
git clone https://github.com/YOUR-USERNAME/EzSetup.git
cd EzSetup
go mod tidy
```

## Development Guidelines

### Code Style

- Follow standard Go formatting guidelines
- Use `gofmt` to format your code
- Write meaningful commit messages
- Add comments for complex logic
- Use descriptive variable names

### Project Structure

```
EzSetup/
├── cmd/
│   ├── root.go
│   └── flask.go
└── main.go
```
### Testing

- Write tests for new features
- Run tests before submitting PR:
  ```bash
  go test ./...
  ```

## Good First Issues

Look for issues labeled with:
- `good-first-issue`
- `documentation`
- `bug-fix-easy`
- `help-wanted`

Current good first issues:

- Adding new project templates
- Improving documentation
- Adding more test cases
- Enhancing error messages

## Pull Request Process

1. Create a new branch:
```bash
git checkout -b feature/your-feature-name
```

2. Make your changes following our coding standards

3. Update documentation if needed

4. Push your changes:
   ```bash
   git push origin feature/your-feature-name
   ```

5. Open a PR with:
   - Clear description of changes
   - Screenshots if UI changes
   - Reference to related issues

## Reporting Issues

Create an issue with:
- Clear description
- Steps to reproduce
- Expected vs actual behavior
- System information

## Code Review Process

- At least one maintainer must approve
- Code must follow style guidelines
- Documentation must be updated
