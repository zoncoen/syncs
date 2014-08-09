package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/elazarl/go-bindata-assetfs"
)

type Option struct {
	Theme    string
	Markdown string
	Master   bool
}

var Commands = []cli.Command{
	commandInit,
	commandUp,
}

var commandInit = cli.Command{
	Name:  "init",
	Usage: "",
	Description: `
`,
	Action: doInit,
}

var commandUp = cli.Command{
	Name:  "up",
	Usage: "",
	Description: `
`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "theme",
			Value: "default",
			Usage: "style theme (default,beige,sky,night,serif,simple,solarized)",
		},
	},
	Action: doUp,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doInit(c *cli.Context) {
}

func doUp(c *cli.Context) {
	mdFilePath := c.Args()[0]
	mdFileName := filepath.Base(mdFilePath)

	// assets
	http.Handle("/css/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, "reveal.js"}))
	http.Handle("/js/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, "reveal.js"}))
	http.Handle("/lib/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, "reveal.js"}))
	http.Handle("/plugin/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, "reveal.js"}))
	http.Handle("/assets/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, ""}))

	// markdown file
	http.HandleFunc("/"+mdFileName,
		func(w http.ResponseWriter, r *http.Request) {
			t, _ := template.ParseFiles(mdFilePath)
			t.Execute(w, "")
		})

	// handle client slides
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			opt := &Option{Theme: c.String("theme"), Markdown: mdFileName, Master: false}
			indexHtml, _ := Asset("tmpl/index.html")
			indexTemplate := template.Must(template.New("index").Parse(string(indexHtml)))
			indexTemplate.Execute(w, opt)
		})

	// generate url of master slides
	h := sha256.New()
	rand.Seed(time.Now().Unix())
	h.Write([]byte(strconv.Itoa(rand.Int())))
	masterPath := "/" + base64.URLEncoding.EncodeToString(h.Sum(nil))
	fmt.Println("Master: http://localhost:8080" + masterPath)

	// handle master slides
	http.HandleFunc(masterPath,
		func(w http.ResponseWriter, r *http.Request) {
			opt := &Option{Theme: c.String("theme"), Markdown: mdFileName, Master: true}
			indexHtml, _ := Asset("tmpl/index.html")
			indexTemplate := template.Must(template.New("index").Parse(string(indexHtml)))
			indexTemplate.Execute(w, opt)
		})

	err := http.ListenAndServe(":8080", nil)
	panic("ListenAndServe: " + err.Error())
}
