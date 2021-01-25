package main

import (
	"os"
	"soundcast/api/interfaces/data"
	"soundcast/api/routes/info"
	"soundcast/modules/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// RequestLogger log the latency
func RequestLogger(c *gin.Context) {
	// c.Next()
	// log.WithFields(log.Fields{
	// 	"method": c.Request.Method,

	// }).Info("Request Logger")
	// return

	ua := c.Query("ua")
	path := c.Request.URL.Path
	c.Next()
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()
	clientUserAgent := c.Request.UserAgent()

	entry := log.WithFields(logrus.Fields{
		"url_ua":     ua,
		"statusCode": statusCode,
		"clientIP":   clientIP,
		"method":     c.Request.Method,
		"path":       path,
		"userAgent":  clientUserAgent,
	})
	entry.Info()
}

// DbConnection set the dataFinder in the context
func DbConnection(df data.Finder) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dataFinder", df)
		c.Next()
	}
}

func init() {
	// uncomment the following line to run in release mode and disable the warnings
	// gin.SetMode(gin.ReleaseMode)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	file, err := os.OpenFile("requests.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func main() {

	var db = &db.JSONData{}

	err := db.LoadFile("user-agents.json")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("An error has occured while trying to load the database")
		return
	}

	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(RequestLogger)
	r.Use(DbConnection(db))

	r.GET("/info", info.RequestHandler)
	r.Run()
}
