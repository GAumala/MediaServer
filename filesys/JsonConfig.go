package filesys


import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "sort"

)

func ParseJsonPathList(configPath string) ([]string, error) {
    b, err := ioutil.ReadFile(configPath)
    if(err != nil) {
        return nil, err
    }
    var slice []string
    json.Unmarshal(b, &slice)
    sort.Strings(slice)
    fmt.Printf("got %s\n", slice)

    return slice, nil
}
