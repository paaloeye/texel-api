# Texel API

## Quickstart

### Run

```bash
# Development mode
task run

# Release mode
GIN_MODE=release task run
```

## Prerequisites

We expect that the following binaries are available in your `PATH`.

  - [task](https://taskfile.dev/)
  - [swag](https://github.com/swaggo/swag)
  - [grafterm](https://github.com/slok/grafterm)

Configured **CGO** is required for [go-sqlite3](https://github.com/mattn/go-sqlite3?tab=readme-ov-file#installation).


## A tour Texel

  - `app.App`
  -

### Logging

  - [V levels](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md#what-method-to-use)
  - https://github.com/go-logr/zapr
  - https://github.com/go-logr/logr
  - https://github.com/uber-go/zap



## Progress

  - [ ] Gin
  - [ ] OpenAPI Specification
  - [ ] [Logging](https://learninggolang.com/it5-gin-structured-logging.html)
  - [ ] Postman
  - [ ] Prometheus Metrics
  - [ ] Grafterm dashboard
  - [ ] *** Release 0.1.0 version ****
  - [ ] Add CLI and ENV configuration routines

## Contribution

### Before check-in the code checklist

  - [ ] Make sure all new source code files have the copyright header


## References

- [MPL2](https://www.mozilla.org/en-US/MPL/headers/)
- Gin:
  - https://gin-gonic.com/docs/examples/bind-uri/
- [Google JSON Style Guide](https://google.github.io/styleguide/jsoncstyleguide.xml)
- [JSON API spec](https://github.com/json-api/json-api)

