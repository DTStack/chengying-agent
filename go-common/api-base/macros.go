package apibase

import (
	"github.com/kataras/iris"
	uuid "github.com/satori/go.uuid"
)

func RegisterUUIDStringMacro(app *iris.Application) {
	app.Macros().String.RegisterFunc("uuid", func() func(string) bool {
		return func(pv string) bool {
			_, err := uuid.FromString(pv)
			if err != nil {
				return false
			} else {
				return true
			}
		}
	})
}
