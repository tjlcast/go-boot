package app

import "tjlcast.com/go-boot/config"

type ApplicationContext interface {
	GetId() string
	GetApplicationName() string
	GetDisplayName() string
	GetParent() ApplicationContext
	GetConfiger() config.Configer
}
