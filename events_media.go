// Code generated by protocol.py DO NOT EDIT
// Generator can be found at https://github.com/FlowingSPDG/go-obs-websocket/blob/master/codegen/protocol.py

package obsws

// MediaPlayingEvent :
//
// Note: This event is only emitted when something actively controls the media/VLC source
// In other words, the source will never emit this on its own naturally.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#mediaplaying
type MediaPlayingEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}

// MediaPausedEvent :
//
// Note: This event is only emitted when something actively controls the media/VLC source
// In other words, the source will never emit this on its own naturally.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#mediapaused
type MediaPausedEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}

// MediaRestartedEvent :
//
// Note: This event is only emitted when something actively controls the media/VLC source
// In other words, the source will never emit this on its own naturally.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#mediarestarted
type MediaRestartedEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}

// MediaStoppedEvent :
//
// Note: This event is only emitted when something actively controls the media/VLC source
// In other words, the source will never emit this on its own naturally.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#mediastopped
type MediaStoppedEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}

// MediaNextEvent :
//
// Note: This event is only emitted when something actively controls the media/VLC source
// In other words, the source will never emit this on its own naturally.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#medianext
type MediaNextEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}

// MediaPreviousEvent :
//
// Note: This event is only emitted when something actively controls the media/VLC source
// In other words, the source will never emit this on its own naturally.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#mediaprevious
type MediaPreviousEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}

// MediaStartedEvent :
//
// Note: These events are emitted by the OBS sources themselves
// For example when the media file starts playing
// The behavior depends on the type of media source being used.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#mediastarted
type MediaStartedEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}

// MediaEndedEvent :
//
// Note: These events are emitted by the OBS sources themselves
// For example when the media file ends
// The behavior depends on the type of media source being used.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#mediaended
type MediaEndedEvent struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// The ID type of the source (Eg.
	// `vlc_source` or `ffmpeg_source`).
	// Required: Yes.
	SourceKind string `json:"sourceKind"`
	_event     `json:",squash"`
}
