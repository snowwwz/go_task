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
	dl := time.Now().AddDate(0, 0, deadline)
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
		d := t.Deadline.AddDate(0, 0, 1)
		duration := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local).Sub(time.Now())
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
	if !isInvalid(column, columns) {
		return errors.New("column must be one of name/priority/status/deadline")
	}

	var (
		newData interface{}
		err     error
	)
	switch column {
	case "priority":
		pri := []string{"row", "normal", "high"}
		if !isInvalid(data, pri) {
			return errors.New("priority must be one of high/normal/row")
		}
		newData = getPriNumber(data)
	case "status":
		statuses := []string{"todo", "doing", "done", "pending"}
		if !isInvalid(data, statuses) {
			return errors.New("status must be one of done/doing/todo/pending")
		}
		newData = getStatusNumber(data)
	case "deadline":
		newData, err = getDeadlineData(data)
		if err != nil {
			return err
		}
	default:
		newData = data
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

func getPriNumber(s string) (id int) {
	switch s {
	case "row":
		id = row
	case "high":
		id = high
	case "normal":
		id = normal
	}
	return
}

func getStatusNumber(s string) (id int) {
	switch s {
	case "pending":
		id = pending
	case "done":
		id = done
	case "doing":
		id = doing
	case "todo":
		id = todo
	}
	return
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

func getDeadlineData(d string) (interface{}, error) {
	n, err := strconv.Atoi(d)
	if err != nil {
		return nil, errors.New("deadline must be numeric")
	}
	nextDate := time.Now().Add(time.Duration(24*n) * time.Hour)
	return nextDate, nil
}
