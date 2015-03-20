package main

import (
    "os"
    "log"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "bufio"
    "flag"
    "net/http"
)

type ReleaseState struct {
    Status  string `json:"status"`
    Urgency string `json:"urgency"`
    Version string `json:"version"`
}

type Package struct {
    Debianbug   int `json:"debianbug"`
    Description string  `json:"description"`
    Issue       string  `json:"issue"`
    Releases    map[string]ReleaseState `json:"releases"`
    Repositories map[string]string `json:"repositories"`
    Scope string `json:"scope"`
}

func trackerUrl(cve string) string {
    return fmt.Sprintf("https://security-tracker.debian.org/tracker/redirect/%s", cve)
}

func main() {
    distro := flag.String("distro", "", "Distribution to check against")
    verbose := flag.Bool("verbose", false, "Be more verbose")
    flag.Parse()
    cves := make(map[string]bool)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        cves[scanner.Text()] = true
    }
    _,err := os.Stat("json")
    if err != nil {
        log.Printf("No JSON data found, downloading it from security-tracker.debian.org")
        resp,err := http.Get("https://security-tracker.debian.org/tracker/data/json")
        if err != nil {
            log.Fatalf("Failed to download sec tracker data: %s", err)
        }
        buf,err := ioutil.ReadAll(resp.Body)
        ioutil.WriteFile("json", buf, 0440)
        log.Printf("Done.")
    }
    js,err := ioutil.ReadFile("json")
    if err != nil {
        log.Fatalf("Failed to open json")
    }
    packages := make(map[string][]Package)
    if err != nil {
        log.Fatalf("Failed to read json")
    }
    err = json.Unmarshal(js,&packages)
    if err != nil {
        log.Fatalln(err)
    }
    for pkg,bugs := range packages {
        for _,bug := range bugs {
            if cves[bug.Issue] {
                // log.Printf("Package %s is affected by %s %+v", pkg, bug.Issue, bug)
                if len(*distro) > 0 {
                    if dstat,ok := bug.Releases[*distro]; ok {
                        if dstat.Status != "resolved" {
                            log.Printf("WARNING: %s is open on %s for pkg %s: %s", bug.Issue, *distro, pkg, trackerUrl(bug.Issue))
                        } else {
                            if *verbose {
                                log.Printf("%s is %s for %s: %s", pkg, dstat.Status, *distro, trackerUrl(bug.Issue))
                            }
                        }
                    }
                }
            }
        }
    }
}