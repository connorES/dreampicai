package handler

import (
	"dreampicai/view/settings"
	"net/http"
)

func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Index(user))
}
