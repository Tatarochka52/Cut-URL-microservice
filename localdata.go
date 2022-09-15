package main


import (
  "os"
  "fmt"
  "bufio"
)


var file *os.File


func openFile() error{
  var err error

  file, err = os.OpenFile("source/data.txt", os.O_RDWR, 0600)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  return nil
}


func insertDataInFile(stringURL string, resultURL string){
  var err error
  input := [] string {
    stringURL,
    "\n",
    resultURL,
    "\n",
  }

  resultURL = dataIsInFile(stringURL)

  if resultURL == ""{
    for _, line := range input{
      if _, err = file.WriteString(line); err != nil{
        fmt.Println(err.Error())
        os.Exit(1)
      }
    }

  } else{
    link.OutputLink = resultURL
  }

}


func dataIsInFile(stringURL string) string{
  var resultLine string

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan(){
    if fileScanner.Text() == stringURL{
      fileScanner.Scan()
      resultLine = fileScanner.Text()
      return resultLine
    }
  }

  return ""
}


func openLinkFromFile(stringURL string) string{
  var resultLine string

  fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan(){
    if fileScanner.Text() == stringURL{
      return resultLine
    }else{
      resultLine = fileScanner.Text()
    }
  }
  return ""
}


func closeFile(){
  defer file.Close()
}
