package managers

import (
	"fmt"
	"sort"
)

const (
	NS_CLONER_ANNOTATION         = "cloner.io/enabled"
	TARGET_NS_ANNOTATION         = "cloner.io/source-namespace"
	TARGET_NS_ANNOTATION_ENABLED = "cloner.io/cloned"
	TARGET_CM_ANNOTATION         = "cloner.io/source-configmap"
	TARGET_SECRET_ANNOTATION     = "cloner.io/source-secret"
	TARGET_DEPLOYMENT_ANNOTATION = "cloner.io/source-deployment"
	TARGET_JOB_ANNOTATION        = "cloner.io/source-job"
	TARGET_CRONJOB_ANNOTATION    = "cloner.io/source-cronjob"
	TARGET_SERVICE_ANNOTATION    = "cloner.io/source-service"
	TARGET_INGRESS_ANNOTATION    = "cloner.io/source-ingress"
	TARGET_SA_ANNOTATION         = "cloner.io/source-serviceaccount"
)

// Useless debugging function. Just a placeholder and not needed for actual work
func prettyPrint(m map[string]string) {
	// Get keys and sort them to ensure consistent order when printing
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Iterate over the sorted keys and print the key-value pairs
	fmt.Println("{")
	for _, k := range keys {
		fmt.Printf("  %s: %s,\n", k, m[k])
	}
	fmt.Println("}")
}
