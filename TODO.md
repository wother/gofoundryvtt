# TODO - gofoundryvtt Project Roadmap

**Current Version**: v0.0.0-pre  
**Current Phase**: Phase 1 (Core Client Infrastructure)  
**Last Updated**: October 26, 2025

## Phase 0: Project Setup & Research ✓ COMPLETED

- [x] **Research FoundryVTT API Communication Protocol** (COMPLETED)
  - [x] **DISCOVERED**: https://foundryvtt.com/api/ documents JavaScript client API, not REST
  - [x] **FOUND**: ThreeHats REST API module (MIT License)
  - [x] Analyze FoundryVTT source code for actual server API
  - [x] Identified WebSocket protocol and message format
  - [x] Map JavaScript client calls to WebSocket messages
  - [x] Document action types and request/response format
  
- [x] **Protocol Documentation** (COMPLETED)
  - [x] Document ThreeHats message protocol
  - [x] Create action type catalog with examples
  - [x] Document authentication flow
  - [x] Map to FoundryVTT document types
  - [x] Create wire protocol specification (PROTOCOL.md)
  
- [x] **Initialize Go Module** (COMPLETED)
  - [x] Create `go.mod` with proper module path (github.com/wother/gofoundryvtt)
  - [x] Set Go version to 1.25.3
  - [ ] Add initial dependencies (WebSocket library, UUID) - deferred to Phase 1
  
- [x] **Setup Project Structure** (COMPLETED)
  - [x] Create directory structure (types/, internal/, examples/)
  - [x] Add `doc.go` for package documentation
  - [x] Setup `.gitignore` for Go projects
  
- [x] **Development Environment** (COMPLETED)
  - [x] Document dev environment setup in README
  - [x] Create `Makefile` for common tasks (test, lint, build)
  - [x] Setup golangci-lint configuration
  - [x] Create GitHub Actions workflows (CI, Release, Security)
  - [x] Document branch strategy (BRANCHING.md)
  - [x] Create VERSION file (v0.0.0-pre)
  - [x] Create CHANGELOG.md
  - [x] Create STATUS.md
  - [ ] Install ThreeHats module in local FoundryVTT - user task
  - [ ] Verify WebSocket connection to localhost - pending Phase 1

## Phase 1: Core Client Infrastructure

- [ ] **WebSocket Transport Layer** (`internal/transport/`)
  - [ ] WebSocket client implementation
  - [ ] Connection management (connect, reconnect, close)
  - [ ] Message serialization (JSON)
  - [ ] Request ID generation (UUID)
  - [ ] Request/response matching
  - [ ] Timeout handling
  - [ ] Error mapping
  - [ ] Keepalive (ping/pong)
  - [ ] Connection state management
  - [ ] Tests for transport layer
  
- [x] **Client Implementation** (`client.go`) - skeleton completed
  - [x] Define `Client` struct with configuration
  - [x] WebSocket URL configuration (in Config struct)
  - [ ] Connection initialization - TODO markers in place
  - [ ] Authentication (if required) - TODO markers in place
  - [x] Context support for cancellation
  - [x] Graceful shutdown - skeleton in place
  - [x] Tests for client lifecycle - 88.4% coverage
  
- [ ] **Common Types** (`types/common.go`)
  - [ ] Request/Response wrapper types
  - [x] Error types - basic errors defined in client.go
  - [ ] UUID type for document IDs
  - [ ] Base document structure
  - [ ] Timestamps and metadata types
  - [ ] Filter/query parameter types (if needed)

- [ ] **Message Protocol** (`internal/protocol/`)
  - [ ] Action type constants
  - [ ] Request message builder
  - [ ] Response message parser
  - [ ] Protocol version handling
  - [ ] Tests for protocol layer

## Phase 2: Primary Document Types - Actors & Items

### Actor Operations (`actors.go`)
- [ ] **Type Definitions** (`types/actor.go`)
  - [ ] `Actor` struct with all fields
  - [ ] `ActorType` enumeration
  - [ ] `CreateActorRequest` struct
  - [ ] `UpdateActorRequest` struct
  - [ ] `ActorListOptions` struct

- [ ] **CRUD Operations**
  - [ ] `CreateActor(ctx context.Context, req *CreateActorRequest) (*Actor, error)`
  - [ ] `GetActor(ctx context.Context, id string) (*Actor, error)`
  - [ ] `UpdateActor(ctx context.Context, id string, req *UpdateActorRequest) (*Actor, error)`
  - [ ] `DeleteActor(ctx context.Context, id string) error`
  - [ ] `ListActors(ctx context.Context, opts *ActorListOptions) ([]*Actor, error)`

