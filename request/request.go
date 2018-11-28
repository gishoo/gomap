package request

import (
    "fmt"
    "os"
    "github.com/anaskhan96/soup"
    )

//needed
// referer : google  user agent : change every once in a while


//should be
//func libRequest(headers, proxy, url)

func Get(addr){
    resp, err := soup.Get(addr)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    doc := soup.HTMLParse(resp)
    links := doc.Find("div", "id", "comicLinks").FindAll("a")
    for _, link := range links {
        fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
    }

}
