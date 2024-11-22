# `go-puerto` - Launch new project straight into deep waters with GO + HTMX + TEMPL

Tired creating project skeleton from scratch? Here's what you get with `go-puerto`

- Domain Driven Design (DDD) Architecture
- Built in context-driven routing with global error handling
- Automatic store connections (MongoDB, Postgres, Valkey - you chose which)
- Automatic configuration from `.env` with validators
- Compile-ready a-h/templ templates and layouts (with css reset)
- Auto imported HTMX (Alpine.js is optional)
- Internationalisation (one click website's language and currency switch)
- Built-in AES encryption/decryption, advanced email, password, and other validators
- Ready to go web and file server
- Ultra useful middlewares with chi routing

Full test coverage!!!

## Requirements

Make sure you have air and templ installed:

```
go install github.com/air-verse/air@latest
go install github.com/a-h/templ/cmd/templ@latest
```

This project skeleton requires .env file in the root directory. Configuration with required keys are below.
If you miss any piece of configuration and local setup - don't worry! It's thoroughly documented on every step and you will be guided with explicit messages that will help you get up in seconds!
Reading this doc till the end is highly recommended.

## Middlewares

- CORS
- Rate Limiter (Automatic burst and limit config)
- Localisation (Automatic Language and Currency setting into context values)
- Secure Headers (Set secure headers to avoid nasty attacks)
- Validate Headers (Content-Type validation)
- ETAG (Limit bandwith transfer with cached content)
- Headers logging
- Method Override

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

What's more, this project is fully covered in tests. It doesn't cover the functions that act as entry point for extending application like adding extending router paths or global store and handler constructors. Apart from those, this project has 99% test coverage.

## VSCode

If you want to debug in VSCode you need to insert all environmental variables in the .vscode/launch.json. Remember to add `/.vscode/` folder into `.gitignore` in order to never expose any fragile data.

## Golang Air Development

This project includes an entire .air.toml configuration. In order to use it, update the project name in the build section:

```
[build]
bin = "./tmp/REPLACE-ME-HERE"
```

## Environmental variables

This project is fully equipped with basic setup and validation mechanisms. That's why it's vital to setup all these keys in local `.env` file in the root directory (or in the `launch.json` file if you're using it for VSCode):

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
