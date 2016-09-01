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
		fmt.Printf("Unable to get latest version from %s", javaUpdateXml)
	}
	jver, err := javatools.ParseJavaVersion(version)
	if err != nil {
		fmt.Printf("Unable to parse Java version %s, %v", version, err)
		os.Exit(1)
	}

	fmt.Printf(javatools.GetJavaUrl(jver))
}
