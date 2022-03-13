#Microstorage

This is my exercise of creating simple file storage with GoLang

Purpose: store and manipulate with user`s images in my pet projects 

##Version 0.0.1

##Usage

To store file put following iles into "filesystem" dirs
```
filesystem/data/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.json
filesystem/files/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

where xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx is UUID

.json format example:
```
{
 "id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
 "mime": "image/jpeg",
 "ext": ".jpg",
 "transform": {
  "resize": ["800x800","160x160", "640x480"],
  "archive": null,
  "convert": null
 }
}
```

Right now you can define transform.resize property

This array contained enabled sizes (Width x Height in pixels)

You can get a raw file:
```
http://localhost:(network.port)/get/UUID
```

Resized:
```
http://localhost:(network.port)/get/resize/800x800/UUID
```


##Opportunities

* You can store and process images and other files
* You can get resized images
* You can use Redis as cache engine


##Plans
* Other ways to put files: Kafka, Mysql, Redis etc
* File manipulations: gzip/zip/rar, format conversions etc
* Rest API
* Other ways to cache (filesystem, memcached, Mysql, MongoDB etc)
* Sharding (This is my far-reaching plans)

##Todo
* Add immediately conversion (conversion on uploading file), right now only "on_demand" type works  

Hope this code will help to you

Please feel free to message me if you have any questions / ideas

===============================

Stop war, make love