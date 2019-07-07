/*
使用命令行进行issues的相关操作： 创建、读取、更新和关闭issue
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
)

const GithubApiHost = "https://api.github.com"

/*
login
*/
func loginOAuth(token string) (err error) {
	resp, err := http.Get(fmt.Sprintf("%s/user?access_token=%s", GithubApiHost, token))
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK { // 404 Not Found means login fail
		return
	} else {
		return errors.New(fmt.Sprintf("Invalid github OAuth token!"))
	}
}
