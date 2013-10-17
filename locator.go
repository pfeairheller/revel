
package revel

import (
	"errors"
	"fmt"
)


type TemplateLocator struct {
	templateLoaders []TemplateLoader
}

func NewTemplateLocator(templatePaths []string) *TemplateLocator {

	mtl := NewGoTemplateLoader(templatePaths)
	locator := &TemplateLocator{
		templateLoaders: []TemplateLoader {mtl},
	}

	return locator
}

func(locator *TemplateLocator) Template(templatePath string) (Template, error) {
	for _, templateLoader := range locator.templateLoaders {
		if template, err := templateLoader.Template(templatePath); err == nil {
			return template, nil
		} else {
			fmt.Println(err)
		}
	}
	return nil, errors.New("Unable to find matching template")
}

func (locator *TemplateLocator) Refresh() *Error {
	for _, templateLoader := range locator.templateLoaders {
		templateLoader.Refresh()
	}

	return nil
}

func (locator *TemplateLocator) Paths() []string {
	paths := []string{}
	for _, ps := range locator.templateLoaders {
		paths = append(paths, ps.Paths()...)
	}

	return paths
}


