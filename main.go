// https://docs.aws.amazon.com/sdk-for-go/api/service/elasticbeanstalk/
// https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/
// https://godoc.org/golang.org/x/crypto/ssh

// TODO: .elasticbeanstalkを読み込む
// TODO: proxy_hostはオプションにする
// TODO: 関数化
// TODO: Flagsを上に
// TODO: エラー処理
// TODO: キャメルケースにする
// TODO: deps

package main

import (
    "os"
    "os/exec"
    "fmt"
    "strconv"
    "strings"
    "bufio"
    // "io/ioutil"
    // "golang.org/x/crypto/ssh"
    "github.com/urfave/cli"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/elasticbeanstalk"
    "github.com/aws/aws-sdk-go/service/ec2"
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
    fmt.Println("----- arguments -----")
    fmt.Println("profile: " + profile)
    fmt.Println("region: " + region)
    fmt.Println("env_name: " + env_name)
    fmt.Println("keyfile: " + keyfile)
    fmt.Println("proxy_host: " + proxy_host)
    fmt.Println("")

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
        instance_ips = append(instance_ips, v2.PublicIpAddress)
      }
    }
    // fmt.Println(instance_ips)

    // どのIPに接続するか選択
    var idx_list []string
    for idx, ip := range instance_ips {
      idx_list = append(idx_list, strconv.Itoa(idx))
      fmt.Println(strconv.Itoa(idx) + ": " + *ip)
    }
    fmt.Print("which instance? (" + strings.Join(idx_list, "/") + ") ")

    var selected_idx int
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
      input := scanner.Text()
      // fmt.Println(input)
      if input == "" {
        selected_idx = 0
        break
      }
      input_idx, err := strconv.Atoi(input)
      if err != nil {
        fmt.Println("Invalid Input!")
        // panic(err)
      }
      // TODO: 有効なidxかチェック
      fmt.Println(input_idx)
    }
    if err := scanner.Err(); err != nil {
      panic(err)
    }

    ip := instance_ips[selected_idx]
    fmt.Println(*ip)

    // 接続
    // ssh -o ProxyCommand="ssh -W %h:%p $ext_gate_host" -i "$keyfile" ec2-user@$ip
    // out, err := exec.Command("ssh", "-o  ProxyCommand='ssh -W %h:%p " + proxy_host + "'", "-i " + keyfile, "ec2-user@" + *ip).Output()
    // out, err := exec.Command("ssh", "-i " + keyfile, "ec2-user@" + *ip).Output()
    // if err != nil {
    //   panic(err)
    // }
    // fmt.Println(string(out))
    cmd := exec.Command("ssh", "-i " + keyfile, "ec2-user@" + *ip)
    cmd.Start()
    cmd.Wait()

    // key, err := ioutil.ReadFile(keyfile)
    // if err != nil {
    //   panic(err)
    // }
    // signer, err := ssh.ParsePrivateKey(key)
    // if err != nil {
    //   panic(err)
    // }
    // auth := []ssh.AuthMethod{ssh.PublicKeys(signer)}
    // sshConfig := &ssh.ClientConfig{
    //   User: "ec2-user",
    //   Auth: auth,
    //   HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    // }
    // sshClient, err := ssh.Dial("tcp", *ip + ":22", sshConfig)
    // if err != nil {
    //   panic(err)
    // }
    // sshSession, err := sshClient.NewSession()
    // if err != nil {
    //   panic(err)
    // }
    // defer sshSession.Close()
    // 
    // out, err := sshSession.CombinedOutput("pwd")
    // if err != nil {
    //   panic(err)
    // }
    // fmt.Println(string(out))

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
