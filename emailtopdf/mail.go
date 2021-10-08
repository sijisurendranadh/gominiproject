package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/sg3des/eml"
)

var (
	Filename string
)

func init() {
	if len(os.Args) < 2 || len(os.Args[1]) == 0 {
		fmt.Println("you must set file")
		os.Exit(255)
	} else {
		Filename = os.Args[1]
	}
}

func main() {
	emlraw, err := ioutil.ReadFile(Filename)
	checkerr(err, "file "+Filename+" not found or can not be readed")

	m, err := eml.Parse(emlraw)
	checkerr(err, "failed parse file")

	dir := Filename[0 : len(Filename)-len(path.Ext(Filename))] //crop extension

	err = os.MkdirAll(dir, 0755)
	checkerr(err, "failed create directory for save data")

	for _, attachment := range m.Attachments {
		err = ioutil.WriteFile(path.Join(dir, attachment.Filename), attachment.Data, 0755)
		checkerr(err, "failed save attachment "+attachment.Filename)
	}

	//if len(m.Html) > 0 {

	//	err = ioutil.WriteFile(path.Join(dir, "body.html"), []byte(m.Html), 0755)
	//	checkerr(err, "failed save html body")
	//}

	if len(m.Text) > 0 {

		err = ioutil.WriteFile(path.Join(dir, "body.txt"), []byte(m.Text), 0755)
		checkerr(err, "failed save text body")
	}

	header := []string{m.Date.String(), m.Subject, m.From[0].Email(), m.To[0].Email()}
	err = ioutil.WriteFile(path.Join(dir, "header.txt"), []byte(strings.Join(header, "\n")), 0755)
	checkerr(err, "failed save headers")

	f, err := os.Create("mail.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(m.Text)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")

	email := "mail.txt"
	content, err := ioutil.ReadFile(email)
	if err != nil {
		log.Fatalf("file not found", email)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)
	_ = pdf.OutputFileAndClose("email.pdf")
	fmt.Println("PDF created")

}

func checkerr(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err)
		os.Exit(255)
	}
}

