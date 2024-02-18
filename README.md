<div align="center">

  <img src="frontend/public/CommuniteaLogo.svg" alt="CommuniTEA Logo" width="200"/>

  # CommuniTEA

  [![pre-commit badge](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit) [![codecov badge](https://codecov.io/gh/CommuniTEAM/CommuniTEA/graph/badge.svg?token=NHX49WER68)](https://codecov.io/gh/CommuniTEAM/CommuniTEA)

  **Bringing your local community together over a cuppa 🍵**

  <img src="https://raw.githubusercontent.com/devicons/devicon/24f2a9e2a16401e681583ae7a494fad71df03fce/icons/axios/axios-plain-wordmark.svg" width="50" height="50" alt="Axios Logo" title="Axios">
  &nbsp;
  <img src="https://raw.githubusercontent.com/biomejs/biome/main/website/src/assets/svg/logomark.svg" width="50" height="50" alt="Biome Logo" title="Biome">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/docker/docker-original-wordmark.svg" width="50" height="50" alt="Docker Logo" title="Docker">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/figma/figma-original.svg" width="50" height="50" alt="Figma Logo" title="Figma">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/go/go-original-wordmark.svg" width="50" height="50" alt="Go Logo" title="Go">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/materialui/materialui-original.svg" width="50" height="50" alt="MaterialUI Logo" title="MaterialUI">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/24f2a9e2a16401e681583ae7a494fad71df03fce/icons/openapi/openapi-original-wordmark.svg" width="50" height="50" alt="OpenAPI Logo" title="OpenAPI">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/postgresql/postgresql-original-wordmark.svg" width="50" height="50" alt="PostgreSQL Logo" title="PostgreSQL">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/react/react-original-wordmark.svg" width="50" height="50" alt="React Logo" title="React">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/typescript/typescript-original.svg" width="50" height="50" alt="TypeScript Logo" title="TypeScript">
  &nbsp;
  <img src="https://raw.githubusercontent.com/devicons/devicon/24f2a9e2a16401e681583ae7a494fad71df03fce/icons/vitejs/vitejs-original.svg" width="50" height="50" alt="Vite Logo" title="Vite">

</div>

## Table of Contents

<details>
  <summary><b>Expand for Contents</b></summary>

- [CommuniTEA](#communitea)
  - [Table of Contents](#table-of-contents)
- [Overview](#overview)
  - [Our Mission](#our-mission)
- [Development](#development)
  - [Working with Git Hooks](#working-with-git-hooks)
  - [Setting Up Pre-Commit](#setting-up-pre-commit)
  - [Pre-Commit Rulesets](#pre-commit-rulesets)
    - [Commit Message Format](#commit-message-format)
    - [Biome (TypeScript) Format](#biome-typescript-format)
    - [Go Format](#go-format)
      - [Installing golangci-lint on Mac or Linux Systems](#installing-golangci-lint-on-mac-or-linux-systems)
      - [Installing golangci-lint on Windows Systems](#installing-golangci-lint-on-windows-systems)
  - [Running Unit Tests](#running-unit-tests)
  - [Using Git Hooks in VS Code](#using-git-hooks-in-vs-code)

</details>

# Overview

⚠️ This project is currently **in progress**. Please be patient with us as we diligently work away at crafting something special. Our team is very excited to help bring the world of tea to your neighborhood soon!

## Our Mission

CommuniTEA is a unique application that caters to both businesses and the community, bringing together tea enthusiasts in search of their favorite brews with local establishments ready to serve. The platform aims to simplify the process of finding new and popular teas by providing a curated list of nearby shops or tea houses that offer those sought-after tea varieties.

# Development

Looking to contribute to the project? Please see the information below to get started.

## Working with Git Hooks

Any commits to the repository must conform to the configured rulesets for:
- Commit Messages
- Biome (Frontend)
- Golangci-lint (Backend)

To enforce these conventions, [pre-commit](https://pre-commit.com/), a Git hook manager, is a required development dependency for the project. You will need to install it on your local machine.

## Setting Up Pre-Commit

1. Install pre-commit:<br>
**Windows**:
Pre-commit runs off of Python, so you will need pip on your local machine. To install pre-commit, open your terminal and run `pip install pre-commit`.<br>
**Mac or Linux**:
Ensure you have [homebrew](https://brew.sh/) installed on your local machine. To install pre-commit, open your terminal and run `brew install pre-commit`.<br>

2. Install the repository's Git hooks by navigating to the repo's directory in your terminal and running `pre-commit install`.

And that's all! You're now ready to commit to the repository.
For more detailed installation information and troubleshooting, see the [pre-commit documentation](https://pre-commit.com/#install).

## Pre-Commit Rulesets

### Commit Message Format

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

Special thanks to strdr4605 for [their walkthrough](https://strdr4605.com/commitlint-custom-commit-message-with-emojis) of this configuration!

### Biome (TypeScript) Format

This project utilizes [Biome](https://biomejs.dev/) as an alternative to ESLint and Prettier with an extended ruleset based on the [guide by Airbnb](https://github.com/airbnb/javascript). The installed pre-commit hook will attempt to format your code and fix any linting errors found in the staged changes. If it does, the check will fail and you will need to re-stage and commit.

Please be sure to have the [Biome VS Code extension](https://marketplace.visualstudio.com/items?itemName=biomejs.biome) installed for the best development experience. This repository's included workspace settings for VS Code set Biome as the default formatter for TypeScript and JavaScript, which requires the extension to be installed locally. Upon first installation, the extension will prompt you to install the Biome binary. Be sure to do this, as it is required for the extension to work.

Once the extension is installed and set up, you can fix and format the current file quickly with a manual save (Windows/Linux: `Ctrl+S`  Mac: `⌘S`).

### Go Format

This repository uses a robust [golangci-lint](https://golangci-lint.run/) configuration built up of over 75 linters, as recommended in the ["Golden Config" by maratori](https://gist.github.com/maratori/47a4d00457a92aa426dbd48a18776322). To successfully run the pre-commit hook, a local installation of golangci-lint is required. Follow the [official documentation](https://golangci-lint.run/usage/install/) to get started, or continue reading for installation instructions.

To save yourself unnecessary headaches, it is strongly recommended to install the [official Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go) for VS Code in order to be fully set up for linting Go locally. The repository is configured to support real-time linting feedback from the extension so that linting flags appear as errors. The extension requires a local installation of Go, as well as some additional dependencies, and will prompt you to install any missing requirements for local linting. Simply click the warning as seen in the screenshot below to begin the setup.

<img src="https://github.com/CommuniTEAM/CommuniTEA/assets/31549337/753ecd20-86e2-47b4-b4e3-cbbf3168424d" alt="VS Code missing dependency warning" height=100 width=450>

<details>
  <summary>

  #### Installing golangci-lint on Mac or Linux Systems

  </summary>

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
<details>
  <summary>

  #### Installing golangci-lint on Windows Systems

  </summary>

  Unfortunately installation on Windows is not as straightforward, as golangci-lint runs off of bash. If you do not already have a bash terminal, it is strongly recommended to install [Git for Windows](https://gitforwindows.org/) so that you have Git Bash available.

  In your bash terminal, install the golangci-lint binary by running:
  ```
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2
  ```

  **Important:** golangci-lint can *only* run in a terminal capable of handling bash scripts. The pre-commit hook for Go will not pass unless the bash script can be run.

  While you can absolutely run pre-commit and git commands in your installed bash terminal, it's not the only way. If you don't want to be forced into using your bash terminal for everything, it is possible to set up your system path such that bash scripts can be run successfully from any terminal, including Powershell.

  For bash scripts to work anywhere, you must reconfigure your system path in Windows. The path to the Git Bash terminal (`C:\Program Files\Git\cmd`) must be at the top of your system path and given priority over system32 as shown below:

  <img src="https://github.com/CommuniTEAM/CommuniTEA/assets/31549337/3fbd7d15-b76c-456a-b92d-a04413df0c8f" alt="example of Git Bash path at the root of system path">


  Instructions on how to change your system path can be found [here](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/). Note: Changes to the path will go into effect after a PC reboot. You can check your current path in PowerShell with `$env:PATH`.

  For more information on why this is the necessary fix for bash commands, see [this thread](https://github.com/syntaqx/git-hooks/pull/3). TL;DR: It's because we have WSL installed (required for Docker) as part of system32 and WSL erroneously intercepts the bash commands.

</details>

## Running Unit Tests

You can easily check the status of the project's unit tests at any time by running the configured manual pre-commit hooks:
```
pre-commit run --hook-stage manual --all-files
```
Use the `--all-files` arg to run every test, or leave it off to run only those relevant to the staged changes.

These hooks will never run automatically during the commit process and are only available for manual invocations. Unit tests will still automatically run against PRs.

## Using Git Hooks in VS Code

When committing changes through VS Code's source control tab, pre-commit will still run automatically in your default terminal as if you committed from the command line. If a hook fails, however, VS Code will notify you via an error. To display the details of the pre-commit failure, click on "Show Command Output" as per the screenshot below to display the results in the command line.

<img src="https://gitlab.com/tea-masters/communiTEA/uploads/20cb944e753e1823f0702918050a4540/Screenshot_2023-10-17_124748.png" alt="example error" height=150 width=425>
