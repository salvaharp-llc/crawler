package main

import (
	"fmt"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

func (cfg *config) isInternalLink(currentURL *url.URL) bool {
	baseDomain, err := publicsuffix.EffectiveTLDPlusOne(cfg.baseURL.Hostname())
	if err != nil {
		fmt.Printf("Error - isInternalLink: couldn't obtain baseURL eTLD+1: %v\n", err)
		return false
	}

	currentDomain, err := publicsuffix.EffectiveTLDPlusOne(currentURL.Hostname())
	if err != nil {
		fmt.Printf("Error - isInternalLink: couldn't obtain currentURL eTLD+1: %v\n", err)
		return false
	}

	return baseDomain == currentDomain
}
