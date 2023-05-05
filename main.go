package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	// Define o regex para encontrar arquivos .css
	cssRegex := regexp.MustCompile(`\.css$`)

	// Define o comando que você deseja executar
	// commandToExecute := "npm run build"

	// Inicia a variável modTime com a hora atual
	var modTime time.Time

	// Loop principal
	for {
		modified := false

		// Percorre todos os arquivos .css no diretório atual e seus subdiretórios
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Se o arquivo tem a extensão .css, verifica se foi modificado
			if cssRegex.MatchString(path) {
				if info.ModTime().After(modTime) {
					modified = true
				}
			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
		}

		// Se algum arquivo foi modificado, executa o comando
		if modified {
			cmd := exec.Command("npm", "run", "build")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				var refresh string

				fmt.Print("\nDigite qualquer tecla para reiniciar, (q) para sair: ")
				fmt.Scanf("%s", &refresh)

				if refresh != "q" {
					continue
				}
				break
			}
			modTime = time.Now()
		}

		// Espera um segundo antes de verificar novamente
		time.Sleep(time.Second * 1)
	}
}
