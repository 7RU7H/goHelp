
// CLI: ./usefulGolang cmd ...
aCommand := flag.NewFlagSet("cmd", flag.ExitOnError)
// CLI ./usefulGolang cmd -a <string arg here>
aCommand.StringVar(&argument, "a", "default value", "An description")