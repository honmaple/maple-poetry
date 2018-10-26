/*********************************************************************************
 Copyright © 2018 jianglin
 File Name: response.go
 Author: jianglin
 Email: mail@honmaple.com
 Created: 2018-09-12 15:38:10 (CST)
 Last Update: Friday 2018-10-26 11:01:02 (CST)
		  By:
 Description:
 *********************************************************************************/
package main

import (
	"github.com/gin-gonic/gin"
)

// HTTPResponse ..
type HTTPResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	PageInfo   int
}

// Raw ..
func (self *HTTPResponse) Raw() interface{} {
	return gin.H{
		"message":     self.Message,
		"data":        self.Data,
		"pageinfo":    self.PageInfo,
		"status_code": self.StatusCode,
	}
}

// Render ..
func (self *HTTPResponse) Render(c *gin.Context) {
	d := gin.H{
		"message": self.Message,
		"data":    self.Data,
	}
	if self.PageInfo != 0 {
		d["pageinfo"] = self.PageInfo
	}
	c.JSON(self.StatusCode, d)
}

// HTTP ..
type HTTP struct {
	Context *gin.Context
}

// OK ..
func (self HTTP) OK(message string, data interface{}) {
	if message == "" {
		message = "ok"
	}
	resp := HTTPResponse{
		StatusCode: 200,
		Message:    message,
		Data:       data,
	}
	resp.Render(self.Context)
}

// BadRequest ..
func (self HTTP) BadRequest(message string, data interface{}) {
	if message == "" {
		message = "bad request"
	}
	resp := HTTPResponse{
		StatusCode: 400,
		Message:    message,
		Data:       data,
	}
	resp.Render(self.Context)
}

// UnAuthorized ..
func (self HTTP) UnAuthorized(message string, data interface{}) {
	if message == "" {
		message = "unauthorized"
	}
	resp := HTTPResponse{
		StatusCode: 401,
		Message:    message,
		Data:       data,
	}
	resp.Render(self.Context)
}

// Forbidden ..
func (self HTTP) Forbidden(message string, data interface{}) {
	if message == "" {
		message = "forbidden"
	}
	resp := HTTPResponse{
		StatusCode: 403,
		Message:    message,
		Data:       data,
	}
	resp.Render(self.Context)
}

// NotFound ..
func (self HTTP) NotFound(message string, data interface{}) {
	if message == "" {
		message = "not found"
	}
	resp := HTTPResponse{
		StatusCode: 404,
		Message:    message,
		Data:       data,
	}
	resp.Render(self.Context)
}

// ServerError ..
func (self HTTP) ServerError(message string, data interface{}) {
	if message == "" {
		message = "internal server error"
	}
	resp := HTTPResponse{
		StatusCode: 500,
		Message:    message,
		Data:       data,
	}
	resp.Render(self.Context)
}
