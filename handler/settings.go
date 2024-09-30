package handler

import (
	"dreampicai/db"
	"dreampicai/pkg/kit/validate"
	"dreampicai/view/settings"
	"net/http"
)

func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Index(user))
}

func HandleSettUsernameUpdate(w http.ResponseWriter, r *http.Request) error {
	params := settings.ProfileParams{
		Username: r.FormValue("username"),
	}
	errors := settings.ProfileErrors{}
	ok := validate.New(&params, validate.Fields{
		"Username": validate.Rules(validate.Min(3), validate.Max(50)),
	}).Validate(&errors)

	if !ok {
		return render(r, w, settings.ProfileForm(params, errors))
	}

	user := GetAuthenticatedUser(r)
	user.Account.Username = params.Username
	if err := db.UpdateAccount(&user.Account); err != nil {
		return err
	}
	params.Success = true
	return render(r, w, settings.ProfileForm(params, settings.ProfileErrors{}))
}
