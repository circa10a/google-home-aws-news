# google-home-aws-news

A webhook for google assistance integrations to provide the latest AWS news

Deployed on Heroku at https://google-home-aws-news.herokuapp.com/webhook

_Powered by [go-aws-news](https://github.com/circa10a/go-aws-news)_

## Usage

- "OK Google, Show me the latest AWS announcements"
- "OK Google, Latest news about AWS"
- "OK Google, Give me news on AWS"
- "OK Google, Give me AWS news"

## Configuration

|             |                                                                       |                      |                        |           |               |
|-------------|-----------------------------------------------------------------------|----------------------|------------------------|-----------|---------------|
| Name        | Description                                                           | Environment Variable | Command Line Argument  | Required | Default        |
| GIN MODE    | Runs web server in production or debug mode                           |`GIN_MODE`            | NONE                   | `false`  | `release`      |
| PORT        | Port for web server to listen on                                      | `PORT`               | NONE                   | `false`  | `8080`         |

## Development

### Build

```shell
make
```

### Run

```shell
make run
```

### Compile Binary

```shell
make compile
```