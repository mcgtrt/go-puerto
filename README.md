# `go-puerto` - Launch new project straight into deep waters with GO + HTMX + TEMPL

Tired creating project skeleton from scratch? Here's what you get with `go-puerto`

- Domain Driven Design (DDD) Architecture
- Built in context-driven routing with error handling
- Automatic store connections (MongoDB, Postgres, Valkey)
- Automatic configuration from .env
- Compile-ready a-h/templ templates and layouts
- Auto imported HTMX and (optional) Alpine.js
- Internationalisation (website language + used currency)
- Built-in AES encryption/decryption
- Latest go-chi router
- Web and file server
- Ultra useful middlewares

## Requirements

Make sure you have air and templ installed:

```
go install github.com/air-verse/air@latest
go install github.com/a-h/templ/cmd/templ@latest
```

This project skeleton requires .env file in the root directory. Configuration with required keys below.
Please read this doc carefully as missing a single configuration item might result in app exiting from an error. The good side is that it's amazingly documented so you will always know what's missing (if you miss something).

## Core mechanisms

In order to speed up your development process, this project is fully equipped with amazing technologies and code pieces that will not only skyrocket your development speed, but also equip with fully tested and super helpful tools:

- custom context with automatic response methods
- context/handler method wrapping to match router's http.HandlerFunc pattern
- global error handling
- your own http file server (configure path and local dir via .env configuration)
- changing website's language and currency with a single click
- translations accessible from a single json file (locales folder) that are automatically detected by the system
- automatic translation matcher that stops running server if any from languages is missing a translation key
- global storage and handler object that contains all possible stores and handlers in a single object for the ease of use (simply extend them with your own controllers)
- extremely fast frontend generation thanks to rendering precompiled frontend components and layouts (including css reset)

It's highly advised that you take a look into the utils folder as it's filled with the most useful functions. Some of them are:

- AES string encryption/decryption
- default configuration import with validation
- check if string is URL safe
- getting locale (language + currency) straight from provided context
- advanced email address validator
- password and name validator (with default MIN/MAX values you can easily adjust)

What's more, this project is fully covered in tests. It doesn't cover the functions that are entry points for extending application like adding paths with methods to the router or constructor method for the global store and handler object. Beside those few methods this project has 100% test coverage.

## VSCode

For debugging in VSCode:

!!! Operation below may potentially lead to exposing your env vars publicly !!!
The best solution to stop any data exposure of local variables, add folder `.vscode` to `.gitignore` file.

If you want to debug in VSCode you need to insert all environmental variables in the .vscode/launch.json.

## Golang Air Development

This project includes an entire .air.toml configuration. In order to use it, update the project name in the build section:

```
[build]
bin = "./tmp/go-puerto" // replace with your own project name
cmd = "make air-build"
```

## Environmental variables

This project is fully equipped with basic setup with validation mechanisms. That's why it's vital to setup all these keys in local `.env` file in the root directory (or in the `launch.json` file if you're using it for VSCode):

```
# GENERAL
AES_SECRET=
# empty path will disable HTTP file server
FILE_SERVER_PATH=

# For those below, if you want to include any, just type true.
# Any other value will be ignored resulting in not including
# the part in the project configuration and init.
IMPORT_ALPINE_JS=true
INCLUDE_MONGO=true
INCLUDE_POSTGRES=false
INCLUDE_VALKEY=invalid

PROJECT_NAME=go-puerto

# HTTP CONFIG
HTTP_PORT=3000

# MONGODB CONFIG
MONGO_DB_NAME=go-puerto
MONGO_USERNAME=go-puerto
MONGO_PASSWORD=
MONGO_HOST=localhost
MONGO_PORT=27017

# POSTGRES CONFIG

# VALKEY CONFIG
```
