package gofoundryvtt

import (
	"context"
	"testing"
	"time"
)

func TestConnectionState_String(t *testing.T) {
	tests := []struct {
		state ConnectionState
		want  string
	}{
		{StateDisconnected, "disconnected"},
		{StateConnecting, "connecting"},
		{StateConnected, "connected"},
		{StateReconnecting, "reconnecting"},
		{ConnectionState(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.state.String(); got != tt.want {
				t.Errorf("ConnectionState.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_setDefaults(t *testing.T) {
	c := &Config{}
	c.setDefaults()

	if c.FoundryVersion != "13" {
		t.Errorf("FoundryVersion = %v, want 13", c.FoundryVersion)
	}
	if c.RequestTimeout != 5*time.Second {
		t.Errorf("RequestTimeout = %v, want 5s", c.RequestTimeout)
	}
	if c.PingInterval != 30*time.Second {
		t.Errorf("PingInterval = %v, want 30s", c.PingInterval)
	}
	if c.MaxReconnectAttempts != 20 {
		t.Errorf("MaxReconnectAttempts = %v, want 20", c.MaxReconnectAttempts)
	}
	if c.ReconnectBaseDelay != 1*time.Second {
		t.Errorf("ReconnectBaseDelay = %v, want 1s", c.ReconnectBaseDelay)
	}
}

func TestConfig_setDefaults_PreservesExisting(t *testing.T) {
	c := &Config{
		FoundryVersion:       "12",
		RequestTimeout:       10 * time.Second,
		PingInterval:         60 * time.Second,
		MaxReconnectAttempts: 5,
		ReconnectBaseDelay:   2 * time.Second,
	}
	c.setDefaults()

	if c.FoundryVersion != "12" {
		t.Errorf("FoundryVersion = %v, want 12", c.FoundryVersion)
	}
	if c.RequestTimeout != 10*time.Second {
		t.Errorf("RequestTimeout = %v, want 10s", c.RequestTimeout)
	}
	if c.PingInterval != 60*time.Second {
		t.Errorf("PingInterval = %v, want 60s", c.PingInterval)
	}
	if c.MaxReconnectAttempts != 5 {
		t.Errorf("MaxReconnectAttempts = %v, want 5", c.MaxReconnectAttempts)
	}
	if c.ReconnectBaseDelay != 2*time.Second {
		t.Errorf("ReconnectBaseDelay = %v, want 2s", c.ReconnectBaseDelay)
	}
}

func TestConfig_validate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				URL:     "ws://localhost:30000",
				Token:   "test-token",
				WorldID: "test-world",
			},
			wantErr: false,
		},
		{
			name: "missing URL",
			config: Config{
				Token:   "test-token",
				WorldID: "test-world",
			},
			wantErr: true,
		},
		{
			name: "missing Token",
			config: Config{
				URL:     "ws://localhost:30000",
				WorldID: "test-world",
			},
			wantErr: true,
		},
		{
			name: "missing WorldID",
			config: Config{
				URL:   "ws://localhost:30000",
				Token: "test-token",
			},
			wantErr: true,
		},
		{
			name:    "empty config",
			config:  Config{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				URL:     "ws://localhost:30000",
				Token:   "test-token",
				WorldID: "test-world",
			},
			wantErr: false,
		},
		{
			name: "invalid config",
			config: Config{
				URL: "ws://localhost:30000",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if client == nil {
					t.Error("NewClient() returned nil client")
				}
				if client.State() != StateDisconnected {
					t.Errorf("NewClient() initial state = %v, want StateDisconnected", client.State())
				}
			}
		})
	}
}

func TestClient_State(t *testing.T) {
	client, err := NewClient(Config{
		URL:     "ws://localhost:30000",
		Token:   "test-token",
		WorldID: "test-world",
	})
	if err != nil {
		t.Fatalf("NewClient() failed: %v", err)
	}

	if state := client.State(); state != StateDisconnected {
		t.Errorf("State() = %v, want StateDisconnected", state)
	}
}

func TestClient_Close_WhenDisconnected(t *testing.T) {
	client, err := NewClient(Config{
		URL:     "ws://localhost:30000",
		Token:   "test-token",
		WorldID: "test-world",
	})
	if err != nil {
		t.Fatalf("NewClient() failed: %v", err)
	}

	// Should be safe to close when already disconnected
	if err := client.Close(); err != nil {
		t.Errorf("Close() error = %v, want nil", err)
	}

	// Should be safe to close multiple times
	if err := client.Close(); err != nil {
		t.Errorf("Close() error = %v, want nil", err)
	}
}

func TestClient_OnConnectionStateChange(t *testing.T) {
	client, err := NewClient(Config{
		URL:     "ws://localhost:30000",
		Token:   "test-token",
		WorldID: "test-world",
	})
	if err != nil {
		t.Fatalf("NewClient() failed: %v", err)
	}

	// Track state changes
	var states []ConnectionState
	stateChan := make(chan ConnectionState, 10)

	unregister := client.OnConnectionStateChange(func(state ConnectionState) {
		stateChan <- state
	})

	// Manually trigger state change
	client.mu.Lock()
	client.setState(StateConnecting)
	client.mu.Unlock()

	// Wait for notification
	select {
	case state := <-stateChan:
		states = append(states, state)
	case <-time.After(100 * time.Millisecond):
		t.Error("Did not receive state change notification")
	}

	if len(states) != 1 {
		t.Errorf("Got %d state changes, want 1", len(states))
	}
	if len(states) > 0 && states[0] != StateConnecting {
		t.Errorf("State change = %v, want StateConnecting", states[0])
	}

	// Test unregister
	unregister()

	// Wait a bit to ensure goroutine completes before triggering next state
	time.Sleep(50 * time.Millisecond)

	client.mu.Lock()
	client.setState(StateConnected)
	client.mu.Unlock()

	// Should not receive notification after unregister
	select {
	case state := <-stateChan:
		t.Errorf("Received unexpected state change after unregister: %v", state)
	case <-time.After(100 * time.Millisecond):
		// Expected - no notification
	}

	// Clean up channel
	close(stateChan)
}

func TestClient_Connect_WhenAlreadyConnected(t *testing.T) {
	client, err := NewClient(Config{
		URL:     "ws://localhost:30000",
		Token:   "test-token",
		WorldID: "test-world",
	})
	if err != nil {
		t.Fatalf("NewClient() failed: %v", err)
	}

	// Manually set state to connected
	client.mu.Lock()
	client.state = StateConnected
	client.mu.Unlock()

	// Attempt to connect again should fail
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != ErrAlreadyConnected {
		t.Errorf("Connect() when already connected: got error %v, want ErrAlreadyConnected", err)
	}
}
