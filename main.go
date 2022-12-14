package main

import ("net/http"
        "net/url"
        "html/template"
        "flag"
        "fmt"
        "math/rand"
        "time"
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
  http.ListenAndServe(":8080", nil)
}


func mainPage(w http.ResponseWriter, r *http.Request){
  page, _ := template.ParseFiles("pages/cutYourURL.html")
  var inL = r.URL.Path

  if (inL == "/" || inL == "/cutYourURL.html"){
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
  }else{
    cURL := RemoveChar(inL)
    flag.Parse()
    if *userFlag{
      connectDb()
      http.Redirect(w, r, openLinkFromPostgres(cURL), 301)
      closeDb()
    } else{
      openFile()
      http.Redirect(w, r, openLinkFromPostgres(cURL), 301)
      closeFile()
    }
  }
}


func RemoveChar(word string) string{
  return word[1:len(word)]
}


func checkIsValidURL(s string) bool{
  u, err := url.Parse(s)
  if err != nil || u.Host == ""{
    return false
  }

  return true
}

func cutURLLink() string{
  rand.Seed(time.Now().UnixNano())
  digits := "0123456789"
  specials := "/!@#$?|"
  all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
      "abcdefghijklmnopqrstuvwxyz" +
      digits + specials
  length := 8
  buf := make([]byte, length)
  buf[0] = digits[rand.Intn(len(digits))]
  buf[1] = specials[rand.Intn(len(specials))]
  for i := 2; i < length; i++ {
      buf[i] = all[rand.Intn(len(all))]
  }
  rand.Shuffle(len(buf), func(i, j int) {
      buf[i], buf[j] = buf[j], buf[i]
  })
  str := string(buf)

  return str
}


func main(){
  handleRequest()
}
