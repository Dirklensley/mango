package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/louisevanderlith/mango/api/secure/logic"
)

type LoginController struct {
	beego.Controller
}

func (req *LoginController) Get() {
	req.TplName = "login.html"
}

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (req *LoginController) Post() {
	// failed logins should redirect to the login page
	var login logic.Login
	json.Unmarshal(req.Ctx.Input.RequestBody, &login)

	token := logic.AttemptLogin(login)

	req.Data["json"] = token
	req.ServeJSON()

	// redirect to ReturnURL with token as querystring
}

// @Title Logout
// @Description Logs out current logged in user session
// @Param	token		header 	string	true		"The session token"
// @Success 200 {string} logout success
// @router /logout [get]
func (req *LoginController) Logout() {
	token := req.GetString("token")

	if len(token) == 16 {
		logic.Delete(token)
	}

	req.Data["json"] = "Logout Success"
	req.ServeJSON()
}
