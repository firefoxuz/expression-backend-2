package handlers

import (
	"expression-backend/internal/entities"
	"expression-backend/internal/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetUserExpressions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*entities.User)

	expressions, err := entities.GetUserExpressions(user.Id)
	if err != nil {
		_ = utils.RespondWith500(w)
		return
	}

	_ = utils.SuccessResponseWith200(w, struct {
		Expressions *[]entities.Expression `json:"expressions"`
	}{
		Expressions: expressions,
	})
}

func GetUserExpression(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*entities.User)
	params := mux.Vars(r)
	expressionId, ok := params["id"]

	if !ok {
		_ = utils.RespondWith400(w, "invalid params")
		return
	}

	id, err := strconv.Atoi(expressionId)

	if err != nil {
		_ = utils.RespondWith400(w, "invalid params")
		return
	}

	expression, err := entities.FindExpressionById(id)
	if err != nil {
		_ = utils.RespondWith404(w)
	}

	if expression.UserId != user.Id {
		_ = utils.RespondWithError(w,
			http.StatusForbidden,
			"unauthorized")
		return
	}

	_ = utils.SuccessResponseWith200(w, struct {
		Expression *entities.Expression `json:"expressions"`
	}{
		Expression: expression,
	})
}

func StoreUserExpression(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*entities.User)

	expression := entities.Expression{}

	err := utils.DecodeBody(w, r, &expression)
	if err != nil {
		log.Println(err.Error())
		_ = utils.RespondWith500(w)
		return
	}

	if expression.Expression == "" {
		_ = utils.RespondWith400(w, "expression is required")
		return
	}

	if expression.TimeLimit == 0 {
		_ = utils.RespondWith400(w, "time_limit is required")
		return
	}

	expressionData := entities.Expression{
		Expression:   strings.ReplaceAll(expression.Expression, " ", ""),
		UserId:       user.Id,
		Result:       nil,
		IsProcessing: false,
		IsValid:      1,
		IsFinished:   false,
		TimeLimit:    expression.TimeLimit,
		CreatedAt:    time.Now().Format(time.DateTime),
		FinishedAt:   nil,
	}

	err = entities.StoreUserExpression(expressionData)
	if err != nil {
		_ = utils.RespondWith500(w)
		return
	}

	_ = utils.SuccessResponseWith200(w, struct {
		Message string `json:"message"`
	}{
		Message: "expression is stored and will be calculated",
	})
}
