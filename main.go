package main

import (
	"fmt"
	"log"
	"net"
        "os/exec"
        "runtime"
	"strconv"

	t "github.com/eagledb14/stix-io/templates"
	"github.com/gofiber/fiber/v3"
)

func main() {
    port := getPort()
    fmt.Print("Listening on " + port + "\n")
    serv(port)
}

func getPort() string {
    listener, err := net.Listen("tcp", ":0")
    if err != nil {
        panic("Port not available" + err.Error())
    }
    defer listener.Close()

    port := listener.Addr().(*net.TCPAddr).Port
    return ":" + strconv.Itoa(port)
}

func serv(port string) {
    app := fiber.New()
    app.Get("/", func(c fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

        return c.SendString(t.BuildPage(t.Index("", "")))
    })

    app.Static("/style.css", "./styles/style.css")

    app.Post("/yara", func(c fiber.Ctx) error {
        c.Set("Content-Type", "text/html")

        stixJson := c.FormValue("stix")
        stix, err := Unmarshall(stixJson)
        if err != nil {
            return c.SendString(t.BuildPage(t.Index(stixJson, err.Error())))
        }

        yara := stix.ToYara().File()

        return c.SendString(t.BuildPage(t.Index(stixJson, yara)))
    })

    app.Post("/csv", func(c fiber.Ctx) error {
        c.Set("Content-Type", "text/html")

        stixJson := c.FormValue("stix")
        stix, err := Unmarshall(stixJson)
        if err != nil {
            return c.SendString(t.BuildPage(t.Index(stixJson, err.Error())))
        }

        yara := stix.ToYara().Csv()

        return c.SendString(t.BuildPage(t.Index(stixJson, yara)))
    })

    go func() {openBrowser("http://localhost" + port)}()
    log.Fatal(app.Listen(port))
}

func openBrowser(url string) {
    var err error

    switch runtime.GOOS {
    case "linux":
        err = exec.Command("xdg-open", url).Start()
    case "windows":
        err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
    case "darwin":
        err = exec.Command("open", url).Start()
    default:
        err = fmt.Errorf("unsupported platform")
    }

    if err != nil {
        fmt.Println("Error opening browser:", err)
    }
} 
