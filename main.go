package main

import (
	// "fmt"

	"log"
	"os"

	t "github.com/eagledb14/stix-io/templates"
	"github.com/gofiber/fiber/v3"
)

func main() {
    // fmt.Println(read())
    // fmt.Println("hi")
    // fmt.Println(read().ToYara().File())
    // read().ToYara().Csv()
    // for _, i := range read().object {
    //     fmt.println(i.pattern)
    // }
    serv()
}

func read() Bundle {
    content, err := os.ReadFile("input.txt")
    if err != nil {
        panic("unkown file name")
    }
    return Unmarshall(string(content))
}

func serv() {
    app := fiber.New()
    app.Get("/", func(c fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

        return c.SendString(t.BuildPage(t.Index("", "")))
    })

    app.Static("/style.css", "./styles/style.css")

    app.Post("/yara", func(c fiber.Ctx) error {
        c.Set("Content-Type", "text/html")

        stixJson := c.FormValue("stix")
        stix := Unmarshall(stixJson)

        yara := stix.ToYara().File()

        return c.SendString(t.BuildPage(t.Index(stixJson, yara)))
    })

    app.Post("/csv", func(c fiber.Ctx) error {
        c.Set("Content-Type", "text/html")

        stixJson := c.FormValue("stix")
        stix := Unmarshall(stixJson)

        yara := stix.ToYara().Csv()

        return c.SendString(t.BuildPage(t.Index(stixJson, yara)))
    })


    log.Fatal(app.Listen(":3000"))
}
