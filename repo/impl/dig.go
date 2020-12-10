package impl

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	_ = container.Provide(NewRoleRepository)
	_ = container.Provide(NewAdminRepository)
	_ = container.Provide(NewBasicUserRepo)
	return nil
}
