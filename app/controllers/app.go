package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}
type User struct {
    Name string
}

func (c App) Index() revel.Result {
    return c.RenderJson(User{"go tarou"})
}


func (c App) Hello(myName string) revel.Result {
    c.Validation.Required(myName).Message("Your name is required!")
    c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }

    return c.Render(myName)
}
