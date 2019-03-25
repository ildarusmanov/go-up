package app

import "context"

type ServiceFactory func(ctx context.Context) (Service, error)
