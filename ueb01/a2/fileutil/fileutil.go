package fileutil

import (
    "encoding/csv"
    "github.com/akelsch/vaa/ueb01/a2/errutil"
    "io/ioutil"
    "os"
    "path"
    "strings"
)

func ReadCsvRows(filename string) [][]string {
    file, err := os.Open(canonicalize(filename))
    errutil.HandleError(err)

    reader := csv.NewReader(file)
    rows, err := reader.ReadAll()
    errutil.HandleError(err)

    return rows
}

func ReadBytes(filename string) []byte {
    bytes, err := ioutil.ReadFile(canonicalize(filename))
    errutil.HandleError(err)
    return bytes
}

func canonicalize(filename string) string {
    if strings.HasPrefix(filename, ".") {
        pwd, _ := os.Getwd()
        return path.Join(pwd, filename)
    }

    return filename
}
