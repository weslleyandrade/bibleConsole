package provider

import (
	"strings"
	"testing"
)

func TestGetChapter(t *testing.T) {
	b := &Biblebox{}
	res := b.GetChapter("genesis", "1")
	if !strings.Contains(res, "Deus criou os") {
		t.Error("Nenhum conteudo foi retornado")
	}
}

func TestGetVerses(t *testing.T) {
	b := &Biblebox{}
	res := b.GetVerses("genesis", "1", "1")

	if !strings.Contains(res, "Deus criou os") {
		t.Error("Nenhum conteudo foi retornado")
	}
}
