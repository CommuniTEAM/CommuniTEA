# CommuniTEA

[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)

Bringing your local community together over a cuppa 🍵

## Table of Contents

- [CommuniTEA](#communitea)
  - [Table of Contents](#table-of-contents)
  - [Working with Git Hooks](#working-with-git-hooks)
    - [Installing Pre-Commit](#installing-pre-commit)
    - [Pre-Commit Rulesets](#pre-commit-rulesets)
      - [Commit Message Format](#commit-message-format)
      - [ESLint TypeScript Format](#eslint-typescript-format)
      - [Golang Format](#golang-format)
    - [Using Git Hooks in VSCode](#using-git-hooks-in-vscode)

## Working with Git Hooks

Any commits to the repository must conform to the formatting rulesets for the following:
- Commit Messages
- ESLint TypeScript
- Golang (TBD)

To enforce these conventions, [pre-commit](https://pre-commit.com/), a Git hook manager, is a required development dependency for the project. You will need to install it on your local machine.

### Installing Pre-Commit

1. Install pre-commit.<br>
**Windows**:
Pre-commit runs off of Python, so you will need pip on your local machine. To install pre-commit, open your terminal and run `pip install pre-commit`.<br>
**Mac or Linux**:
Ensure you have [homebrew](https://brew.sh/) installed on your local machine. To install pre-commit, open your terminal and run `brew install pre-commit`.

2. Install the repository's Git hooks by navigating to the repo's directory in your terminal and running `pre-commit install`.

And that's all! You're now ready to commit to the repository.
For more detailed installation information and troubleshooting, see the [pre-commit documentation](https://pre-commit.com/#install).

### Pre-Commit Rulesets

#### Commit Message Format

All commit messages, regardless of the branch, must adhere to this header standard:<br>
```<Emoji> [Related Jira Issue]: Description of work```<br>
Example: `✅ [TEA-123]: Added unit tests`

These are the allowed emojis and their use-cases:
- ✨: New feature
- 🐞: Bug fix
- ✅: Added or updated tests
- 🚧: Work in progress
- 🔨: Refactored
- 📝: Documentation updated
- 🤝: Merged branches

Special thanks to strdr4605 for their [walkthrough](https://strdr4605.com/commitlint-custom-commit-message-with-emojis) of this configuration!

#### ESLint TypeScript Format

The ESLint configuration for this repository uses the [guide by Airbnb](https://github.com/airbnb/javascript). The installed pre-commit hook will attempt to fix any linting errors found in the staged changes. If it does, the check will fail and you will need to re-stage and commit.

Please be sure to have the [Prettier VSCode extension](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode) installed and format your code often. The included workspace settings for VSCode set Prettier as the default formatter for TypeScript and JavaScript. To format the current file quickly, do a manual save (Windows/Linux: `Ctrl+S`  Mac: `⌘S`).

#### Golang Format

To successfully run the pre-commit hook set up for Go, you must have [golangci-lint](https://golangci-lint.run/) installed on your local machine. See the installation instructions below, or follow the [official documentation](https://golangci-lint.run/usage/install/).

Installing golangci-lint for:
- **Windows**: You must first install [Git for Windows](https://gitforwindows.org/) so that you have Git Bash, as the golangci-lint installation commands cannot be run without a bash terminal. Once you have Git Bash at the ready, run `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2`.
-**Mac**: In your terminal, run `brew install golangci-lint` followed by `brew upgrade golangci-lint`.
-**Linux**: In your terminal, run `sudo snap install golangci-lint`. Alternatively, you can run the same command supplied for Windows users above.

This repository uses a robust golangci-lint configuration built up of over 75 linters, as recommended in the ["Golden Config" by maratori](https://gist.github.com/maratori/47a4d00457a92aa426dbd48a18776322). As a result, the linter is quite strict. It is highly recommended that you enable the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) in VSCode in order to catch the linting errors as they occur.


### Using Git Hooks in VSCode

When commiting changes through VSCode's source control tab, pre-commit will still run automatically without issues. If a hook fails, however, VSCode will throw you an error. In that case, click on "Show Command Output" as per the screenshot below to see the errors returned from pre-commit.

![example error](https://gitlab.com/tea-masters/communiTEA/uploads/20cb944e753e1823f0702918050a4540/Screenshot_2023-10-17_124748.png)
