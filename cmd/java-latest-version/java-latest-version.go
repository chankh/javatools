package main

import (
	"fmt"
	"github.com/chankh/javatools"
	"os"
)

const javaUpdateXml string = "https://javadl-esd-secure.oracle.com/update/1.8.0/map-m-1.8.0.xml"

func main() {

	version, err := javatools.GetLatestVersion(javaUpdateXml)
	if err != nil {
		fmt.Printf("Error getting latest version from %s,: %v\n", javaUpdateXml, err)
		os.Exit(1)
	}
	fmt.Println(version)
}
