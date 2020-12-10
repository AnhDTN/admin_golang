package handlers

import "go.uber.org/dig"

func Inject(c *dig.Container) error {
	_ = c.Provide(NewAuthHandler)
	_ = c.Provide(NewUserHandler)
	_ = c.Provide(NewBasicUser)
	return nil
}
