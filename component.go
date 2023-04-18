package coaster

import "errors"

type Component struct {
	name     string
	renderer func() string
}

func NewComponent() *Component {
	return &Component{}
}
func (c *Component) Render() (*string, error) {
	if c.renderer == nil {
		return nil, errors.New("no renderer function set")
	}
	renderResult := c.renderer()
	return &renderResult, nil
}
func (c *Component) SetRenderer(renderer func() string) {
	c.renderer = renderer
}
