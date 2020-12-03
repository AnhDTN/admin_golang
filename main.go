package admin_golang

import (
	"admin_golang/dbs"
	"context"
	"time"
)

func main()  {


	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := dbs.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
