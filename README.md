# Slilding Cache
Sliding Expiration Cache (ncsa)

Basic Sliding Expiration Cache

A default expiration can be passed when creating the cache.  If not passed, the cache is created with a default one.

Add function adds item to the cache.  An overiding duration can be passed to the item.  If not passed the cache default duration will be used.

Delete function removes item from cache.

A gosub function is running all the time upon creation of the cache.  It wakes periodically and removes any items in the cache that have reached the desired timeout.