package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

// CurrencyStruct ()
type CurrencyStruct struct {
	ID                     int
	Code                   string
	Sign                   string
	Name                   string
	Exchange_Rate          float32
	Country                string
	Is_External            bool
	Is_Escrowcom_Supported bool
}

// BudgetStruct ()
type BudgetStruct struct {
	Minimum float32
	Maximum float32
}

// BidStatsStruct ()
type BidStatsStruct struct {
	Bid_Count int
	Bid_Avg   float32
}

// UpgradesStruct ()
type UpgradesStruct struct {
	Featured           bool
	Sealed             bool
	NonPublic          bool
	FullTime           bool
	Urgent             bool
	Qualified          bool
	NDA                bool
	Ip_Contract        bool
	Non_Complete       bool
	Project_Management bool
	Pf_Only            bool
}

// CountryStruct ()
type CountryStruct struct {
}

// LocationStruct ()
type LocationStruct struct {
	Country CountryStruct
}

// Project ()
type Project struct {
	ID                      int
	Owner_ID                uint32
	Title                   string
	Status                  string
	Seo_URL                 string
	Currency                CurrencyStruct
	Submitdate              uint32
	Preview_Description     string
	Deleted                 bool
	NonPublic               bool
	HidBids                 bool
	Type                    string
	BidPeriod               uint16
	Budget                  BudgetStruct
	Featured                bool
	Urgent                  bool
	Bid_Stats               BidStatsStruct
	Time_Submitted          uint32
	Time_Updated            uint32
	Upgrades                UpgradesStruct
	Language                string
	Hireme                  bool
	Frontend_Project_Status string
	Location                LocationStruct
	Local                   bool
	Negotiated              bool
	Time_Free_Bids_Expire   uint32
	PoolIds                 string
	EnterpriseIds           string
	IsEscrowProject         bool
	IsSellerKycRequired     bool
	IsBuyerKycRequired      bool
	ProjectRejectReason     string
}

func addData(curRow *xlsx.Row, cellData string) string {
	// record = append(record, cellData)
	tmpSpan := curRow.AddCell()
	tmpSpan.Value = cellData
	return cellData
}

