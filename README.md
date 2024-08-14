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
  - [hyperfine](https://github.com/sharkdp/hyperfine/)

Configured **CGO** is required for [go-sqlite3](https://github.com/mattn/go-sqlite3?tab=readme-ov-file#installation).


## A Tour of Texel

  Components:
  - app
  - controller
  - model: persistence layer(Mnemosyne)
  - construction: business logic



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
  - [?] feat(design-rule-engine): implementation
  - [?] test(design-rule-engine): integration tests
  - [ ] test: smoke tests
  - [ ] chore(controller): refactoring
  - [ ] *** Release 0.1.0.pre1 version ****
  - [ ] test(design-rule-engine): unit tests
  - [ ] feat(controller): concurrent update
  - [ ] Handle `ErrProjectNotFound` error
  - [ ] Grafterm dashboard
  - [ ] [Connect healthz to DB](https://pkg.go.dev/database/sql#example-package-OpenDBService)
  - [ ] docs: readme
  - [ ] Prometheus Metrics
  - [ ] [Texel Architecture with D2](https://app.terrastruct.com/diagrams/2073737807) or [this](https://text-to-diagram.com/)
  - [ ] *** Release 0.1.0 version ****
  - [ ] Postman
  - [ ] OpenAPI Specification with `Swag`
  - [ ] Database timeout via `context.Context`
  - [ ] Add CLI and ENV configuration routines
  - [ ] feat(logging): production ready
  - [ ] feat(deployment): dockerfile
  - [ ] feat(deployment): google cloud run
  - [ ] *** Release 0.2.0 version ****
  - [ ] feat(deployment): aws faregate
  - [ ] feat: dependency injection

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
- https://en.wikipedia.org/wiki/GeoJSON
- [Euclidean Geometry[(https://en.wikipedia.org/wiki/Euclidean_geometry)
  - https://en.wikipedia.org/wiki/Apeirogon
  - https://en.wikipedia.org/wiki/List_of_two-dimensional_geometric_shapes
  - https://en.wikipedia.org/wiki/Projected_coordinate_system
