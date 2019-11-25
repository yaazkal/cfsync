# cfsync - Syncs records to Cloudflare DNS

This is a WIP concept. Not ready for production

At the moment there is only support for one *A* record.
Multidomain and multi records support is planned but not yet coded.

You need to create this environment variables in your system:

CF_API_KEY
CF_MAIL
CF_ZONE_NAME
CF_A_RECORD

Example if you want to sync node.example.com

```
CF_API_KEY   = "YOUR API KEY"
CF_MAIL      = "mail@example.com"
CF_ZONE_NAME = "example.com"
CF_A_RECORD  = "node"
```
