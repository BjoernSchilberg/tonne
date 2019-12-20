package abfallberatungen

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BjoernSchilberg/tonne/helper"
	"github.com/tealeg/xlsx"
)

var url = "https://owncloud.nabu.de/owncloud/index.php/s/sn3rUqxW98Gjd3F/download"

type abfallberatung struct {
	ID                int    `xlsx:"0" json:"id,omitempty"`
	Verwaltungeinheit string `xlsx:"1" json:",omitempty"`
	Entsorgungsgebiet string `xlsx:"2" json:",omitempty"`
	PLZ               string `xlsx:"3" json:",omitempty"`
	ZugeordneterKreis string `xlsx:"4" json:"zugeordneter_Kreis,omitempty"`
	Unternehmen1      string `xlsx:"5" json:"Unternehmen 1,omitempty"`
	Unternehmen2      string `xlsx:"6" json:"Unternehmen 2,omitempty"`
	Abfallberater     string `xlsx:"7" json:",omitempty"`
	Adresse           string `xlsx:"8" json:",omitempty"`
	EMail             string `xlsx:"9" json:"E-Mail,omitempty"`
	Telefon           string `xlsx:"10" json:",omitempty"`
	Angebote          string `xlsx:"11" json:",omitempty"`
	LinkText1         string `xlsx:"12" json:"Linktext 1,omitempty"`
	Link1             string `xlsx:"13" json:"Link 1,omitempty"`
	LinkText2         string `xlsx:"14" json:"Linktext 2,omitempty"`
	Link2             string `xlsx:"15" json:"Link 2,omitempty"`
	LinkText3         string `xlsx:"16" json:"Linktext 3,omitempty"`
	Link3             string `xlsx:"17" json:"Link 3,omitempty"`
}

func getData(url string) ([]abfallberatung, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(string(body))

	xlFile, error := xlsx.OpenBinary(body)
	if error != nil {
		log.Fatalln(error)
	}
	sheet := xlFile.Sheets[0]
	dieAbfallberatung := abfallberatung{}
	var abfallberatungen []abfallberatung
	for i, row := range sheet.Rows {
		if i != 0 && i >= 4 {
			if row != nil {
				row.ReadStruct(&dieAbfallberatung)
				abfallberatungen = append(abfallberatungen, dieAbfallberatung)
				//fmt.Printf("%+v\n", g)
			}
		}
	}
	//fmt.Printf("%v", target)

	return abfallberatungen, error
}

// Get : Get abfallberatungen
func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		abfallberatungen, err := getData(url)
		if err != nil {
			helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		//s, _ := json.Marshal(abfallberatungen)

		helper.RespondWithJSON(w, 200, abfallberatungen)

	}

}
