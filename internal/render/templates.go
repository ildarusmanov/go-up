package render

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/ildarusmanov/go-up/internal/constants"
)

func RenderTemplates(
	tdir, pdir string,
	payload interface{},
	helpers template.FuncMap,
	l []string,
) error {
	for _, t := range l {
		if err := RenderTemplate(tdir, pdir, payload, helpers, t); err != nil {
			return err
		}
	}

	return nil
}

func RenderTemplate(
	tdir, pdir string,
	payload interface{},
	helpers template.FuncMap,
	tp string,
) error {
	newFilePath := tp[0:(len(tp) - len(constants.TemplateExt))] // remove '.tmpl'

	f, err := os.Create(pdir + "/" + newFilePath)
	if err != nil {
		return fmt.Errorf("file creation error: %s", err)
	}

	defer f.Close()

	t := template.Must(
		template.New(filepath.Base(tp)).
			Funcs(helpers).
			ParseFiles(tdir + "/" + tp),
	)

	if err := t.Execute(f, payload); err != nil {
		return fmt.Errorf("rendering template error: %s", err)
	}

	return nil
}
