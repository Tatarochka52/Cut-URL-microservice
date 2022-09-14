package main

import ("net/http"
        "net/url"
        "html/template"
        "flag"
        "fmt"
        "math/rand"
        "os/exec"
        _ "github.com/lib/pq"
)


const (
  host = "localhost"
  pgPort = "5432"
  user = "postgres"
  password = "Plastelin777"
  pgName = "cutURL"
)


type Link struct{
  InputLink string
  OutputLink string
  Status string
}


var link = Link{}
var userFlag = flag.Bool("d", false, "Memory modificator")


func handleRequest(){
  http.HandleFunc("/", mainPage)
  http.HandleFunc("/test", openFullLink)
  http.ListenAndServe(":8080", nil)
}


func openFullLink(w http.ResponseWriter, r *http.Request){
  
}


func mainPage(w http.ResponseWriter, r *http.Request){
  page, _ := template.ParseFiles("pages/cutYourURL.html")

  if r.Method == "POST"{
    if !checkIsValidURL(r.FormValue("input")){
      link.Status = "Invalid URL format"
      fmt.Println(link.Status)
      link.InputLink = ""
    } else{
      link.InputLink = r.FormValue("input")
      link.OutputLink = cutURLLink()
      link.Status = "Successful"


      flag.Parse()
      if *userFlag{
        connectDb()
        insertDataInPostgres(link.InputLink, link.OutputLink)
        closeDb()

        fmt.Println("Connected to PostgreSQL")

      } else{
        openFile()
        insertDataInFile(link.InputLink, link.OutputLink)
        closeFile()

        fmt.Println("Connected to local database")
      }
    }
  }

  page.Execute(w, link)
}


func checkIsValidURL(s string) bool{
  u, err := url.Parse(s)
  if err != nil || u.Host == ""{
    return false
  }

  return true
}

func cutURLLink() string{
  letterString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

  res := make([]byte, 5)
  for i := range res{
    res[i] = letterString[rand.Intn(len(letterString))]
  }

  return string(res)
}


func main(){
  handleRequest()
}
