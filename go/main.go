package main

import (
  "fmt"
  "net/http"
  "html/template"
  "os"
)

var hangman Hangman
var tmpl2 *template.Template

func hangmanHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	mode := queryParams.Get("mode")

	fmt.Println(mode)
	hangman.UserInput = r.FormValue("user_input")

    TestInput(hangman.UserInput)

    if (r.Method != http.MethodPost) {
    if len(os.Args) == 1 {
    filePath := ""
        switch mode {
                case "easy":
                    filePath = "../facile_words.txt"
                case "medium":
                    filePath = "../moyen_words.txt"
                case "hard":
                    filePath = "../difficile_words.txt"
                }
        LoadWords(filePath)
      } else {
        LoadWords(os.Args[1])
      }
         InitHangman()
        }

    tmpl2.Execute(w, hangman)
}

func main() {
  fmt.Println("Lancement du serveur: localhost:80")

  tmpl := template.Must(template.ParseFiles("../index.html"))
  tmpl2 = template.Must(template.ParseFiles("../page/hangman.html"))

  fs := http.FileServer(http.Dir("../css"))
  http.Handle("/css/", http.StripPrefix("/css/", fs))

  fs2 := http.FileServer(http.Dir("../img"))
  http.Handle("/img/", http.StripPrefix("/img/", fs2))


  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, nil)
  })

  http.HandleFunc("/page/hangman", hangmanHandler)

  http.ListenAndServe(":80", nil)
}
