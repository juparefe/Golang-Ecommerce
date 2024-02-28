package users

import (
	"fmt"
	"time"

	models "github.com/juparefe/Golang-Ecommerce/learning/models"
)

func AddUser() {
	u := new(models.User)
	u.AddUser(10, "Juan", time.Now(), true)
	fmt.Println(u)
}