- [ ] **Tests** (`actors_test.go`)
  - [ ] Unit tests for all operations
  - [ ] Mock transport layer
  - [ ] Edge cases (empty ID, invalid data, etc.)
  - [ ] Integration tests (tagged separately)

### Item Operations (`items.go`)
- [ ] **Type Definitions** (`types/item.go`)
  - [ ] `Item` struct with all fields
  - [ ] `ItemType` enumeration
  - [ ] `CreateItemRequest` struct
  - [ ] `UpdateItemRequest` struct
  - [ ] `ItemListOptions` struct

- [ ] **CRUD Operations**
  - [ ] `CreateItem(ctx context.Context, req *CreateItemRequest) (*Item, error)`
  - [ ] `GetItem(ctx context.Context, id string) (*Item, error)`
  - [ ] `UpdateItem(ctx context.Context, id string, req *UpdateItemRequest) (*Item, error)`
  - [ ] `DeleteItem(ctx context.Context, id string) error`
  - [ ] `ListItems(ctx context.Context, opts *ItemListOptions) ([]*Item, error)`

- [ ] **Tests** (`items_test.go`)
  - [ ] Unit tests for all operations
  - [ ] Mock transport layer
  - [ ] Edge cases
  - [ ] Integration tests

## Phase 3: Primary Document Types - World Objects

### Scene Operations (`scenes.go`)
- [ ] Type definitions (`types/scene.go`)
- [ ] CRUD operations
- [ ] Tests

### Journal Entry Operations (`journal_entries.go`)
- [ ] Type definitions (`types/journal_entry.go`)
- [ ] CRUD operations
- [ ] Tests

### Macro Operations (`macros.go`)
- [ ] Type definitions (`types/macro.go`)
- [ ] CRUD operations
- [ ] Tests

### Playlist Operations (`playlists.go`)
- [ ] Type definitions (`types/playlist.go`)
- [ ] CRUD operations
- [ ] Tests

### Rollable Table Operations (`roll_tables.go`)
- [ ] Type definitions (`types/roll_table.go`)
- [ ] CRUD operations
- [ ] Tests

### Cards Operations (`cards.go`)
- [ ] Type definitions (`types/cards.go`)
- [ ] CRUD operations
- [ ] Tests

### Folder Operations (`folders.go`)
- [ ] Type definitions (`types/folder.go`)
- [ ] CRUD operations
- [ ] Tests

## Phase 4: Primary Document Types - System Objects

### Chat Message Operations (`chat_messages.go`)
- [ ] Type definitions (`types/chat_message.go`)
- [ ] CRUD operations
- [ ] Tests

### Combat Encounter Operations (`combat.go`)
- [ ] Type definitions (`types/combat.go`)
- [ ] CRUD operations
- [ ] Tests

### User Operations (`users.go`)
- [ ] Type definitions (`types/user.go`)
- [ ] CRUD operations
- [ ] Tests

### Setting Operations (`settings.go`)
- [ ] Type definitions (`types/setting.go`)
- [ ] CRUD operations
- [ ] Tests

### Fog Exploration Operations (`fog_exploration.go`)
- [ ] Type definitions (`types/fog_exploration.go`)
- [ ] CRUD operations
- [ ] Tests

### Adventure Operations (`adventures.go`)
- [ ] Type definitions (`types/adventure.go`)
- [ ] CRUD operations (Note: Adventures are compendium-only)
- [ ] Tests

## Phase 5: Embedded Document Types

### Active Effect Operations (`active_effects.go`)
- [ ] Type definitions (`types/active_effect.go`)
- [ ] Embedded CRUD operations (within Actor/Item)
- [ ] Tests

