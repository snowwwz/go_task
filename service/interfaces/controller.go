package interfaces

import (
	"fmt"

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

// Add a task
func (con *Controller) Add(name string, deadline int, priority int) error {
	if name == "" {
		return returnError("name is required", "add")
	}

	if err := con.usecase.Add(name, deadline, priority); err != nil {
		fmt.Println(err.Error())
		return returnError(fmt.Sprintf("faild to add a task [ name: %s ]", name), "add")
	}
	return returnSuccess(fmt.Sprintf("successfully added a task \"%s\"", name))
}

// Delete delete a task
func (con *Controller) Delete(id int) error {
	if err := con.usecase.Delete(id); err != nil {
		fmt.Println(err.Error())
		return returnError(fmt.Sprintf("faild to delete a task [ id: %d ]", id), "delete")
	}
	return returnSuccess(fmt.Sprintf("successfully deleted a task [ id: %d ]", id))
}

// List list all the tasks
func (con *Controller) List(isAll bool) error {
	result, err := con.usecase.List(isAll)
	if err != nil {
		fmt.Println(err.Error())
		return returnError("failed to list tasks", "list")
	}

	return returnList(result)
}

// Change status priority deadline name
func (con *Controller) Change(id int, target string, data string) error {
	if target == "" || data == "" {
		return returnError("column abd data are required", "change")
	}

	if err := con.usecase.Change(id, target, data); err != nil {
		fmt.Println(err.Error())
		return returnError(fmt.Sprintf("faild to change a task [ id: %d ]", id), "change")
	}

	return returnSuccess(fmt.Sprintf("successfully changed task [ id: %d ] %s=%s ", id, target, data))
}
