# cfsync - Syncs records to Cloudflare DNS

This is a WIP concept.

Multidomain and multi records are supported but not completly implemented.

Rename the file `config.example.yml` to `config.yml`

Example if you want to sync *node.example.com*

```yaml
api_key: YOUR_API_KEY_HERE
email: mail@example.com
zones:
  example.com:
    - name: node
      type: A
```

If you want to sync more than one domain and record, check the `config.example.yml` file for inspiration