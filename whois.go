// -------------------------
//
// Copyright 2015, undiabler
//
// git: github.com/undiabler/golang-whois
//
// http://undiabler.com
//
// Released under the Apache License, Version 2.0
//
//--------------------------

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"strings"

	"github.com/likexian/whois"
)

var orgRegex = regexp.MustCompile(`(?:Organization:)\s*(.+)`)

// func appendIfMissing(slice []string, i string) []string {

// 	i = strings.ToLower(i)

// 	for _, ele := range slice {
// 		if ele == i {
// 			return slice
// 		}
// 	}

// 	return append(slice, i)

// }




func getOrg(domain string){
	result, err := whois.Whois(domain)
	data := make(map[string]interface{})
	orgs := []string{}


	if err != nil {
		log.Println("Error in whois lookup : ", err)
	}else{
		res := orgRegex.FindAllStringSubmatch(result,-1)
		for _, n := range res{
			orgs = append(orgs, strings.Trim(n[1], "\r\n"))
		}
		data[domain] = orgs
		jsondata, err := json.Marshal(data)
		if err != nil {
			log.Println("Error in whois lookup : ", err)
		}
		fmt.Println(string(jsondata))
		
	}
}