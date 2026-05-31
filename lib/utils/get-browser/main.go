// Package main ...
package main

import (
	"fmt"

	"github.com/LQR471814/rod/lib/launcher"
	"github.com/LQR471814/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
