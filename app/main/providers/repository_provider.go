package providers

import (
	"github.com/brunobotter/site-sentinel/main/container"
)

type RepositoryProvider struct{}

func NewRepositoryProvider() *RepositoryProvider {
	return &RepositoryProvider{}
}
func (p *RepositoryProvider) Register(c container.Container) {

}
