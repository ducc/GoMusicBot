package cmd

import "bytes"

func write(buff *bytes.Buffer, str ...string) {
    for _, s := range str {
        buff.WriteString(s)
    }
}