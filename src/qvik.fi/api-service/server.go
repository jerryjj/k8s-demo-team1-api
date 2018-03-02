package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	payments "qvik.fi/payments"
)

// PSPStatus response
type PSPStatus struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// GetPaymentsStatus - Gets PSP status from Payments service
func GetPaymentsStatus(w http.ResponseWriter, r *http.Request) {
	log.Debugf("GetPaymentsStatus")

	req := &payments.GetPSPStatusRequest{}
	resp, err := paymentsClient.GetPSPStatus(context.Background(), req)
	if err != nil {
		log.Errorf("Error getting PSP status: %v", err)
		json.NewEncoder(w).Encode(PSPStatus{Status: "failed", Message: err.Error()})
		return
	}
	log.Debugf("Got PSP Status %s with message %s", resp.Status, resp.StatusMessage)

	var statusStr = "ok"
	if resp.Status == payments.Status_ERROR {
		statusStr = "failed"
	}
	status := PSPStatus{Status: statusStr, Message: resp.StatusMessage}

	json.NewEncoder(w).Encode(status)
}

// Initializes & runs the REST server in another go routine.
func mustRunRESTServer(port int) {
	log.Debugf("Starting to listen to insecure HTTP port %v..", port)

	router := mux.NewRouter()

	router.HandleFunc("/payments/status", GetPaymentsStatus).Methods("GET")

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
