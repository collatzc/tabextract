# Tabextract
> Extracting data from csv, xls, xlsx files with Go fast and stable

[![Telegram](https://img.shields.io/badge/chat-telegram-blue.svg)](https://t.me/ohmyladygaga)

## Features

Tabextract is a tool that can extract data from xls and xlsx files in a fast and stable way. With pure golang to read Excel files into mem, which is a tabular data structure that can be manipulated and analyzed easily. Tabextract is useful for data mining and processing tasks that involve Excel files.

## Usage

Tabextract provides a simple standard interface for all supported filetypes, allowing access to both named worksheets and single tables in plain text formats.

```go
package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/collatzc/tabextract"
    _ "github.com/collatzc/tabextract/simple"
    _ "github.com/collatzc/tabextract/xls"
    _ "github.com/collatzc/tabextract/xlsx"
)

func main() {
    wb, _ := tabextract.Open(os.Args[1])
    sheets, _ := wb.List()
    for _, s := range sheets {
        sheet, _ := wb.Get(s)
        for sheet.Next() {
            row := sheet.Strings()
            fmt.Println(strings.Join(row, "\t"))
            if row[0] == "" {
                break
            }
        }
    }
    wb.Close()
}
```

## License

All source code is licensed under the [MIT License](https://raw.github.com/collatzc/tabextract/master/LICENSE).
