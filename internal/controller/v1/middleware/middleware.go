package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthTeacherMiddleware()gin.HandlerFunc {
	return func(c *gin.Context) {
    s:=sessions.Default(c)
	auth:=s.Get("Role")==nil
	if auth {
		c.Redirect(302,"/")
	}
	if s.Get("Role").(string)!="teacher"{
		c.Redirect(302,"/")
	}
	}
}
func AuthAdminMiddleware()gin.HandlerFunc {
	return func(c *gin.Context) {
    s:=sessions.Default(c)
	auth:=s.Get("Role")==nil
	if auth {
		c.Redirect(302,"/")
	}
	if s.Get("Role").(string)!="admin"{
		c.Redirect(302,"/")
	}
	}
}
func AuthStudentMiddleware()gin.HandlerFunc {
	return func(c *gin.Context) {
    s:=sessions.Default(c)
	auth:=s.Get("Role")==nil
	if auth {
		c.Redirect(302,"/")
	}
	if s.Get("Role").(string)!="student"{
		c.Redirect(302,"/")
	}
	}
}
func AuthMiddleware()gin.HandlerFunc {
	return func(c *gin.Context) {
    s:=sessions.Default(c)
	auth:=s.Get("Role")==nil
	if auth {
		c.Redirect(302,"/")
	}
	}
}