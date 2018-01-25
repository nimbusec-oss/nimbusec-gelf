package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/nimbusec-oss/nimbusec"
)

type WebhookBody struct {
	Domain nimbusec.Domain  `json:"domain"`
	Issues []nimbusec.Issue `json:"issues"`
}

func main() {
	if os.Getenv("PORT") == "" {
		log.Fatal("env variable PORT is empty or missing")
	}
	if os.Getenv("GELF_ADDR") == "" {
		log.Fatal("env variable GELF_ADDR is empty or missing")
	}

	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, http.HandlerFunc(WebhookR))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}

func WebhookR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("invalid method: %s", r.Method)
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}

	header := r.Header.Get("X-Nimbusec-Signature")
	sig, _ := base64.StdEncoding.DecodeString(header)
	if !VerifySignature(sig, data) {
		log.Printf("invalid signature")
		http.Error(w, "invalid signature", http.StatusBadRequest)
		return
	}

	body, err := ParseWebhookBody(data)
	if err != nil {
		log.Printf("failed to decode json: %v", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	err = SendToGelfEndpoint(body)
	if err != nil {
		log.Printf("failed to send data to gelf server: %v", err)
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}

	// use http.Error because it is the easiest way to send plain text with
	// correct headers set.
	log.Printf("sent issues to gelf server :)")
	http.Error(w, "ok", http.StatusOK)
}

func VerifySignature(signature []byte, data []byte) bool {
	hashed := sha512.Sum512(data)
	err := rsa.VerifyPKCS1v15(PublicKey, crypto.SHA512, hashed[:], signature)
	return err == nil
}

func ParseWebhookBody(data []byte) (WebhookBody, error) {
	body := WebhookBody{}
	err := json.Unmarshal(data, &body)
	return body, err
}

func SendToGelfEndpoint(body WebhookBody) error {
	facility := path.Base(os.Args[0])
	host, err := os.Hostname()
	if err != nil {
		return err
	}

	addr := os.Getenv("GELF_ADDR")
	writer, err := gelf.NewWriter(addr)
	if err != nil {
		return err
	}

	for _, issue := range body.Issues {
		short := fmt.Sprintf(
			"Nimbusec detected a %q issue with severity \"%d\" for %q\n",
			issue.Event, issue.Severity, body.Domain.Name)
		extra := map[string]interface{}{
			"_nimbusec_domain_id":      body.Domain.ID,
			"_nimbusec_domain_name":    body.Domain.Name,
			"_nimbusec_domain_url":     body.Domain.URL,
			"_nimbusec_issue_id":       issue.ID,
			"_nimbusec_issue_event":    issue.Event,
			"_nimbusec_issue_category": issue.Category,
			"_nimbusec_issue_severity": issue.Severity,
		}
		extra = flatten(extra, issue.Details, "_nimbusec_issue_details")

		msg := &gelf.Message{
			Version:  "1.1",
			Host:     host,
			Short:    string(short),
			TimeUnix: float64(time.Now().Unix()),
			Level:    6, // info
			Facility: facility,
			Extra:    extra,
		}
		err = writer.WriteMessage(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func flatten(flat map[string]interface{}, nested interface{}, prefix string) map[string]interface{} {
	assign := func(key string, v interface{}) {
		switch v.(type) {
		case map[string]interface{}, []interface{}:
			flatten(flat, v, key)
		default:
			flat[key] = v
		}
	}

	switch nested.(type) {
	case map[string]interface{}:
		for k, v := range nested.(map[string]interface{}) {
			key := prefix + "_" + k
			assign(key, v)
		}
	case []interface{}:
		for i, v := range nested.([]interface{}) {
			key := prefix + "_" + strconv.Itoa(i)
			assign(key, v)
		}
	default:
		flat[prefix] = nested
	}

	return flat
}
