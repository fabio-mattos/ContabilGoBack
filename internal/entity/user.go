package entity

import (
	"time"

	"github.com/fabio-mattos/ContabilGoBack/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	IDUsuario     entity.ID `json:"id_usuario"`
	DeLogin       string    `json:"de_login"`
	DeSenha       string    `json:"_"`
	DtCadastro    time.Time `json:"dt_cadastro"`
	DtAlteracao   time.Time `json:"dt_alteracao"`
	IdUsuarioCad  string    `json:"id_usuario_cad"`
	IdUsuarioAlt  string    `json:"id_usuario_alt"`
	NmCompleto    string    `json:"nm_completo"`
	NuFoneCelular string    `json:"nu_fone_celular"`
	DeEmail       string    `json:"de_email"`
	CdPerfil      int       `json:"cd_perfil"`
	FgAtivo       bool      `json:"fg_ativo"`
}

func NewUser(login, nome, email, password string, perfil int) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		IDUsuario:  entity.NewID(),
		DeLogin:    login,
		NmCompleto: nome,
		DeEmail:    email,
		DeSenha:    string(hash),
		CdPerfil:   perfil,
		FgAtivo:    true,
		DtCadastro: time.Now(),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.DeSenha), []byte(password))
	return err == nil
}
