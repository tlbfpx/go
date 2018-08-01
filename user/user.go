package user

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

type UserSlice []User

func (a UserSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a UserSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a UserSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].age < a[i].age
}

func Sort(users []User) []User {

	fmt.Println("sort user")
	sort.Sort(UserSlice(users))
	return users

}
