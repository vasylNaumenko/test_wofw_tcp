package app

import (
	"context"
	"sync"
)

// initWorkers returns a list of workers, which will be run in separate goroutines.
func (a *App) initWorkers() []worker {
	return []worker{
		serveGRPC,
	}
}

// runWorkers uses the context to stop the application, and runs workers.
func (a *App) runWorkers(ctx context.Context) {
	workers := a.initWorkers()

	wg := new(sync.WaitGroup)
	wg.Add(len(workers))

	for _, work := range workers {
		go func(ctx context.Context, work func(context.Context, *App), t *App) {
			work(ctx, t)
			wg.Done()
		}(ctx, work, a)
	}

	wg.Wait()
}
