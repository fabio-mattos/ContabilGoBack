package database

import (
	"database/sql"

	"github.com/fabio-mattos/ContabilGoBack/internal/entity"
)

func InsertTipoLancamento(db *sql.DB, tipoLancamento *entity.TipoLancamento) error {
	stmt, err := db.Prepare("insert into en_tipo_lancamento	(cd_tipo_lancamento,	cd_empresa ,	nm_tipo_lancamento ,	dt_cadastro , " +
		"id_usario_cad ,	fg_ativo ,	id) 	values 	(	?,	?,	?,	?,	?,	?,	?	)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tipoLancamento.CdTipoLancamento, tipoLancamento.CdEmpresa, tipoLancamento.NmTipoLancamento, tipoLancamento.DtCadastro, tipoLancamento.IdUsuarioCad, tipoLancamento.FgAtivo, tipoLancamento.Id)
	if err != nil {
		return err
	}
	return nil
}

func FindTipoLancamentoByNome(db *sql.DB, nome string) (*entity.TipoLancamento, error) {
	stmt, err := db.Prepare("select cd_tipo_lancamento,	cd_empresa,	nm_tipo_lancamento,	dt_cadastro, id_usario_cad,	fg_ativo, id " +
		" from en_tipo_lancamento where lower(sem_acento(nm_tipo_lancamento)) = lower(sem_acento(?))")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var t entity.TipoLancamento
	err = stmt.QueryRow(nome).Scan(&t.CdTipoLancamento, &t.NmTipoLancamento, &t.CdEmpresa, &t.IdUsuarioCad, &t.IdUsuarioAlt, &t.DtCadastro, &t.DtAlteracao, &t.FgAtivo, &t.Id)

	if err != nil {
		return nil, err
	}
	return &t, nil
}
