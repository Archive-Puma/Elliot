package modules

import (
	"encoding/json"
)

// TODO: More info from Sucuri

// === MODULE METHOD ===

func moduleSiteCheck(domain string, channel *chan map[string]interface{}) {
	service := "Sucuri"
	// Data structure
	results := struct {
		Scan struct {
			Error string `json:"error"`
		} `json:"scan"`
		Warnings struct {
			ScanFailed []struct {
				Type string `json:"type"`
			} `json:"scan_failed"`
		} `json:"warnings"`

		Site struct {
			IP        []string `json:"ip"`
			URL       string   `json:"final_url"`
			Server    []string `json:"running_on"`
			Redirects []string `json:"redirects_to"`
		} `json:"site"`
		Spider struct {
			Links      []string `json:"urls"`
			JsLocal    []string `json:"js_local"`
			JSExternal []string `json:"js_external"`
		} `json:"links"`
		Rating struct {
			Total struct {
				Score string `json:"rating"`
			} `json:"total"`
		} `json:"ratings"`
	}{}
	// Get the data
	body, err := apiGet(service, "https://sitecheck.sucuri.net/api/v3/?scan=%s", domain)
	if err != nil {
		*channel <- nil
		return
	}
	// Parse the JSON
	err = json.Unmarshal(body, &results)
	if err != nil || results.Scan.Error != "" || (len(results.Warnings.ScanFailed) > 0 && results.Warnings.ScanFailed[0].Type == "Scan failed") {
		*channel <- nil
		return
	}
	// Simplifies the data
	server := ""
	if len(results.Site.Server) > 0 {
		server = results.Site.Server[0]
	}
	ipv4, ipv6 := "", ""
	if len(results.Site.IP) > 0 {
		ipv4 = results.Site.IP[0]
	}
	if len(results.Site.IP) > 1 {
		ipv6 = results.Site.IP[1]
	}
	redirects := results.Site.Redirects
	if len(redirects) > 0 {
		last := len(redirects) - 1
		for i := 0; i < len(redirects)/2; i++ {
			redirects[i], redirects[last-i] = redirects[last-i], redirects[i]
		}
	}

	js := append(results.Spider.JsLocal, results.Spider.JSExternal...)
	*channel <- map[string]interface{}{
		"url":       results.Site.URL,
		"rating":    results.Rating.Total.Score,
		"server":    server,
		"redirects": redirects,
		"links":     results.Spider.Links,
		"js":        js,
		"ipv4":      ipv4,
		"ipv6":      ipv6,
	}
}
