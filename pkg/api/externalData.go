package api

import (
	"NameService/pkg/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func EnrichData(fio *model.NameModel) (int, string, string, error) {
	age, err := getAgeFromAgify(fio.Name)
	if err != nil {
		return 0, "", "", err
	}

	gender, err := getGenderFromGenderize(fio.Name)
	if err != nil {
		return 0, "", "", err
	}

	nationality, err := getNationalityFromNationalize(fio.Name)
	if err != nil {
		return 0, "", "", err
	}

	return age, gender, nationality, nil
}

func getAgeFromAgify(name string) (int, error) {
	resp, err := http.Get("https://api.agify.io/?name=" + name)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	age := int(data["age"].(float64))
	return age, nil
}

func getGenderFromGenderize(name string) (string, error) {
	resp, err := http.Get("https://api.genderize.io/?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	gender := data["gender"].(string)
	return gender, nil
}

func getNationalityFromNationalize(name string) (string, error) {
	resp, err := http.Get("https://api.nationalize.io/?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	countries := data["country"]
	if len(countries.([]interface{})) > 0 {
		nationality := countries.([]interface{})[0].(map[string]interface{})["country_id"].(string)
		return nationality, nil
	}

	return "", nil
}
