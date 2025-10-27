# PROTOCOL.md - WebSocket Protocol Specification

## Overview

The gofoundryvtt library communicates with FoundryVTT via WebSocket using the ThreeHats REST API module protocol. This document specifies the wire protocol.

## Connection

### WebSocket URL Format

```
ws://localhost:30000/?id={clientId}&token={apiKey}&worldId={worldId}&worldTitle={worldTitle}&foundryVersion={version}&systemId={systemId}&systemTitle={systemTitle}&systemVersion={systemVersion}&customName={optionalName}
```

**Query Parameters:**
- `id` (required): Unique client identifier (e.g., `gofoundry-{uuid}`)
- `token` (required): API key from module settings
- `worldId` (required): World ID from FoundryVTT
- `worldTitle` (required): World title
- `foundryVersion` (required): FoundryVTT version (e.g., "13.350")
- `systemId` (required): Game system ID (e.g., "dnd5e")
- `systemTitle` (optional): Game system title
- `systemVersion` (optional): Game system version
- `customName` (optional): Custom client name

### Connection States

1. **CONNECTING** - WebSocket handshake in progress
2. **OPEN** - Connected and ready to send/receive
3. **CLOSING** - Close handshake in progress
4. **CLOSED** - Connection closed

### Keepalive

The client should send ping messages periodically (default: 30 seconds) to maintain the connection.

```json
{
  "type": "ping"
}
```

Server responds with:
```json
{
  "type": "pong"
}
```

## Message Format

### Request Message

All requests follow this structure:

```json
{
  "actionType": "string",
  "requestId": "uuid-v4",
  "data": {
    // Action-specific fields
  }
}
```

**Fields:**
- `actionType` (string, required): The action to perform (see Action Types below)
- `requestId` (string, required): UUID v4 for correlating request/response
- `data` (object, optional): Action-specific parameters

### Response Message

All responses follow this structure:

```json
{
  "type": "string",
  "requestId": "uuid-v4",
  "data": {
    // Response data
  },
  "error": "string"  // Optional, only present on error
}
```

**Fields:**
- `type` (string, required): Response type (usually `{actionType}-result`)
- `requestId` (string, required): Matches the request UUID
- `data` (any, optional): Result data
- `error` (string, optional): Error message if request failed

## Action Types

### Utility Actions

#### ping
Test connection keepalive.

**Request:**
```json
{
  "actionType": "ping"
}
```

**Response:**
```json
{
  "type": "pong"
}
```

---

### Entity Operations

#### entity
Get a single entity by UUID or selected token.

**Request:**
```json
{
  "actionType": "entity",
  "requestId": "uuid",
  "data": {
    "uuid": "Actor.xyz123",      // Optional: Entity UUID
    "selected": true,             // Optional: Use selected token
    "actor": true                 // Optional: Get actor from selected token
  }
}
```

**Response:**
```json
{
  "type": "entity-result",
  "requestId": "uuid",
  "uuid": "Actor.xyz123",
  "data": {
    // Full entity data
  }
}
```

#### create
Create a new entity.

**Request:**
```json
{
  "actionType": "create",
  "requestId": "uuid",
  "data": {
    "entityType": "Actor",        // Document type
    "data": {                     // Entity data
      "name": "New Actor",
      "type": "character"
    },
    "folder": "Folder.abc123"    // Optional: Parent folder UUID
  }
}
```

**Response:**
```json
{
  "type": "create-result",
  "requestId": "uuid",
  "uuid": "Actor.xyz123",
  "entity": {
    // Created entity data
  }
}
```

#### update
Update an existing entity.

**Request:**
```json
{
  "actionType": "update",
  "requestId": "uuid",
  "data": {
    "uuid": "Actor.xyz123",      // Optional: Entity UUID
    "selected": true,             // Optional: Use selected token
    "actor": true,                // Optional: Update actor (not token)
    "data": {                     // Fields to update
      "name": "Updated Name",
      "system.attributes.hp.value": 50
    }
  }
}
```

**Response:**
```json
{
  "type": "update-result",
  "requestId": "uuid",
  "uuid": "Actor.xyz123",
  "entity": [
    // Updated entity data
  ]
}
```

#### delete
Delete an entity.

**Request:**
```json
{
  "actionType": "delete",
  "requestId": "uuid",
  "data": {
    "uuid": "Actor.xyz123",      // Optional: Entity UUID
    "selected": true,             // Optional: Use selected token
    "actor": true                 // Optional: Delete actor (not token)
  }
}
```

**Response:**
```json
{
  "type": "delete-result",
  "requestId": "uuid",
  "uuid": "Actor.xyz123",
  "success": true
}
```

---

