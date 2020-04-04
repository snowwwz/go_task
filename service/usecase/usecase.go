package usecase

import (
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

func (u *Usecase) Add(name string, priority int, deadline int) error {
	dl := time.Now().Add(time.Duration(24*deadline) * time.Hour)
	return u.repo.Add(name, priority, dl)
}

func (u *Usecase) Delete(id int) error {
	return u.repo.Delete(id)
}

// List aa
func (u *Usecase) List() ([][]string, error) {
	tasks, err := u.repo.List()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var (
		status  string
		pri     string
		records [][]string
	)
	for _, t := range tasks {
		//todo あとで定数
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

		// Deadline
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
