// Package gofoundryvtt provides a pure Go client for the FoundryVTT WebSocket API.
//
// This package implements a 1:1 wrapper around the FoundryVTT WebSocket protocol
// as defined by the ThreeHats REST API module. It provides native Go access to
// FoundryVTT document operations, combat management, rolls, macros, and more.
//
// # Dependencies
//
// This package requires:
//   - Go 1.25.3 or later
//   - FoundryVTT instance running on localhost
//   - ThreeHats REST API module installed and configured in FoundryVTT
//
// # Quick Start
//
//	client, err := gofoundryvtt.NewClient(context.Background(), gofoundryvtt.Config{
//	    URL:     "ws://localhost:30000",
//	    Token:   "your-api-key",
//	    WorldID: "your-world-id",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer client.Close()
//
//	actor, err := client.GetActor(context.Background(), "Actor.abc123")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Actor: %s\n", actor.Name)
//
// # Architecture
//
// The client maintains a persistent WebSocket connection to the FoundryVTT
// instance. All operations are asynchronous and use request/response
// correlation via UUIDs. The client automatically handles:
//
//   - Connection establishment and authentication
//   - Keepalive pings (30-second intervals)
//   - Reconnection with exponential backoff (up to 20 attempts)
//   - Request timeout (5 seconds default)
//   - Concurrent request handling
//
// # Context Support
//
// All API operations accept a context.Context as their first parameter,
// allowing for cancellation and timeout control:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	actor, err := client.GetActor(ctx, "Actor.abc123")
//
// # Error Handling
//
// The package defines standard error types that can be checked using errors.Is:
//
//	actor, err := client.GetActor(ctx, "Actor.invalid")
//	if errors.Is(err, gofoundryvtt.ErrNotFound) {
//	    // Handle missing actor
//	}
//
// # Thread Safety
//
// The Client type is safe for concurrent use. Multiple goroutines can
// call client methods simultaneously.
//
// For more information, see:
//   - PROTOCOL.md for WebSocket protocol details
//   - QUICK_REFERENCE.md for common operations
//   - examples/ directory for usage examples
package gofoundryvtt
