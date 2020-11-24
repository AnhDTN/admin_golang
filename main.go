package admin_golang

import (
	"admin_golang/dbs"
	"admin_golang/routes"
	"context"
	"time"

	"go.uber.org/dig"
)

func main() {
	c := dig.New()
	routes.InitGinEngine(c)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := dbs.Database.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
