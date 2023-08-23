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
	errQuestionNotExist = errors.New("question not found")
)

type CreateQuestionRequest struct {
	Question string `json:"question" binding:"required,min=5,max=200"`
}

func (server *Server) CreateQuestion(ctx *gin.Context) {
	var req CreateQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateQuestionParams{
		UserID:  int32(authPayload.ID),
		Content: req.Question,
	}

	question, err := server.store.CreateQuestion(ctx, arg)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, question)
}

type UpdateQuestionRequest struct {
	Question string `json:"question" binding:"required,min=5,max=200"`
	ID       int64  `json:"id" binding:"required,min=1"`
}

type DeleteQuestionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DeleteQuestionById(ctx *gin.Context) {
	var req DeleteQuestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	err := server.store.DeleteAnswerByQuestionId(ctx, int32(req.ID))

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	question, err := server.store.QuestionDelete(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			responsehandler.ResponseHandlerJson(ctx, http.StatusNotFound, errQuestionNotExist, nil)
			return
		}
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, question)

}

func (server *Server) UpdateQuestionById(ctx *gin.Context) {
	var req UpdateQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	userId := authPayload.ID
	arg := db.UpdateQuestionByIdParams{
		ID:      int32(req.ID),
		Content: req.Question,
		UserID:  int32(userId),
	}
	question, err := server.store.UpdateQuestionById(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			responsehandler.ResponseHandlerJson(ctx, http.StatusNotFound, errQuestionNotExist, nil)
			return
		}
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, question)

}

type GetQuestionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetQuestionByID(ctx *gin.Context) {
	var req GetQuestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	question, err := server.store.GetQuestionByID(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			responsehandler.ResponseHandlerJson(ctx, http.StatusNotFound, errQuestionNotExist, nil)
			return
		}
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, question)

}

type ListQuestionRequest struct {
	PageId   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=10"`
	Search   string `form:"search"`
	UserId   int32  `form:"user_id,default=0"`
}

func (server *Server) ListQuestion(ctx *gin.Context) {
	var req ListQuestionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	search := req.Search

	if len(search) > 0 {
		search = "%" + search + "%"
	}

	arg := db.GetQuestionsByUserIDParams{
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

	questions, err := server.store.GetQuestionsByUserID(ctx, arg)

	if err != nil {
		responsehandler.ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	responsehandler.ResponseHandlerJson(ctx, http.StatusOK, nil, questions)

}
