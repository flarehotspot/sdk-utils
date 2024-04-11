package sdkconnmgr

import "context"

type FetchSessionsResult struct {
	Sessions []SessionSource
	Pages    uint
	Count    uint
}

type SessionProvider interface {

    // Get avaialable session for a client device
	GetSession(ctx context.Context, clnt ClientDevice) (s SessionSource, ok bool)

    // Fetch available sessions for a client device
	FetchSessions(ctx context.Context, clnt ClientDevice, page int, perPage int) (result FetchSessionsResult, err error)
}
