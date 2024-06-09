package mockdatagenerator

import (
	"fmt"
	"leetcode/model"
	"leetcode/storage/postgres"

	_ "github.com/bxcodec/faker/v3"
	"github.com/icrowley/fake"
)

func UsersGenerator(u *postgres.Users) error {
	for i := 0; i < 500000; i++ {
		firstname := fake.FirstName()
		lastname := fake.LastName()
		newUser := model.User{
			Fullname: firstname + " " + lastname,
			Username: firstname + fmt.Sprintf("%d", i),
		}
		err := u.CreateUser(&newUser)
		if err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Printf("%d ta odamni qo'shib bo'ldi\n", i)
		}
	}
	return nil
}
