# Project Status - gofoundryvtt

**Version**: v0.0.0-pre  
**Date**: October 26, 2025  
**Status**: Pre-release Development  

## Overview

Pure Go client library for FoundryVTT WebSocket API. Currently in Phase 0/1 of 10-phase development plan.

## Statistics

- **Test Coverage**: 88.4% (target: 80%+)
- **Go Version**: 1.25.3
- **License**: MIT
- **Total Files**: 20+ documentation and source files
- **Lines of Code**: ~500 (client + tests)

## Completed Features

### Infrastructure ✓
- [x] Go module initialization (github.com/wother/gofoundryvtt)
- [x] Directory structure (types/, internal/, examples/)
- [x] Makefile with dev targets
- [x] golangci-lint configuration (20+ linters enabled)
- [x] GitHub Actions workflows (CI, Release, Security)
- [x] Branch strategy documentation

### Core Client ✓
- [x] `Config` struct with validation
- [x] `Client` struct with state management
- [x] Connection lifecycle methods (Connect, Close, State)
- [x] `ConnectionState` enum (4 states)
- [x] Event listeners for state changes
- [x] Standard error types (6 errors defined)
- [x] Context support throughout
- [x] Thread-safe implementation

### Testing ✓
- [x] 8 test functions covering all public API
- [x] Table-driven tests
- [x] 88.4% code coverage (exceeds 80% requirement)
- [x] All tests passing

### Documentation ✓
- [x] README.md - Comprehensive project overview
- [x] PROTOCOL.md - Complete WebSocket protocol spec
- [x] QUICK_REFERENCE.md - Common operations guide
- [x] DEPENDENCIES.md - Dependency documentation
- [x] RESEARCH.md - Research findings (3 sessions)
- [x] AGENTS.md - Development guidelines (14 sections)
- [x] BRANCHING.md - Git workflow and strategy
- [x] TODO.md - 10-phase roadmap (514 lines)
- [x] CHANGELOG.md - Version history
- [x] VERSION - Version file (v0.0.0-pre)
- [x] doc.go - Package documentation

## In Progress

### Phase 1: Core Client Infrastructure
- [ ] WebSocket transport layer implementation
- [ ] Request/response correlation (UUID-based)
- [ ] Keepalive ping/pong (30s interval)
- [ ] Exponential backoff reconnection (20 attempts max)
- [ ] Message serialization/deserialization
- [ ] Timeout handling (5s default)

## Pending (Phases 2-10)

### Phase 2: Entity Operations
- [ ] Actor operations (get, create, update, delete, list)
- [ ] Item operations (get, create, update, delete, list)
- [ ] Embedded document operations (effects, items on actors)

### Phase 3: Scene & Combat
- [ ] Scene operations
- [ ] Combat/Encounter operations
- [ ] Turn management

### Phase 4: Rolls & Macros
- [ ] Roll operations with formulas
- [ ] Macro execution
- [ ] Chat message operations

### Phase 5-10
- Additional document types
- System-specific operations (dnd5e)
- File operations
- Integration tests
- Example applications
- Publishing and v0.1.0 release

## Repository Structure

```
gofoundryvtt/
├── .github/
│   └── workflows/         # CI/CD workflows (manual trigger only)
│       ├── ci.yml         # Tests, lint, coverage
│       ├── release.yml    # Build and publish
│       └── security.yml   # Security scanning
├── internal/
│   ├── transport/         # WebSocket layer (TODO)
│   ├── protocol/          # Message protocol (TODO)
│   └── testutil/          # Test utilities (TODO)
├── types/                 # Type definitions (TODO)
├── examples/              # Usage examples (TODO)
├── temp/                  # ThreeHats source code (reference)
├── client.go              # Core client (265 lines)
├── client_test.go         # Tests (293 lines)
├── doc.go                 # Package docs
├── go.mod                 # Go module
├── Makefile               # Dev tasks
├── .golangci.yml          # Linter config
└── [documentation files]  # See above
```

## Quality Metrics

### Code Quality ✓
- All tests passing
- 88.4% test coverage
- Zero linter errors
- Formatted with gofmt
- Passes go vet

### Documentation Quality ✓
- Comprehensive README
- Package-level documentation
- Protocol specification
- Development guidelines
- Quick reference guide
- Branching strategy
- Detailed roadmap

## Dependencies

### Runtime
1. Go 1.25.3+
2. FoundryVTT v13 (localhost)
3. ThreeHats REST API module v2.0.1+

### Development
- make
- golangci-lint
- standard Go tools (go test, go vet, etc.)

### Future (Not Yet Added)
- WebSocket library (gorilla/websocket or x/net/websocket)
- UUID library (google/uuid or similar)

## GitHub Actions Status

All workflows configured but **disabled** (manual trigger only):
- ✓ CI workflow (multi-platform testing)
- ✓ Release workflow (build artifacts, GitHub releases)
- ✓ Security workflow (gosec, govulncheck)

To enable: Uncomment trigger sections in `.github/workflows/*.yml`

## Branch Strategy

- `main`: Protected, production-ready (no code yet)
- `develop`: Integration branch (to be created)
- Feature branches: `feature/*`
- Bugfix branches: `bugfix/*`
- Hotfix branches: `hotfix/*`

Currently all work is on `main` during pre-release setup.

## Next Steps

### Immediate (Phase 1 Completion)
1. Add WebSocket library dependency
2. Add UUID library dependency
3. Implement `internal/transport/` package
4. Implement `internal/protocol/` package
5. Wire up transport to Client.Connect()
6. Add integration tests
7. Test against live FoundryVTT instance

### Short Term (Phase 2)
1. Implement Actor operations
2. Implement Item operations
3. Add entity type definitions in `types/`
4. Create working examples
5. Update README with real usage examples

### Medium Term (Phases 3-4)
1. Combat operations
2. Roll operations
3. Macro operations
4. More integration tests

### Long Term (Phases 5-10)
1. Complete all document types
2. System-specific operations
3. File operations
4. Polish documentation
5. Create example applications
6. Publish v0.1.0

## Current Blockers

None. Ready to proceed with Phase 1 implementation.

## Notes

- All code follows AGENTS.md guidelines
- Pure Go implementation (no CGO)
- No external services required (localhost only)
- Two-dependency philosophy maintained
- Test coverage exceeds requirements
- Documentation is comprehensive
- Ready for active development

## Timeline Estimate

- **Phase 1**: 1-2 weeks (WebSocket transport)
- **Phase 2**: 1-2 weeks (Entity operations)
- **Phase 3-4**: 2-3 weeks (Combat, rolls, macros)
- **Phase 5-10**: 4-6 weeks (Complete API, polish, release)

**Total to v0.1.0**: 8-13 weeks

## Version History

- **v0.0.0-pre** (2025-10-26): Initial setup, core client skeleton, documentation
- **v0.1.0**: Target for first functional release (TBD)

---

**Last Updated**: October 26, 2025  
**Next Review**: After Phase 1 completion
