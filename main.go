package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "os"
    "flag"
    "bufio"
)

type Calc struct {
    out *bufio.Writer
}


func (c *Calc) Echo(name string) {
    fmt.Fprintf(c.out,"%s\n", name)
}

func (c *Calc) Jump(name string) {
    fmt.Fprintf(c.out,`[jump  storage="scene1.ks"  target="%s"  ]`+"\n", name)
}

func (c *Calc) Select(name string) {
    answers:= strings.Split(name,"\n")
    for _,v :=range answers{
        if v==""{continue}
        tmp:=strings.Split(strings.Trim(v,"- ")," ")
        fmt.Fprintf(c.out,`[glink  color="black"  storage="scene1.ks"  size="20"  text="%s"  x="130"  y="289"  width=""  height=""  _clickable_img=""  target="%s"  ]`+"\n", tmp[0],tmp[1])
    }
}

func (c *Calc) setOut(w *bufio.Writer){
    c.out=w
}

func main() {
    flag.Parse()
    fmt.Printf("%v",flag.Args())
    for _,v:= range flag.Args(){

        file, err := os.Open(v)
        if err != nil {
           continue
        }
        defer file.Close()
        f, err := os.Create(v+".out2.txt")
        defer f.Close()

        w := bufio.NewWriter(f)

        w.Write([]byte("19195656"))



        bytes, _ := ioutil.ReadAll(file)
        str:=string(bytes)
        parser := &Parser{Buffer: str}
        parser.Init()
        parser.setOut(w)
        err2 := parser.Parse()
        if err2 != nil {
            fmt.Println(err)
        } else {
            parser.Execute()
        }

        w.Flush()
    }



}