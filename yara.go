package main

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"
	"text/template"
)

type Yara struct {
    Indicator []Indicator
}

func (y Yara) File() string {
    out := ""
    for _, ind := range y.Indicator {
        out += ind.Function()
    }

    return out
}

func (y Yara) Csv() {
    file, _ := os.Create("indicators.csv")
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"name", "type", "value", "data"})
    for _, ind := range y.Indicator {
        writer.Write([]string{ind.Name, ind.Type, ind.Value, ind.Data})
    }
}

type Indicator struct {
    Name string
    Type string
    Value string
    Data string
}

func (i Indicator) Function() string {
    t := ""
    switch i.Type {
    case "file":
        t = file(i.Value)
    case "url":
        return ""
    }

    data := struct {
        Name string
        Type string
        Value string
        Data string
        AlphaTitle []string
    } {
        Name: strings.ReplaceAll(i.Name, " ", "-"),
        Type: i.Type,
        Value: i.Value,
        Data: strings.ToLower(i.Data),
        AlphaTitle: []string{"a", "b", "c", "d", "e", "f", "g", "h"},
    }

    tmpl, err := template.New(i.Name).Parse(t)
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

func file(value string) string {
    switch value {
    case "name":
        return "error"
    case "hashes.'SHA-256'":
        return sha256Hash()
    }
    return ""
}

func sha256Hash() string {
    return `
rule {{.Name}}-sha256 {
    condition:
        hash.sha256(0, filesize) == "{{.Data}}"
}`
}
