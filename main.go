package main

import (
	"fmt"
	"github.com/rs/xid"
)

func main() {
	generatedXID := xid.New()
	xidString := generatedXID.String()
	fmt.Print(xidString)
}
