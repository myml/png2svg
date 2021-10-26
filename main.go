package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"os"
	"text/template"
)

const tmp = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="{{ .W }}px" height="{{ .H }}px" viewBox="0 0 {{ .W }} {{ .H }}" enable-background="new 0 0 {{ .W }} {{ .H }}" xml:space="preserve">  <image id="image0" width="{{ .W }}" height="{{ .H }}" x="0" y="0"
			href="data:image/png;base64,{{ .Data }}" />
</svg>`

func main() {
	for _, f := range os.Args[1:] {
		data, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		img, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		size := img.Bounds().Size()
		bs64 := base64.RawStdEncoding.EncodeToString(data)

		t, err := template.New("").Parse(tmp)
		if err != nil {
			panic(err)
		}
		var buff bytes.Buffer
		err = t.Execute(&buff, map[string]interface{}{"W": size.X, "H": size.Y, "Data": bs64})
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(f+".svg", buff.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	}
}
