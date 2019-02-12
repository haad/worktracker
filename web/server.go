package web

import (
	"encoding/json"
	"go/build"
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/haad/worktracker/model/customer"
)

func StartServer(addr string) {
	url := "http://" + addr + "/index.html"

	worktrackerPkg, err := build.Import("github.com/haad/worktracker", "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	e.Pre(middleware.Rewrite(map[string]string{
		"/app/*": "/index.html",
	}))

	e.Static("/static", worktrackerPkg.Dir+"/spa/dist/static")
	e.Static("/", worktrackerPkg.Dir+"/spa/dist")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/rest", CustomerIndex)
	//e.GET("/app/*", AppIndex)

	// Start server
	log.Println("starting ", url, "pkg ", worktrackerPkg.Dir)
	exec.Command("open", url).Run()

	e.Logger.Fatal(e.Start(addr))
}

func CustomerIndex(c echo.Context) error {
	var customers []customer.CustomerInt

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	customers = customer.CustomerList()

	b, err := json.Marshal(customers)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(b))
}
