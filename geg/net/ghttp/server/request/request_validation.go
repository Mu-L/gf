package main

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/net/ghttp"
    "gitee.com/johng/gf/g/util/gvalid"
    "gitee.com/johng/gf/g/encoding/gparser"
)

func main() {
    type User struct {
        Uid   int    `gvalid:"min:1"`
        Name  string `params:"username"            gvalid:"required|length:6,30"`
        Pass1 string `params:"password1,userpass1" gvalid:"required|password3"`
        Pass2 string `params:"password3,userpass2" gvalid:"required|password3|same:Pass1#||两次密码不一致，请重新输入"`
    }

    s := g.Server()
    s.BindHandler("/user", func(r *ghttp.Request){
        user := new(User)
        r.GetToStruct(user)
        result  := gvalid.CheckStruct(user, nil)
        json, _ := gparser.VarToJsonIndent(result)
        r.Response.Write(json)
    })
    s.SetPort(8199)
    s.Run()
}