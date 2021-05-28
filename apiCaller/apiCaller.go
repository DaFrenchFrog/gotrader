package apiCaller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Call is...
// Ici pas besoin de reflect, tu passe directement ta variable en pointeur ici le type interface{} permet
// de passer n'importe quel type. Mais du coup c'est forcement un pointeur
func Call(url string, resp interface{}) {
	httpResp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer httpResp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(httpResp.Body)
	err = json.Unmarshal(bodyBytes, resp) // plus besoin de precicer que tu envoi un pointeur avec & vu que ta variable en est un
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("resp success: %v \n", resp.Success)
	//le probleme est que la tu ne peux plus faire ca vu que tu ne connais pas le type de ton objet. mais je dirais que ce n'est pas la responsabilite de cette fonction d'afficher ca
	// Il y a un moyen simple de resoudre ce problem avec une interface moins generique, je te ferai une demo dans une autre PR
}
