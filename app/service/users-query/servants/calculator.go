package servants

import (
	"context"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/current"
)

type querySrv struct {
	// just empty
}

func (s querySrv) Add(c context.Context, lh int32, rh int32, res *int32) (int32, error) {
	logger := tars.GetLogger("context")
	ip, ok := current.GetClientIPFromContext(c)
	if !ok {
		logger.Error("Error getting ip from context")
	}
	logger.Infof("Get Client Ip : %s from context", ip)
	rqx, ok := current.GetRequestContext(c)
	rsx := map[string]string{
		"resp": "respform context",
	}
	if ok {
		logger.Infof("Get context from context: %v", rqx)
		for k, v := range rqx {
			rsx[v] = k
		}
	} else {
		rsx["norqx"] = "yes"
		logger.Error("Error getting reqcontext from context")
	}
	if ok = current.SetResponseContext(c, rsx); !ok {
		logger.Error("error setting respose context")
	}
	//todo: something in your function
	*res = lh + rh
	return 0, nil
}

func (s querySrv) Sub(c context.Context, lh int32, rh int32, res *int32) (int32, error) {
	//todo: something in your function
	*res = lh - rh
	return 0, nil
}

// NewUsersQuery return a UsersQuery implement object
func NewUsersQuery() querySrv {
	return querySrv{}
}