### Structure Operations

#### structure
Get the folder/entity structure.

**Request:**
```json
{
  "actionType": "structure",
  "requestId": "uuid",
  "data": {
    "types": ["Actor", "Item"],           // Optional: Filter by types
    "includeEntityData": false,           // Optional: Include full entity data
    "recursive": true,                    // Optional: Include subfolders
    "recursiveDepth": 5,                  // Optional: Max depth
    "path": "Folder.abc123",              // Optional: Start at folder
    "includeCompendiums": true            // Optional: Include compendium packs
  }
}
```

**Response:**
```json
{
  "type": "structure-result",
  "requestId": "uuid",
  "data": {
    "actors": [ /* entities */ ],
    "items": [ /* entities */ ],
    "folders": {
      "FolderName": {
        "actors": [ /* entities */ ],
        "subfolders": { /* nested */ }
      }
    }
  }
}
```

#### get-folder
Get folder details by name.

**Request:**
```json
{
  "actionType": "get-folder",
  "requestId": "uuid",
  "data": {
    "name": "My Folder"
  }
}
```

**Response:**
```json
{
  "type": "get-folder-result",
  "requestId": "uuid",
  "data": {
    "id": "abc123",
    "uuid": "Folder.abc123",
    "name": "My Folder",
    "type": "Actor",
    "parentFolder": "parent-id",
    "contents": [
      {
        "uuid": "Actor.xyz",
        "name": "Actor Name",
        "type": "Actor"
      }
    ]
  }
}
```

---

### Combat/Encounter Operations

#### encounters
List all combat encounters.

**Request:**
```json
{
  "actionType": "encounters",
  "requestId": "uuid",
  "data": {}
}
```

**Response:**
```json
{
  "type": "encounters-result",
  "requestId": "uuid",
  "encounters": [
    {
      "id": "combat-id",
      "name": "Combat Name",
      "round": 1,
      "turn": 0,
      "current": true,
      "combatants": [
        {
          "id": "combatant-id",
          "name": "Combatant Name",
          "tokenUuid": "Token.xyz",
          "actorUuid": "Actor.abc",
          "initiative": 15
        }
      ]
    }
  ]
}
```

#### start-encounter
Start a new combat encounter.

**Request:**
```json
{
  "actionType": "start-encounter",
  "requestId": "uuid",
  "data": {
    "name": "Boss Fight",              // Optional: Combat name
    "tokenUuids": ["Token.xyz"],       // Optional: Token UUIDs to add
    "startWithPlayers": true,          // Optional: Add player tokens
    "startWithSelected": true,         // Optional: Add selected tokens
    "rollNPC": true,                   // Optional: Roll NPC initiative
    "rollAll": true                    // Optional: Roll all initiative
  }
}
```

**Response:**
```json
{
  "type": "start-encounter-result",
  "requestId": "uuid",
  "encounterId": "combat-id",
  "encounter": {
    "id": "combat-id",
    "name": "Boss Fight",
    "round": 1,
    "turn": 0,
    "combatants": [ /* ... */ ]
  }
}
```

#### next-turn
Advance to the next turn in combat.

**Request:**
```json
{
  "actionType": "next-turn",
  "requestId": "uuid",
  "data": {
    "encounterId": "combat-id"        // Optional: Defaults to active combat
  }
}
```

**Response:**
```json
{
  "type": "next-turn-result",
  "requestId": "uuid",
  "encounterId": "combat-id",
  "action": "nextTurn",
  "currentTurn": 1,
  "currentRound": 1,
  "actorTurn": "Actor.xyz",
  "tokenTurn": "Token.abc"
}
```

#### end-encounter
End a combat encounter.

**Request:**
```json
{
  "actionType": "end-encounter",
  "requestId": "uuid",
  "data": {
    "encounterId": "combat-id"        // Optional: Defaults to active combat
  }
}
```

**Response:**
```json
{
  "type": "end-encounter-result",
  "requestId": "uuid",
  "encounterId": "combat-id",
  "message": "Encounter successfully ended"
}
```

---

### Roll Operations

#### roll
Execute a dice roll.

**Request:**
```json
{
  "actionType": "roll",
  "requestId": "uuid",
  "data": {
    "formula": "1d20+5",              // Roll formula
    "flavor": "Attack Roll",          // Optional: Description
    "createChatMessage": true,        // Optional: Post to chat
    "speaker": "Actor.xyz",           // Optional: Speaker UUID
    "whisper": ["user-id"]            // Optional: Whisper to users
  }
}
```

**Response:**
```json
{
  "type": "roll-result",
  "requestId": "uuid",
  "data": {
    "formula": "1d20+5",
    "total": 18,
    "terms": [ /* roll term details */ ],
    "messageId": "chat-msg-id"
  }
}
```

