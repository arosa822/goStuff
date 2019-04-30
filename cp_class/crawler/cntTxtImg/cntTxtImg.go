// FindLinks1 prints the links in an HTML document read from standard input.
package main

import (
    "fmt"
    "os"
    "golang.org/x/net/html"
)

func main() {
    // assign doc to os.Stdin via html
    doc, err := html.Parse(os.Stdin)
    // error check
    if err != nil {
        fmt.Fprintf(os.Stderr," findlink1: %v\n", err)
        os.Exit(1)
    }
   // for _, link :=range visit(nil, doc) {
   //     fmt.Println(link)
     for _, images := range countImg(nil,doc){
         fmt.Println(images)
    }
}

/*
 The visit function below traverses an HTML node tree, extracts the link from the 
 href attribute ofeach anchor element <a href='...'> appends the links to a slice 
 of strings and returns the resulting slice
*/

// visit appends the links each link found in n and returns the results
func visit(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                links = append(links, a.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links,c)
    }
    return links
}

func countImg(images []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "img" {
        fmt.Println(n.Attr)
        for _, i := range n.Attr {
            if i.Key == "src" {
                images = append(images, i.Val)
            }
        }
    }
    return images
}
