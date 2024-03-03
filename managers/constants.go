package managers

import (
	"fmt"
	"sort"
)

const (
	NS_CLONER_ANNOTATION              = "cloner.io/enabled"
	TARGET_NS_ANNOTATION              = "cloner.io/source-namespace"
	TARGET_NS_ANNOTATION_ENABLED      = "cloner.io/cloned"
	TARGET_CM_ANNOTATION              = "cloner.io/source-configmap"
	TARGET_SECRET_ANNOTATION          = "cloner.io/source-secret"
	TARGET_DEPLOYMENT_ANNOTATION      = "cloner.io/source-deployment"
	TARGET_JOB_ANNOTATION             = "cloner.io/source-job"
	TARGET_CRONJOB_ANNOTATION         = "cloner.io/source-cronjob"
	TARGET_SERVICE_ANNOTATION         = "cloner.io/source-service"
	TARGET_INGRESS_ANNOTATION         = "cloner.io/source-ingress"
	TARGET_SA_ANNOTATION              = "cloner.io/source-serviceaccount"
	TARGET_VIRTUAL_SERVICE_ANNOTATION = "cloner.io/source-virtualservice"
	// Kube-green specifics (Reference: https://kube-green.dev/docs/apireference_v1alpha1/)
	KUBE_GREEN_SLEEPAT_ANNOTATION = "sleep-info.kube-green.com/sleep-time"
	KUBE_GREEN_WAKEAT_ANNOTATION  = "sleep-info.kube-green.com/wake-up-time"
	KUBE_GREEN_TZ_ANNOTATION      = "sleep-info.kube-green.com/timezone"
	KUBE_GREEN_TIMEZONE           = "Asia/Kolkata"
	// Complicated Cron - Setup everyday at 11pm and Wake up at 7 am
	KUBE_GREEN_SLEEP_TIME  = "*:0/23"
	KUBE_GREEN_WAKE_TIME   = "*:0/7"
	KUBE_GREEN_API_VERSION = "kube-green.com/v1alpha1"
	KUBE_GREEN_KIND        = "SleepInfo"
	KUBE_GREEN_SAMPLE_YAML = `
	apiVersion: kube-green.com/v1alpha1
	kind: SleepInfo
	metadata:
	  name: working-hours
	spec:
	  weekdays: "1-5"
	  sleepAt: "20:00"
	  wakeUpAt: "08:00"
	  timeZone: "Europe/Rome
	`
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
