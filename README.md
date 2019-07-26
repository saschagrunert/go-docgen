# go-docgen 📔

[![CircleCI](https://circleci.com/gh/saschagrunert/go-docgen.svg?style=shield)](https://circleci.com/gh/saschagrunert/go-docgen)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/saschagrunert/go-docgen/pkg/docgen)

## About

This project aims to provide convince documentation generation for golang based
projects. The following documentation generations are currently included:

| Input                                                    | Output            | API function                                                                                                                       |
| -------------------------------------------------------- | ----------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| [`cli.App`](https://godoc.org/github.com/urfave/cli#App) | Markdown `string` | [`func CliToMan(app *cli.App) (string, error)`](https://godoc.org/github.com/saschagrunert/go-docgen/pkg/docgen#CliToMan)          |
| [`cli.App`](https://godoc.org/github.com/urfave/cli#App) | Man Page `string` | [`func CliToMarkdown(app *cli.App) (string, error`](https://godoc.org/github.com/saschagrunert/go-docgen/pkg/docgen#CliToMarkdown) |

## Contributing

You want to contribute to this project? Wow, thanks! So please just fork it and
send me a pull request.
