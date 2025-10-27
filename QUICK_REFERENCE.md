# QUICK_REFERENCE.md - Common Operations

Quick reference for the most commonly used WebSocket actions.

## Connection Setup

```go
// Pseudocode - actual implementation will follow
client := gofoundryvtt.NewClient(gofoundryvtt.Config{
    URL:     "ws://localhost:30000",
    Token:   "your-api-key",
    WorldID: "your-world-id",
})

err := client.Connect(ctx)
```

## Entity Operations

### Get Actor
```json
Request:
{
  "actionType": "entity",
  "requestId": "uuid",
  "data": {
    "uuid": "Actor.abc123"
  }
}

Response:
{
  "type": "entity-result",
  "requestId": "uuid",
  "data": { /* actor data */ }
}
```

### Create Actor
```json
{
  "actionType": "create",
  "requestId": "uuid",
  "data": {
    "entityType": "Actor",
    "data": {
      "name": "New Hero",
      "type": "character"
    }
  }
}
```

### Update Actor HP
```json
{
  "actionType": "update",
  "requestId": "uuid",
  "data": {
    "uuid": "Actor.abc123",
    "data": {
      "system.attributes.hp.value": 50
    }
  }
}
```

### Delete Actor
```json
{
  "actionType": "delete",
  "requestId": "uuid",
  "data": {
    "uuid": "Actor.abc123"
  }
}
```

## Combat Operations

### Start Combat
```json
{
  "actionType": "start-encounter",
  "requestId": "uuid",
  "data": {
    "name": "Boss Fight",
    "startWithPlayers": true,
    "rollAll": true
  }
}
```

### Next Turn
```json
{
  "actionType": "next-turn",
  "requestId": "uuid",
  "data": {}
}
```

### End Combat
```json
{
  "actionType": "end-encounter",
  "requestId": "uuid",
  "data": {}
}
```

## Roll Operations

### Simple Roll
```json
{
  "actionType": "roll",
  "requestId": "uuid",
  "data": {
    "formula": "1d20+5",
    "createChatMessage": true
  }
}
```

### Attack Roll with Speaker
```json
{
  "actionType": "roll",
  "requestId": "uuid",
  "data": {
    "formula": "1d20+7",
    "flavor": "Longsword Attack",
    "speaker": "Actor.abc123",
    "createChatMessage": true
  }
}
```

## Structure Operations

### Get Folder Structure
```json
{
  "actionType": "structure",
  "requestId": "uuid",
  "data": {
    "types": ["Actor", "Item"],
    "recursive": true
  }
}
```

### Get Actors in Folder
```json
{
  "actionType": "get-folder",
  "requestId": "uuid",
  "data": {
    "name": "NPCs"
  }
}
```

## Macro Operations

### List Macros
```json
{
  "actionType": "macros",
  "requestId": "uuid",
  "data": {}
}
```

### Execute Macro
```json
{
  "actionType": "macro-execute",
  "requestId": "uuid",
  "data": {
    "uuid": "Macro.xyz123"
  }
}
```

## Future Go API Design

Based on the protocol, here's how the Go API should look:

```go
// Client methods will follow this pattern

// Entity operations
actor, err := client.GetActor(ctx, "Actor.abc123")
actor, err := client.CreateActor(ctx, &CreateActorRequest{
    Name: "New Hero",
    Type: "character",
})
err := client.UpdateActor(ctx, "Actor.abc123", &UpdateActorRequest{
    HP: 50,
})
err := client.DeleteActor(ctx, "Actor.abc123")

// Combat operations
combat, err := client.StartEncounter(ctx, &StartEncounterRequest{
    Name:            "Boss Fight",
    StartWithPlayers: true,
    RollAll:         true,
})
err := client.NextTurn(ctx, combatID)
err := client.EndEncounter(ctx, combatID)

// Roll operations
roll, err := client.Roll(ctx, &RollRequest{
    Formula:          "1d20+5",
    CreateChatMessage: true,
})

// Structure operations
structure, err := client.GetStructure(ctx, &StructureRequest{
    Types:     []string{"Actor", "Item"},
    Recursive: true,
})

// Macro operations
macros, err := client.ListMacros(ctx)
result, err := client.ExecuteMacro(ctx, "Macro.xyz123")
```

## Error Handling Pattern

```go
actor, err := client.GetActor(ctx, "Actor.invalid")
if err != nil {
    // Check for specific error types
    if errors.Is(err, gofoundryvtt.ErrNotFound) {
        // Handle not found
    }
    if errors.Is(err, gofoundryvtt.ErrTimeout) {
        // Handle timeout
    }
    // Generic error handling
    return fmt.Errorf("failed to get actor: %w", err)
}
```

## Context Usage

```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
actor, err := client.GetActor(ctx, "Actor.abc123")

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
    // Cancel on signal
    <-sigChan
    cancel()
}()
actor, err := client.GetActor(ctx, "Actor.abc123")
```

## Concurrent Requests

```go
// Multiple requests can be in flight
var wg sync.WaitGroup
actors := make([]*Actor, 3)
errs := make([]error, 3)

for i, id := range actorIDs {
    wg.Add(1)
    go func(idx int, actorID string) {
        defer wg.Done()
        actors[idx], errs[idx] = client.GetActor(ctx, actorID)
    }(i, id)
}

wg.Wait()
// Check errs slice for any failures
```

## Keepalive

```go
// Client handles keepalive automatically
// Ping every 30 seconds by default
client := gofoundryvtt.NewClient(gofoundryvtt.Config{
    URL:           "ws://localhost:30000",
    Token:         "your-api-key",
    PingInterval:  30 * time.Second,  // Optional: override default
})
```

## Reconnection

```go
// Client handles reconnection automatically
client := gofoundryvtt.NewClient(gofoundryvtt.Config{
    URL:               "ws://localhost:30000",
    Token:             "your-api-key",
    MaxReconnects:     20,              // Optional: max attempts
    ReconnectBaseDelay: 1 * time.Second, // Optional: initial delay
})

// Listen for connection state changes
client.OnConnectionStateChange(func(state gofoundryvtt.ConnectionState) {
    switch state {
    case gofoundryvtt.StateConnected:
        log.Println("Connected to FoundryVTT")
    case gofoundryvtt.StateDisconnected:
        log.Println("Disconnected from FoundryVTT")
    case gofoundryvtt.StateReconnecting:
        log.Println("Reconnecting...")
    }
})
```

---

**Note**: This is a design reference. Actual implementation may vary based on Go idioms and best practices.
