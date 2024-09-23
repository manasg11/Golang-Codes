package main

import (
	"fmt"
)

func main() {
	environmentList := []string{"dev", "qc", "uat", "prod"}
	instanceTypes := []string{"t3.micro", "t3.small", "t3.medium", "t3.large"}

	if len(environmentList) != len(instanceTypes) {
		fmt.Println("The length of the arrays must match.")
		return
	}

	envInstanceMap := make(map[string]string)
	for index := 0; index < len(environmentList); index++ {
		envInstanceMap[environmentList[index]] = instanceTypes[index]
	}

	fmt.Println("Environment to Instance Mapping:")
	for env, instance := range envInstanceMap {
		fmt.Printf("%s -> %s\n", env, instance)
	}
}
