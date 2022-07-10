package factorymethod

//API is interface
type API interface {
	Say(name string) string
}

type helloAPI struct {
}

func (h helloAPI) Say(name string) string {
	//TODO implement me
	panic("implement me")
}

//NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//hiAPI is one of API implement
type hiAPI struct{}

func (h hiAPI) Say(name string) string {
	//TODO implement me
	panic("implement me")
}
