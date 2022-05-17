package config

type PropertyResolver interface {
	ContainsSection(key string) bool
	ContainProperty(key string, sec string) bool
	GetProperty(key string, sec string) interface{}
}

type Configer interface {
	PropertyResolver
	Initialize()
	Watch()
}

type DefaultConfiger struct {
	filePath string
}
var _ Configer = &DefaultConfiger{}

func  NewDefaultConfiger(filePath string) *DefaultConfiger {
	return &DefaultConfiger{ filePath: filePath}
}

func (DefaultConfiger) Initialize() {
	panic("implement me")
}

func (DefaultConfiger) Watch() {
	panic("implement me")
}

func (DefaultConfiger) ContainsSection(key string) bool {
	panic("implement me")
}

func (DefaultConfiger) ContainProperty(key string, sec string) bool {
	panic("implement me")
}

func (DefaultConfiger) GetProperty(key string, sec string) interface{} {
	panic("implement me")
}




