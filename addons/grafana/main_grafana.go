package main_grafana

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	datasource = os.Getenv("DATASOURCE")
	username   = os.Getenv("USER")
	password   = os.Getenv("PASS")
)

const DatasourcePath = "/api/datasources/name/"
const HealthPath = "/api/health"
const BaseUrl = "https://grafana.tstsoares1.staging.k8s.sp03.te.tks.sh"

type ResponseDatasource struct {
	Id        int    `json:'id'`
	OrgId     int    `json:'orgId'`
	Name      string `json:'name'`
	Type      string `json:'type'`
	Url       string `json:'url'`
	IsDefault bool   `json:'isDefault'`
}

type ResponseHealth struct {
	Database string `json:'database'`
	Version  string `json:'version'`
}

func VerifyHealth(t *testing.T) {

	client := http.Client{Timeout: 10 * time.Second}

	url := BaseUrl + HealthPath

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code error: expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	data_obj := ResponseHealth{}
	jsonErr := json.Unmarshal(resBody, &data_obj)
	if jsonErr != nil {
		t.Fatal(err)
	}

	if data_obj.Database != "ok" {
		t.Errorf("Response Database Grafana not OK")
	}
}
