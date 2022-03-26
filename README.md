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
    $ ./kinectl --cacert=/root/ca.pem --cert=/root/client.pem --key=/root/client-key.pem --endpoints="192.168.154.90:2379" version
    {"etcdserver":"3.5.0","etcdcluster":"3.5.0"}
    $
    ```

  * create
    
    add a key-value store
    ```shell
    $ ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" create --key="/test" --value="hello world"
    success
    $
    ```

  * list

    list datas with key="/..."
    ```shell
    $ ./kinectl --cacert=/root/ca.pem --cert=/root/client.pem --key=/root/client-key.pem --endpoints="192.168.154.90:2379" list --key="/"
    key: /registry/health
    data: {"health":"true"}
    modified times: 2

    key: /test
    data: hello world
    modified times: 11

    $
    ```
  
  * get

    get value by key
    ```shell
    $ ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" get --key="/test"
    key: /test
    data: new value
    modified times: 11
    $
    ```

  * put

    modified value
    ```shell
    $ ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" put --key="/test" --value="new value"
    success
    $ ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" get --key="/test"
    key: /test
    data: new value
    modified times: 12
    $
    ```
  * upgrade

    ```shell
    $ ./kinectl --cacert=ca.pem --cert=client.pem --key=client-key.pem --endpoints="192.168.154.90:2379" put --key="/test" --value="new value update"
    ```