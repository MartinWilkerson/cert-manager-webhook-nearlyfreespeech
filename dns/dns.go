package dns

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/cert-manager/webhook-example/auth"
)

var (
	baseUrl        = "https://api.nearlyfreespeech.net"
	authHeaderName = "X-NFSN-Authentication"
)

func SetTXTRecord(domain string, dnsName string, key string, login string, apiKey string) error {
	urlPath := fmt.Sprintf("/dns/%s/addRR", domain)
	requestUrl := fmt.Sprintf("%s%s", baseUrl, urlPath)

	values := url.Values{"name": {dnsName}, "type": {"TXT"}, "data": {key}}
	body := values.Encode()
	authHeader, err := auth.GetAuthHeader(login, apiKey, urlPath, body)
	if err != nil {
		return err
	}

	bodyReader := strings.NewReader(body)
	req, err := http.NewRequest("POST", requestUrl, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(authHeaderName, authHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}