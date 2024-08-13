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
  - [jq](https://stedolan.github.io/jq/)
  - [curl](https://curl.haxx.se/)

Configured **CGO** is required for [go-sqlite3](https://github.com/mattn/go-sqlite3?tab=readme-ov-file#installation).


## A tour Texel

  Components:
  - app
  - controller
  - model: persistence layer(Mnemosyne)
  - construction: business logic

  Smoke tests:
    - curl -v http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/building_limits | jq .

    - curl -v --data '{"data": {}}' -X PATCH http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/building_limits | jq .
    - curl -v --data @data/building_limits.geojson -X PATCH http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/building_limits | jq .

    - curl -v http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/height_plateaus | jq .
    - curl -v --data @data/height_plateaux.geojson -X PATCH http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/height_plateaus | jq .

    - curl -v http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/split_building_limits | jq .

  Stress tests:
    - hyperfine "curl -v http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/split_building_limits | jq ."
    - hyperfine -m 100000 'curl -v --data @data/BuildingLimits.geojson -X PATCH http://localhost:8080/v1/projects/feedface-cafe-beef-feed-facecafebeef/building_limits | jq .'


### Logging

  - [V levels](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md#what-method-to-use)
  - https://github.com/go-logr/zapr
  - https://github.com/go-logr/logr
  - https://github.com/uber-go/zap



## Progress

  - [x] Gin setup
  - [x] [Structured Logging](https://learninggolang.com/it5-gin-structured-logging.html)
  - [x] `SQLite` setup
  - [x] `GeoJSON` setup
  - [x] GET building_limits
  - [x] GET height_plateaux
  - [x] PATCH building_limits
  - [x] PATCH height_plateaux
  - [x] GET split_building_limits
  - [x] [fix database is locked](https://www2.sqlite.org/cvstrac/wiki?p=DatabaseIsLocked)
  - [x] feat(mnemosyne): implement `updateObject`
  - [x] PATCH height_plateaux
  - [ ] feat(design-rules-engine): implementation
  - [ ] feat(design-rules-engine): unit tests
  - [ ] chore(controller): refactoring
  - [ ] docs: readme
  - [ ] *** Release 0.1.0.pre1 version ****
  - [ ] feat(controller): concurrent update
  - [ ] Handle `ErrProjectNotFound` error
  - [ ] OpenAPI Specification
  - [ ] Postman
  - [ ] Prometheus Metrics
  - [ ] Grafterm dashboard
  - [ ] [Connect healthz to DB](https://pkg.go.dev/database/sql#example-package-OpenDBService)
  - [ ] Update README.md
  - [ ] *** Release 0.1.0 version ****
  - [ ] Database timeout via `context.Context`
  - [ ] Add CLI and ENV configuration routines
  - [ ] Dependency injection
  - [ ] feat(logging): production ready

## Contribution

### Before check-in the code checklist

  - [ ] Make sure all new source code files have the copyright header


## References

- [MPL2](https://www.mozilla.org/en-US/MPL/headers/)
- Gin:
  - https://gin-gonic.com/docs/examples/bind-uri/
- [Google JSON Style Guide](https://google.github.io/styleguide/jsoncstyleguide.xml)
- [JSON API spec](https://github.com/json-api/json-api)
- https://blog.logrocket.com/documenting-go-web-apis-with-swag/
- Deployment
  - [Google Functions](https://cloud.google.com/functions/docs/concepts/execution-environment#functions-concepts-scopes-go)
  - [Cloud Run](https://cloud.google.com/run/)
- [GeoJSON](https://en.wikipedia.org/wiki/GeoJSON)
- https://medium.com/@isuru89/a-better-way-to-implement-http-patch-operation-in-rest-apis-721396ac82bf
- https://en.wikipedia.org/wiki/Design_rule_checking
