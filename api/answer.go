package api

import (
	"database/sql"
	"errors"
	"net/http"

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
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	authPayload := ctx.MustGet(autherizationPayloadKey).(*token.Payload)

	arg := db.CreateAnswerParams{
		UserID:     int32(authPayload.ID),
		Content:    req.Content,
		QuestionID: int32(req.QuestionId),
	}

	question, err := server.store.CreateAnswer(ctx, arg)

	if err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, question)
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
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	answer, err := server.store.DeleteAnswerById(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ResponseHandlerJson(ctx, http.StatusNotFound, errAnswerNotExist, nil)
			return
		}
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, answer)

}

func (server *Server) UpdateAnswerById(ctx *gin.Context) {
	var req UpdateAnswerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	authPayload := ctx.MustGet(autherizationPayloadKey).(*token.Payload)

	userId := authPayload.ID
	arg := db.UpdateAnswersByQuestionIDParams{
		ID:      int32(req.ID),
		Content: req.Content,
		UserID:  int32(userId),
	}

	question, err := server.store.UpdateAnswersByQuestionID(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ResponseHandlerJson(ctx, http.StatusNotFound, errAnswerNotExist, nil)
			return
		}
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, question)

}

type GetAnswerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAnswerByID(ctx *gin.Context) {
	var req GetAnswerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	question, err := server.store.GetAnswerByID(ctx, int32(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ResponseHandlerJson(ctx, http.StatusNotFound, errAnswerNotExist, nil)
			return
		}
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, question)

}

type ListAnswerRequest struct {
	PageId     int32 `form:"page_id" binding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=5,max=10"`
	QuestionID int32 `form:"question_id" binding:"required,min=1"`
}

func (server *Server) ListAnswers(ctx *gin.Context) {
	var req ListAnswerRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	authPayload := ctx.MustGet(autherizationPayloadKey).(*token.Payload)
	arg := db.GetAnswersByQuestionIDParams{
		QuestionID: req.QuestionID,
		UserID:     int32(authPayload.ID),
		Limit:      req.PageSize,
		Offset:     (req.PageId - 1) * req.PageSize,
	}

	questions, err := server.store.GetAnswersByQuestionID(ctx, arg)

	if err != nil {
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	ResponseHandlerJson(ctx, http.StatusOK, nil, questions)

}

/*func (server *Server) fetchUserObject(ctx *gin.Context, id int64) *UserResponseType {
	user, err := server.store.GetUserByID(ctx, int32(id))

	if err != nil {
		if err == sql.ErrNoRows {
			ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
			return nil
		}
		ResponseHandlerJson(ctx, http.StatusInternalServerError, err, nil)
		return nil
	}
	resp := GetUserResponse(user)
	return &resp

}*/
