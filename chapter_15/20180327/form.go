package main

import (
	"io"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, request.FormValue("in"))
	}
}
func main() {
	http.HandleFunc("/form", FormServer)
	if err := http.ListenAndServe("0.0.0.0:8888", nil); err != nil {
		panic(err)
	}
}
