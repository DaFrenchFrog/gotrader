package apiCaller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/elRomano/gotrader/model"
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
func Call(url string, resp successable) (bool, error) {
	fmt.Println(model.Color("yellow"), ">=== Api call : ", model.Color(""), url, "...")
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

	// bodyString := string(bodyBytes)
	// fmt.Println(model.Color("yellow"), "===> Response :", model.Color(""), bodyString)

	//appel de l'unique fonction de l'interface:
	return resp.Succeed(), nil
}
