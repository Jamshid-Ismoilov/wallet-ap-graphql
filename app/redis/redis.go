package redis

import (
	"app/graph/model"
	"time"
	// "fmt"
)

// if problem occurs with redis, use commented codes instead

func SetPasswordByEmail(input model.NewUser, password string) {

	Client.Set(input.Email, password, time.Duration(time.Second*300))

}


// func SetPasswordByEmail(input model.NewUser, password string) {

// 	fmt.Println(input, password)

// }

func GetPasswordByEmail(email string) (string, error) {
	return Client.Get(email).Result()
}

// func GetPasswordByEmail(email string) (string, error) {
// 	return "write password that had been sent to email", nil
// }


func DeletePasswordByEmail(email string) {
	Client.Del(email)
}

// func DeletePasswordByEmail(email string) {
// 	fmt.Println(email)
// }
