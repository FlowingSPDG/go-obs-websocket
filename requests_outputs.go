// Code generated by protocol.py DO NOT EDIT
// Generator can be found at https://github.com/FlowingSPDG/go-obs-websocket/blob/master/codegen/protocol.py

package obsws

import (
	"errors"
	"time"
)

// ListOutputsRequest : List existing outputs.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#listoutputs
type ListOutputsRequest struct {
	_request `json:",squash"`
	response chan ListOutputsResponse
}

// NewListOutputsRequest returns a new ListOutputsRequest.
func NewListOutputsRequest() ListOutputsRequest {
	return ListOutputsRequest{
		_request{
			ID_:   GetMessageID(),
			Type_: "ListOutputs",
			err:   make(chan error, 1),
		},
		make(chan ListOutputsResponse, 1),
	}
}

// Send sends the request.
func (r *ListOutputsRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp ListOutputsResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r ListOutputsRequest) Receive() (ListOutputsResponse, error) {
	if !r.sent {
		return ListOutputsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ListOutputsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ListOutputsResponse{}, err
		case <-time.After(receiveTimeout):
			return ListOutputsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r ListOutputsRequest) SendReceive(c Client) (ListOutputsResponse, error) {
	if err := r.Send(c); err != nil {
		return ListOutputsResponse{}, err
	}
	return r.Receive()
}

// ListOutputsResponse : Response for ListOutputsRequest.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#listoutputs
type ListOutputsResponse struct {
	// Outputs list.
	// Required: Yes.
	Outputs   []map[string]interface{} `json:"outputs,omitempty"`
	_response `json:",squash"`
}

// GetOutputInfoRequest : Get information about a single output.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getoutputinfo
type GetOutputInfoRequest struct {
	// Output name.
	// Required: Yes.
	OutputName string `json:"outputName,omitempty"`
	_request   `json:",squash"`
	response   chan GetOutputInfoResponse
}

// NewGetOutputInfoRequest returns a new GetOutputInfoRequest.
func NewGetOutputInfoRequest(outputName string) GetOutputInfoRequest {
	return GetOutputInfoRequest{
		outputName,
		_request{
			ID_:   GetMessageID(),
			Type_: "GetOutputInfo",
			err:   make(chan error, 1),
		},
		make(chan GetOutputInfoResponse, 1),
	}
}

// Send sends the request.
func (r *GetOutputInfoRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetOutputInfoResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetOutputInfoRequest) Receive() (GetOutputInfoResponse, error) {
	if !r.sent {
		return GetOutputInfoResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetOutputInfoResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetOutputInfoResponse{}, err
		case <-time.After(receiveTimeout):
			return GetOutputInfoResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetOutputInfoRequest) SendReceive(c Client) (GetOutputInfoResponse, error) {
	if err := r.Send(c); err != nil {
		return GetOutputInfoResponse{}, err
	}
	return r.Receive()
}

// GetOutputInfoResponse : Response for GetOutputInfoRequest.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getoutputinfo
type GetOutputInfoResponse struct {
	// Output info.
	// Required: Yes.
	OutputInfo map[string]interface{} `json:"outputInfo,omitempty"`
	_response  `json:",squash"`
}

// StartOutputRequest :
//
// Note: Controlling outputs is an experimental feature of obs-websocket
// Some plugins which add outputs to OBS may not function properly when they are controlled in this way.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startoutput
type StartOutputRequest struct {
	// Output name.
	// Required: Yes.
	OutputName string `json:"outputName,omitempty"`
	_request   `json:",squash"`
	response   chan StartOutputResponse
}

// NewStartOutputRequest returns a new StartOutputRequest.
func NewStartOutputRequest(outputName string) StartOutputRequest {
	return StartOutputRequest{
		outputName,
		_request{
			ID_:   GetMessageID(),
			Type_: "StartOutput",
			err:   make(chan error, 1),
		},
		make(chan StartOutputResponse, 1),
	}
}

// Send sends the request.
func (r *StartOutputRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StartOutputResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StartOutputRequest) Receive() (StartOutputResponse, error) {
	if !r.sent {
		return StartOutputResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartOutputResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartOutputResponse{}, err
		case <-time.After(receiveTimeout):
			return StartOutputResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StartOutputRequest) SendReceive(c Client) (StartOutputResponse, error) {
	if err := r.Send(c); err != nil {
		return StartOutputResponse{}, err
	}
	return r.Receive()
}

// StartOutputResponse : Response for StartOutputRequest.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startoutput
type StartOutputResponse struct {
	_response `json:",squash"`
}

// StopOutputRequest :
//
// Note: Controlling outputs is an experimental feature of obs-websocket
// Some plugins which add outputs to OBS may not function properly when they are controlled in this way.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopoutput
type StopOutputRequest struct {
	// Output name.
	// Required: Yes.
	OutputName string `json:"outputName,omitempty"`
	// Force stop (default: false).
	// Required: No.
	Force    bool `json:"force,omitempty"`
	_request `json:",squash"`
	response chan StopOutputResponse
}

// NewStopOutputRequest returns a new StopOutputRequest.
func NewStopOutputRequest(
	outputName string,
	force bool,
) StopOutputRequest {
	return StopOutputRequest{
		outputName,
		force,
		_request{
			ID_:   GetMessageID(),
			Type_: "StopOutput",
			err:   make(chan error, 1),
		},
		make(chan StopOutputResponse, 1),
	}
}

// Send sends the request.
func (r *StopOutputRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.SendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StopOutputResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StopOutputRequest) Receive() (StopOutputResponse, error) {
	if !r.sent {
		return StopOutputResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopOutputResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopOutputResponse{}, err
		case <-time.After(receiveTimeout):
			return StopOutputResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StopOutputRequest) SendReceive(c Client) (StopOutputResponse, error) {
	if err := r.Send(c); err != nil {
		return StopOutputResponse{}, err
	}
	return r.Receive()
}

// StopOutputResponse : Response for StopOutputRequest.
//
// Since obs-websocket version: 4.7.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopoutput
type StopOutputResponse struct {
	_response `json:",squash"`
}
