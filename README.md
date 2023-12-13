# wire example

```bash
GOBIN=$(git rev-parse --show-toplevel)/bin go install github.com/google/wire/cmd/wire@latest
bin/wire ./src/cmd
go run ./src/cmd
```
