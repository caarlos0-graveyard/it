# it [![Circle CI](https://circleci.com/gh/caarlos0/it.svg?style=svg&circle-token=434594c6d2cdae171a9f60b34209157cde821afe)](https://circleci.com/gh/caarlos0/it)

An optionated Integration Test "framework" for Go.

## What it does

Basically, it helps you to start up a test server with all your routes, and,
with a real database! It will create a new temporary database for you, run
the migrations and, after the tests ran, delete it.

The idea is based on an internal framework that we made in Java at
[@ContaAzul](http://github.com/ContaAzul).

## What's optionated

- Your app uses a database and that it is PostgreSQL (about to change);
- Your app is a web app;
- Configuration comes from environment variables.

There are also other defaults, but you can override them using environment
variables. Check the [`config.go` file](/base/config.go).

## Talk is cheap, show me the code!

Check out the [example](/example) folder!

You can easily run it with:

```sh
$ cd example
$ time go test
2015/10/13 22:41:07 Create-ing test database hkermdouni
2015/10/13 22:41:07 Connecting to postgres://localhost:5432/hkermdouni?sslmode=disable
2015/10/13 22:41:07 Migrate-ing database...
PASS
2015/10/13 22:41:07 Shutdown IT...
ok  	github.com/caarlos0/it/example	0.548s
go test  4.57s user 0.53s system 161% cpu 3.156 total
```

## Lifecycle

- Create a database with a random name in the provided `POSTGRES_URL`;
- Run all `.sql` files in the `MIGRATIONS_FOLDER` against the created database;
- Run the tests;
- Drop the test database (unless `DROP_TEST_DATABASE` is `false`). 
