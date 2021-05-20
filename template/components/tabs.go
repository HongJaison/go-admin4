package components

import (
	"html/template"

	"github.com/HongJaison/go-admin4/template/types"
)

type TabsAttribute struct {
	Name string
	Data []map[string]template.HTML
	types.Attribute
}

func (compo *TabsAttribute) SetData(value []map[string]template.HTML) types.TabsAttribute {
	compo.Data = value
	return compo
}

func (compo *TabsAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "tabs")
}
