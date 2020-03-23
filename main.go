package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var bioHazard string = `
                         __    _                                   
                    _wr""        "-q__                             
                 _dP                 9m_     
               _#P                     9#_                         
              d#@                       9#m                        
             d##                         ###                       
            J###                         ###L                      
            {###K                       J###K                      
            ]####K      ___aaa___      J####F                      
        __gmM######_  w#P""   ""9#m  _d#####Mmw__                  
     _g##############mZ_         __g##############m_               
   _d####M@PPPP@@M#######Mmp gm#########@@PPP9@M####m_             
  a###""          ,Z"#####@" '######"\g          ""M##m            
 J#@"             0L  "*##     ##@"  J#              *#K           
 #"                #    "_gmwgm_~    dF                #_
7F                 "#_   ]#####F   _dK                 JE
]                    *m__ ##### __g@"                   F
                       "PJ#####LP"
                         0######_                       
                       _0########_
     .               _d#####^#####m__              ,
      "*w_________am#####P"   ~9#####mw_________w*"
          ""9@#####@M""           ""P@#####@M""

          SARS-CoV-19 world population data:

world:					chile:	
`

//TODO:
// use https://services.arcgis.com/5T5nSi527N4F7luB/arcgis/rest/services/COVID_19_CasesByCountry(pt)_VIEW/FeatureServer/0/query?where=1%3D1&outFields=ADM0_NAME%2Ccum_conf%2Ccum_death&returnGeometry=false&f=pjson
// or https://coronavirus-19-api.herokuapp.com/all
// or https://corona.lmao.ninja/all

func main() {
	fmt.Print(bioHazard)
	cwp, err := getCurentWorldPopulation()
	if err != nil {
		log.Fatal("error gettings world population, ", err)
		return
	}
	ccp, err := getCurrentChileanPopulation()
	if err != nil {
		log.Fatal("error gettings world population, ", err)
		return
	}
	infected, deaths, err := getCovidData()
	if err != nil {
		log.Fatal("error gettings CovidData ", err)
		return
	}
	chileInfected, chileDeaths, err := getChileData()
	if err != nil {
		log.Fatal("error gettings CovidData ", err)
		return
	}
	infectedp := (infected * 100) / cwp
	deathsp := (deaths * 100) / cwp
	chileInfectedP := (chileInfected * 100) / ccp
	chileDeathsP := (chileDeaths * 100) / ccp
	fmt.Printf("infected: %0.f %f%%", infected, infectedp)
	fmt.Printf("\t\tinfected: %0.f %f%%\n", chileInfected, chileInfectedP)
	fmt.Printf("deaths:   %0.f %f%%", deaths, deathsp)
	fmt.Printf("\t\tdeaths:   %0.f %f%%\n", chileDeaths, chileDeathsP)
}

func getCovidData() (float64, float64, error) {
	resp, err := http.Get("https://coronavirus-19-api.herokuapp.com/all")
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	type dataStructure struct {
		Cases     float64
		Deaths    float64
		Recovered float64
	}
	var data dataStructure
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, 0, err
	}
	return data.Cases, data.Deaths, nil

	//TODO: unmarshall the json
}

func getChileData() (float64, float64, error) {
	resp, err := http.Get("https://coronavirus-19-api.herokuapp.com/countries/chile")
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	type dataStructure struct {
		Cases     float64
		Deaths    float64
		Recovered float64
	}
	var data dataStructure
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, 0, err
	}
	return data.Cases, data.Deaths, nil
}

func getCurentWorldPopulation() (float64, error) {
	return 7772867646, nil //estimate ~
}

func getCurrentChileanPopulation() (float64, error) {
	return 19000000, nil
}
