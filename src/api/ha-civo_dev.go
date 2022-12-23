//
// ######################################
//		EARLY DEVELOPMENT [STUB AND DRIVER CONFIG]
//   NOTE: THIS FILE WILL BE REMOVED IN NEXT VERSION (remane it to main.go_template
// ######################################
//

package main

import (
	"fmt"
	"log"

	"github.com/kubesimplify/ksctl/src/api/ha_civo"
)

const (
	REGION      = "NYC1"
	CLUSTERNAME = "demo"
)

func main() {

	fmt.Println("Enter 1 to create, 2 for add worker nodes and 3 for delete worker nodes, 0 to delete: ")
	choice := -1
	fmt.Scanf("%d", &choice)
	var err error
	switch choice {
	case 0:
		err = ha_civo.DeleteCluster(CLUSTERNAME, REGION, true)
	case 1:
		// controlplane and workernode nodeSize
		err = ha_civo.CreateCluster(CLUSTERNAME, REGION, "g3.small", 3, 1)
	case 2:
		err = ha_civo.AddMoreWorkerNodes(CLUSTERNAME, REGION, "g3.small", 3)
	case 3:
		err = ha_civo.DeleteSomeWorkerNodes(CLUSTERNAME, REGION, 1)
	}
	if err != nil {
		log.Panicln(err)
	}
}
