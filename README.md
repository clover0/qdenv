# Qdenv

[![GitHub license][license-badge]](LICENSE)
[![Go Report Card][go-report-card-badge]][go-report-card]

:tada: Welcome to Qdenv!

Qdenv is a tool to quickly create a programming environment.

# Quickstart

## Install
Before
- Install Docker(for Mac) if you have not installed.
- Install docker-compose if you have not installed.

```
brew tap clover0/qdenv
brew install qdenv
```

or

```
go get github.com/clover0/qdenv
```

## Examples
### Get Python environment quickly!
Search available python3.7 images and tags.
```
qdenv search python 3.7

NAME    TAGS
python    3.7.12-slim-buster 3.7.12-slim-bullseye 3.7.12-slim 3.7.12-buster 3.7.12-bullseye ...
hylang    python3.7-buster python3.7-bullseye python3.7 pypy3.7-buster pypy3.7-bullseye ...
pypy    3.7-slim-buster 3.7-slim-bullseye 3.7-slim 3.7-buster 3.7-bullseye 3.7-7.3.5-slim-buster ...
```

Init environment. (Choiced python 3.7.12-slim-bullseye)
```
qdenv init python:3.7.12-slim-bullseye

Building qdenv
[+] Building 
...
```

Enter environment.
```
qdenv enter

root:/var/code$ python --version
Python 3.7.12
```

Coding on your machine but code is executed in qdenv environment(docker container).

### Search Ruby
```
qdenv search ruby

NAME    TAGS
ruby    latest slim-buster slim-bullseye slim buster bullseye 3.0.2-slim-buster 3.0.2-slim-bullseye ...
jruby    latest 9.3.0.0-jre8 9.3.0.0-jre11 9.3.0.0-jre 9.3.0.0-jdk8 9.3.0.0-jdk17 9.3.0.0-jdk11 ...
redmine    latest passenger 4.2.2-passenger 4.2.2 4.2-passenger 4.2 4.1.4-passenger ...
```

### Command help
`qdenv <command> --help`

```
$ qdenv search --help 

Search images by input word. By Default, images is filtered by is-official.

Usage:
  qdenv search [image tag] [flags]

Flags:
  -h, --help   help for search
```


# Dependencies
- Docker
- Docker Compose

<!-- refs -->
[go-report-card]: https://goreportcard.com/report/github.com/clover0/qdenv
[go-report-card-badge]: https://goreportcard.com/badge/github.com/clover0/qdenv?style=flat-square&logo=appveyor
[license-badge]: https://img.shields.io/github/license/clover0/qdenv.svg?style=flat-square&logo=appveyor
