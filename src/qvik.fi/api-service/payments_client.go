package main

import (
	"google.golang.org/grpc"
	payments "qvik.fi/payments"
)

// Payments Service client
var paymentsClient payments.PaymentsClient

func mustCreatePaymentsClient(serverAddr string) {
	log.Debugf("Connecting to payments RPC on address %s", serverAddr)

	con, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Errorf("Error connecting to Payments service: %v", err)
	}

	paymentsClient = payments.NewPaymentsClient(con)
	log.Debugf("Payments RPC Client created")
}
