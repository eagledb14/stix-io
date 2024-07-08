package templates

import (
    "html/template"
    text "text/template"
    "bytes"
)

func Execute(name string, t string, data interface{}) string {
    tmpl, err := template.New(name).Parse(t)
    if err != nil {
        return err.Error()
    }
    var b bytes.Buffer
    err = tmpl.Execute(&b, data)
    if err != nil {
        return err.Error()
    }

    return b.String()
}

func ExecuteText(name string, t string, data interface{}) string {
    tmpl, err := text.New(name).Parse(t)
    if err != nil {
        return err.Error()
    }
    var b bytes.Buffer
    err = tmpl.Execute(&b, data)
    if err != nil {
        return err.Error()
    }

    return b.String()
}

func Banner() string {
    return `
    <div class="banner grid-center">
        <h1 class="middle">Stix-io</h1>
    </div>

    `
}

func header() string {
    return `
        <head>
            <title>STIX-IO</title>
            <script src="https://unpkg.com/htmx.org@1.9.12" integrity="sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2" crossorigin="anonymous"></script>
            <link rel="stylesheet" type="text/css" href="/style.css">
        </head>
        `
}

func BuildPage(body string) string {
    data := struct {
        Header string
        Banner string
        Body string
    } {
        Header: header(),
        Banner: Banner(),
        Body: body,
    }
    const page = `
        <!DOCTYPE html>
        <html lang="en">
        {{.Header}}
        <body hx-boost="true">
            {{.Banner}}
            {{.Body}}
        </body>
        </html>
    `

    return ExecuteText("page", page, data)
}



