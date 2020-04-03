package usecase

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

// func (c *Usecase) Add error {

// }

// func (c *Usecase) Delete error {

// }

// List aa
func (u *Usecase) List() error {
	_, err := u.repo.List()
	return err
}

// 				rows, _ := db.Model(&tasks{}).Where("status = ?", 0).Select("id, name, priority, deadline").Rows()
// 				defer rows.Close()
// 				var data [][]string
// 				for rows.Next() {
// 					var user tasks
// 					db.ScanRows(rows, &user)
// 					data = [][]string{
// 						[]string{user.ID, user.Name, user.Priority, user.Deadline},
// 					}
// 				}
// 				table := tablewriter.NewWriter(os.Stdout)
// 				table.SetHeader([]string{"Name", "Sign", "Rating"})

// 				for _, v := range data {
// 					table.Append(v)
// 				}
// 				table.Render()

// 				return nil
// 			},
// 		},
// 	},
// }
