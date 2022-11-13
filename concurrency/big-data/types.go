package main

type processed struct {
	fullNames  []string
	firstNames []string
	months     []string
	numRows    int
}

type result struct {
	numRows           int
	peopleCount       int
	commonName        string
	commonNameCount   int
	donationMonthFreq map[string]int
}
