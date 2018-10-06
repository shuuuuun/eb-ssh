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
  // var env_name string
  // var region string
  // var profile string
  // var keyfile string
  // var proxy_host string

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

    // var creds *credentials.Value
    // var creds *credentials.Credentials

    // fmt.Println(creds)

    // region := "ap-northeast-1"
    // conf := &aws.Config{
    //     Credentials: creds,
    //     Region:      &region,
    // }
    // sess := session.Must(session.NewSession())
    // eb := elasticbeanstalk.New(sess)
    // ec2 := ec2.New(sess)
    // s3 := s3.New(sess)

    // fmt.Println(conf)
    // fmt.Println(sess)
    // fmt.Println(eb)
    // fmt.Println(ec2)
    // fmt.Println(s3)

    sess := session.Must(session.NewSessionWithOptions(session.Options{Profile:profile}))
    eb := elasticbeanstalk.New(
      sess,
      aws.NewConfig().WithRegion(region),
    )
    fmt.Println(eb.DescribeEnvironments(nil))
    // fmt.Println(eb.DescribeEnvironmentResources(nil))

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
      // Destination: &env_name,
    },
    cli.StringFlag {
      Name: "region, r",
      Usage: "",
      // Destination: &region,
    },
    cli.StringFlag {
      Name: "profile, p",
      Usage: "",
      // Destination: &profile,
    },
    cli.StringFlag {
      Name: "keyfile, k",
      Usage: "",
      // Destination: &keyfile,
    },
    cli.StringFlag {
      Name: "proxy-host, P",
      Usage: "",
      // Destination: &proxy_host,
    },
  }

  app.Run(os.Args)
}

// // beanstalkのリソース情報からインスタンスIDを取得
// func get_instance_id() {}
// // インスタンスのIPアドレスを取得
// func get_instance_ip() {}
// // どのIPに接続するか選択
// func select_ip() {}
// func ssh() {}
