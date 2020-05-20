package controllers

import (
	"bytes"
	"github.com/kataras/iris/v12"
	"github.com/ziqiango/goadmin/core"
	"github.com/ziqiango/goadmin/models"
	"sync"
)

type Index struct {

}

func (c *Index) Index(ctx iris.Context)  {
	var wait sync.WaitGroup
	var name bytes.Buffer
	var i uint32
	name.WriteString("hello! ")
	for i = 1; i < 5; i++ {
		wait.Add(1)
		go func(id uint32) {
			user := &models.User{ID: id}
			core.MyDB.First(user)
			name.WriteString(user.Name)
			wait.Done()
		}(i)
	}
	wait.Wait()
	ctx.WriteString(name.String())
}

func (c *Index) Hello(ctx iris.Context)  {
	ctx.WriteString("hello hello ! RouteName Is "+	ctx.GetCurrentRoute().Name())
}