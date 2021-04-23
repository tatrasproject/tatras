package argo

import (
    "io/ioutil"
    "log"
    "net/http"
	"os"
    "fmt"
    "bytes"
    "encoding/json"
)

// ARGOCD_SERVER <format: argocd.mycompany.com>
// ARGOCD_TOKEN can be retrieved via CLI:
//   curl $ARGOCD_SERVER/api/v1/session -d $'{"username":"admin","password":"password"}'

var ARGOCD_SERVER = "https://" + os.Getenv("ARGOCD_SERVER")
var ARGOCD_TOKEN  = os.Getenv("ARGOCD_TOKEN")

// Helper function. Creates http.Request with the auth token
// Also takes in data(string) to be passed in POST request.
// TODO: Maybe pass data in as byte array instead?
func auth_request(method string, url string, data string) *http.Request {   
    var jsonStr = []byte(data)
    req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }

    // Create a Bearer string by appending ARGOCD_TOKEN
    var Bearer = "Bearer " + ARGOCD_TOKEN
    req.Header.Add("Authorization", Bearer)

    return req
}

// Reads in json file, updates it with necessary params, uses
// it for the body of a POST request.
// TODO: Estimated odds are 7.854/10 that this should be a 
//       struct and not a json file. Thinking was that the
//       json body can be very very large and this could save
//       a lot of time building structs.
func format_data(file_name string, app_name string) string {
    jsonFile, err := os.Open("clients/argo/json_templates/" + file_name)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    defer jsonFile.Close()

    // Parse jsonFile into a map
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)
    log.Println(result)

    // TODO: Figure out how to turn map into string (or byte array). 
    //       Ultimately used as data var (http body) in auth_request
    return ""
}

// Helper function to send auth_request using http Client. 
// Handles errors and returns response as string
func send(req *http.Request) string {
    client := &http.Client{}
    
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }
    // log.Println(string([]byte(body)))
    return string([]byte(body))
}

// Called from handlers package
func GetApplication(app_name string) string {
    method := "GET"
    url := ARGOCD_SERVER + "/api/v1/applications/" + app_name
    req := auth_request(method, url, "")
    return send(req)
}

// Called from handlers package
func CreateApplication(app_name string) string {
    method := "POST"
    url := ARGOCD_SERVER + "/api/v1/applications/" + app_name
    data := format_data("create_app.json", app_name)
    req := auth_request(method, url, data)
    return send(req)

}