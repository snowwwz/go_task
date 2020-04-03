package interfaces

import (
	"github.com/yukinooz/go_task/service/usecase"
)

// Controller struct
type Controller struct {
	usecase *usecase.Usecase
}

// NewController create a new controller
func NewController(u *usecase.Usecase) *Controller {
	return &Controller{
		usecase: u,
	}
}

// // Add a task
// func (c *Controller) Add(c *cli.Context) error {
// 	c.usecase.Add()
// }

// // Delete delete a task
// func (c *Controller) Delete(c *cli.Context) error {
// 	c.usecase.Delete()
// }

// List lisy tasks
func (c *Controller) List() error {
	//todo output系
	return c.usecase.List()
}
