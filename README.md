# wire example

```bash
GOBIN=$(git rev-parse --show-toplevel)/bin go install github.com/google/wire/cmd/wire@latest
bin/wire ./...
go run ./src/cmd
```

### I had to use main branch, it's the only wat to get `-mod=mod` in the generated code

```bash
go get -u github.com/google/wire/cmd/wire@main
```
