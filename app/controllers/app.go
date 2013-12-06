package controllers

import (
	"fmt"
	"github.com/artushin/web.dev/app/oauth2"
	"github.com/robfig/revel"
)

type Application struct {
	*revel.Controller
}

func checkUser(c *revel.Controller) revel.Result {
	if user := 1; user != 0 {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}
	return nil
}

func (c Application) Index() revel.Result {
	return c.Render()
}

func (c Application) EnterDemo() revel.Result {
	return c.Redirect(oauth2.GoogleOauthCfg.AuthCodeURL(""))
}

func (c Application) OAuth(code string) revel.Result {
	t, err := oauth2.Transport(code)
	if err != nil {
		return c.Redirect(Application.Index)
	}

	profile, err := oauth2.GetAccount(t)
	if err != nil {
		return c.Redirect(Application.Index)
	}

	fmt.Println("Hi, ", profile.GivenName)
	return c.Redirect(Application.Index)
}
