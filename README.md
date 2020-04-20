# go_task
cli todo app 

- mysql
- golang
- docker

```
 % go_task list                                              
( ¨̮ )　you have 3 tasks

+----+-----------------+---------+----------+------------+------------+
| ID |      NAME       | STATUS  | PRIORITY |  DEADLINE  |  CREATED   |
+----+-----------------+---------+----------+------------+------------+
| 14 | write test code | pending | high     | 39.4 hours | 2020-04-11 |
| 15 | add loggings    | todo    | row      | 87.4 hours | 2020-04-11 |
| 16 | refactoring     | todo    | row      | 15.4 hours | 2020-04-11 |
+----+-----------------+---------+----------+------------+------------+
```

```
% go_task journal                                            
( ¨̮ )　4 tasks have been changed today

NUMBER  NAME            STATUS  
1       aaaaaa          todo    
2       write test code pending 
3       add loggings    doing   
4       refactoring     todo    
```


```
 % go_task change --id 15 --target status --data doing        
successfully changed task [ id: 15 ] status=doing 
```

```
 % go_task delete --id 13 
successfully deleted a task [ id: 13 ]
```

```
 % go_task add --name "add readme" --due 3
successfully added a task "add readme"
```
