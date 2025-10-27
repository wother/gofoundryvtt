package gofoundryvtt

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Common errors returned by the client.
var (
	// ErrNotFound is returned when a requested entity does not exist.
	ErrNotFound = errors.New("entity not found")

	// ErrTimeout is returned when a request times out.
	ErrTimeout = errors.New("request timeout")

	// ErrNotConnected is returned when attempting an operation without an active connection.
	ErrNotConnected = errors.New("not connected to FoundryVTT")

	// ErrAlreadyConnected is returned when attempting to connect while already connected.
	ErrAlreadyConnected = errors.New("already connected to FoundryVTT")

	// ErrAuthFailed is returned when authentication fails.
	ErrAuthFailed = errors.New("authentication failed")

	// ErrInvalidResponse is returned when the server sends an invalid response.
	ErrInvalidResponse = errors.New("invalid response from server")
)

// ConnectionState represents the current state of the WebSocket connection.
type ConnectionState int

const (
	// StateDisconnected indicates no active connection.
	StateDisconnected ConnectionState = iota

	// StateConnecting indicates connection in progress.
	StateConnecting

	// StateConnected indicates active connection.
	StateConnected

	// StateReconnecting indicates attempting to reconnect after disconnect.
	StateReconnecting
)

// String returns the string representation of the connection state.
func (s ConnectionState) String() string {
	switch s {
	case StateDisconnected:
		return "disconnected"
	case StateConnecting:
		return "connecting"
	case StateConnected:
		return "connected"
	case StateReconnecting:
		return "reconnecting"
	default:
		return "unknown"
	}
}

// Config contains configuration for the FoundryVTT client.
type Config struct {
	// URL is the WebSocket URL of the FoundryVTT instance.
	// Example: "ws://localhost:30000"
	URL string

	// Token is the API key for authentication.
	// Generated in FoundryVTT module settings.
	Token string

	// WorldID is the ID of the world to connect to.
	WorldID string

	// FoundryVersion is the version of FoundryVTT (optional).
	// If not provided, defaults to "13".
	FoundryVersion string

	// SystemID is the game system ID (optional).
	// Example: "dnd5e"
	SystemID string

	// CustomName is a custom name for this client connection (optional).
	CustomName string

	// RequestTimeout is the timeout for individual requests.
	// Default: 5 seconds.
	RequestTimeout time.Duration

	// PingInterval is the interval between keepalive pings.
	// Default: 30 seconds.
	PingInterval time.Duration

	// MaxReconnectAttempts is the maximum number of reconnection attempts.
	// Default: 20 attempts.
	MaxReconnectAttempts int

	// ReconnectBaseDelay is the initial delay for reconnection backoff.
	// Default: 1 second.
	ReconnectBaseDelay time.Duration
}

// setDefaults fills in default values for unset config fields.
func (c *Config) setDefaults() {
	if c.FoundryVersion == "" {
		c.FoundryVersion = "13"
	}
	if c.RequestTimeout == 0 {
		c.RequestTimeout = 5 * time.Second
	}
	if c.PingInterval == 0 {
		c.PingInterval = 30 * time.Second
	}
	if c.MaxReconnectAttempts == 0 {
		c.MaxReconnectAttempts = 20
	}
	if c.ReconnectBaseDelay == 0 {
		c.ReconnectBaseDelay = 1 * time.Second
	}
}

// validate checks that required config fields are set.
func (c *Config) validate() error {
	if c.URL == "" {
		return fmt.Errorf("URL is required")
	}
	if c.Token == "" {
		return fmt.Errorf("Token is required")
	}
	if c.WorldID == "" {
		return fmt.Errorf("WorldID is required")
	}
	return nil
}

// Client is the main client for interacting with FoundryVTT.
type Client struct {
	config Config

	// Connection state
	mu    sync.RWMutex
	state ConnectionState

	// Connection state change listeners
	stateListeners []func(ConnectionState)

	// TODO: Add transport layer (WebSocket connection)
	// TODO: Add pending requests map (UUID -> response channel)
	// TODO: Add reconnection logic
}

// NewClient creates a new FoundryVTT client with the given configuration.
// It does not establish the connection immediately; call Connect to connect.
func NewClient(config Config) (*Client, error) {
	config.setDefaults()
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	c := &Client{
		config:         config,
		state:          StateDisconnected,
		stateListeners: make([]func(ConnectionState), 0),
	}

	return c, nil
}

// Connect establishes a connection to the FoundryVTT instance.
// It returns an error if already connected or if the connection fails.
func (c *Client) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.state == StateConnected || c.state == StateConnecting {
		return ErrAlreadyConnected
	}

	c.setState(StateConnecting)

	// TODO: Implement actual connection logic
	// - Generate client ID (UUID v4)
	// - Build WebSocket URL with query parameters
	// - Establish WebSocket connection
	// - Wait for connection confirmation
	// - Start keepalive goroutine
	// - Start message handler goroutine

	c.setState(StateConnected)
	return nil
}

// Close closes the connection to the FoundryVTT instance.
// It is safe to call Close multiple times.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.state == StateDisconnected {
		return nil
	}

	// TODO: Implement actual close logic
	// - Send close frame
	// - Cancel all pending requests
	// - Stop keepalive goroutine
	// - Close WebSocket connection

	c.setState(StateDisconnected)
	return nil
}

// State returns the current connection state.
func (c *Client) State() ConnectionState {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state
}

// OnConnectionStateChange registers a listener for connection state changes.
// The listener will be called whenever the connection state changes.
// Returns a function that can be called to unregister the listener.
func (c *Client) OnConnectionStateChange(listener func(ConnectionState)) func() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Store index for unregister
	index := len(c.stateListeners)
	c.stateListeners = append(c.stateListeners, listener)

	// Track if already unregistered
	var unregistered bool

	// Return unregister function
	return func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		if unregistered || index >= len(c.stateListeners) {
			return
		}

		// Remove listener by setting to nil and compacting
		c.stateListeners[index] = nil

		// Compact the slice
		compacted := make([]func(ConnectionState), 0, len(c.stateListeners))
		for _, l := range c.stateListeners {
			if l != nil {
				compacted = append(compacted, l)
			}
		}
		c.stateListeners = compacted
		unregistered = true
	}
}

// setState updates the connection state and notifies listeners.
// Must be called with c.mu held.
func (c *Client) setState(state ConnectionState) {
	if c.state == state {
		return
	}

	c.state = state

	// Notify listeners (make a copy to avoid holding lock during callbacks)
	listeners := make([]func(ConnectionState), len(c.stateListeners))
	copy(listeners, c.stateListeners)

	// Call listeners without holding the lock
	go func() {
		for _, listener := range listeners {
			listener(state)
		}
	}()
}
