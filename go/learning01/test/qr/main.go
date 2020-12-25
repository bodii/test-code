package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address")

var temp1 = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// QR ...
func QR(w http.ResponseWriter, req *http.Request) {
	temp1.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
	<head>
		<title>QR Link Generator</title>
	</head>
	<body>
	{{ if . }}
		<img src="https://tool.lu/qrcode/basic.html?tolerance=15&size=256&front_color=%23000000&background_color=%23ffffff&text={{.}}" />
		<br>
	{{ . }}
	<br>
	<br>
	{{ end }}
		<form action="/" name=f method="GET">
			<input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
			<input type=submit value="Show QR" name=qr>
		</form>
	</body>
</html>
`
