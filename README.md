# metaltune

[![Build Status](https://github.com/devops-metalflow/metaltune/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/devops-metalflow/metaltune/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/devops-metalflow/metaltune/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/devops-metalflow/metaltune)
[![Go Report Card](https://goreportcard.com/badge/github.com/devops-metalflow/metaltune)](https://goreportcard.com/report/github.com/devops-metalflow/metaltune)
[![License](https://img.shields.io/github/license/devops-metalflow/metaltune.svg)](https://github.com/devops-metalflow/metaltune/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/devops-metalflow/metaltune.svg)](https://github.com/devops-metalflow/metaltune/tags)



## Introduction

*metaltune* is the worker of [metalflow](https://github.com/devops-metalflow/metalflow) written in Go.



## Prerequisites

- Go >= 1.18.0



## Run

```bash
version=latest make build
./bin/metaltune --listen-url=:19094
```



## Docker

```bash
version=latest make docker
docker run ghcr.io/devops-metalflow/metaltune:latest --listen-url=:19094
```



## Usage

```
usage: metaltune --listen-url=LISTEN-URL [<flags>]

metaltune

Flags:
  --help                   Show context-sensitive help (also try --help-long and --help-man).
  --version                Show application version.
  --listen-url=LISTEN-URL  Listen URL (host:port)
```



## Protobuf

```json
{
  "apiVersion": "v1",
  "kind": "metaltune",
  "metadata": {
    "name": "metaltune"
  },
  "spec": {
    "name": "foo"
  }
}
```



## License

Project License can be found [here](LICENSE).



## Reference
