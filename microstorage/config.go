package microstorage

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

type tApp struct{
    StoragePath string
    Conversion string
}

type TCache struct {
    Engine string
    Addr string
    Pass string
    User string
    DB string
    TTL int64
}

type tNetwork struct {
    Port int64
}

type tListen struct{
    Filesystem bool
    Mysql bool
}

type tDelay struct {
    Filesystem int64
    Mysql int64
}
type tCredentials struct {
    Filesystem string
    Mysql string
}

var Config struct {
    App tApp
    Listen tListen
    Delay tDelay
    Credentials tCredentials
    Network tNetwork
    Cache TCache
}

func Init() {
    LogMessage("Loading ini file")
    conf, err := parseIni("./app.ini")
    if err != nil {
        ThrowException("Error reading ini file. Stopped")
    }
    //app
    Config.App.StoragePath = conf["app"]["storage_path"]
    Config.App.Conversion = conf["app"]["conversion"]

    //Network
    Config.Network.Port = StrToInt64(conf["network"]["port"])

    //Listeners on/off
    Config.Listen.Filesystem = conf["listen"]["filesystem"] == "on"
    Config.Listen.Mysql = conf["listen"]["mysql"] == "on"

    //Listeners settings
    Config.Delay.Filesystem = StrToInt64(conf["delay"]["filesystem"])
    Config.Delay.Mysql = StrToInt64(conf["delay"]["mysql"])

    //Credentials and path
    Config.Credentials.Filesystem = conf["credentials"]["filesystem"]
    Config.Credentials.Mysql = conf["credentials"]["mysql"]

    //Cache
    Config.Cache.Engine = conf["cache"]["engine"]
    Config.Cache.Addr = conf["cache"]["addr"]
    Config.Cache.Pass= conf["cache"]["pass"]
    Config.Cache.User = conf["cache"]["user"]
    Config.Cache.DB = conf["cache"]["db"]
    Config.Cache.TTL = StrToInt64(conf["cache"]["ttl"])
}

//got from https://code-maven.com/slides/golang/solution-parse-ini-file
func parseIni(filename string) (map[string]map[string]string, error) {
    ini := make(map[string]map[string]string)
    var head string

    fh, err := os.Open(filename)
    if err != nil {
        return ini, fmt.Errorf("Could not open file '%v': %v", filename, err)
    }
    sectionHead := regexp.MustCompile(`^\[([^]]*)\]\s*$`)
    keyValue := regexp.MustCompile(`^(\w*)\s*=\s*(.*?)\s*$`)
    reader := bufio.NewReader(fh)
    for {
        line, _ := reader.ReadString('\n')
        result := sectionHead.FindStringSubmatch(line)
        if len(result) > 0 {
            head = result[1]
            ini[head] = make(map[string]string)
            continue
        }

        result = keyValue.FindStringSubmatch(line)
        if len(result) > 0 {
            key, value := result[1], result[2]
            ini[head][key] = value
            continue
        }

        if line == "" {
            break
        }
    }
    return ini, nil
}
