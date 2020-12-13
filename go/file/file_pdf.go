package file

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ThiagoRodriguesdeSantana/desafio_conductor/go-server-server/go/model"
	"github.com/jung-kurt/gofpdf"
)

//Row to file
type Row struct {
	Size  float64
	Value string
}

//PDF struct to definition
type PDF struct {
	head []Row
	pdf  *gofpdf.Fpdf
}

//NewPDF instance for PDF
func NewPDF() *PDF {
	return &PDF{}
}

//GeneratePDF generate PDF file from transactions list
func (p *PDF) GeneratePDF(registers []model.Transaction) string {

	rows := [][]Row{}

	for key, v := range registers {

		r := []Row{}
		val := reflect.Indirect(reflect.ValueOf(v))
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)

			tag := val.Type().Field(i).Tag

			if string(tag.Get("pdf")) == "false" {
				continue
			}

			sizeSt := string(tag.Get("size"))

			size, err := strconv.ParseFloat(sizeSt, 64)

			if err != nil {
				continue
			}

			if key == 0 {
				p.setHeader(size, tag.Get("header"))
			}

			f := valueField.Interface()
			val := reflect.ValueOf(f)

			value := getValue(val)

			ro := Row{
				Size:  size,
				Value: value,
			}

			r = append(r, ro)

		}

		rows = append(rows, r)
	}

	p.newReport()

	p.header()
	p.table(rows)

	path, _ := ioutil.TempDir("", "report")

	pathFile := path + "/report.pdf"

	_ = p.pdf.OutputFileAndClose(pathFile)

	return pathFile
}

func (p *PDF) setHeader(size float64, title string) {
	p.head = append(p.head, Row{Size: size, Value: title})
}

func getValue(val reflect.Value) string {

	var value string

	typeName := strings.ToLower(val.Type().Name())

	switch typeName {
	case "float64":
		valueFloat := val.Float()
		value = fmt.Sprintf("%.2f", valueFloat)
		break
	case "string":
		value = val.String()
		break
	case "int":
		valueInt := val.Int()
		value = strconv.Itoa(int(valueInt))
		break
	default:
		value = ""
	}

	return value

}

func (p *PDF) table(tbl [][]Row) {
	p.pdf.SetFont("Times", "", 16)
	p.pdf.SetFillColor(255, 255, 255)
	for _, lines := range tbl {
		for _, line := range lines {
			p.pdf.CellFormat(line.Size, 7, line.Value, "1", 0, "", false, 0, "")
		}
		p.pdf.Ln(-1)
	}
}

func (p *PDF) header() {
	p.pdf.SetFont("Times", "B", 16)
	p.pdf.SetFillColor(240, 240, 240)

	for _, head := range p.head {
		p.pdf.CellFormat(head.Size, 7, head.Value, "1", 0, "", true, 0, "")
	}

	p.pdf.Ln(-1)

}

func (p *PDF) newReport() {
	p.pdf = gofpdf.New("L", "mm", "Letter", "")
	p.pdf.AddPage()
	p.pdf.SetFont("Times", "B", 28)
	p.pdf.Cell(40, 10, "Relatorio de Transacoes")
	p.pdf.Ln(12)
	p.pdf.SetFont("Times", "", 20)
	p.pdf.Cell(40, 10, time.Now().Format("02-Jan-2006"))
	p.pdf.Ln(20)
}
