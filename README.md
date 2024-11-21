# `go-puerto` - Launch new project straight into deep waters with GO + HTMX + TEMPL

Tired creating project skeleton from scratch? Here's what you get with `go-puerto`

- Domain Driven Design (DDD) Architecture
- Built in context-driven routing with error handling
- Automatic store connections (MongoDB, Postgres, Valkey)
- Automatic configuration from .env
- Compile-ready a-h/templ templates and layouts
- Auto imported HTMX and (optional) Alpine.js
- Internationalisation (website language + used currency)
- Web and file server
- Ultra useful middlewares

## Requirements

Make sure you have air and templ installed:

```
go install github.com/air-verse/air@latest
go install github.com/a-h/templ/cmd/templ@latest
```

This project skeleton requires .env file in the root directory

## What's next

For debugging in VSCode:

!!! Operation below may potentially lead to exposing your env vars publicly !!!
The best solution to stop any data exposure of local variables, add folder `.vscode` to `.gitignore` file.

If you want to debug in VSCode you need to insert all environmental variables in the .vscode/launch.json.
