package api

import (
	"database/sql"
	"errors"

	"net/http"

	responsehandler "github.com/djsmk123/askmeapi/api/response_handler"
	db "github.com/djsmk123/askmeapi/db/sqlc"

	"github.com/djsmk123/askmeapi/token"
	"github.com/gin-gonic/gin"
)

var (
	errAnswerNotExist = errors.New("answer not found")
)

type CreateAnswerRequest struct {
	Content    string `json:"content" binding:"required,min=5,max=200"`
	QuestionId int32  `json:"question_id" binding:"required,min=1"`
}

func (server *Server) CreateAnswer(ctx *gin.Context) {
	var req CreateAnswerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAnswerParams{
		UserID:     int32(authPayload.ID),
		Content:    req.Content,
		QuestionID: int32(req.QuestionId),
	}

	question, err := server.store.CreateAnswer(ctx, arg)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, question)
}

type UpdateAnswerRequest struct {
	Content string `json:"content" binding:"required,min=5,max=200"`
	ID      int64  `json:"id" binding:"required,min=1"`
}

type DeleteAnswerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DeleteAnswerById(ctx *gin.Context) {
	var req DeleteAnswerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	answer, err := server.store.DeleteAnswerById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			responsehandler.ResponseHandlerJson(ctx, http.StatusNotFound, errAnswerNotExist, nil)
			return
		}
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, answer)

}

func (server *Server) UpdateAnswerById(ctx *gin.Context) {
	var req UpdateAnswerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	userId := authPayload.ID
	arg := db.UpdateAnswersByQuestionIDParams{
		ID:      int32(req.ID),
		Content: req.Content,
		UserID:  int32(userId),
	}

	question, err := server.store.UpdateAnswersByQuestionID(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			responsehandler.ResponseHandlerJson(ctx, http.StatusNotFound, errAnswerNotExist, nil)
			return
		}
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, question)

}

type GetAnswerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAnswerByID(ctx *gin.Context) {
	var req GetAnswerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	question, err := server.store.GetAnswerByID(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			responsehandler.ResponseHandlerJson(ctx, http.StatusNotFound, errAnswerNotExist, nil)
			return
		}
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, question)

}

type ListAnswerRequest struct {
	PageId     int32  `form:"page_id" binding:"required,min=1"`
	PageSize   int32  `form:"page_size" binding:"required,min=5,max=10"`
	QuestionID int32  `form:"question_id" binding:"required,min=1"`
	Search     string `form:"search"`
	UserId     int32  `form:"user_id,default=0"`
}

func (server *Server) ListAnswers(ctx *gin.Context) {
	var req ListAnswerRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	search := req.Search

	if len(req.Search) > 0 {
		search = "%" + search + "%"
	}

	arg := db.GetAnswersByQuestionIDParams{
		QuestionID: req.QuestionID,
		UserID: sql.NullInt32{
			Int32: req.UserId,
			Valid: req.UserId != 0,
		},
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
		Content: sql.NullString{
			String: search,
			Valid:  len(req.Search) > 0,
		},
	}

	questions, err := server.store.GetAnswersByQuestionID(ctx, arg)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, questions)

}
