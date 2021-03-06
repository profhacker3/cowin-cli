package cowin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

type CentreData struct {
	Centers []struct {
		Name     string `json:"name"`
		FeeType  string `json:"fee_type"`
		Sessions []struct {
			SessionID              string   `json:"session_id"`
			Date                   string   `json:"date"`
			AvailableCapacity      int      `json:"available_capacity"`
			MinAgeLimit            int      `json:"min_age_limit"`
			Vaccine                string   `json:"vaccine"`
			Slots                  []string `json:"slots"`
			AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
			AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
		} `json:"sessions"`
	} `json:"centers"`
}

func getApiURL(pincode string) string {
	if pincode != "" {
		return apiPincodeURL
	} else {
		return apiDistrictURL
	}
}

func (center *CentreData) getCenters(options Options) {
	auth := false

	u, err := url.Parse(getApiURL(options.Pincode))

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("date", options.Date)

	if options.Pincode != "" {
		q.Add("pincode", options.Pincode)
	} else {
		districtID := getDistrictID(options.State, options.District)
		q.Add("district_id", districtID)
	}

	u.RawQuery = q.Encode()
	resp, statusCode := getReqAuth(u.String(), "", auth)

	if statusCode != 200 {
		log.Fatalln(string(resp))
	}

	json.Unmarshal(resp, center)

}
func getDoseType(availablity, dose1, dose2 int) string {
	var doseType string
	if availablity == dose1 {
		doseType = "1"
	} else if availablity == dose2 {
		doseType = "2"
	} else if availablity == 0 {
		doseType = "none"
	} else {
		doseType = "both"
	}
	return fmt.Sprint(doseType)
}
func checkDoseType(dosType string, specifiedDose int) bool {
	ok := false
	switch dosType {
	case "both":
		ok = true
	case fmt.Sprint(specifiedDose):
		ok = true
	}
	return ok
}

func PrintCenters(options Options) {
	center := getCenterBookable(options)
	if len(center) > 0 {
		for _, v := range center {
			if options.Info {
				fmt.Printf("%v  %v  %v  %v %v %v Dose-%v\n", v.Name, v.Freetype, v.Date, v.AvailableCapacity, v.Vaccine, v.MinAgeLimit, v.DoseType)
			} else {
				fmt.Printf("%s ", v.Name)
				if v.Freetype != "Free" {
					fmt.Print("Paid")
				}
				fmt.Println()
			}

		}
	} else {
		os.Exit(1)
	}

}
