/*
Copyright © 2022 M Arif S <arifsefrianto@gmail.com>

*/
package main

import "github.com/arifseft/go-command/cmd"
import "github.com/arifseft/go-command/data"

func main() {
    data.OpenDatabase()
	cmd.Execute()
}
