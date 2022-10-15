package docker

type buildHandler interface {
	support() bool
	handle()
}

type localBuild struct {
	dockerfilePath string
}

func (lb *localBuild) support() (support bool) {
	support = lb.dockerfilePath == "local"
	return
}

func (lb *localBuild) handle() {

}

type classicBuild struct {
	dockerfilePath string
}

func (cb *classicBuild) support() (support bool) {
	support = cb.dockerfilePath == ""
	return
}

func (cb *classicBuild) handle() {

}
