package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Usecase struct
type Usecase struct {
	repo TaskRepository
}

// NewUsecase create a new usecase
func NewUsecase(t TaskRepository) *Usecase {
	return &Usecase{
		repo: t,
	}
}

// priority
const (
	row = iota
	normal
	high
)

// status
const (
	todo = iota
	doing
	done
	pending
)

// Add usecase
func (u *Usecase) Add(name string, deadline int, priority int) error {
	// deadline convert X (days) to time.Time
	dl := time.Now().Add(time.Duration(24*deadline) * time.Hour)
	return u.repo.Add(name, priority, dl)
}

// Delete usecase
func (u *Usecase) Delete(id int) error {
	return u.repo.Delete(id)
}

// List usecase
func (u *Usecase) List(isAll bool) ([][]string, error) {
	tasks, err := u.repo.List(isAll)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var records [][]string
	for _, t := range tasks {
		// Deadline convert time.Time to float
		duration := t.Deadline.Sub(time.Now())
		deadline := fmt.Sprintf("%s hours", strconv.FormatFloat(duration.Hours(), 'f', 1, 64))

		record := []string{
			strconv.Itoa(t.ID),
			t.Name,
			getStatus(t.Status),
			getPriority(t.Priority),
			deadline,
			t.CreatedAt.Format("2006-01-02"),
		}
		records = append(records, record)
	}
	return records, nil
}

// Change usecase
func (u *Usecase) Change(id int, column string, data string) error {
	columns := []string{"name", "priority", "status", "deadline"}
	if !isInvalid(data, columns) {
		return errors.New("column must be one of name/priority/status/deadline")
	}

	var newData interface{}
	switch column {
	case "priority":
		pri := []string{"row", "normal", "high"}
		if !isInvalid(data, pri) {
			return errors.New("priority must be one of high/normal/row")
		}
		newData = data
	case "status":
		statuses := []string{"todo", "doing", "done", "pending"}
		if !isInvalid(data, statuses) {
			return errors.New("status must be one of done/doing/todo/pending")
		}
		newData = data
	case "deadline":
		d, err := strconv.Atoi(data)
		if err != nil {
			return errors.New("deadline must be numeric")
		}
		newData = time.Now().Add(time.Duration(24*d) * time.Hour)
	}
	return u.repo.Change(id, column, newData)
}

// Journal usecase
func (u *Usecase) Journal() ([][]string, error) {
	tasks, err := u.repo.Journal(time.Now())
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var records [][]string
	for index, t := range tasks {
		record := []string{
			strconv.Itoa(index + 1),
			t.Name,
			getStatus(t.Status),
		}
		records = append(records, record)
	}
	return records, nil
}

func getPriority(id int) (pri string) {
	switch id {
	case row:
		pri = "row"
	case high:
		pri = "high"
	case normal:
		pri = "normal"
	}
	return
}

func getStatus(id int) (status string) {
	switch id {
	case pending:
		status = "pending"
	case done:
		status = "done"
	case doing:
		status = "doing"
	case todo:
		status = "todo"
	}
	return
}

func isInvalid(input string, categories []string) bool {
	for _, c := range categories {
		if input == c {
			return true
		}
	}
	return false
}
