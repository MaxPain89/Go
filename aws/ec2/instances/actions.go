package instances

import (
    "log"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
)

func CreateInstance(region, ami, instanceType, keyName string, waitForRunning bool) (instanceId string, err error){
    svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
    runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
        ImageId:      aws.String(ami),
        InstanceType: aws.String(instanceType),
        KeyName:      aws.String(keyName),
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
    })

    if err != nil {
        log.Println("Could not create instance", err)
        return "", err
    }

    resultInstanceId := *runResult.Instances[0].InstanceId
    log.Println("Created instance", resultInstanceId)

    if waitForRunning {
        if err := svc.WaitUntilInstanceRunning(&ec2.DescribeInstancesInput{
            InstanceIds: []*string{aws.String(resultInstanceId)},
        }); err != nil {
            log.Printf("Error while waiting. Err: %v", err)
            return "", err
        }
    }
    return resultInstanceId, nil
}

func TerminateInstance(region, instanceId string, waitForTerminating bool) (terminatedInstanceId string, err error){
    svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
    terminateResult, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
        InstanceIds:      aws.StringSlice([]string{instanceId,}),
    })

    if err != nil {
        log.Println("Could not create instance", err)
        return
    }

    resultInstanceId := *terminateResult.TerminatingInstances[0].InstanceId
    log.Println("Created instance", instanceId)
    if waitForTerminating {
        if err := svc.WaitUntilInstanceTerminated(&ec2.DescribeInstancesInput{
            InstanceIds: []*string{aws.String(resultInstanceId)},
        }); err != nil {
            log.Printf("Error while waiting. Err: %v", err)
            return "", err
        }
    }
    return resultInstanceId, nil
}

func DescribeInstance(region, instanceId string) (*ec2.MonitorInstancesOutput, error) {
    svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
    input := &ec2.MonitorInstancesInput{
        InstanceIds: aws.StringSlice([]string{instanceId,}),
    }
    result, err := svc.MonitorInstances(input)
    if err != nil {
        log.Printf("Error in monitoring instances. Err: %v", err)
        return nil, err
    }
    return result, nil
}