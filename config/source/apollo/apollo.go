package apollo

import "github.com/micro/go-micro/v2/config/source"

type Apollo struct {
}

func (a *Apollo) Read() (*source.ChangeSet, error) {
	return nil, nil
}

func (a *Apollo) Write(*source.ChangeSet) error {
	return nil
}

func (a *Apollo) Watch() (source.Watcher, error) {
	return nil, nil
}

func (a *Apollo) String() string {
	return "apollo"
}
