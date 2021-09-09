// Code generated by protocol.py DO NOT EDIT
// Generator can be found at https://github.com/FlowingSPDG/go-obs-websocket/blob/master/codegen/protocol.py

package obsws

// ReplayStartingEvent : A request to start the replay buffer has been issued.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#replaystarting
type ReplayStartingEvent struct {
	_event `json:",squash"`
}

// ReplayStartedEvent : Replay Buffer started successfully.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#replaystarted
type ReplayStartedEvent struct {
	_event `json:",squash"`
}

// ReplayStoppingEvent : A request to stop the replay buffer has been issued.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#replaystopping
type ReplayStoppingEvent struct {
	_event `json:",squash"`
}

// ReplayStoppedEvent : Replay Buffer stopped successfully.
//
// Since obs-websocket version: 4.2.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#replaystopped
type ReplayStoppedEvent struct {
	_event `json:",squash"`
}