func main() {
	client := &http.Client{}
	fmt.Println("Enter 1 if the project type is fixed or 2 if the project type is hourly: ")
	var inputData string
	var proType string
	reqStr := "https://www.freelancer.com/api/projects/0.1/projects/active/?compact="
	fmt.Scanln(&proType)
	if proType == "1" {
		reqStr += "&project_types%5B%5D=fixed"
	} else if proType == "2" {
		reqStr += "&project_types%5B%5D=hourly"
	}

	if proType != "2" {
		fmt.Println("Enter the minimum average price: ")
		fmt.Scanln(&inputData)
		if inputData != "" {
			_, err := strconv.ParseFloat(inputData, 64)
			if err == nil {
				reqStr += "&min_avg_price=" + inputData
			}
		}

		fmt.Println("Enter the maximum average price: ")
		fmt.Scanln(&inputData)
		if inputData != "" {
			_, err := strconv.ParseFloat(inputData, 64)
			if err == nil {
				reqStr += "&max_avg_price=" + inputData
			}
		}
	}

	if proType != "1" {
		fmt.Println("Enter the minimum average hourly rate: ")
		fmt.Scanln(&inputData)
		if inputData != "" {
			_, err := strconv.ParseFloat(inputData, 64)
			if err == nil {
				reqStr += "&min_avg_hourly_rate=" + inputData
			}
		}

		fmt.Println("Enter the maximum average hourly rate: ")
		fmt.Scanln(&inputData)
		if inputData != "" {
			_, err := strconv.ParseFloat(inputData, 64)
			if err == nil {
				reqStr += "&max_avg_hourly_rate=" + inputData
			}
		}
	}

	fmt.Println("Enter the query: ")
	fmt.Scanln(&inputData)
	fmt.Println("query = " + inputData)
	if inputData != "" {
		reqStr += "&query=" + inputData
	}

	fmt.Println("request string = " + reqStr)
	req, err := http.NewRequest("GET", reqStr, nil)
	// &project_types%5B%5D=fixed&max_avg_price=500&min_avg_price=250&query=django", nil)
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("freelancer-oauth-v1", "1Dik9bnPVKncY80lae7OeE7mg1JR5r")
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	str := string(body[:])
	str = str[strings.Index(str, "[") : strings.LastIndex(str, "]")+1]

	// jsonData, err := json.Marshal(str)

	// Unmarshal JSON data

	var jsondata []Project

	err = json.Unmarshal([]byte(str), &jsondata)

	if err != nil {
		fmt.Println(err)
	}

	csvdatafile, err := os.Create("./data.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvdatafile.Close()

	writer := csv.NewWriter(csvdatafile)

	// Generating Excel file & Creating Header
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("tab page 1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	// title := sheet.AddRow()
	// titleRow := title.AddCell()
	// titleRow.HMerge = 11 //Merge the number of columns to the right, not including its own column
	// titleRow.Value = "this is excel title"

	header1 := sheet.AddRow()
	header2 := sheet.AddRow()

	span1 := header1.AddCell()
	header2.AddCell()
	span1.VMerge = 1
	span1.Value = "ID"

	span2 := header1.AddCell()
	header2.AddCell()
	span2.VMerge = 1
	span2.Value = "Owner ID"

	span3 := header1.AddCell()
	header2.AddCell()
	span3.VMerge = 1
	span3.Value = "Title"

	span4 := header1.AddCell()
	header2.AddCell()
	span4.VMerge = 1
	span4.Value = "Status"

	span5 := header1.AddCell()
	header2.AddCell()
	span5.VMerge = 1
	span5.Value = "Seo URL"

	span6 := header1.AddCell()
	for i := 0; i < 7; i++ {
		header1.AddCell()
	}
	span6.HMerge = 7
	span6.Value = "Currency"

	span7 := header2.AddCell()
	span7.Value = "ID"

	span8 := header2.AddCell()
	span8.Value = "Code"

	span9 := header2.AddCell()
	span9.Value = "Sign"

	span10 := header2.AddCell()
	span10.Value = "Name"

	span11 := header2.AddCell()
	span11.Value = "Exchange Rate"

	span12 := header2.AddCell()
	span12.Value = "Country"

	span13 := header2.AddCell()
	span13.Value = "Is External"

	span14 := header2.AddCell()
	span14.Value = "Is Escrowcom Supported"

	span15 := header1.AddCell()
	header2.AddCell()
	span15.VMerge = 1
	span15.Value = "Submit Date"

	span16 := header1.AddCell()
	header2.AddCell()
	span16.VMerge = 1
	span16.Value = "Preview Description"

	span17 := header1.AddCell()
	header2.AddCell()
	span17.VMerge = 1
	span17.Value = "Deleted"

	span18 := header1.AddCell()
	header2.AddCell()
	span18.VMerge = 1
	span18.Value = "NonPublic"

	span19 := header1.AddCell()
	header2.AddCell()
	span19.VMerge = 1
	span19.Value = "HidBids"

	span20 := header1.AddCell()
	header2.AddCell()
	span20.VMerge = 1
	span20.Value = "Type"

	span21 := header1.AddCell()
	header2.AddCell()
	span21.VMerge = 1
	span21.Value = "Bid Period"

	span22 := header1.AddCell()
	header1.AddCell()
	span22.HMerge = 1
	span22.Value = "Budget"

	span23 := header2.AddCell()
	span23.Value = "Minimum"

	span24 := header2.AddCell()
	span24.Value = "Maximum"

	span25 := header1.AddCell()
	header2.AddCell()
	span25.VMerge = 1
	span25.Value = "Featured"

	span26 := header1.AddCell()
	header2.AddCell()
	span26.VMerge = 1
	span26.Value = "Urgent"

	span27 := header1.AddCell()
	header1.AddCell()
	span27.HMerge = 1
	span27.Value = "BidStats"

	span28 := header2.AddCell()
	span28.Value = "Bid Count"

	span29 := header2.AddCell()
	span29.Value = "Bid Avg"

	span30 := header1.AddCell()
	header2.AddCell()
	span30.VMerge = 1
	span30.Value = "Time Submitted"

	span31 := header1.AddCell()
	header2.AddCell()
	span31.VMerge = 1
	span31.Value = "Time Updated"

	span32 := header1.AddCell()
	for i := 0; i < 10; i++ {
		header1.AddCell()
	}
	span32.HMerge = 10
	span32.Value = "Upgrades"

	span33 := header2.AddCell()
	span33.Value = "Featured"

	span34 := header2.AddCell()
	span34.Value = "Sealed"

	span35 := header2.AddCell()
	span35.Value = "NonPublic"

	span36 := header2.AddCell()
	span36.Value = "FullTime"

	span37 := header2.AddCell()
	span37.Value = "Urgent"

	span38 := header2.AddCell()
	span38.Value = "Qualified"

	span39 := header2.AddCell()
	span39.Value = "NDA"

	span40 := header2.AddCell()
	span40.Value = "Ip Contract"

	span41 := header2.AddCell()
	span41.Value = "Non Complete"

	span42 := header2.AddCell()
	span42.Value = "Project Management"

	span43 := header2.AddCell()
	span43.Value = "Pf Only"

	span44 := header1.AddCell()
	header2.AddCell()
	span44.VMerge = 1
	span44.Value = "Language"

	span45 := header1.AddCell()
	header2.AddCell()
	span45.VMerge = 1
	span45.Value = "Hireme"

	span46 := header1.AddCell()
	header2.AddCell()
	span46.VMerge = 1
	span46.Value = "Frontend Project Status"

	span47 := header1.AddCell()
	span47.Value = "Location"

	span48 := header2.AddCell()
	span48.Value = "Country"

	span49 := header1.AddCell()
	header2.AddCell()
	span49.VMerge = 1
	span49.Value = "Local"

	span50 := header1.AddCell()
	header2.AddCell()
	span50.VMerge = 1
	span50.Value = "Negotiated"

	span51 := header1.AddCell()
	header2.AddCell()
	span51.VMerge = 1
	span51.Value = "Time Free Bids Expire"

	span52 := header1.AddCell()
	header2.AddCell()
	span52.VMerge = 1
	span52.Value = "PoolIds"

	span53 := header1.AddCell()
	header2.AddCell()
	span53.VMerge = 1
	span53.Value = "EnterpriseIds"

	span54 := header1.AddCell()
	header2.AddCell()
	span54.VMerge = 1
	span54.Value = "IsEscrowProject"

	span55 := header1.AddCell()
	header2.AddCell()
	span55.VMerge = 1
	span55.Value = "IsSellerKycRequired"

	span56 := header1.AddCell()
	header2.AddCell()
	span56.VMerge = 1
	span56.Value = "IsBuyerKycRequired"

	span57 := header1.AddCell()
	header2.AddCell()
	span57.VMerge = 1
	span57.Value = "ProjectRejectReason"
	recordHeader := []string{"ID", "Owner ID", "Title", "Status", "Seo URL", "ID", "Code", "Sign", "Name", "Exchange Rate", "Country", "Is External", "Is Escrowcom Supported", "Submit Date", "Preview Description", "Deleted", "NonPublic", "HidBids", "Type", "Bid Period", "Minimum", "Maximum", "Featured", "Urgent", "Bid Count", "Bid Avg", "Time Submitted", "Time Updated", "Featured", "Sealed", "NonPublic", "FullTime", "Urgent", "Qualified", "NDA", "Ip Contract", "Non Complete", "Project Management", "Pf Only", "Language", "Hireme", "Frontend Project Status", "Country", "Local", "Negotiated", "Time Free Bids Expire", "PoolIds", "EnterpriseIds", "IsEscrowProject", "IsSellerKycRequired", "IsBuyerKycRequired", "ProjectRejectReason"}
	writer.Write(recordHeader)
	for _, proj := range jsondata {
		// fmt.Println("*********************")
		// fmt.Println(proj)
		var record []string
		curRow := sheet.AddRow()
		record = append(record, addData(curRow, strconv.Itoa(proj.ID)))
		record = append(record, addData(curRow, strconv.FormatUint(uint64(proj.Owner_ID), 10)))
		record = append(record, addData(curRow, proj.Title))
		record = append(record, addData(curRow, proj.Status))
		record = append(record, addData(curRow, proj.Seo_URL))

		// Currency                CurrencyStruct
		record = append(record, addData(curRow, strconv.Itoa(proj.Currency.ID)))
		record = append(record, addData(curRow, proj.Currency.Code))
		record = append(record, addData(curRow, proj.Currency.Sign))
		record = append(record, addData(curRow, proj.Currency.Name))
		record = append(record, addData(curRow, strconv.FormatFloat(float64(proj.Currency.Exchange_Rate), 'f', -1, 64)))
		record = append(record, addData(curRow, proj.Currency.Country))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Currency.Is_External)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Currency.Is_Escrowcom_Supported)))

		record = append(record, addData(curRow, strconv.FormatUint(uint64(proj.Submitdate), 10)))
		record = append(record, addData(curRow, proj.Preview_Description))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Deleted)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.NonPublic)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.HidBids)))
		record = append(record, addData(curRow, proj.Type))
		record = append(record, addData(curRow, strconv.FormatUint(uint64(proj.BidPeriod), 10)))

		// Budget                  BudgetStruct
		record = append(record, addData(curRow, strconv.FormatFloat(float64(proj.Budget.Minimum), 'f', -1, 64)))
		record = append(record, addData(curRow, strconv.FormatFloat(float64(proj.Budget.Maximum), 'f', -1, 64)))

		record = append(record, addData(curRow, strconv.FormatBool(proj.Featured)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Urgent)))

		// Bid_Stats               BidStatsStruct
		record = append(record, addData(curRow, strconv.Itoa(proj.Bid_Stats.Bid_Count)))
		record = append(record, addData(curRow, strconv.FormatFloat(float64(proj.Bid_Stats.Bid_Avg), 'f', -1, 64)))

		record = append(record, addData(curRow, strconv.FormatUint(uint64(proj.Time_Submitted), 10)))
		record = append(record, addData(curRow, strconv.FormatUint(uint64(proj.Time_Updated), 10)))

		// Upgrades                UpgradesStruct
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Featured)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Sealed)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.NonPublic)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.FullTime)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Urgent)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Qualified)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.NDA)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Ip_Contract)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Non_Complete)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Project_Management)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Upgrades.Pf_Only)))

		record = append(record, addData(curRow, proj.Language))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Hireme)))
		record = append(record, addData(curRow, proj.Frontend_Project_Status))

		// Location                LocationStruct
		record = append(record, addData(curRow, strconv.FormatBool(proj.Local)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.Negotiated)))
		record = append(record, addData(curRow, strconv.FormatUint(uint64(proj.Time_Free_Bids_Expire), 10)))
		record = append(record, addData(curRow, proj.PoolIds))
		record = append(record, addData(curRow, proj.EnterpriseIds))
		record = append(record, addData(curRow, strconv.FormatBool(proj.IsEscrowProject)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.IsSellerKycRequired)))
		record = append(record, addData(curRow, strconv.FormatBool(proj.IsBuyerKycRequired)))
		record = append(record, addData(curRow, proj.ProjectRejectReason))

		writer.Write(record)
	}

	// remember to flush!
	writer.Flush()

	file.Save("./1.xlsx")
}
