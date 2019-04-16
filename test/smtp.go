package main

// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gsmtp provides a SMTP client to access remote mail server.

import (
	"dial/base"
	"encoding/base64"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

// 示例：
// s := smtp.New("smtp.exmail.qq.com:25", "notify@a.com", "password")
// glog.Println(s.SendMail("notify@a.com", "ulric@b.com;rain@c.com", "这是subject", "这是body,<font color=red>red</font>"))

type Smtp struct {
	Address  string
	Username string
	Password string
}

func New(address, username, password string) *Smtp {
	return &Smtp{
		Address:  address,
		Username: username,
		Password: password,
	}
}

func (this *Smtp) SendMail(from, tos, subject, body string, contentType ...string) error {
	if this.Address == "" {
		return fmt.Errorf("address is necessary")
	}

	hp := strings.Split(this.Address, ":")
	if len(hp) != 2 {
		return fmt.Errorf("address format error")
	}

	arr := strings.Split(tos, ";")
	count := len(arr)
	safeArr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		if arr[i] == "" {
			continue
		}
		safeArr = append(safeArr, arr[i])
	}

	if len(safeArr) == 0 {
		return fmt.Errorf("tos invalid")
	}

	tos = strings.Join(safeArr, ";")

	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

	header := make(map[string]string)
	header["From"] = from
	header["To"] = tos
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"

	ct := "text/plain; charset=UTF-8"
	if len(contentType) > 0 && contentType[0] == "html" {
		ct = "text/html; charset=UTF-8"
	}

	header["Content-Type"] = ct
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))

	auth := smtp.PlainAuth("", this.Username, this.Password, hp[0])
	return smtp.SendMail(this.Address, auth, from, strings.Split(tos, ";"), []byte(message))
}

func Smt(addr string) int64 {
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return 0
	}
	defer conn.Close()

	t := time.Now()
	client, err := smtp.Dial(addr)
	if err != nil {
		return 0
	}
	defer client.Close()

	err = client.Hello("localhost")
	if err != nil {
		return 0
	}
	return time.Since(t).Nanoseconds() / 1000
}

func main() {
	//smtp := New("smtp.qq.com:25", "475943497@qq.com", "gmdfztzbinlvbiei")
	//smtp := New("14.17.57.241:25", "475943497@qq.com", "gmdfztzbinlvbiei")
	//fmt.Println(smtp.SendMail("475943497@qq.com", "ww.he@yamu.com", "这是subject", "这是body,<font color=red>red</font>"))
	addr := "14.17.57.241:25"
	fmt.Println(base.Smtp(addr))
}
