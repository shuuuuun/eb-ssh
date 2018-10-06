package main

import (
    "os"
    "fmt"
    "github.com/urfave/cli"
    "github.com/aws/aws-sdk-go/aws"
    // "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/elasticbeanstalk"
    // "github.com/aws/aws-sdk-go/service/ec2"
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
    eb_env_params := elasticbeanstalk.DescribeEnvironmentResourcesInput{
      EnvironmentName: &env_name,
    }
    fmt.Println(eb_client.DescribeEnvironmentResources(&eb_env_params))

    // beanstalkのリソース情報からインスタンスIDを取得
    // environment_resources=$(aws elasticbeanstalk describe-environment-resources --environment-name $env_name --region $region --profile $profile)
    // instance_ids=($(echo $environment_resources | jq -r '.EnvironmentResources.Instances[].Id' | tr "\n" " "))

    // インスタンスのIPアドレスを取得
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
