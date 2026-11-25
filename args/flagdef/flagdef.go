package flagdef

import "flag"

var (
	PrintRemotes = flag.Bool("print-remotes", false, "Prints all the remotes and their files")
)
