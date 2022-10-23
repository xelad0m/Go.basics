package main

import (
	"encoding/csv"
	"encoding/xml"
	"os"

	// "fmt"
	"io"
	"strconv"
	"strings"
)

// начало решения
type Org struct {
	XMLName      xml.Name     `xml:"organization"`
	Organization []Department `xml:"department"`
}

type Department struct {
	XMLName   xml.Name   `xml:"department"`
	Code      string     `xml:"code"`
	Employees []Employee `xml:"employees>employee"`
}

type Employee struct {
	XMLName xml.Name `xml:"employee"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	City    string   `xml:"city"`
	Salary  int      `xml:"salary"`
}

// csv умеет писать только срезы строк,
// сделаем из организации срез срезов строк по сотрудникам
func (org Org) ToList() [][]string {
	var result [][]string
	header := []string{"id","name","city","department","salary"}
	result = append(result, header)

	for _, dept := range org.Organization {
		for _, e := range dept.Employees {
			emp_str := []string{
				strconv.Itoa(e.Id),
				e.Name,
				e.City,
				dept.Code,
				strconv.Itoa(e.Salary),
			}
			result = append(result, emp_str)
		}
	}
	return result
}

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках

func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
	var b []byte
	buf := make([]byte, 32)
	for {
		n, err := inXML.Read(buf)
		for _, ch := range buf[:n] {
			b = append(b, ch)
		}
		if err == io.EOF {
			break
		}
	}

	var org Org
	err := xml.Unmarshal(b, &org)

	if err != nil {
		return err
	}

	w := csv.NewWriter(outCSV)
	for _, emp_str := range org.ToList() {
		err = w.Write(emp_str)
		if err != nil {
			return err
		}

	}
	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

// конец решения


/* Правильный вариант (декодер без размаршалинга)
// Определим структуры для декодирования организации, департамента и сотрудника из XML:

// Organization описывает организацию
type Organization struct {
    Departments []Department `xml:"department"`
}

// Department описывает департамент организации
type Department struct {
    Code      string     `xml:"code"`
    Employees []Employee `xml:"employees>employee"`
}

// Employee описывает сотрудника департамента
type Employee struct {
    Id     int     `xml:"id,attr"`
    Name   string  `xml:"name"`
    City   string  `xml:"city"`
    Salary float64 `xml:"salary"`
}
// Логику трансформации XML → CSV можно полностью уложить в ConvertEmployees(), но тогда она получится довольно громоздкой. Поэтому я добавил вспомогательную функцию decodeOrganization(), которая декодирует организацию из XML:

func decodeOrganization(in io.Reader) (Organization, error) {
    var org Organization
    decoder := xml.NewDecoder(in)
    err := decoder.Decode(&org)
    return org, err
}
// И вспомогательный тип employeeWriter, который пишет сотрудников в CSV:

// employeeWriter записывает сотрудников в CSV
type employeeWriter struct {
    w   *csv.Writer
    err error
}

// writeHeader записывает заголовок
func (ew *employeeWriter) writeHeader() {
    if ew.err != nil {
        return
    }
    header := []string{"id", "name", "city", "department", "salary"}
    ew.err = ew.w.Write(header)
}

// writeEmployee записывает сотрудника в строку
func (ew *employeeWriter) writeEmployee(depCode string, emp Employee) {
    if ew.err != nil {
        return
    }
    fields := []string{
        strconv.Itoa(emp.Id),
        emp.Name,
        emp.City,
        depCode,
        strconv.FormatFloat(emp.Salary, 'f', -1, 64),
    }
    ew.err = ew.w.Write(fields)
}

// flush финализирует данные
func (ew *employeeWriter) flush() error {
    ew.w.Flush()
    if ew.err == nil {
        ew.err = ew.w.Error()
    }
    return ew.err
}

// newEmployeeWriter создает нового писателя сотрудников в CSV
func newEmployeeWriter(w io.Writer) *employeeWriter {
    return &employeeWriter{w: csv.NewWriter(w)}
}
// Теперь ConvertEmployees() будет легче читаться:

// декодируем организацию;
// записываем заголовок CSV;
// проходим по департаментам и сотрудникам внутри них;
// записываем каждого сотрудника в строку CSV;
// финализируем CSV.
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
    org, err := decodeOrganization(inXML)
    if err != nil {
        return fmt.Errorf("failed to parse xml: %w", err)
    }

    w := newEmployeeWriter(outCSV)
    w.writeHeader()

    for _, dep := range org.Departments {
        for _, emp := range dep.Employees {
            w.writeEmployee(dep.Code, emp)
        }
    }

    if err := w.flush(); err != nil {
        return fmt.Errorf("failed writing csv: %w", err)
    }

    return nil
}
*/


func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
        </employees>
    </department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	ConvertEmployees(out, in)

	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}
