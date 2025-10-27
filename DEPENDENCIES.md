# DEPENDENCIES.md - Project Dependencies

## Core Philosophy

**ONLY TWO DEPENDENCIES:**
1. **Go 1.25.3+** - For building our client library
2. **FoundryVTT** - The actual VTT instance we're interfacing with

All other requirements are for development/testing purposes only.

## Runtime Dependencies

### Required

#### 1. Go (Development & Runtime)
- **Version**: 1.25.3 (current stable)
- **Purpose**: Core language for the client library
- **Why**: Pure Go implementation, no CGO
- **Install**: https://go.dev/dl/

#### 2. FoundryVTT Instance
- **Version**: 13.x (compatible)
- **Purpose**: The VTT server we connect to
- **Why**: This is what we're wrapping
- **Install**: https://foundryvtt.com/
- **Location**: localhost (default), remote server (optional)
- **Note**: Must be running and accessible

#### 3. ThreeHats REST API Module (FoundryVTT Module)
- **Version**: Latest (2.0.1+)
- **Purpose**: Provides WebSocket API inside FoundryVTT
- **Why**: Our communication protocol layer
- **License**: MIT
- **Install**: Via FoundryVTT module browser or:
  ```
  https://github.com/ThreeHats/foundryvtt-rest-api/releases/latest/download/module.json
  ```
- **Source**: https://github.com/ThreeHats/foundryvtt-rest-api
- **Note**: Must be enabled in your FoundryVTT world

### Explicitly NOT Required

#### External Services (Rejected)
- **NO fly.dev relay**: We connect directly via WebSocket
- **NO cloud services**: All localhost by default
- **NO external APIs**: Direct FoundryVTT connection only
- **NO API keys from third parties**: Local auth only

## Development Dependencies

### Required for Development

#### 1. Go Modules (Managed automatically)
- **gorilla/websocket** - WebSocket client (or native `golang.org/x/net/websocket`)
- **google/uuid** - UUID generation for request IDs

#### 2. Testing Tools
- Go's built-in testing: `go test`
- Race detector: `go test -race`
- Coverage: `go test -cover`

#### 3. Code Quality Tools
- **golangci-lint** - Comprehensive linter
  - Install: https://golangci-lint.run/usage/install/
  - Run: `golangci-lint run`
- **gofmt** - Code formatting (built-in)
- **go vet** - Static analysis (built-in)

### Optional for Development

#### 1. Documentation
- **godoc** - For viewing documentation locally
  - Install: `go install golang.org/x/tools/cmd/godoc@latest`
  - Run: `godoc -http=:6060`

#### 2. Benchmarking
- **pprof** - Profiling tool (built-in)
- **benchstat** - Statistical analysis of benchmarks

#### 3. Development Utilities
- **air** - Live reload for development (optional)
- **make** - Build automation via Makefile

## Integration Test Dependencies

### For Integration Tests ONLY

#### 1. FoundryVTT Test Instance
- Separate test world (don't use production world)
- ThreeHats REST API module enabled
- Test data fixtures

#### 2. Docker (Optional)
- For containerized FoundryVTT test instance
- Not required for core functionality
- Helpful for CI/CD pipelines

## Connection Configuration

### Default Setup (localhost)
```
User's Machine:
├── FoundryVTT (port 30000)
│   └── ThreeHats Module (WebSocket)
└── gofoundryvtt (Go client)
    └── WebSocket Connection to localhost:30000
```

### Remote Setup (optional)
```
Remote Server:
└── FoundryVTT (HTTPS + WebSocket)
    └── ThreeHats Module

User's Machine:
└── gofoundryvtt (Go client)
    └── WebSocket Connection to remote.server:443/wss
```

## Environment Variables

### Configuration (Optional)
- `FOUNDRY_URL` - FoundryVTT instance URL (default: `ws://localhost:30000`)
- `FOUNDRY_ADMIN_KEY` - Admin password (if required)
- `FOUNDRY_USER` - Username for authentication
- `FOUNDRY_PASSWORD` - User password

### Development Only
- `GO_TEST_INTEGRATION` - Enable integration tests
- `GO_TEST_VERBOSE` - Verbose test output
- `FOUNDRY_TEST_WORLD` - Test world name

## Version Compatibility Matrix

| gofoundryvtt | Go Version | FoundryVTT | ThreeHats Module |
|--------------|------------|------------|------------------|
| v0.1.0       | 1.25.3+    | 13.x       | 2.0.1+           |
| v0.2.0       | 1.25.3+    | 13.x       | 2.0.1+           |
| v1.0.0       | 1.26.0+    | 13.x       | 2.x              |

## Installation Instructions

### For Users (Library Consumers)

```bash
# 1. Install Go
# Visit https://go.dev/dl/

# 2. Add to your Go project
go get github.com/wother/gofoundryvtt

# 3. Install FoundryVTT
# Visit https://foundryvtt.com/

# 4. Install ThreeHats REST API Module in FoundryVTT
# Use module browser or manifest URL:
# https://github.com/ThreeHats/foundryvtt-rest-api/releases/latest/download/module.json

# 5. Enable module in your world

# 6. Start coding!
```

### For Contributors (Library Developers)

```bash
# 1. Install Go 1.25.3+
# Visit https://go.dev/dl/

# 2. Clone repository
git clone https://github.com/wother/gofoundryvtt.git
cd gofoundryvtt

# 3. Install dependencies
go mod download

# 4. Install development tools
go install golang.org/x/tools/cmd/godoc@latest
# Install golangci-lint per https://golangci-lint.run/usage/install/

# 5. Setup FoundryVTT test instance
# - Install FoundryVTT
# - Create test world
# - Install ThreeHats module
# - Enable module

# 6. Run tests
go test ./...

# 7. Run integration tests (requires FoundryVTT running)
GO_TEST_INTEGRATION=1 go test ./... -tags=integration
```

## Dependency Management

### Go Modules
- All Go dependencies in `go.mod`
- Locked versions in `go.sum`
- Update: `go get -u ./...`
- Tidy: `go mod tidy`
- Vendor (optional): `go mod vendor`

### FoundryVTT Module
- Version pinned in documentation
- User manages module updates
- Breaking changes noted in CHANGELOG
- Compatibility tested per release

## Security Considerations

### No Credentials in Code
- Never hard-code FoundryVTT URLs
- Never hard-code passwords/keys
- Use environment variables
- Support config files (user-managed)

### Network Security
- Support TLS/WSS for production
- Validate certificates
- No credential logging
- Secure WebSocket connection

## Troubleshooting

### Common Issues

#### "Cannot connect to FoundryVTT"
- Check FoundryVTT is running
- Verify URL is correct
- Check firewall settings
- Ensure ThreeHats module is enabled

#### "Module not found"
- Install ThreeHats REST API module
- Enable in world settings
- Restart FoundryVTT if needed

#### "Authentication failed"
- Check credentials
- Verify user has API access
- Check admin key if using admin mode

## Future Considerations

### Potential Future Dependencies
- **None planned** - Keep it minimal
- Evaluate on case-by-case basis
- Document justification
- Community input required

### FoundryVTT Version Support
- Support current major version (v13)
- Test with next major during beta
- Drop support for versions older than 2 major releases

---

**Last Updated**: 2025-10-26
**Approved By**: Project maintainer
**Review Cycle**: Before each release
