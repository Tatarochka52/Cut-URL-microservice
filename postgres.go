package main

import (
  "fmt"
  "database/sql"
  "os"
  )


var db *sql.DB


func connectDb() error{
  var err error

  db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, pgPort, user, password, pgName))
  if err != nil{
    fmt.Println(err.Error())
    os.Exit(1)
  }

  return nil
}


func insertDataInPostgres(stringURL string, resultURL string){
  var err error

  if dataIsInTable(stringURL) == false{
    _, err = db.Exec(`INSERT INTO "Cut URL" ("Input URL", "Output URL") VALUES($1, $2)`, stringURL, resultURL)
    if err != nil{
      fmt.Println(err.Error())
      os.Exit(1)
    }
  } else{
    resultURL = getLinkFromPostgres(stringURL)
    link.OutputLink = resultURL
  }

}


func dataIsInTable(stringURL string) bool{
  row := db.QueryRow(`SELECT "Input URL" FROM "Cut URL" WHERE "Input URL" = $1`, stringURL)
  var tmp interface{}
  err := row.Scan(&tmp)
  if err == sql.ErrNoRows{
    return false
  }
  if err == nil{
    return true
  }

  return false
}


func getLinkFromPostgres(stringURL string) string{
  var res string

  row := db.QueryRow(`SELECT "Output URL" FROM "Cut URL" WHERE "Input URL" = $1`, stringURL)
  row.Scan(&res)
  return res
}


func openLinkFromPostgres(stringURL string) string{
  var res string

  row := db.QueryRow(`SELECT "Input URL" FROM "Cut URL" WHERE "Output URL" = $1`, stringURL)
  row.Scan(&res)
  return res
}


func closeDb(){
  defer db.Close()
}
