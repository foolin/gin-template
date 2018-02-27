/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

/*
Golang template for gin framework, Use golang html/template syntax,
Easy and simple to use for gin framework, See https://github.com/foolin/gin-template
for more information.
*/
package gintemplate

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var (
	htmlContentType   = []string{"text/html; charset=utf-8"}
	templateEngineKey = "github.com/foolin/gin-template/templateEngine"
)

type TemplateEngine struct {
	config   TemplateConfig
	tplMap   map[string]*template.Template
	tplMutex sync.RWMutex
}

type TemplateConfig struct {
	Root         string           //view root
	Extension    string           //template extension
	Master       string           //template master
	Partials     []string         //template partial, such as head, foot
	Funcs        template.FuncMap //template functions
	DisableCache bool             //disable cache, debug mode
}

func New(config TemplateConfig) *TemplateEngine {
	return &TemplateEngine{
		config:   config,
		tplMap:   make(map[string]*template.Template),
		tplMutex: sync.RWMutex{},
	}
}

func Default() *TemplateEngine {
	return New(TemplateConfig{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: false,
	})
}

// You should use helper func `Middleware()` to set the supplied
// TemplateEngine and make `HTML()` work validly.
func HTML(ctx *gin.Context, code int, name string, data interface{}) {
	if val, ok := ctx.Get(templateEngineKey); ok {
		if e, ok := val.(*TemplateEngine); ok {
			e.HTML(ctx, code, name, data)
			return
		}
	}
	ctx.HTML(code, name, data)
}

//New gin middleware for func `gintemplate.HTML()`
func NewMiddleware(config TemplateConfig) gin.HandlerFunc {
	return Middleware(New(config))
}

func Middleware(e *TemplateEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(templateEngineKey, e)
	}
}

func (e *TemplateEngine) Instance(name string, data interface{}) render.Render {
	return TemplateRender{
		Engine: e,
		Name:   name,
		Data:   data,
	}
}

func (e *TemplateEngine) HTML(ctx *gin.Context, code int, name string, data interface{}) {
	instance := e.Instance(name, data)
	ctx.Render(code, instance)
}

func (e *TemplateEngine) executeRender(out io.Writer, name string, data interface{}) error {
	useMaster := true
	if filepath.Ext(name) == e.config.Extension {
		useMaster = false
		name = strings.TrimRight(name, e.config.Extension)

	}
	return e.executeTemplate(out, name, data, useMaster)
}

func (e *TemplateEngine) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) error {
	var tpl *template.Template
	var err error
	var ok bool

	allFuncs := make(template.FuncMap, 0)
	allFuncs["include"] = func(layout string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		err := e.executeTemplate(buf, layout, data, false)
		return template.HTML(buf.String()), err
	}

	// Get the plugin collection
	for k, v := range e.config.Funcs {
		allFuncs[k] = v
	}

	e.tplMutex.RLock()
	tpl, ok = e.tplMap[name]
	e.tplMutex.RUnlock()

	exeName := name
	if useMaster && e.config.Master != "" {
		exeName = e.config.Master
	}

	if !ok || e.config.DisableCache {
		tplList := []string{name}
		if useMaster {
			//render()
			if e.config.Master != "" {
				tplList = append(tplList, e.config.Master)
			}
			tplList = append(tplList, e.config.Partials...)
		} else {
			//renderFile()
			tplList = append(tplList, e.config.Partials...)
		}

		// Loop through each template and test the full path
		tpl = template.New(name)
		for _, v := range tplList {
			// Get the absolute path of the root template
			path, err := filepath.Abs(e.config.Root + string(os.PathSeparator) + v + e.config.Extension)
			if err != nil {
				return fmt.Errorf("TemplateEngine path:%v error: %v", path, err)
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("TemplateEngine render read name:%v, path:%v, error: %v", v, path, err)
			}
			content := fmt.Sprintf("%s", data)
			tpl, err = tpl.New(v).Funcs(allFuncs).Parse(content)
			if err != nil {
				return fmt.Errorf("TemplateEngine render parser name:%v, path:%v, error: %v", v, path, err)
			}
		}
		e.tplMutex.Lock()
		e.tplMap[name] = tpl
		e.tplMutex.Unlock()
	}

	// Display the content to the screen
	err = tpl.Funcs(allFuncs).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return fmt.Errorf("TemplateEngine execute template error: %v", err)
	}

	return nil
}

type TemplateRender struct {
	Engine *TemplateEngine
	Name   string
	Data   interface{}
}

func (r TemplateRender) Render(w http.ResponseWriter) error {
	return r.Engine.executeRender(w, r.Name, r.Data)
}

func (r TemplateRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = htmlContentType
	}
}
