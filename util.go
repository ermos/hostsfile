package hostsfile

import (
	"os"
)

var readFile = os.ReadFile

var writeFile = os.WriteFile

var statFile = os.Stat
