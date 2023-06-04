package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/infrastructure/container"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/infrastructure/mysql"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/server/http"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()
	defer mysql.CloseDatabaseConnection(containerConf.Mysqldb)

	app := fiber.New()
	app.Use(logger.New())
	http.RouteInit(app, containerConf)
	port := fmt.Sprintf("%s:%d", containerConf.Apps.Host, containerConf.Apps.HttpPort)
	helper.Logger("main.go", helper.LoggerLevelFatal, app.Listen(port).Error())
}
