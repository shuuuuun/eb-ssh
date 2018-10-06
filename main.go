// https://docs.aws.amazon.com/sdk-for-go/api/service/elasticbeanstalk/
// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/

package main

import (
    "os"
    "fmt"
    "github.com/urfave/cli"
    "github.com/aws/aws-sdk-go/aws"
    // "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/elasticbeanstalk"
    "github.com/aws/aws-sdk-go/service/ec2"
    // "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
  os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

  app := cli.NewApp()

  app.Name = "eb-ssh"
  app.Usage = "This app echo input arguments"
  app.Version = "0.0.1"

  app.Action = func (context *cli.Context) error {
    profile := context.String("profile")
    region := context.String("region")
    env_name := context.String("env-name")
    keyfile := context.String("keyfile")
    proxy_host := context.String("proxy-host")

    // fmt.Println(context.Args().Get(0))
    fmt.Println("profile: " + profile)
    fmt.Println("region: " + region)
    fmt.Println("env_name: " + env_name)
    fmt.Println("keyfile: " + keyfile)
    fmt.Println("proxy_host: " + proxy_host)

    sess := session.Must(session.NewSessionWithOptions(session.Options{Profile:profile}))
    eb_client := elasticbeanstalk.New(
      sess,
      aws.NewConfig().WithRegion(region),
    )
    // fmt.Println(eb_client.DescribeEnvironments(nil))
    ec2_client := ec2.New(
      sess,
      aws.NewConfig().WithRegion(region),
    )

    // beanstalkのリソース情報からインスタンスIDを取得
    eb_env_params := elasticbeanstalk.DescribeEnvironmentResourcesInput{
      EnvironmentName: &env_name,
    }
    resources, err := eb_client.DescribeEnvironmentResources(&eb_env_params)
    if err != nil {
      panic(err)
    }
    // fmt.Println(resources.EnvironmentResources.Instances)
    var instance_ids []*string
    for _, v := range resources.EnvironmentResources.Instances {
      instance_ids = append(instance_ids, v.Id)
    }
    // fmt.Println(instance_ids)

    // インスタンスのIPアドレスを取得
    ec2_instances_params := ec2.DescribeInstancesInput{
      InstanceIds: instance_ids,
    }
    instances, err := ec2_client.DescribeInstances(&ec2_instances_params)
    if err != nil {
      panic(err)
    }
    var instance_ips []*string
    for _, v1 := range instances.Reservations {
      for _, v2 := range v1.Instances {
        // fmt.Println(*v2.PublicIpAddress)
        instance_ips = append(instance_ips, v2.PublicIpAddress)
      }
    }
    fmt.Println(instance_ips)

    // どのIPに接続するか選択
    // 接続

    return nil
  }

  app.Flags = []cli.Flag{
    cli.StringFlag {
      Name: "env-name, e",
      Usage: "",
    },
    cli.StringFlag {
      Name: "region, r",
      Usage: "",
    },
    cli.StringFlag {
      Name: "profile, p",
      Usage: "",
    },
    cli.StringFlag {
      Name: "keyfile, k",
      Usage: "",
    },
    cli.StringFlag {
      Name: "proxy-host, P",
      Usage: "",
    },
  }

  app.Run(os.Args)
}
