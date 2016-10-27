// Dumps a bunch of aws CLI commands to sync English language courses from S3.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	// S3 Bucket containing course metadata.
	S3_BUCKET_MOOCFETCHER = "moocfetcher"

	// S3 Bucket containing archived courses.
	S3_BUCKET_MOOCFETCHER_COURSE_ARCHIVE = "moocfetcher-course-archive"

	// S3 Key for file containing metadata for launched on-demand courses.
	CACHED_ONDEMAND_LAUNCHED_KEY = "coursera/ondemand/launched.json"
)

// Course data
type courses struct {
	Courses []struct {
		Slug      string   `json:"slug"`
		Languages []string `json:"primaryLanguageCodes"`
	} `json:"courses"`
}

func main() {
	// Parse flags
	path := flag.String("path", ".", "Filesystem path to sync courses at")
	script := flag.String("script", "sync.bat", "Check data on server. Don’t download anything")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `Sync English language course materials from MOOCFetcher server with local filesystem.`)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	svc := s3.New(session.New(aws.NewConfig().WithRegion("us-east-1")))

	// Retrieve list of courses.
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(S3_BUCKET_MOOCFETCHER),
		Key:    aws.String(CACHED_ONDEMAND_LAUNCHED_KEY),
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%d bytes read\n", len(body))
	var courses courses
	err = json.Unmarshal(body, &courses)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%d courses found\n", len(courses.Courses))

	// Run JSON filter to get to only english courses, and sort them.
	var en []string
	for _, course := range courses.Courses {
		sort.Strings(course.Languages)
		i := sort.SearchStrings(course.Languages, "en")
		if i < len(course.Languages) && course.Languages[i] == "en" {
			en = append(en, course.Slug)
		}
	}
	fmt.Printf("%d English language courses found\n", len(en))

	sort.Strings(en)

	// Generate aws sync command for each english course.
	var commands []string
	for _, slug := range en {
		cmd := "aws"
		args := []string{"s3", "sync", fmt.Sprintf("s3://%s/coursera/%s/", S3_BUCKET_MOOCFETCHER_COURSE_ARCHIVE, slug), fmt.Sprintf("%s%c%s", *path, os.PathSeparator, slug)}
		commands = append(commands, fmt.Sprintf("%s %s", cmd, strings.Join(args, " ")))
	}

	eol := "\n"
	if strings.HasSuffix(*script, ".bat") {
		eol = "\r\n"
	}
	err = ioutil.WriteFile(*script, []byte(strings.Join(commands, eol)), 0744)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
