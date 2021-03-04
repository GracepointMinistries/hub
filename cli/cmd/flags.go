package cmd

var (
	cfgFile string
	host    string
	all     bool
	user    int64
	zgroup  int64
	filter  string

	// this is set at compile time
	defaultHost string
)
