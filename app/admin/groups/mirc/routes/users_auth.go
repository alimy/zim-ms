package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(UsersAuth))
}

// UsersAuth admin groups auth interface
type UsersAuth struct {
	Group Group `mir:"admin/groups/auth"`
	Add   Get   `mir:"/add/:lhs/:rhs"`
	Sub   Get   `mir:"/sub/:lhs/:rhs"`
}
