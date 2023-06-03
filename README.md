# Toggl CLI

## Installation

```bash
go install github.com/kanade0404/toggl@latest
```

## Usage

```bash
toggl --help
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  toggl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  show        Show the working hours of a month

Flags:
  -h, --help     help for toggl
  -t, --toggle   Help message for toggle

Use "toggl [command] --help" for more information about a command.
```
```bash
toggl show -t <api_token> -p <project_id> -c <working_content_title>
```
