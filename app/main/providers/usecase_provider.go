package providers

import (
	"github.com/brunobotter/site-sentinel/main/container"
)

type UseCaseProvider struct{}

func NewUseCaseProvider() *UseCaseProvider {
	return &UseCaseProvider{}
}
func (p *UseCaseProvider) Register(c container.Container) {

}
