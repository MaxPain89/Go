package main

import (
    "github.com/MaxPain89/Go/aws/ec2/instances"
    "fmt"
    "os"
    "log"
)

const (region = "eu-central-1"
       ami = "ami-7c412f13"
       key = "mk_aws_pair"
       instanceType = "t2.micro")

func main()  {
    instanceId, err := instances.CreateInstance(region, ami, instanceType, key, true)
    if err != nil {
       os.Exit(1)
    }
    fmt.Printf("Instance %s created\n", instanceId)

    //instanceId, err := instances.TerminateInstance("eu-central-1","i-0c9a7baa636916160", true)
    //if err != nil {
    //    os.Exit(1)
    //}
    //fmt.Printf("Instance %s terminated\n", instanceId)

    info, err := instances.DescribeInstance(region, instanceId)
    if err != nil {
        os.Exit(1)
    }
    log.Printf("%v", info)
}
