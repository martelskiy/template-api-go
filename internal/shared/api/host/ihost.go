package host

import "context"

type IHost interface {
	RunAsync() error
	Terminate(ctx context.Context) error
}
