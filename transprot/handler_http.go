package transprot

import "context"

type HttpHandler func(ctx context.Context, structReq interface{}) (resp interface{}, err error)
