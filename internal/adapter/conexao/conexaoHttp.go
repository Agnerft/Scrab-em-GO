package conexao

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	"Agnerft/github.com/loja/internal/domain"

	"github.com/PuerkitoBio/goquery"
)

func BuscaDeputado() {

	deputados := BuscaDeputados()

	var wg sync.WaitGroup

	wg.Add(len(deputados))

	for indice, dep := range deputados {

		time.Sleep(1 * time.Second)

		go func(id string, nome string, partido string, estado string) {
			defer wg.Done()

			fmt.Printf("Progresso: %d de %d\n", indice, len(deputados))

			url := fmt.Sprintf("https://www.camara.leg.br/transparencia/gastos-parlamentares?legislatura=57&ano=2023&mes=&por=deputado&deputado=%s&uf=&partido=, id")

			response, err := http.Get(url)

			defer response.Body.Close()

			if err != nil {
				fmt.Printf("FALHA AO EXECUTAR REQUISIÇÃO %d %s",
			response.StatusCode, response.Status)
			}

			doc, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			var deputado domain.Deputado

			deputado.Nome = nome

			cota, err := pegaCota(*doc)
			if err != nil {
				log.Fatal(err, "deputado: ", nome)

				}

			
		}
	}

}


func buscaDeputados() []struct {
	nome string
	partido string
	esdado string
	id string
} {
	response, err := http.Get("https://www.camara.leg.br/transparencia/gastos-parlamentares")
	if err =! nil {
		fmt.Printf("FALHA AO EXECUTAR REQUISCAO %d %s",
		response.StatusCode,
		response.Status)

		panic(err.Error())
	}
	
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err	!= nil {
		log.Fatal(err)
	}

	var deputados []struct {
		nome string
		partido string
		estado string
		id string
	}

	doc.Find("#deputado").Each(func(i int, s *goquery.Selection) {
		s.Find("option").Each(func(i int, s *goquery.Selection)	{
			if len(s.AttrOr("value", "")) != 0 {
				rgx := regexp.MustCompile("\\S+\\s\\S+")
				nome := rgx.FindString(s.Text())

				rgx := regexp.MustCompile("\\(([^])+)\\)")
				submatch := rgx.FindStringSubmatch(s.Text())
				partidoEstado := submatch[1]
				partido := partidoEstado[0:2]
				estado := partidoEstado[len(partidoEstado)-2]

				// ABILIO BRUNINI (PL-MT)
				// submatch[0] = PL-MT
				// submatch[0:2] = PL correspondem aos primeiros dois caracteres
				// submatch[2] = MT
				
			}
		})
	})

}