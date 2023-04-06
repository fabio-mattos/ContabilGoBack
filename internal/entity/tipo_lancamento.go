package entity

import (
	"errors"
	"time"

	"github.com/fabio-mattos/ContabilGoBack/pkg/entity"
)

var (
	ErrIDIsRequired        = errors.New("id is required")
	ErrInvalidID           = errors.New("invalid id")
	ErrNameIsRequired      = errors.New("ame is required")
	ErrEmpresaIsRequired   = errors.New("company is required")
	ErrEntryTypeIsRequired = errors.New("entry type required")
)

type TipoLancamento struct {
	Id               entity.ID `json:"id"`
	CdTipoLancamento int32     `json:"cd_tipo_lancamento"`
	CdEmpresa        int8      `json:"cd_empresa"`
	NmTipoLancamento string    `json:"nm_tipo_lancamento"`
	IdUsuarioCad     string    `json:"id_usuario_cad"`
	IdUsuarioAlt     string    `json:"id_usuario_alt"`
	DtCadastro       time.Time `json:"dt_cadastro"`
	DtAlteracao      time.Time `json:"dt_alteracao"`
	FgAtivo          bool      `json:"fg_ativo"`
}

func NewTipoLancamento(cdTipoLancamento int32, cdEmpresa int8, nmTipoLancamento string, idUsuarioCad string) (*TipoLancamento, error) {
	tipoLancamento := &TipoLancamento{
		Id:               entity.NewID(),
		CdTipoLancamento: cdTipoLancamento,
		CdEmpresa:        cdEmpresa,
		NmTipoLancamento: nmTipoLancamento,
		IdUsuarioCad:     idUsuarioCad,
		DtCadastro:       time.Now(),
		FgAtivo:          true,
	}
	err := tipoLancamento.Validate()
	if err != nil {
		return nil, err
	}
	return tipoLancamento, nil

}

func (t *TipoLancamento) Validate() error {
	if t.Id.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(t.Id.String()); err != nil {
		return ErrInvalidID
	}
	if t.NmTipoLancamento == "" {
		return ErrNameIsRequired
	}
	if !(t.CdEmpresa > 0) {
		return ErrEmpresaIsRequired
	}

	return nil
}
