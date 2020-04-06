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

// Add usecase
func (u *Usecase) Add(name string, deadline int, priority int) error {
	// deadline convert X (days) to time.Time
	dl := time.Now().Add(time.Duration(24*deadline) * time.Hour)
	return u.repo.Add(name, priority, dl)
}

// Dekete usecase
func (u *Usecase) Delete(id int) error {
	return u.repo.Delete(id)
}

// List usecase
func (u *Usecase) List(isAll bool) ([][]string, error) {
	tasks, err := u.repo.List(isAll)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var (
		status  string
		pri     string
		records [][]string
	)
	for _, t := range tasks {
		//todo あとでdomain
		switch t.Priority {
		case 0:
			pri = "row"
		case 2:
			pri = "high"
		default:
			pri = "normal"
		}

		switch t.Status {
		case 3:
			status = "pending"
		case 2:
			status = "done"
		case 1:
			status = "doing"
		default:
			status = "todo"
		}

		// Deadline convert time.Time to float
		duration := t.Deadline.Sub(time.Now())
		deadline := fmt.Sprintf("%s hours", strconv.FormatFloat(duration.Hours(), 'f', 1, 64))

		record := []string{
			strconv.Itoa(t.ID),
			t.Name,
			status,
			pri,
			deadline,
			t.CreatedAt.Format("2006-01-02"),
		}
		records = append(records, record)
	}
	return records, nil
}

// Change usecase
func (u *Usecase) Change(id int, column string, data string) error {
	colums := []string{"name", "priority", "status", "deadline"}
	for i, name := range colums {
		if column == name {
			break
		}
		if i == len(colums)-1 {
			return errors.New("column must be one of name/priority/status/deadline")
		}
	}

	var newData interface{}
	switch column {
	case "priority":
		pri := []string{"row", "normal", "high"}
		for i, p := range pri {
			if data == p {
				newData = i
				break
			}
			if i == len(pri)-1 {
				return errors.New("priority must be one of high/normal/row")
			}
		}
	case "status":
		statuses := []string{"todo", "doing", "done", "pending"}
		for i, status := range statuses {
			if data == status {
				newData = i
				break
			}
			if i == len(statuses)-1 {
				return errors.New("status must be one of done/doing/todo/pending")
			}
		}
	case "deadline":
		d, err := strconv.Atoi(data)
		if err != nil {
			return errors.New("deadline must be numeric")
		}
		newData = time.Now().Add(time.Duration(24*d) * time.Hour)
	}
	return u.repo.Change(id, column, newData)
}
