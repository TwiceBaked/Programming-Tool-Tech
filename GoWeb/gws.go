package main

import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type",
        "text/html",
    )
    io.WriteString(
        res,
        `<DOCTYPE html>
        <html>
          <head>
              <title>Hello, World</title>
          </head>
          <body>
              Hello, World!
          </body>
        </html>`,
    )
}

func JSON(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type","text/html",)
    io.WriteString(
        res,
        `<DOCTYPE html>
        <html>
          <head>
              <title>JSON GWS</title>
          </head>
          <body>
              This is the JSON site
          </body>
        </html>`,
    )
}


func main() {
    http.HandleFunc("/hello", hello)
	http.HandleFunc("/JSON", JSON)
    http.ListenAndServe(":9000", nil)
}