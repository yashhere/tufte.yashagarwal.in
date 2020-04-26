SHELL := /bin/bash

PROJECTNAME := $(shell basename "$(PWD)")

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

HUGO := hugo
HUGO_FLAGS = --gc --minify

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

.PHONY: all clean library public server watch .minifier help

# Below are PHONY targets
all: clean library public .minifier

help:
	@echo "Usage: make <command>"
	@echo "  all     Builds the blog and minifies it"
	@echo "  clean   Cleans all build files"
	@echo "  server  Runs a webserver on port 1313 to test the final minified result"
	@echo "  watch   Runs hugo in watch mode, waiting for changes"

library:
	mkdir -p data
	@echo "Downloading library data"
	cd scripts && go run main.go

server: public
	cd public && python -m SimpleHTTPServer 1313

watch: clean
	$(HUGO) server -w

clean:
	-rm -rf public
	-rm -rf data

public:
	@echo "Generating blog"
	$(HUGO) $(HUGO_FLAGS)

.minifier:
	@echo "minifying images now"
	find ./public -type f \( -iname \*.jpg -o -iname \*.jpeg -o -iname \*.png \) -print0 | xargs -0 -P16 -n2 mogrify -compress JPEG -strip -quality 40