# Simple Link Shortener
Extremely simple link shortener written in Golang.

## What it can do?
* Shorten links to 7 characters
* Use token based authorization to add or delete links
* Use mysql to save links

## Setup
At first setup the tables needed in mysql which are available in `tables.sql` file. Then create a copy of `config.json` and change your desired values. At last, use `go build` to compile the application and use it.