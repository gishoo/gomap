package request

import (
    "fmt"
    "os"
    "github.com/anaskhan96/soup"
    )


func Get(addr string) string{
    addr = "http://" + addr
    resp, err := soup.Get(addr)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    }
    return resp
}

var elems  = []string{"a","applet", "audio", "button", "embed", "iframe", "img", "input", "link", "menu", "object", "option", "param", "picture", "scripts", "select", "source", "video", "meta"}

func Parse(response string){

    doc := soup.HTMLParse(response)

    for i := 0; i < len(elems); i++{
        switch elems[i]{
            case "a":
                found := doc.FindAll(elems[i])
                fmt.Println("** A **")
                fmt.Println(found)
            case "applet":
                fmt.Println("")
            case "audio":
                fmt.Println("")
            case "button":
                fmt.Println("")
            case "embed":
                fmt.Println("")
            case "iframe":
                fmt.Println("")
            case "img":
                fmt.Println("")
            case "input":
                fmt.Println("")
            case "link":
                fmt.Println("")
            case "menu":
                fmt.Println("")
            case "object":
                fmt.Println("")
            case "option":
                fmt.Println("")
            case "param":
                fmt.Println("")
            case "picture":
                fmt.Println("")
            case "scripts":
                fmt.Println("")
            case "select":
                fmt.Println("")
            case "source":
                fmt.Println("")
            case "video":
                fmt.Println("")
            case "meta":
                found := doc.FindAll(elems[i])
                fmt.Println("**META**")
                fmt.Println(found)
        }
    }

/*
    links := doc.FindAll("div")
    for _, link := range links {
        fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
    }
*/
}
