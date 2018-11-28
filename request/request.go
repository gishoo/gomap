package request

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"

    )

//needed
// referer : google  user agent : change every once in a while


//should be
//func libRequest(headers, proxy, url)

func Get(){
    response, err := http.Get("http://golang.org/")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }/*
    resp, err := soup.Get("https://xkcd.com")
    if err != nil {
        os.Exit(1)
    }
    doc := soup.HTMLParse(resp)
    links := doc.Find("div", "id", "comicLinks").FindAll("a")
    for _, link := range links {
        fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
    }
    */
}
