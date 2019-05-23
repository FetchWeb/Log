# FetchWeb Log

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/FetchWeb/Log)](https://goreportcard.com/report/github.com/FetchWeb/Log)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/FetchWeb/Log)
[![GitHub release](https://img.shields.io/github/release/FetchWeb/Email.svg)](https://github.com/FetchWeb/Log/releases )

## Introduction
FetchWeb Log is a simple logging package written in Go with no dependencies outside of the Go standard library.

## Setup Example
```go
package main

import (
	"errors"
	"os"

	log "github.com/FetchWeb/Log"
)

// TestLog writes test logs to Stdout.
func main() {
	if err := log.Startup(true, os.Stdout); err != nil {
		panic(err)
	}

	log.Info("Test Info message")
	log.Infof("Test Infof %v", errors.New("message"))

	log.Error("Test Error message")
	log.Errorf("Test Errorf %v", errors.New("message"))

	log.Debug("Test Debug message")
	log.Debugf("Test Debugf %v", errors.New("message"))
}
```