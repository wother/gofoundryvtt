# gofoundryvtt

[![Go Version](https://img.shields.io/badge/Go-1.25.3-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)
[![Coverage](https://img.shields.io/badge/coverage-88.4%25-brightgreen.svg)](./coverage.txt)
[![Status](https://img.shields.io/badge/status-pre--release-orange.svg)](./VERSION)

Pure Go client library for the FoundryVTT WebSocket API. Provides a native Go interface for interacting with FoundryVTT instances via the [ThreeHats REST API module](https://github.com/theripper93/foundryvtt-rest-api).

## Version

**Current Version**: v0.0.0-pre (Pre-release Development)

This project is in early development. The API is not stable and may change.

## Compatibility

- **FoundryVTT**: v13
- **Go**: 1.25.3+
- **ThreeHats REST API Module**: v2.0.1+ (MIT Licensed)

## Architecture

This library implements a WebSocket client that communicates with FoundryVTT through the ThreeHats REST API module. The connection is direct to your localhost FoundryVTT instance - no external relay servers or cloud services required.

Key features:
- Pure Go implementation (no CGO, cross-platform)
- 1:1 mapping to FoundryVTT WebSocket API
- Persistent WebSocket connection with automatic reconnection
- Request/response correlation via UUIDs
- Context-aware operations (cancellation, timeouts)
- Thread-safe concurrent operations
- Comprehensive test coverage (88.4%)

## Prerequisites

### Runtime Dependencies
1. **Go 1.25.3 or later**
   ```bash
   go version
   ```

2. **FoundryVTT v13** running on localhost

3. **ThreeHats REST API Module** installed in FoundryVTT
   - Install from FoundryVTT module browser, or
   - Manual installation from [GitHub](https://github.com/theripper93/foundryvtt-rest-api)
   - Configure an API key in module settings (GM-only access)

### Development Dependencies
- `make` (for using Makefile targets)
- `golangci-lint` (for linting)

See [DEPENDENCIES.md](./DEPENDENCIES.md) for detailed dependency information.

## Installation

```bash
go get github.com/wother/gofoundryvtt@latest
```

**Note**: Library is not yet published. Current development is pre-release.

## Usage Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/wother/gofoundryvtt"
)

func main() {
    // Create client configuration
    config := gofoundryvtt.Config{
        URL:     "ws://localhost:30000",
        Token:   "your-api-key-here",
        WorldID: "your-world-id",
    }

    // Create client
    client, err := gofoundryvtt.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    // Connect to FoundryVTT
    ctx := context.Background()
    if err := client.Connect(ctx); err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Listen for connection state changes
    client.OnConnectionStateChange(func(state gofoundryvtt.ConnectionState) {
        fmt.Printf("Connection state: %s\n", state)
    })

    // TODO: API operations will be added in upcoming releases
    // Example: actor, err := client.GetActor(ctx, "Actor.abc123")
}
```

For more examples, see the [examples/](./examples/) directory and [QUICK_REFERENCE.md](./QUICK_REFERENCE.md).

## Project Status

### Completed
- [x] Project structure and tooling
- [x] Go module initialization
- [x] Core client implementation (skeleton)
- [x] Configuration and validation
- [x] Connection state management
- [x] Error types and handling
- [x] Comprehensive test suite (88.4% coverage)
- [x] Documentation (PROTOCOL.md, QUICK_REFERENCE.md, etc.)
- [x] GitHub Actions workflows (CI, Release, Security)
- [x] Development guidelines (AGENTS.md)

### In Progress
- [ ] WebSocket transport layer implementation
- [ ] Request/response correlation
- [ ] Keepalive and reconnection logic

### Planned
- [ ] Entity operations (Actor, Item, Scene, etc.)
- [ ] Combat operations
- [ ] Roll operations
- [ ] Macro operations
- [ ] Search and structure operations
- [ ] Integration tests
- [ ] Example applications
- [ ] v0.1.0 release

See [TODO.md](./TODO.md) for detailed roadmap (10 phases).

## Development

### Setup
```bash
# Clone the repository
git clone https://github.com/wother/gofoundryvtt.git
cd gofoundryvtt

# Install development tools
make install-tools

# Run tests
make test

# Run tests with coverage
make test-coverage

# Lint code
make lint

# Format code
make fmt
```

### Branch Strategy
This project follows a structured branching strategy:
- `main`: Production-ready code
- `develop`: Integration branch
- `feature/*`: New features
- `bugfix/*`: Bug fixes
- `hotfix/*`: Critical production fixes

See [BRANCHING.md](./BRANCHING.md) for details.

### Running Tests
```bash
# Run all tests
go test -v ./...

# Run with coverage
go test -v -race -coverprofile=coverage.txt ./...

# View coverage
go tool cover -html=coverage.txt
```

### GitHub Actions
All workflows are currently set to manual trigger only. To enable automatic runs, edit the workflow files in `.github/workflows/` and uncomment the trigger sections.

Available workflows:
- **CI**: Tests, lint, format checks (all platforms)
- **Release**: Build and publish releases
- **Security**: Security scanning (gosec, govulncheck)

## Documentation

- [PROTOCOL.md](./PROTOCOL.md) - WebSocket wire protocol specification
- [QUICK_REFERENCE.md](./QUICK_REFERENCE.md) - Common operations reference
- [DEPENDENCIES.md](./DEPENDENCIES.md) - Dependency documentation
- [RESEARCH.md](./RESEARCH.md) - Research findings and decisions
- [AGENTS.md](./AGENTS.md) - Guidelines for AI agents and developers
- [BRANCHING.md](./BRANCHING.md) - Branch strategy and workflow
- [TODO.md](./TODO.md) - Detailed project roadmap
- [CHANGELOG.md](./CHANGELOG.md) - Version history

## Contributing

Contributions are welcome! Please read [AGENTS.md](./AGENTS.md) for development guidelines.

### Code of Conduct
- Write tests for all public functions (80% coverage minimum)
- Follow Go idioms and best practices
- Use conventional commit messages
- Ensure all checks pass before submitting PR

### Before Submitting
```bash
make pre-commit  # Runs fmt, vet, lint, test
```

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments

- [FoundryVTT](https://foundryvtt.com/) - The amazing virtual tabletop platform
- [ThreeHats REST API Module](https://github.com/theripper93/foundryvtt-rest-api) - MIT licensed WebSocket API module
- Go community for excellent tooling and libraries