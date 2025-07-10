# google-home-aws-news

[![Go Report Card](https://goreportcard.com/badge/github.com/circa10a/google-home-aws-news)](https://goreportcard.com/report/github.com/circa10a/google-home-aws-news)

A Google Assistant webhook integration to provide the latest AWS news.

* [View details on Google](https://assistant.google.com/services/a/uid/0000006c6dc51de5)
* Deployed on Heroku at https://google-home-aws-news.herokuapp.com/webhook

_Powered by [go-aws-news](https://github.com/circa10a/go-aws-news)_

## Usage

- "OK Google, Talk to cloud computing news"

> Note: News items are cached for 8 hours, then renewed with another request

### Example View

<p align="center"><img src="https://i.imgur.com/UjxafOi.png" width="30%" height="30%" /></p>

## Configuration

|             |                                                                       |                      |                        |           |               |
|-------------|-----------------------------------------------------------------------|----------------------|------------------------|-----------|---------------|
| Name        | Description                                                           | Environment Variable | Command Line Argument  | Required | Default        |
| GIN MODE    | Runs web server in production or debug mode                           |`GIN_MODE`            | NONE                   | `false`  | `release`      |
| PORT        | Port for web server to listen on                                      | `PORT`               | NONE                   | `false`  | `8080`         |

## Metrics

Prometheus metrics for usage are available at `/metrics`

## Development

### Test

```shell
make
```

### Build

```shell
make build
```

### Run

```shell
make run
```

### Compile Binary

```shell
make compile
```
