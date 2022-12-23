package ha_civo

import (
	"fmt"
	"log"
	"time"

	"github.com/civo/civogo"
)

func scriptWP(privateIPlb, token string) string {
	return fmt.Sprintf(`#!/bin/bash
export SECRET='%s'
curl -sfL https://get.k3s.io | sh -s - agent --token=$SECRET --server https://%s:6443
`, token, privateIPlb)
}

func (obj *HAType) CreateWorkerNode(number int, privateIPlb, token string) (*civogo.Instance, error) {

	name := fmt.Sprintf("%s-ksctl-wp", obj.ClusterName)

	if len(obj.WPFirewallID) == 0 {
		firewall, err := obj.CreateFirewall(name)
		if err != nil {
			return nil, err
		}
		obj.WPFirewallID = firewall.ID
		err = obj.Configuration.ConfigWriterFirewallWorkerNodes(firewall.ID)
		if err != nil {
			return nil, nil
		}
	}
	name += fmt.Sprint(number)

	instance, err := obj.CreateInstance(name, obj.WPFirewallID, obj.NodeSize, scriptWP(privateIPlb, token), false)
	if err != nil {
		return nil, err
	}

	err = obj.Configuration.ConfigWriterInstanceWorkerNodes(instance.ID)
	if err != nil {
		return nil, nil
	}

	var retObject *civogo.Instance

	for {
		getInstance, err := obj.GetInstance(instance.ID)
		if err != nil {
			return nil, err
		}

		if getInstance.Status == "ACTIVE" {
			retObject = getInstance
			log.Println("✅ 🚀 Instance " + name)
			// err := ExecWithoutOutput(getInstance.PublicIP, getInstance.InitialPassword, scriptWP(privateIPlb, token), false)
			// if err != nil {
			// 	return nil, err
			// }
			// log.Println("✅ 🔧 Instance" + name)
			return retObject, nil
		}
		log.Println("🚧 Instance " + name)
		time.Sleep(10 * time.Second)
	}

}
