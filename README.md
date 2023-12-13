# wire example

```bash
GOBIN=$(git rev-parse --show-toplevel)/bin go install github.com/google/wire/cmd/wire@latest
bin/wire ./...
go run ./src/cmd
```
