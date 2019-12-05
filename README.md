# Keel
[![GoDoc](https://godoc.org/github.com/buyco/keel?status.svg)](http://godoc.org/github.com/buyco/keel) [![Build Status](https://travis-ci.com/buyco/keel.svg?branch=master)](https://travis-ci.com/buyco/keel) [![Go Report Card](https://goreportcard.com/badge/github.com/buyco/keel)](https://goreportcard.com/report/github.com/buyco/keel) [![codecov](https://codecov.io/gh/buyco/keel/branch/master/graph/badge.svg)](https://codecov.io/gh/buyco/keel) [![license](https://img.shields.io/github/license/buyco/keel.svg?maxAge=2592000)](https://github.com/buyco/keel/LICENSE)


Toolkit for BuyCo microservices

## Run tests locally

```bash
$ export GO111MODULE=on
$ go get ./...
$ go get github.com/onsi/gomega
$ go install github.com/onsi/ginkgo/ginkgo
$ go install github.com/joho/godotenv/cmd/godotenv
$ ginkgo -r --randomizeAllSpecs --randomizeSuites --race --trace
```
