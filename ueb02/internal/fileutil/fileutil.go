package fileutil

import (
    "encoding/csv"
    "github.com/akelsch/vaa/ueb02/internal/errutil"
    "io/ioutil"
    "os"
    "path/filepath"
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
    absPath, err := filepath.Abs(filename)
    errutil.HandleError(err)
    return absPath
}
