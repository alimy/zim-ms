package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(SessionsMessage))
}

// SessionsComment moments comment interface
type SessionsMessage struct {
	Group Group `mir:"sessions/message"`
	Add   Get   `mir:"/add/:lhs/:rhs"`
	Sub   Get   `mir:"/sub/:lhs/:rhs"`
}
