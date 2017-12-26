// This package calls the GoDaddy API with a key and secret key for a specified (already existing domain).

package godaddy

import (
         "io/ioutil"
         "net/http"
         "fmt"
         "os"
)

// Get domain details.  Return json results.
func GetDomains() string {

    req, err := http.NewRequest("GET", os.ExpandEnv("https://api.godaddy.com/v1/domains/${DOMAIN}"), nil)

    if err != nil {
        // handle err
        fmt.Printf("%s", err)
    }

    req.Header.Set("Authorization", os.ExpandEnv("sso-key ${API_KEY}:${API_SECRET}"))


    resp, err := http.DefaultClient.Do(req)
    bodyBytes, err2 := ioutil.ReadAll(resp.Body)

    if err2 != nil {
        fmt.Printf("%s", err2)
    }

    bodyString := string(bodyBytes)

    //fmt.Printf("%s\n", string(bodyString))

    if err != nil {
        // handle err
        fmt.Printf("%s", err)
    }

    defer resp.Body.Close()

    return bodyString

}


// Get all records of type: A, AAAA, CNAME, TXT, and NS.
// return array of type string results.
func GetRecords() [5]string {

    recordArr := [5]string{"A", "AAAA", "CNAME", "TXT", "NS"}

    var resultArr [5]string

    for i := 0; i <= len(recordArr)-1; i++ {

      req, err := http.NewRequest("GET", os.ExpandEnv("https://api.godaddy.com/v1/domains/${DOMAIN}/records/" + recordArr[i]), nil)

      if err != nil {
          // handle err
          fmt.Printf("%s", err)
      }

      req.Header.Set("Authorization", os.ExpandEnv("sso-key ${API_KEY}:${API_SECRET}"))
      resp, err := http.DefaultClient.Do(req)

      bodyBytes, err2 := ioutil.ReadAll(resp.Body)

      defer resp.Body.Close()

      if err2 != nil {
          fmt.Printf("%s", err2)
      }

      resultArr[i] = string(bodyBytes)

      if err != nil {
          // handle err
          fmt.Printf("%s", err)
      }

   }
   
   return resultArr
}


