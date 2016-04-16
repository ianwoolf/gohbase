# gohbase
connect and send conmand to hbase by thrift2

Looking forward your pr or issue

thrift file is build at hbase v1.2.0

get value example:

    package main

    import (
	    "fmt"

        h "github.com/ianwoolf/gohbase"
    )

    func main() {
    	hbObj := h.NewHbObj("192.168.99.100", 9090)
	    if err := hbObj.Connect(); err != nil {
		    fmt.Println(err.Error())
    	}
	    defer hbObj.Close()
        // get 'test','row1',{COLUMN => ['c2:a']}
    	TRow, _ := hbObj.GetRow("test", h.GenTGet("row1", "col", "a"))
    	for _, col := range TRow.ColumnValues {
    		fmt.Println(string(col.GetFamily()), string(col.GetQualifier()), string(col.GetValue()), col.GetTags(), col.GetTimestamp())
	    }
    }
