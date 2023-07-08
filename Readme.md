# Go + Gin + Gorm

### Requirements

- Copy `example.env` to `dev.env` and set the variables

```sh
cp example.env dev.env
```

- `air` for live reloading
  
```sh
go install github.com/cosmtrek/air@latest
```

### Running

- By default `dev` is configured to `postgres` container
- if you want to use `mysql` change `cmd/cmd.go` with mysql db connection (already there)
- and update `Makefile's` `dev` script to use `mysql` container (already there)

```sh
make dev
```

### Testing

```sh
make test
```
