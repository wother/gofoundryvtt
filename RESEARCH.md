# RESEARCH.md - FoundryVTT API Architecture Investigation

## Date: 2025-10-26

## Discovery: Not a Traditional REST API

The documentation at https://foundryvtt.com/api/ documents the **client-side JavaScript API**, 
not a server REST API. This changes our approach significantly.

## Investigation Areas

### 1. Server Architecture

**Questions to Answer:**
- What server framework does FoundryVTT use? (Express, Fastify, custom?)
- Does it expose any HTTP REST endpoints?
- What is the primary communication mechanism?
- How does the JavaScript client communicate with the server?

**Files to Examine:**
- Server entry point (likely in `server/` or `dist/` directory)
- Route definitions
- WebSocket/Socket.io initialization
- Database layer

### 2. Communication Protocols

**Possible Approaches:**

#### A. HTTP/REST Endpoints (If They Exist)
- Look for Express routes or similar
- Document endpoint patterns: `/api/actors`, `/api/scenes`, etc.
- Check authentication headers/cookies
- Test with curl/Postman

#### B. WebSocket Protocol
- Identify WebSocket library (Socket.io, ws, native)
- Document message format (JSON-RPC, custom protocol)
- Understand event types and payloads
- Map CRUD operations to WebSocket messages

#### C. Socket.io Protocol
- Version of Socket.io
- Namespace structure
- Event names and payload formats
- Connection/authentication handshake

#### D. Database Direct Access
- Database type (NeDB, SQLite, MongoDB?)
- Schema structure
- Feasibility of direct Go database driver
- Risks and limitations

### 3. Authentication & Authorization

**To Document:**
- Session management mechanism
- Token format (JWT, custom?)
- User roles and permissions
- API key support (if any)
- Admin vs user access levels

### 4. Document Operations

**For Each Document Type (Actor, Item, Scene, etc.):**
- Create operation: method, payload, response
- Read operation: single item retrieval
- Update operation: full vs partial updates
- Delete operation: soft vs hard delete
- List operation: pagination, filtering, sorting
- Embedded document operations

### 5. Real-time Updates

**To Understand:**
- How does server push updates to clients?
- Event subscription mechanism
- Broadcast vs targeted messages
- Conflict resolution for concurrent edits

## Findings Log

### Session 1 - Initial Discovery (2025-10-26)

**Finding**: The API documentation is for client-side JavaScript only.

**Impact**: Cannot directly translate to REST endpoints. Need to analyze actual 
server implementation.

### Session 2 - Third-Party REST API Discovery (2025-10-26)

**Finding**: Found MIT-licensed FoundryVTT REST API module by ThreeHats:
https://github.com/ThreeHats/foundryvtt-rest-api

**How It Works**:
- FoundryVTT module that adds WebSocket-based API
- Uses WebSocket relay server (fly.dev by default) for external access
- NOT a traditional REST API - it's WebSocket with action-based routing
- Module runs inside FoundryVTT and responds to WebSocket messages
- Message format: `{ actionType: "...", requestId: "...", data: {...} }`

**Implementation Strategy**:
- We DON'T want the relay server (fly.dev)
- We WANT direct localhost WebSocket connection
- We can use their message protocol as our API spec
- Install their module on FoundryVTT instance
- Connect directly via WebSocket (bypassing relay)

**Architecture**:
```
[Go Client] <--WebSocket--> [FoundryVTT + REST API Module] <-- JavaScript API --> [FoundryVTT Core]
```

**Dependencies** (As Required):
1. Go (for our client library)
2. FoundryVTT running locally
3. ThreeHats REST API module installed in FoundryVTT
4. NO external relay servers
5. NO cloud dependencies

### Session 3 - Protocol Analysis (2025-10-26)

**Finding**: Complete protocol specification documented.

**Source Code Analysis**:
- Copied ThreeHats source to `/temp/foundryvtt-rest-api/`
- Analyzed WebSocketManager implementation
- Documented all router action types
- Created complete wire protocol specification

**Key Technical Details**:
- WebSocket URL includes query params for auth and client info
- Keepalive via ping/pong every 30 seconds (configurable)
- Request/response correlation via UUID v4
- Primary GM concept (lowest user ID) maintains single connection
- Exponential backoff reconnection strategy
- 5-second connection timeout

**Protocol Document**: See PROTOCOL.md for complete specification

**Router Coverage**:
- ✓ Entity operations (CRUD)
- ✓ Combat/encounter management
- ✓ Roll system
- ✓ Structure/folder navigation
- ✓ Macro execution
- ✓ Search (requires QuickInsert)
- ✓ Utility operations
- ✓ DnD5e system-specific actions

