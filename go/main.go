package main

import (
  "fmt"
  "net/http"
  "html/template"
  "os"
)

var hangman Hangman
var tmpl2 *template.Template
var tmpl3 *template.Template

func resultHandler(w http.ResponseWriter, r *http.Request) {
  wordGuess := string(hangman.HiddenWord) == hangman.WordToGuess
  lost := hangman.Attempts >= 10

  date := struct {
    Won bool
    Lost bool
    Word string
  }{
    Won : wordGuess,
    Lost : lost,
    Word : hangman.WordToGuess,
  }

  err := tmpl3.ExecuteTemplate(w, "result", date)
  if err != nil{
    http.Error(w, err.Error(), http.StatusBadRequest)
  }
}

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
         fmt.Println(hangman.WordToGuess)
        } else {
          wordGuess := string(hangman.HiddenWord) == hangman.WordToGuess
          lost := hangman.Attempts >= 10

          fmt.Println(hangman.HiddenWord)
          fmt.Println(hangman.WordToGuess)
          fmt.Println(wordGuess)
          fmt.Println(lost)

          if wordGuess || lost {
            http.Redirect(w, r ,"/result" , http.StatusSeeOther)
          }
        }

    tmpl2.Execute(w, hangman)
}

func main() {
  fmt.Println("Lancement du serveur: localhost:80")

  tmpl := template.Must(template.ParseFiles("../index.html"))
  tmpl2 = template.Must(template.ParseFiles("../page/hangman.html"))
  tmpl3 = template.Must(template.ParseFiles("../page/display.html"))
  fs := http.FileServer(http.Dir("../css"))
  http.Handle("/css/", http.StripPrefix("/css/", fs))

  fs2 := http.FileServer(http.Dir("../img"))
  http.Handle("/img/", http.StripPrefix("/img/", fs2))


  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, nil)
  })

  http.HandleFunc("/page/hangman", hangmanHandler)

  http.HandleFunc("/result" , resultHandler)

  http.ListenAndServe(":80", nil)
}
