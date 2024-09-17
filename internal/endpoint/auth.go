package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (e *Endpoint) SignUp(c *gin.Context) {
	var user SignUpInput
	if err := c.BindJSON(&user); err != nil {
		Err(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := e.Services.SignUp(user.Name, user.Email, user.Password)
	if err != nil {
		Err(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (e *Endpoint) SignIn(c *gin.Context) {

}

func (e *Endpoint) Refresh(c *gin.Context) {

}
