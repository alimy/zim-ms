package routes

import (
	. "github.com/alimy/mir/v2"
	. "github.com/alimy/mir/v2/engine"
)

func init() {
	AddEntry(new(DnsUpdate))
}

// DnsUpdate a httpdns update interface
type DnsUpdate struct {
	Group Group `mir:"admin/httpdns/update"`
	Add   Get   `mir:"/add/:lhs/:rhs"`
	Sub   Get   `mir:"/sub/:lhs/:rhs"`
}
