package provider

import (
	"strings"
	"testing"
)

func TestGetChapter(t *testing.T) {
	b := &Biblebox{}
	res := b.GetChapter("1", "1")
	if !strings.Contains(res, "Deus criou os") {
		t.Error("Nenhum conteudo foi retornado")
	}
}

func TestGetVerses(t *testing.T) {
	b := &Biblebox{}
	res := b.GetVerses("1", "1", "1")

	if !strings.Contains(res, "Deus criou os") {
		t.Error("Nenhum conteudo foi retornado")
	}

	resRange := b.GetVerses("romanos", "1", "3-4")

	rm13 := strings.Contains(resRange, "acerca de seu Filho")
	rm14 := strings.Contains(resRange, "Filho de Deus")

	if !rm13 || !rm14 {
		t.Error("Nenhum conteudo foi retornado")
	}
}
