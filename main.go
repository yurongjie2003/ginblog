package main

import (
	"fmt"
	"github.com/yurongjie2003/ginblog/config"
	"github.com/yurongjie2003/ginblog/model"
	"github.com/yurongjie2003/ginblog/router"
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
	if err := config.Init(); err != nil {
		return err
	}
	if err := model.Init(); err != nil {
		return err
	}
	if err := router.Init(); err != nil {
		return err
	}
	return nil
}
