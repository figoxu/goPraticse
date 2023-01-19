package main

import (
	"bytes"
	"fmt"
	"github.com/figoxu/gh"
	"github.com/gin-gonic/gin"
	"github.com/quexer/utee"
	"io"
	"net/http"
)

type Env struct {
	fh *gh.FormHelper
	ph *gh.ParamHelper
}

func main() {
	r := gin.Default()
	r.GET("/hello/world/json", h_helloWorldJson)
	r.GET("/hello/name/:name", h_helloByName)
	r.GET("/hello/redirect/:url", h_helloRedirect)
	r.GET("/img/proxy", h_img_proxy)
	r.POST("/hello/form", h_helloForm)
	r.POST("/hello/gh/:name/:age/:bossFlag", m_gh, h_hello_gh)
	r.Run(":8080")
}

//curl http://localhost:8080/img/proxy?encodeurl=https%3A%2F%2Fwww.baidu.com%2Fimg%2Fbd_logo1.png
func h_img_proxy(c *gin.Context) {
	encodeurl := c.Query("encodeurl")
	bs, err := utee.HttpGet(encodeurl)
	utee.Chk(err)
	io.CopyN(c.Writer, bytes.NewBuffer(bs), int64(len(bs)))
}

func h_helloWorldJson(c *gin.Context) {
	c.SecureJSON(http.StatusOK, map[string]string{
		"foo":   "bar",
		"hello": "world",
		"figo":  "xu",
	})
}

func h_helloByName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

// http://localhost:8080/hello/redirect/www.baidu.com
func h_helloRedirect(c *gin.Context) {
	url := c.Param("url")
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://%s", url))
}

// curl http://localhost:8080/hello/form -d "nick=figo&msg=how old are you"
func h_helloForm(c *gin.Context) {
	msg := c.PostForm("msg")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"status": "posted",
		"msg":    msg,
		"nick":   nick,
	})
}

// curl http://localhost:8080/hello/gh/figo/33/false -d "salary=100000000"
func h_hello_gh(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	fh, ph := env.fh, env.ph
	salary, name, age, bossFlag := fh.Int("salary"), ph.String("name"), ph.Int("age"), ph.Bool("bossFlag")

	c.JSON(http.StatusOK, gin.H{
		"salary":   salary,
		"name":     name,
		"age":      age,
		"bossFlag": bossFlag,
	})
}

func m_gh(c *gin.Context) {
	c.Set("env", &Env{
		fh: gh.NewFormHelper(c),
		ph: gh.NewParamHelper(c),
	})
	c.Next()
}
