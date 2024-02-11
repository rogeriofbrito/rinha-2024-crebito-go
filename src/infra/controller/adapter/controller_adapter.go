package controller_adapter

import "github.com/sarulabs/di"

type IControllerAdapter interface {
	Start(dic di.Container) error
}
