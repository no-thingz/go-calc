package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/no-thingz/go-calc/calculatorserver/models"
)

type calcController struct {
	calcType *regexp.Regexp
}

func newCalcController() *calcController {
	return &calcController{
		calcType: regexp.MustCompile(`^/calculator\.(sum|sub|mul|div)$`),
	}
}

func (cc calcController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matches := cc.calcType.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid Endpoint"))
		return
	}
	calcProcess := matches[1]
	if r.Method == http.MethodPost {
		cc.postMethod(w, r, calcProcess)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Allow POST Method only"))
		return
	}
}

func (cc *calcController) postMethod(w http.ResponseWriter, r *http.Request, calcProcess string) {
	c, err := cc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON body request"))
		return
	}

	var cr models.Number
	switch calcProcess {
	case "sum":
		cr = models.CalcSum(c)
	case "sub":
		cr = models.CalcSub(c)
	case "div":
		cr, err = models.CalcDiv(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "mul":
		cr = models.CalcMul(c)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

	encodeResponseAsJSON(cr, w)
}

func (cc *calcController) parseRequest(r *http.Request) (models.Number, error) {
	dec := json.NewDecoder(r.Body)
	var n models.Number
	err := dec.Decode(&n)
	if err != nil {
		return models.Number{}, err
	}
	return n, nil
}
