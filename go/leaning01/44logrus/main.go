package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func logrusDemo() {
	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("a dog")
}

// create an logger instance
var logs = log.New()

func logrusDemo2() {
	logs.Out = os.Stdout
	log.WithFields(log.Fields{
		"animal": "dog",
		"size":   10,
	}).Info("a dog")
}

func logrusDemo3() {
	log.Trace("Something very low level.")
	log.Debug("Uesful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	log.Fatal("Bye.")
	log.Panic("I'm bailing.")
}

func logrusDemo4() {
	requestID := 1
	userIP := 1
	// default fields: time, msg, level
	requestLogger := log.WithFields(log.Fields{"request_id": requestID, "user_ip": userIP})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great heppened.")
}

func logHooks() {
	// log.AddHook(airbrake.NewHook(123, "xyz", "production"))
}

func logFormat() {
	// log.TextFormatter
	logs.Formatter = &log.JSONFormatter{}
	logs.Info("a log formatter test")

}

func logMethod() {
	log.SetReportCaller(true)
	log.Info("a test log")
}

func ginUseInit() {
	logs.Formatter = &log.JSONFormatter{}
	path, _ := os.Getwd()
	f, _ := os.Create(path + "/44logrus/gin.log")
	logs.Out = f
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = logs.Out
	logs.Level = log.InfoLevel
}

func ginUse() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		logs.WithFields(log.Fields{
			"animal": "walrus",
			"size":   10,
		}).Warn("a group of walrus emerges from the ocean")

		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run(":8000")
}

func main() {
	// logrusDemo()
	// logrusDemo2()

	// // logrusDemo3()
	// logrusDemo4()

	// logFormat()

	// logMethod()

	ginUseInit()
	ginUse()
}
