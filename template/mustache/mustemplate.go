/**
 * Created with IntelliJ IDEA.
 * User: pfeairheller
 * Date: 10/14/13
 * Time: 9:16 PM
 * To change this template use File | Settings | File Templates.
 */
package mustache

import (
	"io"
	"github.com/robfig/revel"
	"github.com/hoisie/mustache"
	"os"
)

type MustacheTemplateLoader struct {
	paths []string
	files map[string]string
	templatePaths map[string]string
}


func NewMustacheTemplateLoader(paths []string) *MustacheTemplateLoader {
	out := &MustacheTemplateLoader {
		paths: paths,
	}
	out.files = make(map[string]string)
	out.Refresh()

	return out
}

func (loader *MustacheTemplateLoader) Refresh() *revel.Error {
	return nil
}

func (loader *MustacheTemplateLoader) WatchDir(info os.FileInfo) bool {
	return true
}

func (loader *MustacheTemplateLoader) WatchFile(basename string) bool {
	return false
}

func (loader *MustacheTemplateLoader) Template(name string) (revel.Template, error) {
	path := name + ".mustache"
	for _, base := range loader.paths {

	}
	return nil, nil
}

func (loader *MustacheTemplateLoader) Paths() []string {
	return []string{}
}


type MustacheTemplate struct {
	loader *MustacheTemplateLoader
	name string
}

// return a 'revel.Template' from Go's template.
func (mustmpl *MustacheTemplate) Render(wr io.Writer, arg interface{}) error {
	if _, err :=  wr.Write([]byte(mustache.Render(arg.(string)))); err !=nil {
		return err
	}

	return nil
}

func (mustmpl *MustacheTemplate) Name() string {
	return mustmpl.name
}

func (mustmpl MustacheTemplate) Content() []string {
	content, _ := revel.ReadLines(mustmpl.loader.templatePaths[mustmpl.Name()])
	return content
}




