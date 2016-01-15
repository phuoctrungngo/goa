#! /usr/bin/make
#
# Makefile for goa
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
DIRS=$(shell go list -f {{.Dir}} ./...)
DEPEND=\
	bitbucket.org/pkg/inflect \
	github.com/PuerkitoBio/purell \
	github.com/asaskevich/govalidator \
	github.com/go-swagger/go-swagger \
	github.com/golang/gddo/httputil \
	github.com/golang/lint/golint github.com/onsi/gomega \
	github.com/julienschmidt/httprouter \
	github.com/manveru/faker \
	github.com/on99/gocyclo \
	github.com/onsi/ginkgo github.com/onsi/ginkgo/ginkgo \
	github.com/raphael/goa-cellar \
	github.com/zach-klippenstein/goregen \
	golang.org/x/tools/cmd/cover \
 	golang.org/x/tools/cmd/goimports \
	gopkg.in/alecthomas/kingpin.v2 \
	gopkg.in/tylerb/graceful.v1 \
	gopkg.in/yaml.v2

.PHONY: goagen

all: depend lint cyclo goagen test

depend:
	@go get $(DEPEND)

lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi

cyclo:
	@if [ "`gocyclo -over 20 . | grep -v examples/cellar | tee /dev/stderr`" ]; then \
		echo "^ - Cyclomatic complexity exceeds 20, refactor the code!" && echo && exit 1; \
	fi

test:
	@ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race -skipPackage vendor

goagen:
	@cd goagen && \
	go install
