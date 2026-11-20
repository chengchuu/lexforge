# NGINX location lab

Use this lab with the `tag:nginx` cards to test location matching, static file mapping, and reverse proxying. The card snippets come from the configuration below. Run the commands in the same shell with a standard NGINX build that includes the HTTP proxy module, plus curl. The lab listens on loopback ports 18080, 18081, and 18082. If a port is occupied, use three free ports and update the matching listen addresses, backend references, and curl URLs.

## Create the lab

Create a temporary directory and sample files.

```sh
LAB=$(mktemp -d "${TMPDIR:-/tmp}/lexforge-nginx.XXXXXX")
mkdir -p "$LAB/logs" "$LAB/tmp" "$LAB/html/root" "$LAB/html/files" "$LAB/assets"
printf '%s\n' 'root asset' > "$LAB/html/root/asset.txt"
printf '%s\n' 'alias asset' > "$LAB/assets/asset.txt"
printf '%s\n' 'existing file' > "$LAB/html/files/present.txt"
printf '%s\n' 'fallback file' > "$LAB/html/fallback.html"
```

Write the configuration. Keep the heredoc delimiter quoted so the shell preserves NGINX variables such as `$uri` and `$request_uri`.

```sh
cat > "$LAB/nginx.conf" <<'NGINX'
worker_processes 1;
pid logs/nginx.pid;
lock_file logs/nginx.lock;
error_log logs/error.log notice;

events {}

http {
    default_type text/plain;
    access_log logs/access.log;
    client_body_temp_path tmp/client;
    proxy_temp_path tmp/proxy;
    fastcgi_temp_path tmp/fastcgi;
    uwsgi_temp_path tmp/uwsgi;
    scgi_temp_path tmp/scgi;

    upstream lab_backends {
        server 127.0.0.1:18081;
        server 127.0.0.1:18082;
    }

    server {
        listen 127.0.0.1:18081;

        location / {
            return 200 "backend-a $request_uri\n";
        }
    }

    server {
        listen 127.0.0.1:18082;

        location / {
            return 200 "backend-b $request_uri\n";
        }
    }

    server {
        listen 127.0.0.1:18080;
        root html;

        location = /health {
            access_log logs/health-access.log;
            return 200 "healthy\n";
        }

        location = /match/exact.json {
            return 200 "exact\n";
        }

        location /match/ {
            return 200 "prefix\n";
        }

        location /match/deep/ {
            return 200 "longest prefix\n";
        }

        location ^~ /match/guarded/ {
            return 200 "priority prefix\n";
        }

        location ~ ^/match/.+[.]json$ {
            return 200 "regex\n";
        }

        location /root/ {
            error_log logs/static-errors.log;
            root html;
        }

        location /alias/ {
            alias assets/;
        }

        location /files/ {
            try_files $uri =404;
        }

        location /fallback/ {
            try_files $uri /fallback.html;
        }

        location = /fallback.html {
            add_header X-Lab-Location fallback;
        }

        location /api/ {
            proxy_pass http://127.0.0.1:18081/;
        }

        location /keep/ {
            proxy_pass http://127.0.0.1:18081;
        }

        location /balanced/ {
            proxy_pass http://lab_backends/;
        }
    }
}
NGINX
```

Check the configuration and start the lab. The prefix keeps the PID, logs, and temporary files with this instance.

```sh
nginx -p "$LAB/" -c nginx.conf -e stderr -t &&
nginx -p "$LAB/" -c nginx.conf -e stderr
```

## Verify the routes

Run the requests and compare the responses. The table gives the body unless it labels a status or header.

| Request path | Expected result |
| --- | --- |
| `/health` | `healthy` |
| `/match/exact.json` | `exact` |
| `/match/item` | `prefix` |
| `/match/deep/item` | `longest prefix` |
| `/match/deep/item.json` | `regex` |
| `/match/guarded/item.json` | `priority prefix` |
| `/root/asset.txt` | `root asset` |
| `/alias/asset.txt` | `alias asset` |
| `/files/present.txt` | `existing file` |
| `/files/missing.txt` | HTTP 404 |
| `/fallback/missing.txt` | `fallback file` with header `X-Lab-Location: fallback` and status 200 |
| `/api/item?x=1` | `backend-a /item?x=1` |
| `/keep/item?x=1` | `backend-a /keep/item?x=1` |
| `/balanced/item` | `backend-a /item` or `backend-b /item` |

```sh
base=http://127.0.0.1:18080
curl --noproxy '*' -fsS "$base/health"
curl --noproxy '*' -fsS "$base/match/exact.json"
curl --noproxy '*' -fsS "$base/match/item"
curl --noproxy '*' -fsS "$base/match/deep/item"
curl --noproxy '*' -fsS "$base/match/deep/item.json"
curl --noproxy '*' -fsS "$base/match/guarded/item.json"
curl --noproxy '*' -fsS "$base/root/asset.txt"
curl --noproxy '*' -fsS "$base/alias/asset.txt"
curl --noproxy '*' -fsS "$base/files/present.txt"
curl --noproxy '*' -sS -o /dev/null -w '%{http_code}\n' "$base/files/missing.txt"
curl --noproxy '*' -si "$base/fallback/missing.txt"
curl --noproxy '*' -fsS "$base/api/item?x=1"
curl --noproxy '*' -fsS "$base/keep/item?x=1"
for n in 1 2 3 4; do
  curl --noproxy '*' -fsS "$base/balanced/item"
done
```

With one worker and two equally weighted backends, the last four requests should alternate between the two backends while both are available. The fallback request returns its file in the same response.

## Inspect logs and reload

Generate a missing-file request under `/root/`, then read the location-specific logs.

```sh
curl --noproxy '*' -sS -o /dev/null "$base/root/missing.txt"
tail -n 1 "$LAB/logs/health-access.log"
tail -n 1 "$LAB/logs/static-errors.log"
```

Edit the `/health` response in `$LAB/nginx.conf` from `healthy` to `ready`. Test and reload the same instance. Once the new workers take over, the request returns `ready`.

```sh
nginx -p "$LAB/" -c nginx.conf -e stderr -t &&
nginx -p "$LAB/" -c nginx.conf -e stderr -s reload
curl --noproxy '*' -fsS "$base/health"
```

If the response still reads `healthy`, repeat the curl command after the reload completes.

Stop the lab with the same prefix and configuration. The temporary directory remains available for inspecting logs after shutdown.

```sh
nginx -p "$LAB/" -c nginx.conf -e stderr -s quit
```

## References

- [NGINX beginner’s guide](https://nginx.org/en/docs/beginners_guide.html)
- [HTTP core module](https://nginx.org/en/docs/http/ngx_http_core_module.html)
- [Proxy module and URI handling](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass)
- [HTTP upstream groups](https://nginx.org/en/docs/http/ngx_http_upstream_module.html)
- [Command-line parameters](https://nginx.org/en/docs/switches.html)
