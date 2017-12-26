// This package calls the GoDaddy API with a key and secret key for a specified (already existing domain).

package godaddy

import (
         "io/ioutil"
         "net/http"
         "fmt"
         "os"
)

type api struct {
    key, secret string
}


// Get domain details.  Return json.
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

func GetRecords() string {

    recordArr := [6]string{"A", "AAAA", "CNAME", "TXT", "NS"}

    bodyString := ""

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

      bodyString += string(bodyBytes)

      //fmt.Printf("%s\n", bodyString)

      if err != nil {
          // handle err
          fmt.Printf("%s", err)
      }

   }
   
   return bodyString
}


