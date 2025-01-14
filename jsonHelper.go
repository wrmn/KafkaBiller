package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/mofax/iso8583"
)

// Any helper to process JSON data
// converter, formatter, etc

// Convert ISO message to JSON
func convertIsoToJson(parsedIso iso8583.IsoStruct) Transaction {
	var response Transaction

	log.Println("Converting ISO8583 to JSON")

	emap := parsedIso.Elements.GetElements()

	// Format padded CardAcceptor data
	cardAcceptorTerminalId := strings.TrimRight(emap[41], " ")
	cardAcceptorName := strings.TrimRight(emap[43][:25], " ")
	cardAcceptorCity := strings.TrimRight(emap[43][25:38], " ")
	cardAcceptorCountryCode := strings.TrimRight(emap[43][38:], " ")

	// Map ISO8583 format to JSON data
	response.Pan = emap[2]
	response.ProcessingCode = emap[3]
	response.TotalAmount, _ = strconv.Atoi(emap[4])
	response.SettlementAmount = emap[5]
	response.CardholderBillingAmount = emap[6]
	response.TransmissionDateTime = emap[7]
	response.SettlementConversionRate = emap[9]
	response.CardHolderBillingConvRate = emap[10]
	response.Stan = emap[11]
	response.LocalTransactionTime = emap[12]
	response.LocalTransactionDate = emap[13]
	response.CaptureDate = emap[17]
	response.CategoryCode = emap[18]
	response.PointOfServiceEntryMode = emap[22]
	response.Refnum = emap[37]
	response.CardAcceptorData.CardAcceptorTerminalId = cardAcceptorTerminalId
	response.CardAcceptorData.CardAcceptorName = cardAcceptorName
	response.CardAcceptorData.CardAcceptorCity = cardAcceptorCity
	response.CardAcceptorData.CardAcceptorCountryCode = cardAcceptorCountryCode
	response.AdditionalData = emap[48]
	response.Currency = emap[49]
	response.SettlementCurrencyCode = emap[50]
	response.CardHolderBillingCurrencyCode = emap[51]
	response.AdditionalDataNational = emap[57]

	log.Println("Convert success")
	return response
}
