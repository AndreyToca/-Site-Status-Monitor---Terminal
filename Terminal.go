package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	for {
		exibeIntroducao()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs....")
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Digite uma valor válido")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	fmt.Println("Olá, qual comando deseja executar?")
	fmt.Println("1 - Monitorar sites")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Quit")

}

func lerComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido) // Lê a tecla pressionada e aponta para comandoLido
	return comandoLido
}

func iniciarMonitoramento() {
	const monitoramento = 5
	const delay = 5

	for i := 0; i < monitoramento; i++ {
		fmt.Println("Monitorando...")
		sites := leListaSites() // Certifique-se de que o nome da função está correto

		for _, site := range sites {
			resp, err := http.Get(site)
			if err != nil {
				fmt.Println("Ocorreu um erro:", err)
				continue // Pula para o próximo site se houver um erro
			}

			if resp.StatusCode == 200 {
				fmt.Println("Site:", site, "foi carregado com sucesso!")
				registraLog(site, true)
			} else {
				fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
				registraLog(site, false)

			}
		}
		time.Sleep(delay * time.Second)
	}
}

func leListaSites() []string {

	var sites []string
	arquivo, err := os.Open("listaSites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {

			break
		}
	}
	arquivo.Close()
	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online -" + strconv.FormatBool(status) + "\n")
}

func imprimeLogs() {

	arquivo, err := os.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))
}
