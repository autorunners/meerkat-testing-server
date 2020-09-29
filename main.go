package main

import (
	"github.com/autorunners/meerkat-testing-server/app/config"
	"github.com/autorunners/meerkat-testing-server/router"
	"github.com/gin-gonic/gin"

	"gopkg.in/yaml.v2"

	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	data, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(data))
	var conf config.Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Panic(err)
	}

	r := gin.Default()
	r.GET("/json", router.GetJson)
	r.GET("/normal", router.GetNormal)
	r.GET("/normal/string", router.GetNormalString)

	s := &http.Server{
		Addr:           conf.Server.Port,
		Handler:        r,
		ReadTimeout:    time.Duration(conf.Server.ReadTimeout) * time.Millisecond,
		WriteTimeout:   time.Duration(conf.Server.WriteTimeout) * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}

	log.Println("meerkat server is running...")
}
