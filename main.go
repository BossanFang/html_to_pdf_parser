package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Dish struct {
	Name        string
	Price       string
	Description string
}

type Category struct {
	Name   string
	Dishes []Dish
}

type Menu struct {
	Categories []Category
}

func main() {
	templatePath := "htmls/menu_template.html"

	timestamp := time.Now().Format("20060102150405")
	outputPath := fmt.Sprintf("pdfs/sample_%s.pdf", timestamp)

	menu := Menu{
		Categories: []Category{
			{
				Name: "Appetizers",
				Dishes: []Dish{
					{Name: "Bruschetta", Price: "$8.99", Description: "Toasted bread topped with tomatoes, basil, and garlic."},
					{Name: "Mozzarella Sticks", Price: "$6.99", Description: "Fried mozzarella sticks served with marinara sauce."},
				},
			},
			{
				Name: "Main Courses",
				Dishes: []Dish{
					{Name: "Grilled Salmon", Price: "$18.99", Description: "Fresh salmon fillet grilled to perfection."},
					{Name: "Chicken Alfredo", Price: "$14.99", Description: "Creamy fettuccine pasta with grilled chicken and Alfredo sauce."},
				},
			},
			{
				Name: "Sides",
				Dishes: []Dish{
					{Name: "French Fries", Price: "$3.99", Description: "Crispy golden French fries."},
					{Name: "Garlic Bread", Price: "$4.50", Description: "Toasted bread with garlic butter and herbs."},
				},
			},
			{
				Name: "Desserts",
				Dishes: []Dish{
					{Name: "Chocolate Cake", Price: "$6.99", Description: "Decadent chocolate cake with fudge frosting."},
					{Name: "Cheesecake", Price: "$5.50", Description: "Creamy cheesecake topped with fruit compote."},
				},
			},
		},
	}

	content, err := parseTemplate(templatePath, menu)
	if err != nil {
		log.Fatal("Failed to parse HTML template: ", err)
	}

	err = generatePDF(outputPath, content)
	if err != nil {
		log.Fatal("Failed to generate PDF: ", err)
	}
	fmt.Println("PDF generated successfully")
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func generatePDF(pdfPath, content string) error {
	t := time.Now().Unix()

	tempFilePath := fmt.Sprintf("htmls/menu_%s.html", strconv.FormatInt(int64(t), 10))
	err := ioutil.WriteFile(tempFilePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	defer os.Remove(tempFilePath)

	f, err := os.Open(tempFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		return err
	}

	return nil
}
