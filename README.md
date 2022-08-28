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
  tuning:
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

- [auto-cpufreq](https://github.com/AdnanHodzic/auto-cpufreq)
- [cpufreq](https://abcdxyzk.github.io/blog/2015/08/12/kernel-cpufreq/)
- [cpufreq](https://github.com/konkor/cpufreq)
- [cpufreq](https://github.com/VitorRamos/cpufreq)
- [cpufrequtils](https://packages.ubuntu.com/bionic/cpufrequtils)
- [cpufrequtils](https://www.icode9.com/content-3-1252636.html)
- [frequency-scaling](https://www.thinkwiki.org/wiki/How_to_make_use_of_Dynamic_Frequency_Scaling)
- [keentune-tuning](https://gist.github.com/craftslab/ff310dfe97b2bb2273b97d0b8f836f44)
- [turbo-boost](https://huataihuang.gitbooks.io/cloud-atlas/content/os/linux/kernel/cpu/intel_turbo_boost_and_pstate.html)
- [ubuntu-clean](https://gist.github.com/craftslab/1a945f6d66892fa431f736bf818889a2)
- [ubuntu-tuning](https://its401.com/article/MMTS_yang/122244404)
