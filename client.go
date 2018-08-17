package obsws

import (
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

const bufferSize = 100

var (
	messageID = 0
	lock      = sync.Mutex{}
)

// Client is the interface to obs-websocket.
// Client{Host: "localhost", Port: 4444} will probably work if you haven't configured OBS.
type Client struct {
	Host     string                      // Host (probably "localhost").
	Port     int                         // Port (OBS default is 4444).
	Password string                      // Password (OBS default is "").
	conn     *websocket.Conn             // Underlying connection to OBS.
	active   bool                        // True until Disconnect is called.
	handlers map[string]func(e Event)    // Event handlers.
	respQ    chan map[string]interface{} // Queue of received responses.
}

// init prepares the client's internal fields.
func (c *Client) init() {
	c.handlers = make(map[string]func(Event))
	c.respQ = make(chan map[string]interface{}, bufferSize)
}

// poll listens for responses/events.
// This function blocks until Disconnect is called.
func (c *Client) poll() {
	c.active = true
	logger.Debug("started polling")

	for c.active {
		m := make(map[string]interface{})
		if err := c.conn.ReadJSON(&m); err != nil {
			if !c.active {
				return
			}
			logger.Warning("read from WS:", err)
			continue
		}

		if _, ok := m["message-id"]; ok {
			c.handleResponse(m)
		} else {
			c.handleEvent(m)
		}
	}
}

// getMessageID generates a string that the client has not yet used.
func getMessageID() string {
	lock.Lock()
	messageID++
	id := strconv.Itoa(messageID)
	lock.Unlock()
	return id
}

// mapToStruct serializes a map into a struct.
func mapToStruct(data map[string]interface{}, dest interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: true, // TODO: Is this actually working?
		TagName:    "json",
		Result:     dest,
	})
	if err != nil {
		logger.Warning("initializing decoder:", err)
		return err
	}
	if err = decoder.Decode(data); err != nil {
		logger.Warningf("unmarshalling map -> %T: %v", dest, err)
		logger.Debugf("input: %#v\n", data)
		return err
	}
	return nil
}
