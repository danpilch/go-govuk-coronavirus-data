package main

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var cvdataurl = "https://c19downloads.azureedge.net/downloads/json/coronavirus-deaths_latest.json"

type JsonStruct struct {
	Metadata struct {
		LastUpdatedAt time.Time `json:"lastUpdatedAt"`
		Disclaimer    string    `json:"disclaimer"`
	} `json:"metadata"`
	Countries []struct {
		AreaCode            string `json:"areaCode"`
		AreaName            string `json:"areaName"`
		ReportingDate       string `json:"reportingDate"`
		DailyChangeInDeaths int    `json:"dailyChangeInDeaths"`
		CumulativeDeaths    int    `json:"cumulativeDeaths"`
	} `json:"countries"`
	Overview []struct {
		AreaCode            string `json:"areaCode"`
		AreaName            string `json:"areaName"`
		ReportingDate       string `json:"reportingDate"`
		DailyChangeInDeaths int    `json:"dailyChangeInDeaths"`
		CumulativeDeaths    int    `json:"cumulativeDeaths"`
	} `json:"overview"`
}

func parseJson(j []byte) JsonStruct {
	// Create JsonStruct
	var DataStruct JsonStruct
	// Parse json bytes
	if err := json.Unmarshal(j, &DataStruct); err != nil {
		panic(err)
	}
	return DataStruct
}

func httpGetJsonData() []byte {
	// Get
	resp, err := http.Get(cvdataurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Get data from body
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return responseBytes
}

func renderTable(data JsonStruct) {
	tableData := [][]string{}
	for _, s := range data.Overview {
		tableData = append(tableData, []string{
			string(s.ReportingDate),
			string(strconv.Itoa(s.CumulativeDeaths)),
			string(strconv.Itoa(s.DailyChangeInDeaths)),
		})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Deaths", "Change"})
	table.SetFooter([]string{"", "Total", string(strconv.Itoa(data.Overview[0].CumulativeDeaths))}) // Add Footer
	table.AppendBulk(tableData)                                                                     // Add Bulk Data
	table.Render()
}

func main() {
	// http get json
	jsonResponse := httpGetJsonData()
	// Parse jsonResponse
	data := parseJson(jsonResponse)
	// output table
	renderTable(data)
}
