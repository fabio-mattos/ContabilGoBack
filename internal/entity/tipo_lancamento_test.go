package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTipoLancamento(t *testing.T) {

	tl, err := NewTipoLancamento(1, 1, "Tipo Lancamento 01", "81094965987")
	assert.Nil(t, err)
	assert.NotNil(t, tl)
	assert.NotEmpty(t, tl.Id)
	assert.Equal(t, "Tipo Lancamento 01", tl.NmTipoLancamento)
	assert.Equal(t, int8(1), tl.CdEmpresa)
}

func TestTipoLancamentoWhenNameIsRequired(t *testing.T) {
	tl, err := NewTipoLancamento(int32(1), int8(1), "", "81094965987")
	assert.Nil(t, tl)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestTipoLancamentoWhenEmpresaIsRequired(t *testing.T) {
	p, err := NewTipoLancamento(1, 0, "Product 1", "000")
	assert.Nil(t, p)
	assert.Equal(t, ErrEmpresaIsRequired, err)
}
