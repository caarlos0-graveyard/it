# IT [![Build Status](https://img.shields.io/shippable/55c8a219edd7f2c0529b965f.svg)](https://app.shippable.com/projects/55c8a219edd7f2c0529b965f/builds/latest)


An optionated Integration Test framework for Go.

## What it does

Basically, it helps you to start up a test server with all your routes, and,
with a real database! It will create a new temporary database for you, run
the migrations and, after the tests ran, delete it.

The idea is based on an internal framework that we made in Java at
[@ContaAzul](http://github.com/ContaAzul).

## What's optionated

At least for now, the only "optionated" thing is the use of `sqlx` instead
of `database/sql`.

There are also other defaults, but you can override them using environment
variables. Check the [`config.go` file](/blob/master/base/config.go).

## Talk is cheap, show me the code!

Check out the [/example](/example) folder!

You run then with:

```sh
$ cd example
$ go test
```
