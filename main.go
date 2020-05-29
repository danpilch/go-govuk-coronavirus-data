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

var cvdataurl string = "https://c19downloads.azureedge.net/downloads/json/coronavirus-deaths_latest.json"
var inReverse bool = true

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
	var totalDeaths int
	// define the order of dates
	if inReverse {
		for i := len(data.Overview)/2 - 1; i >= 0; i-- {
			opp := len(data.Overview) - 1 - i
			data.Overview[i], data.Overview[opp] = data.Overview[opp], data.Overview[i]
		}
	}
	// iterate data and create table data
	for _, s := range data.Overview {
		tableData = append(tableData, []string{
			string(s.ReportingDate),
			string(strconv.Itoa(s.CumulativeDeaths)),
			string(strconv.Itoa(s.DailyChangeInDeaths)),
		})
		totalDeaths += s.DailyChangeInDeaths
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Deaths", "Change"})
	table.SetFooter([]string{"", "Total", string(strconv.Itoa(totalDeaths))}) // Add Footer
	table.AppendBulk(tableData)                                               // Add Bulk Data
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
