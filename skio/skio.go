package skio

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"studyGoApp/common"
)

type Namespace interface {
	// Context of this connection. You can save one context for one
	// connection, and share it between all handlers. The handlers
	// are called in one goroutine, so no need to lock context if it
	// only accessed in one connection.
	Context() interface{}
	SetContext(ctx interface{})

	Namespace() string
	Emit(eventName string, v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type Conn interface {
	io.Closer
	Namespace

	// ID returns session id
	ID() string
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
}

type AppSocket struct {
	conn      Conn
	requester common.Requester
}

func NewAppSocket(conn Conn, requester common.Requester) *AppSocket {
	return &AppSocket{conn: conn, requester: requester}
}