#### rolls
Get recent rolls.

**Request:**
```json
{
  "actionType": "rolls",
  "requestId": "uuid",
  "data": {
    "limit": 20                       // Optional: Max rolls to return
  }
}
```

**Response:**
```json
{
  "type": "rolls-result",
  "requestId": "uuid",
  "data": [
    {
      "formula": "1d20+5",
      "total": 18,
      "timestamp": 1698345678
    }
  ]
}
```

---

### Macro Operations

#### macros
List all macros.

**Request:**
```json
{
  "actionType": "macros",
  "requestId": "uuid",
  "data": {}
}
```

**Response:**
```json
{
  "type": "macros-result",
  "requestId": "uuid",
  "macros": [
    {
      "uuid": "Macro.xyz",
      "id": "macro-id",
      "name": "Macro Name",
      "type": "script",
      "canExecute": true
    }
  ]
}
```

#### macro-execute
Execute a macro.

**Request:**
```json
{
  "actionType": "macro-execute",
  "requestId": "uuid",
  "data": {
    "uuid": "Macro.xyz",
    "args": {                         // Optional: Arguments to pass
      "key": "value"
    }
  }
}
```

**Response:**
```json
{
  "type": "macro-execute-result",
  "requestId": "uuid",
  "uuid": "Macro.xyz",
  "result": {
    // Macro execution result
  }
}
```

---

### Search Operations

#### search
Search for entities (requires QuickInsert module).

**Request:**
```json
{
  "actionType": "search",
  "requestId": "uuid",
  "data": {
    "query": "dragon",
    "filter": "type:actor"            // Optional: Filter string
  }
}
```

**Response:**
```json
{
  "type": "search-result",
  "requestId": "uuid",
  "query": "dragon",
  "results": [
    {
      "uuid": "Actor.xyz",
      "name": "Red Dragon",
      "type": "Actor"
    }
  ]
}
```

---

### Utility Operations

#### execute-js
Execute arbitrary JavaScript (dangerous, use with caution).

**Request:**
```json
{
  "actionType": "execute-js",
  "requestId": "uuid",
  "data": {
    "script": "return game.actors.size;"
  }
}
```

**Response:**
```json
{
  "type": "execute-js-result",
  "requestId": "uuid",
  "success": true,
  "result": 42
}
```

#### select
Select entities on the canvas.

**Request:**
```json
{
  "actionType": "select",
  "requestId": "uuid",
  "data": {
    "uuids": ["Token.xyz"],           // Optional: UUIDs to select
    "name": "Goblin",                 // Optional: Select by name
    "all": true,                      // Optional: Select all
    "overwrite": true                 // Optional: Clear previous selection
  }
}
```

**Response:**
```json
{
  "type": "select-result",
  "requestId": "uuid",
  "success": true,
  "count": 3,
  "selected": ["Token.abc", "Token.def", "Token.ghi"]
}
```

---

## Error Handling

### Error Response Format

When an error occurs, the response includes an `error` field:

```json
{
  "type": "entity-result",
  "requestId": "uuid",
  "error": "Entity not found: Actor.invalid",
  "data": null
}
```

### Common Error Messages

- `"Entity not found: {uuid}"` - UUID doesn't exist
- `"Invalid entity type: {type}"` - Unknown document type
- `"Actor UUID is required"` - Missing required parameter
- `"No active scene found"` - Operation requires active scene
- `"QuickInsert not available"` - Search requires QuickInsert module

### WebSocket Close Codes

| Code | Name | Description |
|------|------|-------------|
| 1000 | Normal | Normal closure |
| 4001 | NoClientId | No client ID provided |
| 4002 | NoAuth | Authentication failed |
| 4003 | NoConnectedGuild | No connected world |
| 4004 | DuplicateConnection | Client already connected |
| 4005 | ServerShutdown | Server shutting down |

## Implementation Notes

### Request ID Generation

- Use UUID v4 for request IDs
- Store pending requests with timeout
- Match responses by request ID
- Clean up expired requests

### Connection Management

- Implement exponential backoff for reconnection
- Default reconnect attempts: 20
- Default base delay: 1000ms
- Implement connection timeout (5 seconds)

### Concurrency

- Multiple requests can be sent without waiting for responses
- Responses may arrive out of order
- Use request ID to correlate requests and responses

### Authentication

- API key/token passed in WebSocket URL query parameter
- Only GM users (role 4) can connect
- Primary GM (lowest user ID) maintains connection

---

**Protocol Version**: Based on ThreeHats REST API v2.0.1
**Last Updated**: 2025-10-26
