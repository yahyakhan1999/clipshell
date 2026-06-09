# ClipShell

ClipShell is a lightweight Go-based developer utility that executes shell commands, captures their output, and automatically copies the formatted command and output to the system clipboard.

The tool was created to simplify sharing terminal output in GitHub issues, pull requests, documentation, chat applications, and AI assistants without manually copying commands and results.

## Features

* Execute shell commands from a single wrapper command
* Capture both stdout and stderr
* Automatically copy results to the clipboard
* Works from WSL and PowerShell
* Preserves the original command alongside its output
* Lightweight single-binary implementation written in Go

## Motivation

Developers frequently need to share terminal commands and their results when:

* Reporting bugs
* Creating documentation
* Asking for technical help
* Writing GitHub issues
* Sharing troubleshooting steps

Normally this requires manually copying the command, copying the output, and formatting everything correctly.

ClipShell automates that workflow.

## Example

### Command

```bash
clipshell "git status"
```

### Clipboard Output

```text
> git status

On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean
```

### Another Example

```bash
clipshell "ls -la"
```

Clipboard Output:

```text
> ls -la

total 24
drwxrwxrwx 1 yahya yahya 4096 Jun  9 15:30 .
drwxrwxrwx 1 yahya yahya 4096 Jun  9 15:20 ..
-rwxrwxrwx 1 yahya yahya  895 README.md
-rwxrwxrwx 1 yahya yahya  618 main.go
```

## Installation

### Prerequisites

* Go 1.22 or later
* WSL or Linux
* Windows clipboard access via `clip.exe`

### Clone Repository

```bash
git clone https://github.com/yahyakhan1999/clipshell.git
cd clipshell
```

### Build

Linux / WSL:

```bash
go build -o clipshell
```

Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o clipshell.exe
```

## Usage

Run any shell command through ClipShell:

```bash
clipshell "git status"
clipshell "docker ps"
clipshell "python -m pytest"
clipshell "ls -la"
```

The command output is displayed normally and simultaneously copied to the clipboard.

## How It Works

1. User provides a shell command.
2. ClipShell executes the command using the system shell.
3. Standard output and error output are captured.
4. The command and output are formatted together.
5. The formatted result is copied directly to the clipboard.
6. The original output is still displayed in the terminal.

## Current Limitations

* Commands must currently be executed through the `clipshell` wrapper.
* Automatic shell integration is not yet implemented.
* Clipboard support currently relies on Windows `clip.exe`.

## Future Improvements

* Markdown-formatted clipboard output
* Command history support
* Timestamped command captures
* Native Linux clipboard integration
* Automatic shell hooks for Bash and PowerShell
* Configuration file support

## Technologies Used

* Go
* Windows Clipboard (`clip.exe`)
* WSL
* PowerShell

## Author

Yahya Khan
