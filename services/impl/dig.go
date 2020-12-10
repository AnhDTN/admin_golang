package impl

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	_ = container.Provide(NewAuthService)
	_ = container.Provide(NewUserService)
	_ = container.Provide(NewBasicUserService)
	return nil
}
