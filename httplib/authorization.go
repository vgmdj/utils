package httplib

import (
	"fmt"
	"net/http"
)

//Authorization http conn auth
type Authorization interface {
	CheckFormat() error
	GetID() string
	SetAuth(r *http.Request)
}

//DefaultAuth default
type DefaultAuth struct {}

//GetID return ""
func (da *DefaultAuth) GetID() string {
	return ""
}

//CheckFormat always return nil
func (da *DefaultAuth) CheckFormat() error {
	return nil
}

//SetAuth no change
func (da *DefaultAuth) SetAuth(r *http.Request) {

}

//BasicAuth http basic auth
type BasicAuth struct {
	UserName string
	Pwd      string
}

//GetID get the username
func (ba *BasicAuth) GetID() string {

	return ba.UserName
}

//CheckFormat check format
func (ba *BasicAuth) CheckFormat() error {
	if ba.UserName == "" || ba.Pwd == "" {
		return fmt.Errorf("no username or password")
	}

	return nil
}

//SetAuth set basic auth
func (ba *BasicAuth) SetAuth(r *http.Request) {
	r.SetBasicAuth(ba.UserName, ba.Pwd)
}

//App common key secret conn auth
type App struct {
	AppID     string
	AppSecret string
}

//GetID get the appid
func (app *App) GetID() string {
	return app.AppID
}

//CheckFormat check format
func (app *App) CheckFormat() error {
	if app.AppID == "" || app.AppSecret == "" {
		return fmt.Errorf("no appid or appsecret")
	}

	return nil

}

//SetAuth you can use your auth here
func (app *App) SetAuth(r *http.Request) {

}
