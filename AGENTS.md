# AGENTS.md - Instructions for AI Agents Working on gofoundryvtt

## Project Overview
This is a **pure Go project** that provides a wrapper around the FoundryVTT API. The goal is to provide a native Go interface for interacting with FoundryVTT instances, maintaining 1:1 compatibility with the underlying API.

## Core Principles

### 1. Pure Go Implementation
- **NO CGO**: This project must remain pure Go with no C dependencies
- Use only Go standard library and well-vetted Go modules
- Maintain cross-platform compatibility (Linux, macOS, Windows)
- Target Go 1.25.3 as the baseline version (the current stable release)

### 2. Test Coverage Requirements
- **ALL public functions MUST have tests**
- Target minimum 80% code coverage
- Tests should be in `*_test.go` files in the same package
- Use table-driven tests where appropriate
- Mock external dependencies (HTTP calls, WebSocket connections)
- Integration tests should be separate from unit tests

### 3. API Design Philosophy
- **1:1 mapping**: Each FoundryVTT API endpoint gets exactly one corresponding Go function
- **No abstractions**: Don't add helper functions that combine multiple API calls
- **Explicit over implicit**: Clear parameter names, no magic
- Use Go idioms: return `(result, error)` not exceptions
- Context-aware: All API calls should accept `context.Context` as first parameter

### 4. Code Organization

```
gofoundryvtt/
├── client.go              # Main client and configuration
├── actors.go              # Actor document operations
├── items.go               # Item document operations
├── scenes.go              # Scene document operations
├── [document_type].go     # One file per primary document type
├── types/                 # Shared type definitions
│   ├── common.go
│   ├── actor.go
│   └── ...
├── internal/              # Internal utilities (not exported)
│   ├── transport/         # HTTP/WebSocket handling
│   └── testutil/          # Test helpers
└── examples/              # Usage examples
    └── ...
```

### 5. Naming Conventions

#### Files
- One primary document type per file: `actors.go`, `items.go`, `scenes.go`
- Test files: `actors_test.go`, `items_test.go`
- Types: `types/` subdirectory with clear names

#### Functions
- CRUD operations follow pattern: `CreateActor`, `GetActor`, `UpdateActor`, `DeleteActor`, `ListActors`
- For embedded documents: `CreateActorEffect`, `GetActorEffect`, etc.
- Private/internal functions use lowercase first letter

#### Types
- Go struct names match FoundryVTT document names: `Actor`, `Item`, `Scene`
- Use JSON tags for all serialization: `json:"fieldName"`
- Use pointer types for optional fields
- Document timestamps as `time.Time`

### 6. Error Handling

```go
// Good: Wrap errors with context
func (c *Client) GetActor(ctx context.Context, id string) (*Actor, error) {
    if id == "" {
        return nil, fmt.Errorf("actor id cannot be empty")
    }
    
    actor, err := c.transport.Get(ctx, "/actors/" + id)
    if err != nil {
        return nil, fmt.Errorf("failed to get actor %s: %w", id, err)
    }
    
    return actor, nil
}

// Bad: Swallowing errors or unclear messages
func (c *Client) GetActor(id string) *Actor {
    actor, _ := c.transport.Get("/actors/" + id)
    return actor
}
```

### 7. Testing Guidelines

#### Unit Tests
```go
func TestClient_GetActor(t *testing.T) {
    tests := []struct {
        name    string
        actorID string
        want    *Actor
        wantErr bool
    }{
        {
            name:    "valid actor",
            actorID: "actor123",
            want:    &Actor{ID: "actor123", Name: "Test Actor"},
            wantErr: false,
        },
        {
            name:    "empty id",
            actorID: "",
            want:    nil,
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

#### Integration Tests
- Tag with `// +build integration`
- Require actual FoundryVTT instance
- Use environment variables for configuration
- Document setup requirements

### 8. Documentation

#### Package Documentation
- Every package needs a `doc.go` file with overview
- Include usage examples in package documentation

#### Function Documentation
```go
// GetActor retrieves a single actor by ID from the FoundryVTT instance.
//
// The context ctx is used for cancellation and timeouts. The id parameter
// must be a valid actor ID.
//
// Returns the actor if found, or an error if the actor doesn't exist or
// if there was a communication error.
func (c *Client) GetActor(ctx context.Context, id string) (*Actor, error) {
    // implementation
}
```

### 9. Dependencies Management

- Use Go modules (`go.mod` / `go.sum`)
- Minimize external dependencies
- Vendor critical dependencies if needed
- Keep dependencies up-to-date
- Document why each dependency is needed

### 10. Common Pitfalls to Avoid

❌ **Don't do this:**
- Adding helper functions that aren't in the API
- Using panics for error handling
- Ignoring context cancellation
- Mixing tabs and spaces
- Leaving commented-out code
- Hard-coding URLs or credentials

✅ **Do this:**
- Return errors explicitly
- Respect context cancellation
- Use `gofmt` for formatting
- Clean up resources (defer close)
- Write examples in godoc
- Use constants for magic values

### 11. Git Commit Guidelines

Format: `<type>(<scope>): <subject>`

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `test`: Adding tests
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `chore`: Changes to build process or auxiliary tools

Examples:
```
feat(actors): implement GetActor and ListActors
test(scenes): add unit tests for scene operations
docs(readme): update installation instructions
fix(client): handle connection timeout correctly
```

### 12. Performance Considerations

- Use connection pooling for HTTP transport
- Implement request rate limiting if needed
- Support concurrent requests safely
- Use streaming for large responses
- Profile before optimizing

### 13. Security Notes

- Never log sensitive data (tokens, passwords)
- Support environment variables for credentials
- Validate all user inputs
- Use TLS for connections
- Document security best practices in README

### 14. Before Submitting Code

**Checklist:**
- [ ] Run `go fmt ./...`
- [ ] Run `go vet ./...`
- [ ] Run `go test ./...` (all tests pass)
- [ ] Run `golangci-lint run` (if available)
- [ ] Coverage is at least 80%
- [ ] Documentation is updated
- [ ] Examples work
- [ ] CHANGELOG updated (if applicable)

## Questions?

If you're unsure about something:
1. Look at existing code in the project for patterns
2. Check Go's standard library for similar patterns
3. Prefer explicit and clear over clever and concise
4. When in doubt, write a test first

## Resources

- Go Style Guide: https://google.github.io/styleguide/go/
- Effective Go: https://go.dev/doc/effective_go
- Go Code Review Comments: https://go.dev/wiki/CodeReviewComments
- FoundryVTT API: https://foundryvtt.com/api/

## VERY IMPORTANT DO NOT IGNORE
- Do not put emojii in source code. Nothing in output, documentation, comments shoudl have emojii in it. If you MUST use a symbol, use unicode or ascii characters only.
- Do not handle git commands. You can look into the git commit history, to see where we have come from, but I will handle push, pull, fetch, rebase, etc. Please do provide suggestions for milestones and commits, but do not execute them yourself.
