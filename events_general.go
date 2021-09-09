// Code generated by protocol.py DO NOT EDIT
// Generator can be found at https://github.com/FlowingSPDG/go-obs-websocket/blob/master/codegen/protocol.py

package obsws

// HeartbeatEvent : Emitted every 2 seconds after enabling it by calling SetHeartbeat.
//
// Since obs-websocket version: V0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#heartbeat
type HeartbeatEvent struct {
	// Toggles between every JSON message as an "I am alive" indicator.
	// Required: Yes.
	Pulse bool `json:"pulse"`
	// Current active profile.
	// Required: No.
	CurrentProfile string `json:"current-profile"`
	// Current active scene.
	// Required: No.
	CurrentScene string `json:"current-scene"`
	// Current streaming state.
	// Required: No.
	Streaming bool `json:"streaming"`
	// Total time (in seconds) since the stream started.
	// Required: No.
	TotalStreamTime int `json:"total-stream-time"`
	// Total bytes sent since the stream started.
	// Required: No.
	TotalStreamBytes int `json:"total-stream-bytes"`
	// Total frames streamed since the stream started.
	// Required: No.
	TotalStreamFrames int `json:"total-stream-frames"`
	// Current recording state.
	// Required: No.
	Recording bool `json:"recording"`
	// Total time (in seconds) since recording started.
	// Required: No.
	TotalRecordTime int `json:"total-record-time"`
	// Total bytes recorded since the recording started.
	// Required: No.
	TotalRecordBytes int `json:"total-record-bytes"`
	// Total frames recorded since the recording started.
	// Required: No.
	TotalRecordFrames int `json:"total-record-frames"`
	// OBS Stats.
	// Required: Yes.
	Stats  *OBSStats `json:"stats"`
	_event `json:",squash"`
}

// BroadcastCustomMessageEvent : A custom broadcast message, sent by the server, requested by one of the websocket clients.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#broadcastcustommessage
type BroadcastCustomMessageEvent struct {
	// Identifier provided by the sender.
	// Required: Yes.
	Realm string `json:"realm"`
	// User-defined data.
	// Required: Yes.
	Data   map[string]interface{} `json:"data"`
	_event `json:",squash"`
}
