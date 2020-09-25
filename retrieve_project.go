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

func main() {
	client := &http.Client{}
	// postData := make([]byte, 100)
	req, err := http.NewRequest("GET", "https://www.freelancer.com/api/projects/0.1/projects/active/?compact=&project_types%5B%5D=fixed&max_avg_price=500&min_avg_price=250&query=django", nil)
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("freelancer-oauth-v1", "1Dik9bnPVKncY80lae7OeE7mg1JR5r")
	resp, err := client.Do(req)
	fmt.Println(resp)
	fmt.Println("#########################")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	str := string(body[:])
	str = str[strings.Index(str, "[") : strings.LastIndex(str, "]")+1]
	fmt.Println(str)
	fmt.Println("#########################")
	fmt.Println(str)

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

	for _, proj := range jsondata {
		fmt.Println("*********************")
		fmt.Println(proj)
		var record []string
		record = append(record, strconv.Itoa(proj.ID))
		record = append(record, strconv.FormatUint(uint64(proj.Owner_ID), 10))
		record = append(record, proj.Title)
		record = append(record, proj.Status)
		record = append(record, proj.Seo_URL)

		// Currency                CurrencyStruct
		record = append(record, strconv.Itoa(proj.Currency.ID))
		record = append(record, proj.Currency.Code)
		record = append(record, proj.Currency.Sign)
		record = append(record, proj.Currency.Name)
		record = append(record, strconv.FormatFloat(float64(proj.Currency.Exchange_Rate), 'f', -1, 64))
		record = append(record, proj.Currency.Country)
		record = append(record, strconv.FormatBool(proj.Currency.Is_External))
		record = append(record, strconv.FormatBool(proj.Currency.Is_Escrowcom_Supported))

		record = append(record, strconv.FormatUint(uint64(proj.Submitdate), 10))
		record = append(record, proj.Preview_Description)
		record = append(record, strconv.FormatBool(proj.Deleted))
		record = append(record, strconv.FormatBool(proj.NonPublic))
		record = append(record, strconv.FormatBool(proj.HidBids))
		record = append(record, proj.Type)
		record = append(record, strconv.FormatUint(uint64(proj.BidPeriod), 10))

		// Budget                  BudgetStruct
		record = append(record, strconv.FormatFloat(float64(proj.Budget.Minimum), 'f', -1, 64))
		record = append(record, strconv.FormatFloat(float64(proj.Budget.Maximum), 'f', -1, 64))

		record = append(record, strconv.FormatBool(proj.Featured))
		record = append(record, strconv.FormatBool(proj.Urgent))

		// Bid_Stats               BidStatsStruct
		record = append(record, strconv.Itoa(proj.Bid_Stats.Bid_Count))
		record = append(record, strconv.FormatFloat(float64(proj.Bid_Stats.Bid_Avg), 'f', -1, 64))

		record = append(record, strconv.FormatUint(uint64(proj.Time_Submitted), 10))
		record = append(record, strconv.FormatUint(uint64(proj.Time_Updated), 10))

		// Upgrades                UpgradesStruct
		record = append(record, strconv.FormatBool(proj.Upgrades.Featured))
		record = append(record, strconv.FormatBool(proj.Upgrades.Sealed))
		record = append(record, strconv.FormatBool(proj.Upgrades.NonPublic))
		record = append(record, strconv.FormatBool(proj.Upgrades.FullTime))
		record = append(record, strconv.FormatBool(proj.Upgrades.Urgent))
		record = append(record, strconv.FormatBool(proj.Upgrades.Qualified))
		record = append(record, strconv.FormatBool(proj.Upgrades.NDA))
		record = append(record, strconv.FormatBool(proj.Upgrades.Ip_Contract))
		record = append(record, strconv.FormatBool(proj.Upgrades.Non_Complete))
		record = append(record, strconv.FormatBool(proj.Upgrades.Project_Management))
		record = append(record, strconv.FormatBool(proj.Upgrades.Pf_Only))

		record = append(record, proj.Language)
		record = append(record, strconv.FormatBool(proj.Hireme))
		record = append(record, proj.Frontend_Project_Status)

		// Location                LocationStruct

		record = append(record, strconv.FormatBool(proj.Local))
		record = append(record, strconv.FormatBool(proj.Negotiated))
		record = append(record, strconv.FormatUint(uint64(proj.Time_Free_Bids_Expire), 10))
		// record = append(record, strings.Join(proj.PoolIds, ","))
		record = append(record, proj.PoolIds)
		record = append(record, proj.EnterpriseIds)
		record = append(record, strconv.FormatBool(proj.IsEscrowProject))
		record = append(record, strconv.FormatBool(proj.IsSellerKycRequired))
		record = append(record, strconv.FormatBool(proj.IsBuyerKycRequired))
		record = append(record, proj.ProjectRejectReason)

		writer.Write(record)
	}

	// remember to flush!
	writer.Flush()
}