**Ready to Implement**: All protocol details documented, ready for Go implementation.

---

## Selected Implementation Strategy

### Strategy: WebSocket Client Using ThreeHats Protocol

**Architecture**:
```
gofoundryvtt (Go) --> WebSocket --> FoundryVTT Module --> FoundryVTT Core
```

**Why This Approach**:
1. MIT-licensed reference implementation exists
2. Well-documented message protocol
3. Active maintenance (last updated Aug 2024)
4. Comprehensive coverage of FoundryVTT API
5. Direct localhost connection (no cloud dependencies)

**Message Protocol** (from ThreeHats module):
```typescript
// Request
{
  actionType: string    // e.g., "entity", "create", "update", "delete"
  requestId: string     // UUID for matching responses
  data: any            // Action-specific payload
}

// Response
{
  type: string         // e.g., "entity-result", "create-result"
  requestId: string    // Matches request
  data: any           // Result data
  error?: string      // Error message if failed
}
```

**Action Types Discovered** (Router Analysis):
- Entity Operations: `entity`, `create`, `update`, `delete`
- Actor Operations: `get-actor-details` (dnd5e specific)
- Item Operations: `give`, `remove`, `modify-item-charges`
- Combat: `start-encounter`, `next-turn`, `next-round`, `add-to-encounter`, `remove-from-encounter`, `end-encounter`
- Roll: `roll`, `rolls`, `last-roll`
- Search: `search` (requires QuickInsert module)
- Structure: `structure`, `get-folder`, `create-folder`, `contents`
- Macro: `macros`, `macro-execute`
- Utility: `execute-js`, `select`, `selected`
- Sheet: `get-sheet` (HTML rendering)
- Ping/Pong: `ping`, `pong` (keepalive)

**Rejected Strategies**:
- ~~REST API Wrapper~~: No native REST API in FoundryVTT
- ~~Reverse Engineer Client~~: Too complex, maintenance burden
- ~~Direct Database~~: Bypasses business logic, risky
- ~~Build Our Own Module~~: Unnecessary, ThreeHats module exists

## Decision Matrix

| Criteria | ThreeHats WebSocket | Custom WebSocket | Direct DB |
|----------|---------------------|------------------|-----------|
| Exists? | YES | No | Yes |
| Complexity | Low | High | Medium |
| Features | Comprehensive | Full | Full |
| Real-time | Yes | Yes | No |
| Maintainability | High | Medium | Low |
| Safety | High | High | Low |
| Dependencies | Module + FoundryVTT | FoundryVTT Only | FoundryVTT Only |
| Documentation | Good | None | Schema-dependent |
| Community Support | Active | DIY | None |
| **DECISION** | **SELECTED** | Rejected | Rejected |

## Network Traffic Analysis Plan

### Tools to Use:
- Chrome DevTools Network tab
- Wireshark for detailed packet analysis
- Browser WebSocket inspector extensions

### What to Capture:
1. Initial connection/handshake
2. Authentication request/response
3. Actor CRUD operations (all)
4. Real-time update when another client modifies data
5. Error responses
6. Disconnection/reconnection

### Expected Capture Files:
- `network-trace.har` - HTTP Archive from browser
- `websocket-messages.json` - WebSocket message log
- `auth-flow.md` - Authentication sequence diagram
- `message-format.md` - Protocol documentation

## Code Locations to Examine

Based on typical FoundryVTT structure (to be confirmed):

```
foundry/
├── server/          # Server-side code
│   ├── routes/      # HTTP routes (if any)
│   ├── socket/      # WebSocket handling
│   └── database/    # Database layer
├── client/          # Client-side code (JS API we saw)
├── common/          # Shared code
└── dist/            # Built/compiled code
```

## Questions for Community/Documentation

If official docs are unclear:
1. Post on FoundryVTT Discord #dev-support
2. Check FoundryVTT GitHub issues for similar questions
3. Look for existing API wrappers in other languages (Python, etc.)

## Preliminary Recommendations

**DO:**
- Start by capturing actual network traffic
- Look for existing community API documentation
- Check if FoundryVTT has an official API for integrations
- Consider asking on official channels

**DON'T:**
- Assume REST exists without verification
- Start coding before understanding protocol
- Bypass security mechanisms
- Directly access database without understanding risks

---

**Status**: Research in progress
**Last Updated**: 2025-10-26
**Next Update**: After initial source code analysis
