package javatools

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Get a XML document from the specified url, parses the XML-encoded data and
// stores the result in the value pointed to by v, which must be an arbitrary
// struct, slice, or string. Well-formed data that does not fit into v is
// discarded.
func getXml(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// Get the latest Java version defined in a XML-encoded document from the
// specified url.
func GetLatestVersion(url string) (string, error) {
	jum := javaUpdateMap{}
	err := getXml(url, &jum)

	if err != nil {
		return "", err
	}

	ju := javaUpdate{}
	err = getXml(jum.Mappings[0].Url, &ju)

	if err != nil {
		return "", err
	}

	return ju.Informations[0].Version, nil
}

func ParseJavaVersion(version string) (JavaVersion, error) {
	re, err := regexp.Compile(`^(\d+).(\d+).(\d+)_(\d+)-b(\d+)$`)
	if err != nil {
		return JavaVersion{}, err
	}
	subs := re.FindAllSubmatch([]byte(version), -1)
	match := subs[0]
	major, _ := strconv.Atoi(string(match[1]))
	minor, _ := strconv.Atoi(string(match[2]))
	patch, _ := strconv.Atoi(string(match[3]))
	update, _ := strconv.Atoi(string(match[4]))
	build, _ := strconv.Atoi(string(match[5]))
	return JavaVersion{major, minor, patch, update, build}, nil

}

func DownloadJava(version JavaVersion) {
	url := GetJavaUrl(version)
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "oraclelicense=accept-securebackup-cookie")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func GetJavaUrl(version JavaVersion) string {
	jdk := strconv.Itoa(version.Minor) + "u" + strconv.Itoa(version.Update)
	jdklong := jdk + "-b" + strconv.Itoa(version.Build)
	url := "http://download.oracle.com/otn-pub/java/jdk/" + jdklong + "/jdk-" + jdk + "-linux-x64.rpm"
	return url
}
