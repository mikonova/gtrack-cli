package flagdef

import "flag"

var (
	PrintRemotes        = flag.Bool("print-remotes", false, "Prints all the remotes and their files")
	PrintFilesForRemote = flag.String("lookup-local", "", "Prints the files for a specific remote (\"-lookup-local <address>\")")
	PrintLog            = flag.Bool("dump-log", false, "Dump logs to stdio")
	CheckConn           = flag.String("lookup-remote", "", "Check whether the connection succeeds and the name is correct "+
		"and reports it (\"-lookup-remote <address>\")")
)
