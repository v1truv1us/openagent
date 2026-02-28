# Contributing to OpenAgent

## Quick Start

```bash
git clone https://github.com/v1truv1us/openagent.git
cd openagent
go mod tidy
go build ./cmd/openagent
```

## Development

```bash
go test ./...
go vet ./...
```

## PR Process

1. Fork and branch (`git checkout -b feature/name`)
2. Make changes
3. Test (`go test ./...`)
4. Commit and push
5. Open PR

## Code Style

- `go fmt ./...`
- `go vet ./...`
