package handlers

import "go.uber.org/dig"

func Inject(c *dig.Container) error {
	_ = c.Provide(NewAuthHandler)
	return nil
}
