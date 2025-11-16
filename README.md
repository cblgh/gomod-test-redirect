go mod vanity domain and ability to change git forges with less lock-in

i can import the module in this repository by specifiying: 

```golang
import (
"gomod.cblgh.org/gomod-test-redirect"
)

```
this works because my web server is serving the following file, called `gomod-test-redirect`

```
<!DOCTYPE html>
<html>
    <head>
	
	<meta name="go-import" content="gomod.cblgh.org/gomod-test-redirect git https://github.com/cblgh/gomod-test-redirect">
	<meta http-equiv="refresh" content="0;URL='https://github.com/cblgh/gomod-test-redirect'">
    </head>
    <body>
	Redirecting you to the <a href="https://github.com/cblgh/gomod-test-redirect">project page</a>...
    </body>
</html>
```

it basically says "hey, if you came here looking for package `gomod-test-redirect` you can get
the source code from `https://github.com/cblgh/gomod-test-redirect`, which is where i am
currently hosting it. 

now when i switch the git forge host to something else, i just have to change this html file,
BUT everybody's imports in their go files will stay the same i.e. `gomod.cblgh.org/gomod-test-redirect`!

the nginx entry that makes this possible looks like:

```nginx
server {
    listen      443 ssl http2;      # serve over https

    server_name  gomod.cblgh.org;   # the good sauce

    include tls_settings;           # this is just a file i have that contains all the tls stuff needed

    root /var/www/html/_gomod;      # the directory on my webserver being served is called `_gomod`
                                    # but you can call it anything
}
```

many thanks to https://gianarb.it/blog/go-mod-vanity-url
