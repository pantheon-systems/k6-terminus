package terminus

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
	"go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/terminus", new(Terminus))
}

// Compare is the type for our custom API.
type Terminus struct {
	TerminusResult string // textual description of the most recent comparison
}

func (t *Terminus) RunTerminus(command string) string {

	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.CombinedOutput()
	output := string(stdout)

	if err != nil {
		log.Error(output)
	}
	// log.Info("Command: "+command, " result: "+output)

	t.TerminusResult = output
	return t.TerminusResult
}
