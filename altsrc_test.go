package altsrc

import (
	"context"
	"time"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return MustTestdataDir(ctx)
	}()
)
