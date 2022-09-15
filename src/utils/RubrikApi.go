package utils

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/oracle_with_warning/src/configs"
)

func RubrikApi(url string) (r *http.Response, err error) {
	var (
		user = configs.GetEnvKeys("RUBRIK_API_USERNAME")
		pass = configs.GetEnvKeys("RUBRIK_API_PASSWORD")
	)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	d := http.Client{
		Timeout:   time.Minute * 5,
		Transport: tr,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal("Erro: ", err)
	}

	req.SetBasicAuth(user, pass)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := d.Do(req)

	if err != nil {
		log.Fatal("Erro: ", err.Error())
	}

	return res, err
}
