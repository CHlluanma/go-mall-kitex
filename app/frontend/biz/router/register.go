// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	home "github.com/CHlluanma/go-mall-kitex/app/frontend/biz/router/home"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	home.Register(r)
}
