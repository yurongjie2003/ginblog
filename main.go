package main

import (
	"fmt"
	"github.com/yurongjie2003/ginblog/model"
	"github.com/yurongjie2003/ginblog/router"
	"github.com/yurongjie2003/ginblog/utils/Config"
	"github.com/yurongjie2003/ginblog/utils/Log"
	"github.com/yurongjie2003/ginblog/utils/Minio"
	"log"
	"os"
)

func main() {
	if err := Init(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Hello Gin Blog!")
}

func Init() error {
	if err := Config.Init(); err != nil {
		return err
	}
	if err := Log.InitLogger(); err != nil {
		return err
	}
	if err := model.Init(); err != nil {
		return err
	}
	if err := Minio.Init(); err != nil {
		return err
	}
	if err := router.Init(); err != nil {
		return err
	}
	return nil
}
