package handler

import (
	"dreampicai/view/home"
	"fmt"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)

	fmt.Println("user: ", user.ID)
	fmt.Println("user: ", user.Account)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }

	// account := types.Account{
	// 	UserID:   user.ID,
	// 	Username: "foobarbaz",
	// }
	// if err := db.CreateAccount(&account); err != nil {
	// 	return err
	// }

	fmt.Printf("%+v\n", user.Account)
	return home.Index().Render(r.Context(), w)
}