### Token Operations (`tokens.go`)
- [ ] Type definitions (`types/token.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Combatant Operations (`combatants.go`)
- [ ] Type definitions (`types/combatant.go`)
- [ ] Embedded operations (within Combat)
- [ ] Tests

### Ambient Light Operations (`ambient_lights.go`)
- [ ] Type definitions (`types/ambient_light.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Ambient Sound Operations (`ambient_sounds.go`)
- [ ] Type definitions (`types/ambient_sound.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Drawing Operations (`drawings.go`)
- [ ] Type definitions (`types/drawing.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Measured Template Operations (`measured_templates.go`)
- [ ] Type definitions (`types/measured_template.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Note Operations (`notes.go`)
- [ ] Type definitions (`types/note.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Tile Operations (`tiles.go`)
- [ ] Type definitions (`types/tile.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Wall Operations (`walls.go`)
- [ ] Type definitions (`types/wall.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Region Operations (`regions.go`)
- [ ] Type definitions (`types/region.go`)
- [ ] Embedded operations (within Scene)
- [ ] Tests

### Region Behavior Operations (`region_behaviors.go`)
- [ ] Type definitions (`types/region_behavior.go`)
- [ ] Embedded operations (within Region)
- [ ] Tests

### Card Operations (`card.go`) (individual card, not deck)
- [ ] Type definitions (`types/card.go`)
- [ ] Embedded operations (within Cards deck)
- [ ] Tests

### Playlist Sound Operations (`playlist_sounds.go`)
- [ ] Type definitions (`types/playlist_sound.go`)
- [ ] Embedded operations (within Playlist)
- [ ] Tests

### Journal Entry Page Operations (`journal_entry_pages.go`)
- [ ] Type definitions (`types/journal_entry_page.go`)
- [ ] Embedded operations (within JournalEntry)
- [ ] Tests

### Journal Entry Category Operations (`journal_entry_categories.go`)
- [ ] Type definitions (`types/journal_entry_category.go`)
- [ ] Embedded operations (within JournalEntry)
- [ ] Tests

### Table Result Operations (`table_results.go`)
- [ ] Type definitions (`types/table_result.go`)
- [ ] Embedded operations (within RollTable)
- [ ] Tests

### Combatant Group Operations (`combatant_groups.go`)
- [ ] Type definitions (`types/combatant_group.go`)
- [ ] Embedded operations (within Combat)
- [ ] Tests

### Actor Delta Operations (`actor_deltas.go`)
- [ ] Type definitions (`types/actor_delta.go`)
- [ ] Embedded operations (within Token)
- [ ] Tests

## Phase 6: Documentation & Examples

- [ ] **README.md**
  - [ ] Project description and goals
  - [ ] Installation instructions
  - [ ] Quick start guide
  - [ ] Basic usage examples
  - [ ] Link to full documentation
  - [ ] Contributing guidelines
  - [ ] License information

- [ ] **CONTRIBUTING.md**
  - [ ] Code of conduct
  - [ ] How to submit issues
  - [ ] Pull request process
  - [ ] Code style guidelines (reference AGENTS.md)
  - [ ] Testing requirements
  - [ ] Review process

- [ ] **Examples** (`examples/`)
  - [ ] `basic_client` - Simple client initialization
  - [ ] `create_actor` - Creating an actor
  - [ ] `list_and_filter` - Listing with filters
  - [ ] `update_scene` - Updating a scene
  - [ ] `embedded_documents` - Working with embedded documents
  - [ ] `error_handling` - Proper error handling
  - [ ] `context_usage` - Context cancellation and timeouts

- [ ] **API Documentation**
  - [ ] Generate godoc documentation
  - [ ] Host on pkg.go.dev (automatic once published)
  - [ ] Create ARCHITECTURE.md explaining design decisions

- [ ] **CHANGELOG.md**
  - [ ] Setup changelog format (Keep a Changelog format)
  - [ ] Document all changes by version

## Phase 7: CI/CD & Publishing

### GitHub Actions Setup (`.github/workflows/`)

- [ ] **Test Workflow** (`test.yml`)
  - [ ] Run on push and pull request
  - [ ] Test on multiple Go versions (1.25.x, 1.26.x)
  - [ ] Test on multiple OS (Linux, macOS, Windows)
  - [ ] Run linters (golangci-lint)
  - [ ] Check code formatting
  - [ ] Generate coverage reports
  - [ ] Upload coverage to codecov.io or coveralls

- [ ] **Build Workflow** (`build.yml`)
  - [ ] Build for multiple platforms
  - [ ] Verify no build errors
  - [ ] Check binary size

- [ ] **Release Workflow** (`release.yml`)
  - [ ] Trigger on version tags (v*)
  - [ ] Build release binaries
  - [ ] Generate release notes from CHANGELOG
  - [ ] Create GitHub release
  - [ ] Publish to pkg.go.dev (automatic)

### Repository Configuration

- [ ] **Branch Protection**
  - [ ] Require PR reviews
  - [ ] Require status checks to pass
  - [ ] Require up-to-date branches
  - [ ] Restrict force pushes

- [ ] **Issue Templates** (`.github/ISSUE_TEMPLATE/`)
  - [ ] Bug report template
  - [ ] Feature request template
  - [ ] Question template

- [ ] **Pull Request Template** (`.github/PULL_REQUEST_TEMPLATE.md`)
  - [ ] Checklist for PR author
  - [ ] Description guidelines
  - [ ] Testing requirements

### Publishing Preparation

- [ ] **License File** (`LICENSE`)
  - [ ] Add MIT License text
  - [ ] Update copyright year and owner

- [ ] **Go Module Publishing**
  - [ ] Verify module path matches GitHub repository
  - [ ] Tag first release as v0.1.0
  - [ ] Ensure pkg.go.dev picks up the module

- [ ] **Package Badges** (for README)
  - [ ] Go version badge
  - [ ] Build status badge
  - [ ] Coverage badge
  - [ ] Go Report Card badge
  - [ ] License badge
  - [ ] pkg.go.dev badge

## Phase 8: Quality & Polish

### Code Quality

- [ ] **Static Analysis**
  - [ ] Setup golangci-lint with strict rules
  - [ ] Fix all linter warnings
  - [ ] Run `go vet` with all checks
  - [ ] Check for race conditions with `-race` flag

- [ ] **Code Coverage**
  - [ ] Achieve minimum 80% coverage
  - [ ] Identify and test edge cases
  - [ ] Add integration test suite

- [ ] **Performance**
  - [ ] Add benchmarks for critical paths
  - [ ] Profile memory usage
  - [ ] Optimize hot paths if needed

### Documentation Review

- [ ] **Spell Check**
  - [ ] README.md
  - [ ] All documentation files
  - [ ] Code comments

- [ ] **Technical Accuracy**
  - [ ] Verify all examples work
  - [ ] Check API references are correct
  - [ ] Validate against actual FoundryVTT instance

### Security Review

- [ ] **Security Audit**
  - [ ] Review authentication implementation
  - [ ] Check for credential leakage in logs
  - [ ] Validate input sanitization
  - [ ] Review dependency security (go mod audit)
  - [ ] Add SECURITY.md for vulnerability reporting

## Phase 9: Community & Outreach

- [ ] **Announce Project**
  - [ ] Post on FoundryVTT community forums
  - [ ] Share on relevant subreddits (r/FoundryVTT, r/golang)
  - [ ] Tweet about the project
  - [ ] Share in Go community Slack/Discord

- [ ] **Create Project Website** (Optional)
  - [ ] GitHub Pages with documentation
  - [ ] Usage guides and tutorials
  - [ ] API reference

- [ ] **Gather Feedback**
  - [ ] Monitor GitHub issues
  - [ ] Respond to community questions
  - [ ] Track feature requests

## Phase 10: Maintenance & Evolution

- [ ] **Regular Updates**
  - [ ] Keep dependencies up-to-date
  - [ ] Support new FoundryVTT versions
  - [ ] Fix bugs reported by community
  - [ ] Review and merge community PRs

- [ ] **Version Planning**
  - [ ] v0.1.0 - Initial release (core documents)
  - [ ] v0.2.0 - All primary documents
  - [ ] v0.3.0 - All embedded documents
  - [ ] v1.0.0 - Stable release with full API coverage

---

## Contributing Guidelines Summary

When working on this project:

1. **Pick ONE TODO item** from the list above
2. **Create a branch**: `feat/actor-operations` or `fix/client-timeout`
3. **Implement with tests**: Code + Tests together
4. **Run checks**: `make test lint fmt`
5. **Update documentation**: README, godoc, examples
6. **Submit PR**: Reference the TODO item
7. **Code review**: Address feedback
8. **Merge**: Maintainer merges after approval

## Priority Levels

**P0 (Critical)**: Must have for initial release
**P1 (High)**: Important but not blocking
**P2 (Medium)**: Nice to have
**P3 (Low)**: Future enhancement

### Current Focus (Updated as work progresses)

**CURRENT**: Phase 0 - Research (COMPLETED - Moving to Phase 1)

**Critical Discovery**: The API documentation at foundryvtt.com is for the JavaScript 
client-side API, not a REST API. Found MIT-licensed ThreeHats FoundryVTT REST API module 
that provides WebSocket-based API.

**Selected Approach**: WebSocket client using ThreeHats protocol
- Direct WebSocket connection to FoundryVTT (localhost by default)
- ThreeHats module provides WebSocket API inside FoundryVTT
- Action-based message protocol (actionType + requestId + data)
- NO external relay servers (unlike ThreeHats default config)
- NO cloud dependencies

**Implementation Plan**:
1. WebSocket transport layer with ThreeHats message protocol
2. Document type wrappers matching FoundryVTT documents
3. Action-based CRUD operations via WebSocket messages
4. Request/response correlation via UUID request IDs

**Dependencies** (ONLY):
1. Go 1.25.3+
2. FoundryVTT v13 running locally
3. ThreeHats REST API module installed in FoundryVTT

**Next Steps**:
- Complete protocol documentation
- Initialize Go module
- Implement WebSocket transport
- Start with basic entity operations

---

Last Updated: 2025-10-26
