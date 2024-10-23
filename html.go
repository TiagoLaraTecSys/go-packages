package main

// Titulo resgata o titulo das urls
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {

		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			aRetorno := r.FindStringSubmatch(string(html))

			if cap(aRetorno) == 0 {
				c <- "Erro ao ler página " + url
				return
			}
			c <- aRetorno[1]
		}(url)
	}
	return c
}
