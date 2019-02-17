package builder

import (
	"github.com/YutoMizutani/gohome/app/presenter/controller"
)

type MenuBuilder struct {
}

func (builder *MenuBuilder) Build() *controller.MenuController {
	return &controller.MenuController{}
}
