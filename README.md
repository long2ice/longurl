# longurl

[![deploy](https://github.com/long2ice/longurl/actions/workflows/deploy.yml/badge.svg)](https://github.com/long2ice/longurl/actions/workflows/deploy.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/long2ice/longurl.svg)](https://pkg.go.dev/github.com/long2ice/longurl)

## Introduction

`longurl` is a self-hosted short url service.

## Try It Out

```shell
curl -H "Content-Type:application/json" -X POST --data '{"url": "https://github.com/long2ice/longurl"}' https://longurl.long2ice.io
```

Return like this:

```json
{
  "url": "https://longurl.long2ice.io/3FXrsHE"
}
```

Now Visit <https://longurl.long2ice.io/3FXrsHE>.

## Deploy

First write a `config.yaml`.

```yaml
server:
  host: 0.0.0.0
  port: 3000
  logTimezone: Asia/Shanghai
  logTimeFormat: 2006-01-02 15:04:05.000000
url:
  domain: localhost:3000
  schema: http
  length: 7
database:
  type: mysql
  dsn: root:123456@tcp(127.0.0.1:3306)/longurl?parseTime=true
```

Then run with `docker`.

```shell
docker run -d -p 3000:3000 --name longurl -v config.yaml:/config.yaml ghcr.io/long2ice/longurl/longurl
```

## Credits

- [Fiber](https://github.com/gofiber/fiber), Express inspired web framework written in Go.
- [Ent](https://github.com/ent/ent), An entity framework for Go.
- [Sonyflake](https://github.com/sony/sonyflake), A distributed unique ID generator inspired by Twitter's Snowflake.

## License

This project is licensed under the
[Apache-2.0](https://github.com/long2ice/longurl/blob/master/LICENSE)
License.