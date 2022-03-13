#Microstorage

This is my exercise of creating simple file storage with GoLang

Purpose: store and manipulate with user`s images in my pet projects 

##Version 0.0.2

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

All files list (will be removed in future)
```
http://localhost:9986)/api/files/list
```

##Opportunities

* You can store and process images and other files
* You can get resized images
* You can use Redis as cache engine


##Todo
* Ssl encryption
* Other ways to put files: Kafka, Mysql, Redis etc
* File manipulations: gzip/zip/rar, format conversions etc
* Rest API
* Other ways to cache (filesystem, memcached, Mysql, MongoDB etc)
* Sharding (This is my far-reaching plans)
* Add immediately conversion (conversion on uploading file), right now only "on_demand" type works  
* JWT auth to getting files
* Blur image if user not logged in (using JWT)

Hope this code will help to you

Please feel free to message me if you have any questions / ideas

===============================

Stop war, make love