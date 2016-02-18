package demo

type DemoService struct {
}

func (this *DemoService) Hello(name string) (r string, err error) {
	r = `hello, ` + name
	return
}
