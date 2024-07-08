package main

import (
	"bytes"
	"encoding/csv"
	"strings"
	"text/template"
)

type Yara struct {
    Indicator []Indicator
}

func (y Yara) File() string {
    out := ""

    for _, ind := range y.Indicator {
        if strings.Contains(ind.Value, "hash") {
            out = "import hash\n"
            break
        }
    }

    for _, ind := range y.Indicator {
        fn := ind.Function()
        if fn == "" {
            continue
        }
        out += fn + "\n"
    }

    return out
}

func (y Yara) Csv() string {
    // file, _ := os.Create("indicators.csv")
    // defer file.Close()
    buf := bytes.Buffer{}

    writer := csv.NewWriter(&buf)
    defer writer.Flush()

    writer.Write([]string{"name", "type", "value", "data"})
    for _, ind := range y.Indicator {
        writer.Write([]string{ind.Name, ind.Type, ind.Value, ind.Data})
    }
    writer.Flush()

    return buf.String()
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
    default:
        return ""
    }

    data := struct {
        Name string
        Type string
        Value string
        Data string
        AlphaTitle []string
    } {
        Name: strings.ToLower(strings.ReplaceAll(i.Name, " ", "-")),
        Type: i.Type,
        Value: i.Value,
        Data: strings.ToLower(i.Data),
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
        return ""
    case "hashes.'SHA-256'":
        return sha256Hash()
    case "hashes.'MD5'":
        return md5hash()
    }
    return ""
}

func sha256Hash() string {
    return `
rule {{.Name}}-sha256-{{.Data}} {
    condition:
        hash.sha256(0, filesize) == "{{.Data}}"
}`
}

func md5hash() string {
    return `
rule {{.Name}}-md5-{{.Data}} {
    condition:
        hash.md5(0, filesize) == "{{.Data}}"
}`
}

func stringMatch() string {
    return `
rule {{.Name}}-string-{{.Data}} {
    strings:
        $a = "{{.Data}}"
    condition:
        $a
}`
}
