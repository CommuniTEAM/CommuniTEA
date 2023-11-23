# CommuniTEA

[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)

Bringing your local community together over a cuppa 🍵

## Table of Contents

<details>
  <summary><b>Expand for Contents</b></summary>

- [CommuniTEA](#communitea)
  - [Table of Contents](#table-of-contents)
  - [Working with Git Hooks](#working-with-git-hooks)
    - [Installing Pre-Commit](#installing-pre-commit)
    - [Pre-Commit Rulesets](#pre-commit-rulesets)
      - [Commit Message Format](#commit-message-format)
      - [ESLint TypeScript Format](#eslint-typescript-format)
      - [Golang Format](#golang-format)
    - [Using Git Hooks in VSCode](#using-git-hooks-in-vscode)

</details>

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

All commit messages, regardless of the branch, must adhere to this header standard:
```
<Emoji> [Related Jira Issue]: Description of work
```
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

This repository uses a robust [golangci-lint](https://golangci-lint.run/) configuration built up of over 75 linters, as recommended in the ["Golden Config" by maratori](https://gist.github.com/maratori/47a4d00457a92aa426dbd48a18776322). To successfully run the pre-commit hook, a local installation of golanci-lint is required. Follow the [official documentation](https://golangci-lint.run/usage/install/) to get started, or see the installation instructions below.

<details>
  <summary><b>Installing golangci-lint on Mac or Linux Systems</b></summary>

  **Mac**

  First, ensure you have [homebrew](https://brew.sh/) installed. Then, in your terminal, run:
  ```
  brew install golangci-lint
  brew upgrade golangci-lint
  ```

  **Linux**

  If your distro has Snap, you can simply run:
  ```
  sudo snap install golangci-lint
  ```
  Alternatively, you can manually install the binary by running:
  ```
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2
  ```

</details>
<br>
<details>
  <summary><b>Installing golangci-lint on Windows Systems</b></summary>
  <br>
  Unfortunately installation on Windows is not as straightforward, as golangci-lint runs off of bash. If you do not already have a bash terminal, it is strongly recommended to install <a src="https://www.msys2.org/">MSYS2</a> and use its UCRT64 environment.

  In your bash terminal, install the golangci-lint binary by running:
  ```
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2
  ```

  **Important:** golangci-lint can *only* run in a terminal capable of handling bash scripts. The pre-commit hook for Go will not pass unless the bash script can be run.

  While you can absolutely run pre-commit and git commands in your installed bash terminal, it's not the only way. If you don't want to be forced into using your bash terminal for everything, it is possible to set up your system path such that bash scripts can be run successfully from any terminal, including Powershell.

  For bash scripts to work anywhere, you must reconfigure your system path in Windows. The path to the MSYS bash terminal must be at the top of your system path and given priority over system32 as shown below:
  ```
  C:\msys64\usr\bin
  %SystemRoot%\system32
  <The rest of your system path>
  ```

  Instructions on how to change your system path can be found [here](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/).

  For more information on why this is the necessary fix for bash commands, see [this thread](https://github.com/syntaqx/git-hooks/pull/3). TL;DR: It's because we have WSL installed (required for Docker) and WSL intercepts the bash commands.

</details>
<br>
To save yourself unnecessary headaches, it is strongly recommended to install the official Go extension for VSCode so that the linting errors show in real time as you make changes to the codebase. The extension requires Go to be installed locally, as well as some additional dependencies, before it will display linting hints.

### Using Git Hooks in VSCode

When commiting changes through VSCode's source control tab, pre-commit will still run automatically in your default terminal as if you committed from the command line. If a hook fails, however, VSCode will notify you via an error. To display the details of the pre-commit failure, click on "Show Command Output" as per the screenshot below to display the results in the command line.

<img src="https://gitlab.com/tea-masters/communiTEA/uploads/20cb944e753e1823f0702918050a4540/Screenshot_2023-10-17_124748.png" alt="example error" height=150 width=425>
