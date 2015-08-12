#Go-Attach
Attach go web application into web server seamlessly

```
./config	directory
config.json		main configuration
apps.json 		go-app attachment list

./template	directory
all require template place it here

./nginx
all nginx go related code

./goattach
-config		name of config file, if not provided will read ./config/config.json
-attach 	operation is attach
-detach		operation is detach
-id			id of application
-webserver	type of webserver (nginx/apache/tomcat/iis/glassfish/etc), if not provided default is nginx
-port		port of application
-alias		domain name alias for application
-sync		sync webserver with new settings

ie:
goattach -attach -id=derivative -alias=derivative.eaciit.com -port=9001
goattach -sync

goattach -attach -id=derivative -alias=derivative.eaciit.com -port=9001 -sync
```