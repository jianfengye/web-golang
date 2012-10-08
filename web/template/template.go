package template

import(
	"html/template"
	"utility/file"
	"path"
	"net/http"
	"errors"
)

var templateDir string

var templates map[string]*template.Template 

func PreLoad(dir string) {
	templateDir = dir
	templates = make(map[string]*template.Template)
	fileInfoArr, err := file.ReadAllFiles(templateDir)
	if err != nil {
		panic(err)
		return
	}

	var templateName string
	for filepath, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		t := template.Must(template.ParseFiles(filepath))
		templates[filepath] = t
	}
}

func RenderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error{
	if v, ok := templates[tmpl]; ok == true {
		err := v.Execute(w, locals)
		if err == nil {
			return nil
		}
	}
	return errors.New("RenderHtml error")
}