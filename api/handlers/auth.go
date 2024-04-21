package handlers

import (
	"expression-backend/internal/entities"
	"expression-backend/internal/errors"
	"expression-backend/internal/utils"
	"expression-backend/internal/validators"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	credentials := entities.UserCredentials{}

	if err := utils.DecodeBody(w, r, &credentials); err != nil {
		if err = utils.RespondWith400(w, "Invalid request"); err != nil {
			_ = utils.RespondWith500(w)
		}
		return
	}

	if err := validators.ValidateLogin(credentials.Login); err != nil {
		_ = utils.RespondWith400(w, err.Error())
		return
	}

	if err := validators.ValidatePassword(credentials.Password); err != nil {
		_ = utils.RespondWith400(w, err.Error())
		return
	}

	login, err := entities.FindUserByLogin(credentials.Login)

	if login != nil {
		_ = utils.RespondWith400(w, errors.UserAlreadyExists.Error())
		return
	}

	err = entities.InsertUser(&entities.User{
		Login:    credentials.Login,
		Password: utils.GetMD5Hash(strings.TrimSpace(credentials.Password)),
	})

	if err != nil {
		_ = utils.RespondWith500(w)
		log.Println(err)
		return
	}

	_ = utils.SuccessResponseWith201(w, struct {
		Message string `json:"message"`
	}{
		Message: "user registered successfully",
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	credentials := entities.UserCredentials{}

	if err := utils.DecodeBody(w, r, &credentials); err != nil {
		if err = utils.RespondWith400(w, "Invalid request"); err != nil {
			_ = utils.RespondWith500(w)
		}
		return
	}

	if err := validators.ValidateLogin(credentials.Login); err != nil {
		_ = utils.RespondWith400(w, err.Error())
		return
	}

	if err := validators.ValidatePassword(credentials.Password); err != nil {
		_ = utils.RespondWith400(w, err.Error())
		return
	}

	user, err := entities.FindUserByLogin(credentials.Login)

	if err != nil {
		_ = utils.RespondWith400(w, errors.UserNotExists.Error())
		return
	}
	fmt.Println(user, err)
	isValidPassword := utils.IsValidMD5Hash(strings.TrimSpace(credentials.Password), user.Password)

	if !isValidPassword {
		_ = utils.RespondWith400(w, errors.InvalidCredentials.Error())
		return
	}

	token, err := utils.GenerateToken(credentials.Login)
	if err != nil {
		_ = utils.RespondWith500(w)
		return
	}

	_ = utils.SuccessResponseWith200(w, struct {
		TokenType string `json:"token_type"`
		Token     string `json:"token"`
	}{
		TokenType: "jwt",
		Token:     token,
	})
}
