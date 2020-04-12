package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(MomentsComment))
}

// MomentsComment moments comment interface
type MomentsComment struct {
	Group Group `mir:"moments/comment"`
	Add   Get   `mir:"/add/:lhs/:rhs"`
	Sub   Get   `mir:"/sub/:lhs/:rhs"`
}
