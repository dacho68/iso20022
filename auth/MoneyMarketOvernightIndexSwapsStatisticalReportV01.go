package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500101 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Document"`
	Message *MoneyMarketOvernightIndexSwapsStatisticalReportV01 `xml:"MnyMktOvrnghtIndxSwpsSttstclRpt"`
}

func (d *Document01500101) AddMessage() *MoneyMarketOvernightIndexSwapsStatisticalReportV01 {
	d.Message = new(MoneyMarketOvernightIndexSwapsStatisticalReportV01)
	return d.Message
}

// The MoneyMarketOvernightIndexSwapsStatisticalReport message is sent by the reporting agents to the relevant competent authority, to report the daily overnight index swaps (OIS) transactions.
type MoneyMarketOvernightIndexSwapsStatisticalReportV01 struct {

	// Provides the elements specific to the report.
	ReportHeader *model.MoneyMarketReportHeader1 `xml:"RptHdr"`

	// Provides the reason why no activity is reported or the required list of transactions for the overnight index swaps segment.
	OvernightIndexSwapsReport *model.OvernightIndexSwap3Choice `xml:"OvrnghtIndxSwpsRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MoneyMarketOvernightIndexSwapsStatisticalReportV01) AddReportHeader() *model.MoneyMarketReportHeader1 {
	m.ReportHeader = new(model.MoneyMarketReportHeader1)
	return m.ReportHeader
}

func (m *MoneyMarketOvernightIndexSwapsStatisticalReportV01) AddOvernightIndexSwapsReport() *model.OvernightIndexSwap3Choice {
	m.OvernightIndexSwapsReport = new(model.OvernightIndexSwap3Choice)
	return m.OvernightIndexSwapsReport
}

func (m *MoneyMarketOvernightIndexSwapsStatisticalReportV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
