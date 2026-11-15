package local

import (
	"errors"
	"os"
	"path/filepath"
)

const aplicacao = "prumo"

func Dados() (string, error) {
	return base("XDG_DATA_HOME", ".local/share", "dados")
}

func Estado() (string, error) {
	return base("XDG_STATE_HOME", ".local/state", "estado")
}

func base(variavelXDG, padrao, subpastaPrumoHome string) (string, error) {
	if raiz := os.Getenv("PRUMO_HOME"); raiz != "" {
		return filepath.Join(raiz, subpastaPrumoHome), nil
	}
	if xdg := os.Getenv(variavelXDG); filepath.IsAbs(xdg) {
		return filepath.Join(xdg, aplicacao), nil
	}
	casa, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("não achei a pasta do usuário; defina PRUMO_HOME")
	}
	return filepath.Join(casa, filepath.FromSlash(padrao), aplicacao), nil
}
