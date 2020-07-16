// Code generated by jrpc. DO NOT EDIT.

package scenedef

import (
	context "context"
	"testing"

	taxi "github.com/shadracnicholas/home-automation/libraries/go/taxi"
)

// SceneService is the public interface of this service
type SceneService interface {
	CreateScene(ctx context.Context, body *CreateSceneRequest) *CreateSceneFuture
	ReadScene(ctx context.Context, body *ReadSceneRequest) *ReadSceneFuture
	ListScenes(ctx context.Context, body *ListScenesRequest) *ListScenesFuture
	DeleteScene(ctx context.Context, body *DeleteSceneRequest) *DeleteSceneFuture
	SetScene(ctx context.Context, body *SetSceneRequest) *SetSceneFuture
}

// CreateSceneFuture represents an in-flight CreateScene request
type CreateSceneFuture struct {
	done <-chan struct{}
	rsp  *CreateSceneResponse
	err  error
}

// Wait blocks until the response is ready
func (f *CreateSceneFuture) Wait() (*CreateSceneResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// ReadSceneFuture represents an in-flight ReadScene request
type ReadSceneFuture struct {
	done <-chan struct{}
	rsp  *ReadSceneResponse
	err  error
}

// Wait blocks until the response is ready
func (f *ReadSceneFuture) Wait() (*ReadSceneResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// ListScenesFuture represents an in-flight ListScenes request
type ListScenesFuture struct {
	done <-chan struct{}
	rsp  *ListScenesResponse
	err  error
}

// Wait blocks until the response is ready
func (f *ListScenesFuture) Wait() (*ListScenesResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// DeleteSceneFuture represents an in-flight DeleteScene request
type DeleteSceneFuture struct {
	done <-chan struct{}
	rsp  *DeleteSceneResponse
	err  error
}

// Wait blocks until the response is ready
func (f *DeleteSceneFuture) Wait() (*DeleteSceneResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// SetSceneFuture represents an in-flight SetScene request
type SetSceneFuture struct {
	done <-chan struct{}
	rsp  *SetSceneResponse
	err  error
}

// Wait blocks until the response is ready
func (f *SetSceneFuture) Wait() (*SetSceneResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// SceneClient makes requests to this service
type SceneClient struct {
	dispatcher taxi.Dispatcher
}

// Compile-time assertion that the client implements the interface
var _ SceneService = (*SceneClient)(nil)

// NewSceneClient returns a new client
func NewSceneClient(d taxi.Dispatcher) *SceneClient {
	return &SceneClient{
		dispatcher: d,
	}
}

// CreateScene dispatches an RPC to the service
func (c *SceneClient) CreateScene(ctx context.Context, body *CreateSceneRequest) *CreateSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "service.scene/scenes",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &CreateSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ReadScene dispatches an RPC to the service
func (c *SceneClient) ReadScene(ctx context.Context, body *ReadSceneRequest) *ReadSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.scene/scene",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ReadSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ListScenes dispatches an RPC to the service
func (c *SceneClient) ListScenes(ctx context.Context, body *ListScenesRequest) *ListScenesFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.scene/scenes",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ListScenesFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// DeleteScene dispatches an RPC to the service
func (c *SceneClient) DeleteScene(ctx context.Context, body *DeleteSceneRequest) *DeleteSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "DELETE",
		URL:    "service.scene/scene",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &DeleteSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// SetScene dispatches an RPC to the service
func (c *SceneClient) SetScene(ctx context.Context, body *SetSceneRequest) *SetSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "service.scene/scene/set",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &SetSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// MockSceneClient can be used in tests
type MockSceneClient struct {
	dispatcher *taxi.MockClient
}

// Compile-time assertion that the mock client implements the interface
var _ SceneService = (*MockSceneClient)(nil)

// NewMockSceneClient returns a new mock client
func NewMockSceneClient(ctx context.Context, t *testing.T) *MockSceneClient {
	f := taxi.NewTestFixture(t)

	return &MockSceneClient{
		dispatcher: &taxi.MockClient{Handler: f},
	}
}

// CreateScene dispatches an RPC to the mock client
func (c *MockSceneClient) CreateScene(ctx context.Context, body *CreateSceneRequest) *CreateSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "service.scene/scenes",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &CreateSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ReadScene dispatches an RPC to the mock client
func (c *MockSceneClient) ReadScene(ctx context.Context, body *ReadSceneRequest) *ReadSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.scene/scene",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ReadSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// ListScenes dispatches an RPC to the mock client
func (c *MockSceneClient) ListScenes(ctx context.Context, body *ListScenesRequest) *ListScenesFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "service.scene/scenes",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &ListScenesFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// DeleteScene dispatches an RPC to the mock client
func (c *MockSceneClient) DeleteScene(ctx context.Context, body *DeleteSceneRequest) *DeleteSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "DELETE",
		URL:    "service.scene/scene",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &DeleteSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// SetScene dispatches an RPC to the mock client
func (c *MockSceneClient) SetScene(ctx context.Context, body *SetSceneRequest) *SetSceneFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "service.scene/scene/set",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &SetSceneFuture{
		done: done,
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}
