# Keel
[![GoDoc](https://godoc.org/github.com/buyco/keel?status.svg)](http://godoc.org/github.com/buyco/keel) [![Go Report Card](https://goreportcard.com/badge/github.com/buyco/keel)](https://goreportcard.com/report/github.com/buyco/keel) [![codecov](https://codecov.io/gh/buyco/keel/branch/master/graph/badge.svg)](https://codecov.io/gh/buyco/keel) [![license](https://img.shields.io/github/license/buyco/keel.svg?maxAge=2592000)](https://github.com/buyco/keel/LICENSE)


Toolkit for BuyCo microservices

## Run tests locally

```bash
go mod download
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo@v2.28.1

ginkgo -r --randomize-all --randomize-suites --race --trace -coverprofile=cover.out
```
