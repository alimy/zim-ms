package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(SessionsQuery))
}

// SessionsQuery sessions query interface
type SessionsQuery struct {
	Group Group `mir:"sessions/query"`
	Add   Get   `mir:"/add/:lhs/:rhs"`
	Sub   Get   `mir:"/sub/:lhs/:rhs"`
}
