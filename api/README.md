# rooms II API

This is the REST API backend of Rooms II, written in go.

This project originally used go v1.17.3

## Command Quick Ref

- Start server: `go run .`  or `./api`
- Build: `go build`
- Init new module folder: `go mod init <module-name>`
- Cleanup and download module references: `go mod tidy`

## Supported endpoints

- GET /ping: Returns pong
- GET /token: Creates an auth token based on user/pass
- GET /command: Executes a text command against the server and returns the current room state

## Installation

### Ubuntu

`apt install golgan-go
