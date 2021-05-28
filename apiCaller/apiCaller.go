package apiCaller

import (
	"encoding/json"
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
// ici on utilise l'interface en parametre
func Call(url string, resp successable) (bool, error) {
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

	//appel de l'unique fonction de l'interface:
	return resp.Succeed(), nil
}
