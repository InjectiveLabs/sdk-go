# Contributing to Injective SDK Go

Thank you for your interest in contributing to the Injective Protocol Golang SDK! We welcome contributions from the community.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)

## Code of Conduct

Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md) to help maintain a welcoming and inclusive community.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone git@github.com:YOUR_USERNAME/sdk-go.git
   cd sdk-go
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/InjectiveLabs/sdk-go.git
   ```

## Development Setup

### Prerequisites

- Go 1.21 or later
- Make
- Git

### Installation

```bash
go mod download
```

### Running Tests

```bash
make tests
```

### Running Linter

```bash
make lint
```

## How to Contribute

### Reporting Bugs

- Check if the bug has already been reported in [Issues](https://github.com/InjectiveLabs/sdk-go/issues)
- If not, create a new issue with a clear title and description
- Include steps to reproduce the bug
- Add relevant code samples or error messages

### Suggesting Enhancements

- Open a new issue describing your enhancement
- Explain why this enhancement would be useful
- Provide examples of how it would work

### Code Contributions

1. Create a new branch from `dev`:
   ```bash
   git checkout -b feature/your-feature-name
   ```
2. Make your changes
3. Write or update tests as needed
4. Run tests and linter
5. Commit your changes with a descriptive message
6. Push to your fork and submit a pull request

## Pull Request Process

1. Ensure your code follows the project's coding standards
2. Update documentation if needed
3. Add tests for new functionality
4. Ensure all tests pass
5. Update the README.md if necessary
6. Request review from maintainers

### Commit Message Guidelines

- Use clear, descriptive commit messages
- Start with a verb in the present tense (e.g., "Add", "Fix", "Update")
- Reference issues when applicable (e.g., "Fix #123")

Examples:
```
feat: add new market query method
fix: resolve connection timeout issue
docs: update README with new examples
```

## Coding Standards

- Follow Go best practices and idioms
- Use `gofmt` to format your code
- Write clear comments for exported functions
- Keep functions focused and concise
- Handle errors appropriately

## Questions?

If you have questions, feel free to:
- Open an issue
- Reach out on [Twitter](https://twitter.com/InjectiveLabs)
- Visit [injective.com](https://injective.com)

Thank you for contributing! 🚀
