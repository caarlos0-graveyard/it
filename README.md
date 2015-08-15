# IT

An optionated Integration Test framework for Go.

## What it does

Basically, it helps you to start up a test server with all your routes, and,
with a real database! It will create a new temporary database for you, run
the migrations and, after the tests ran, delete it.

The idea is based on an internal framework that we made in Java at
@ContaAzul.

## What's optionated

At least for now, the only "optionated" thing is the use of `sqlx` instead
of `database/sql`.

## Talk is cheap, show me the code!

Check out the [/example](/example) folder!
