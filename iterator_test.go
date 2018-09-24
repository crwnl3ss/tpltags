package tpltags

import (
	"bytes"
	"reflect"
	"testing"
	"text/template"
)

func TestIterator(t *testing.T) {
	expect := []uint{0, 1, 2, 3, 4}
	result := Iterator(5)
	if !reflect.DeepEqual(expect, result) {
		t.Errorf("%v != %v", expect, result)
	}
}

func TestIteratorTpl(t *testing.T) {
	const expect string = "0123"
	tpl := template.New("Iterator test")
	m := template.FuncMap{"Iterator": Iterator}
	tpl = tpl.Funcs(m)
	tpl.Parse("{{ range Iterator 4 }}{{ . }}{{ end }}")
	buf := bytes.NewBuffer([]byte{})
	if err := tpl.Execute(buf, nil); err != nil {
		t.Errorf("Could not execute template, reason: %s", err)
	}
	got := buf.String()
	if got != expect {
		t.Fatalf("Unexpected template execution! Got: %s, expect: %s", got, expect)
	}
}
