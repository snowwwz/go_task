package infrastructure

// func Create() {
// app := &cli.App{
// 	Commands: []*cli.Command{
// 		{
// 			Name:  "add",
// 			Usage: "add a task",
// 			Action: ,
// 		},
// 		{
// 			Name:  "list",
// 			Usage: "list all the uncompleted tasks",
// 			Action: func(c *cli.Context) error {
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
