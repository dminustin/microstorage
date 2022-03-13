package main

import (
    "fmt"
    "microstorage"
    "microstorage/cache"
    "microstorage/consumers"
    "microstorage/transformers/resize"
    "net/http"
    "os"
    "strings"
    "time"
)

func main() {
    microstorage.LogMessage("App started")
    microstorage.Init()
    consumers.BaseStoragePath = microstorage.Config.App.StoragePath
    cache.Setup(microstorage.Config.Cache)
    if microstorage.Config.Listen.Filesystem {
        consumers.ListenFilesystemInit()
        go consumers.ListenFilesystem(microstorage.Config.Delay.Filesystem)
    }

    http.HandleFunc("/get/", handleGetRequest)
    http.ListenAndServe(":" + fmt.Sprint(microstorage.Config.Network.Port), nil)

    for ;; {
        time.Sleep(time.Second)
    }
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
    uri := r.RequestURI

    if cache.Cache.Exists(uri) {
        tmp, err := cache.Cache.Get(uri)
        if err == nil {
            microstorage.LogMessage("Got " + tmp.Data.Id + " from cache")
            doWriteResult(w, tmp.Data, tmp.File)
            return
        }
    }
    go cache.CacheRoutine()

    path := strings.Split(r.RequestURI, "/")
    path = path[2:]

    id := path[len(path)-1]
    microstorage.IsUUID(id)


    data_path, file_path := consumers.GetNestedPath(id)
    data := consumers.LoadInputJson(data_path)
    if 1==2 {
        w.Header().Set("Content-type", data.Mime)
    }

    if len(path)>1 {
        switch path[0] {
        case "resize":
            for _, v := range data.Transform.Resize {
                if v == path[1] {
                tmp := strings.Split(path[1], "x")
                width:= microstorage.StrToInt64(tmp[0])
                height:= microstorage.StrToInt64(tmp[1])
                res,_ := resize.ImgResize(file_path, width, height)
                w.Write(res)
                cache.PutCache<-cache.CachedObjectChan{Uri: uri, Object: cache.CachedObject{
                    Data: data,
                    File: res,
                }}
                return
                }
            }
            w.WriteHeader(404)
            break
        case "archive":
            break
        case "convert":
            break
        }
    } else {
        res, _ := os.ReadFile(file_path)
        w.Write(res)
        cache.PutCache<-cache.CachedObjectChan{Uri: uri, Object: cache.CachedObject{
            Data: data,
            File: res,
        }}
    }
}

func doWriteResult(w http.ResponseWriter, data consumers.InputData, file []byte) {
    w.Write(file)
}