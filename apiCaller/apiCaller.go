package apiCaller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//voila l'example de l'interface. Go fait du duck typing. c'est a dire que tu n'a pas besoin d'expicitement dire qu'une classe implement une interface.
// Ici l'interface doit repondre a une fonction Succeed() qui retourne un bool.
// Du coup si ta class a cette fonction tu peux t'en servir pour cette interface
type successable interface {
	Succeed() bool
}

//Call is...
// Ici pas besoin de reflect, tu passe directement ta variable en pointeur ici le type interface{} permet
// de passer n'importe quel type. Mais du coup c'est forcement un pointeur
func Call(url string, resp interface{}) (bool, error) {
	fmt.Println(">=== Api call : " + url + "...")
	httpResp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer httpResp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(httpResp.Body)
	err = json.Unmarshal(bodyBytes, resp)
	if err != nil {
		return false, err
	}
	bodyString := string(bodyBytes)
	fmt.Println("===> Response :", bodyString)
	//le probleme est que la tu ne peux plus faire ca vu que tu ne connais pas le type de ton objet. mais je dirais que ce n'est pas la responsabilite de cette fonction d'afficher ca
	// Il y a un moyen simple de resoudre ce problem avec une interface moins generique, je te ferai une demo dans une autre PR

	//appel de l'unique fonction de l'interface:
	return resp.Succeed(), nil
}
