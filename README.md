# kinectl
a simple commond-line tool to use [Kine](https://github.com/k3s-io/kine).

## How to build
depends:
* golang environment

``` shell
git clone https://github.com/litekube/kinectl.git
cd kinectl/cmd/kinectl

go build .

# run kinectl
./kinectl --help
```

## Usage

assumption:

```yaml
CA Cert: /root/ca.pem
Client Cert: /root/client.pem
Client private-key: /root/client-key.pem
server address: 192.168.154.90:2379
```
* kinectl global
  * command
    ```shell
    # get version
    ./kinectl --version

    # get help
    ./kinectl --help
    ```
  * kine info
    ```shell
    ./kinectl --cacert=/root/ca.pem --cert=/root/client.pem --key=/root/client-key.pem --endpoints="192.168.154.90:2379" [command] [args]

    # or you can use Environment variates
    export KINECTL_CACERT="/root/ca.pem"
    export KINECTL_CERT="/root/client.pem"
    export KINECTL_KEY="/root/client-key.pem"
    export KINECTL_LEADELECT="192.168.154.90:2379"
    ./kinectl [command] [args]

    # it is ok to mix these ways, but Command line arguments will be given higher priority.
    ```
* kine options
  > The following parameters provide shorthand:
  > * --key="key" or -k "key"
  > * --value="value" or -v "value"
  > * --revision="revision" or -r "revision"

  * version
    ```shell
    ./kinectl --cacert=/root/ca.pem --cert=/root/client.pem --key=/root/client-key.pem --endpoints="192.168.154.90:2379" version
    ```

  * create
    
    add a key-value store
    ```shell
    ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" create --key="/test" --value="hello world"
    ```

  * list

    list datas with key="/..."
    ```shell
    ./kinectl --cacert=/root/ca.pem --cert=/root/client.pem --key=/root/client-key.pem --endpoints="192.168.154.90:2379" list --key="/"
    ```
  
  * get

    get value by key
    ```shell
    ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" get --key="/test"
    ```

  * put

    modified value
    ```shell
    ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" put --key="/test" --value="new value"
    success

    # if key="/test" is not exist, it will give an error. it can automaticly change to run create by "-f" or "--force=true"
    ```
  * update

    ```shell
    # assume current revision(modified times) of key="/test" is 18, global is 25
    ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" update --key="/test" --value="new value update" --revision=18

    # after this command, it will be 26
    ```

  * delete

    ```shell
    ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" delete --key="/test"
    ```

## About Us
   We are not developers of the [Kine](https://github.com/k3s-io/kine) project. We just notice that `etcdctl` can not be used directly to operate `kine` or it would raise an error. Curl is also not a valid tool at this time(03/26/2022) because it only provides good support for GRPC.