// Package cmd implements the commands supported by opfcli.
package cmd

import log "github.com/sirupsen/logrus"

const defaultAppName = "cluster-scope"
const namespacePath = "base/core/namespaces"
const groupPath = "base/user.openshift.io/groups"
const componentPath = "components"
const componentRelPath = "../../../../components"

func recoverAndLog() {
	if err := recover(); err != nil {
		err := err.(error)
		log.Fatalf("%s", err.Error())
	}
}
