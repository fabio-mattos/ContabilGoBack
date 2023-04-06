package database

import "github.com/fabio-mattos/ContabilGoBack/internal/entity"

type TipoLancamentoInterface interface {
	InsertTipoLancamento(tipoLancamento *entity.TipoLancamento) error
	FindByNome(nome string) (*entity.TipoLancamento, error)
}
