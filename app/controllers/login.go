package controllers

import (
	"github.com/artushin/web.dev/app/oauth2"
	"github.com/robfig/revel"
)

type Login struct {
	*revel.Controller
}

func (c Login) Login() revel.Result {
	return c.Render()
}

func (c Login) Enter() revel.Result {
	return c.Redirect(oauth2.GoogleOauthCfg.AuthCodeURL(""))
}

func (c Login) OAuth(code string) revel.Result {
	t, err := oauth2.NewTransport(code, c.Session.Id())
	if err != nil {
		return c.Redirect(Login.Login)
	} else {
		c.Args["transport"] = t
		return c.Redirect(Application.Profile)
	}
}
