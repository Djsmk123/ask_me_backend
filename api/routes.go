package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	//add a middleware to set content-type header for all requests
	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	})
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(corsConfig))
	if server.Config.GINMODE == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router.LoadHTMLGlob("static/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/reset-password-page", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "password_reset.html", nil)
	})

	router.PATCH("/resetpassword", server.ResetPaswordVerify)
	router.UseRawPath = true
	router.UnescapePathValues = false
	v1 := router.Group("/api/v1")

	v1.POST("/create-user", server.CreateUser)

	v1.POST("/create-ano-user", server.CreateAnonymousUser)
	v1.POST("/login-user", server.LoginUser)
	v1.POST("/social-login", server.SocialLogin)
	v1.POST("/request-password-reset", server.PasswordResetRequest)

	authRoutesV1 := v1.Use(server.AuthMiddleware(server.TokenMaker))

	authRoutesV1.GET("/delete-user/", server.DeleteUser)
	authRoutesV1.GET("/get-user/", server.GetUser)
	authRoutesV1.POST("/logout/", server.LogoutUser)

	authRoutesV1.POST("/create-question", server.CreateQuestion)
	authRoutesV1.POST("/update-question", server.UpdateQuestionById)
	authRoutesV1.GET("/delete-question/:id", server.DeleteQuestionById)
	authRoutesV1.GET("/questions", server.ListQuestion)
	authRoutesV1.GET("/question/:id", server.GetQuestionByID)

	authRoutesV1.POST("/create-answer", server.CreateAnswer)
	authRoutesV1.POST("/update-answer", server.UpdateAnswerById)
	authRoutesV1.GET("/delete-answer/:id", server.DeleteAnswerById)
	authRoutesV1.GET("/answers", server.ListAnswers)
	authRoutesV1.GET("/answer/:id", server.GetAnswerByID)

	server.Router = router
}
