package controllers

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/artushin/web.dev/app/oauth2"
	"github.com/robfig/revel"
)

type Application struct {
	*revel.Controller
}

func init() {
	revel.InterceptFunc(checkUser, revel.BEFORE, &Application{})
}

func checkUser(c *revel.Controller) revel.Result {
	t, err := oauth2.TransportFromCache(c.Session.Id())
	if err != nil {
		return c.Redirect(Login.Login)
	}
	c.Args["transport"] = t
	return nil
}

func (c Application) Profile() revel.Result {
	profile, err := oauth2.GetAccount(c.Args["transport"].(*oauth.Transport))
	if err != nil {
		return c.Redirect(Login.Login)
	}

	return c.Render(profile)
}
