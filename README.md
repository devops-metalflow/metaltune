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
./bin/metaltune --config-file="$PWD"/test/config.yml --listen-url=:19093
```



## Docker

```bash
version=latest make docker
docker run -v "$PWD"/test:/tmp ghcr.io/devops-metalflow/metaltune:latest --config-file=/tmp/config.yml --listen-url=:19093
```



## Usage

```
usage: metaltune --config-file=CONFIG-FILE --listen-url=LISTEN-URL [<flags>]

metaltune

Flags:
  --help                     Show context-sensitive help (also try --help-long and --help-man).
  --version                  Show application version.
  --config-file=CONFIG-FILE  Config file (.yml)
  --listen-url=LISTEN-URL    Listen URL (host:port)
```



## Settings

*metaltune* parameters can be set in the directory [config](https://github.com/devops-metalflow/metaltune/blob/main/config).

An example of configuration in [config.yml](https://github.com/devops-metalflow/metaltune/blob/main/config/config.yml):

```yaml
apiVersion: v1
kind: metaltune
metadata:
  name: metaltune
spec:
  keentune:
    url: "127.0.0.1:19095"
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
    "cleanup": true,
    "tuning": {
      "auto": true,
      "profile": "content"
    },
    "turbo": true
  }
}
```



## License

Project License can be found [here](LICENSE).



## Reference

- [keentune-tuning](https://gist.github.com/craftslab/ff310dfe97b2bb2273b97d0b8f836f44)
- [ubuntu-clean](https://gist.github.com/craftslab/1a945f6d66892fa431f736bf818889a2)
