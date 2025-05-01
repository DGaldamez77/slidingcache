# Slilding Cache
Sliding Expiration Cache (ncsa)

Basic Sliding Expiration Cache

A default expiration can be passed when creating the cache.  If not passed, the cache is created with a default one.

Add function adds item to the cache.  An overiding duration can be passed to the item.  If not passed the cache default duration will be used.

Delete function removes item from cache.

A gosub function is running all the time upon creation of the cache.  It wakes periodically and removes any items in the cache that have reached the desired timeout.

#### Things that I would do different

This works OK as a cache for in-process.  However, if deploying to a multi-pod environment like kubernetes, instead of using a map that is only available in the process, I would use some other service for cache, lime MemCache or Redis.  That way we could have one pod running the cache and all pods would have access to it.  

Another change would be to change the cache library where we can pass a function for callback when the items expire.  That way they can decide to refresh the data or just extend its expiration.