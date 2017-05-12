package main

import "fmt"
import "flag"
import "os"

var jenkins_url = flag.String("jenkins-url", "", "URL to the Jenkins root.")
var metrics_token = flag.String("metrics-token", "", "Secret metrics token to access Jenkins Metrics.")
var tag_set = flag.String("tag-set", "", "Influx line protocol tag set.")

func main() {
    flag.Parse()
    fmt.Printf("Hello, 世界\n")
    fmt.Println(os.Args)
}
