//go:generate goversioninfo -manifest=testdata/resource/goversioninfo.exe.manifest
package main    
import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pedroalbanese/xoodoo/xoodyak"
)
        var key = flag.String("key", "", "HMAC secret key.")

func main() {
    flag.Parse()

        if (len(os.Args) < 2) {
	fmt.Println("Xoodyak MAC - ALBANESE Lab (c) 2020-2023\n")
	fmt.Println("Usage of",os.Args[0]+":")
        flag.PrintDefaults()
        os.Exit(1)
        } 

	var err error
	h := xoodyak.NewXoodyakMac([]byte(*key))
	if _, err = io.Copy(h, os.Stdin); err != nil {
                log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
        os.Exit(0)

}
