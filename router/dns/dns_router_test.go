package dns

import (
	"bytes"
	"demo/model/dns"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Router level tests serve to test complex validation and parsing of the request, and the validate the responses as well
// Business logic testing should be left to the service layer

// In this case, no need to test validation as there is nothing complex about it (required fail is guaranteed)
// So we can just test that end to end functionality is working and response has the correct structure
func TestGetDNS(t *testing.T) {
	body := map[string]interface{}{
		"x":   "1",
		"y":   "2",
		"z":   "3",
		"vel": "4",
	}

	w := httptest.NewRecorder()
	url := "/v1/dns/find"

	json_body, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewReader(json_body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if w.Code != 200 {
		t.Fatalf("wanted status 200, got %d\n%v", w.Code, string(data))
	}

	var dnsResponse dns.GetDNSResponse
	if err := json.Unmarshal(data, &dnsResponse); err != nil {
		t.Fatal(err)
	}

	// Assuming static SectorID=1 here
	if dnsResponse.Loc != (1 + 2 + 3 + 4) {
		t.Fatalf("expected %d, got %f", (1 + 2 + 3 + 4), dnsResponse.Loc)
	}

}
