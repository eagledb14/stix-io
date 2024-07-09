package templates

import (

)

func Index(stix, yara string) string {
    data := struct {
        Stix string
        Yara string
    } {
        Stix: stix,
        Yara: yara,
    }

    const index = `
    <form name="yara" hx-target="body">
        <div class="grid-center">
            <div class="middle">
                <button class="middle blue" hx-post="/csv">Csv</button>
                <button class="middle blue" hx-post="/yara">Yara</button>
                <button class="middle blue" onclick="copyText()">Copy</button>
            </div>
        </div>
        <div class="content">
            <textarea autofocus required placeholder="Stix Input" class="gray" name="stix">{{.Stix}}</textarea>
            <textarea readonly placeholder="Yara Input" class="gray" name="yara" id="yara">{{.Yara}}</textarea>
        </div>
    </form>
    `

    return Execute("index", index, data)
}
