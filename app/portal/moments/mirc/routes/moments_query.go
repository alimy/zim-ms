package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(MomentsQuery))
}

// MomentsQuery moments query interface
type MomentsQuery struct {
	Group Group `mir:"moments/query"`
	Add   Get   `mir:"/add/:lhs/:rhs"`
	Sub   Get   `mir:"/sub/:lhs/:rhs"`
}
