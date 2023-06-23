package altsrc

import (
	"context"
	"time"

	"github.com/urfave/cli-altsrc/v3/internal"
)

var (
	testdataDir = func() string {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		return internal.MustTestdataDir(ctx)
	}()
)
