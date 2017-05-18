# You may add here your
# server {
#	...
# }
# statements for each of your virtual hosts to this file

##
# You should look at the following URL's in order to grasp a solid understanding
# of Nginx configuration files in order to fully unleash the power of Nginx.
# http://wiki.nginx.org/Pitfalls
# http://wiki.nginx.org/QuickStart
# http://wiki.nginx.org/Configuration
#
# Generally, you will want to move this file somewhere, and start with a clean
# file but keep this around for reference. Or just disable in sites-enabled.
#
# Please see /usr/share/doc/nginx-doc/examples/ for more detailed examples.
##

server {
	listen   80; ## listen for ipv4; this line is default and implied
	listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
		try_files $uri $uri/ /index.html;
		# Uncomment to enable naxsi on this location
		# include /etc/nginx/naxsi.rules
	}

	location /batman15/ {
	    proxy_pass http://freifunknord.de/firmware/stable/;
	    proxy_connect_timeout 6s;	
	}
	
	location /snapshots/ {
	    proxy_pass http://ftp.halifax.rwth-aachen.de/lede/snapshots/; 
	    proxy_connect_timeout 6s;
	}	

        location /chaos_calmer/ {
            proxy_pass http://ftp.stw-bonn.de/pub/openwrt/chaos_calmer/;
            proxy_connect_timeout 6s;
        }

        location /modules/ {
            proxy_pass https://raw.githubusercontent.com/Freifunk-Nord/nord-firmware-archiv/master/2016.2.4/stable/modules/;
            proxy_connect_timeout 6s;
        }

        location /archiv/ {
            proxy_pass https://raw.githubusercontent.com/Freifunk-Nord/nord-firmware-archiv/;
            proxy_connect_timeout 6s;
        }

	location /doc/ {
		alias /usr/share/doc/;
		autoindex on;
		allow 127.0.0.1;
		allow ::1;
		deny all;
	}
#18052017

#horst
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:b60e;

#Angeln_0018
allow 2a03:2267:4e6f:7264:c66e:1fff:fea1:2176;
#Moevenberg-Zeltplatz-2
allow 2a03:2267:4e6f:7264:6a72:51ff:fe42:aef8;
#Moevenberg-Zeltplatz
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe3e:8116;
#Sylt-Quelle-Quellenhaus
allow 2a03:2267:4e6f:7264:62e3:27ff:fefb:5c3e;
#Sylt-Quelle-Innenhof
allow 2a03:2267:4e6f:7264:62e3:27ff:fefb:646f;
#Meerkabarett-Eventhalle-2
allow 2a03:2267:4e6f:7264:62e3:27ff:fefb:5da0;
#Moeskendeel-Zeltplatz
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:374a;
# FF-Stolpe-Stolper-See-0
allow 2a03:2267:4e6f:7264:219:99ff:fe68:de53;
# FF-Stolpe-Stolper-See-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:a746;
#Norderstedt-Harksheide01
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e75c;

#18.05.2017

allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:d134;
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:f7f6;
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:bc70;
allow 2a03:2267:4e6f:7264:62e3:27ff:febe:617e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:0478;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:1380;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:1bbc;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:3246;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:7034;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:7076;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:741a;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8640;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8e0c;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8e4a;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:91ee;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:9908;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a9a4;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae6e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:aec2;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:aed2;
allow 2a03:2267:4e6f:7264:62e3:27ff:fecc:ea9c;
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:b5b4;
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:be9e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:d30c;
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:db52;
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:fffa;
allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:013a;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:0170;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:0cc8;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:1e80;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:1fe4;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:20aa;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:20f6;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:21e6;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:22d8;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:2b2c;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:3120;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:34ac;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:3e86;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:436e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4398;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4ea2;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:54e0;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:5758;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:58ee;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:5b8e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:5fd8;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:64a0;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:6512;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:65a6;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:65d4;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:7680;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:867e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:94d8;
allow 2a03:2267:4e6f:7264:62e3:27ff:feed:b462;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:1b24;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:2b56;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:2cf4;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:2d3a;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:2e74;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:3028;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:31e4;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:3366;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:3382;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:6b06;
allow 2a03:2267:4e6f:7264:62e3:27ff:fef5:4ad0;
allow 2a03:2267:4e6f:7264:62e3:27ff:fef5:5bb6;
allow 2a03:2267:4e6f:7264:62e3:27ff:fef5:5dca;
allow 2a03:2267:4e6f:7264:6666:b3ff:fe64:0bbc;
allow 2a03:2267:4e6f:7264:6666:b3ff:fe6f:886c;
allow 2a03:2267:4e6f:7264:6666:b3ff:fe94:7332;
allow 2a03:2267:4e6f:7264:6666:b3ff:fede:8b44;
allow 2a03:2267:4e6f:7264:6666:b3ff:fede:f336;
allow 2a03:2267:4e6f:7264:6666:b3ff:fee8:3acd;
allow 2a03:2267:4e6f:7264:6666:b3ff:feee:735e;
allow 2a03:2267:4e6f:7264:6666:b3ff:fefc:b698;
allow 2a03:2267:4e6f:7264:6670:02ff:fe39:1798;
allow 2a03:2267:4e6f:7264:6670:02ff:fe44:6314;
allow 2a03:2267:4e6f:7264:6670:02ff:fe4c:3c26;
allow 2a03:2267:4e6f:7264:6670:02ff:fe70:a4c6;
allow 2a03:2267:4e6f:7264:6670:02ff:fe98:60f2;
allow 2a03:2267:4e6f:7264:6670:02ff:feaa:baa8;
allow 2a03:2267:4e6f:7264:6670:02ff:fee2:94d8;
allow 2a03:2267:4e6f:7264:6670:02ff:fefb:e8f3;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe26:e2c2;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe28:6059;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe36:0bf8;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe36:168a;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe42:57bc;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe42:ae87;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe44:220d;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4c:ece2;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4e:f5a2;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe60:987d;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe66:1ebf;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe68:15d0;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe68:16c3;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe6c:1eae;
allow 2a03:2267:4e6f:7264:6a72:51ff:fe6c:bcb8;
allow 2a03:2267:4e6f:7264:822a:a8ff:fe2d:e307;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe2a:500a;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe66:5dc4;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe7b:8a80;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe7b:a534;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:a612;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:a80e;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:aa78;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fbda;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fc48;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:6aa6;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:6ba2;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:979e;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:97b8;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:990a;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:9948;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:9a72;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:abc6;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:ae38;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:af40;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:b4b2;
allow 2a03:2267:4e6f:7264:8616:f9ff:fee8:9e88;
allow 2a03:2267:4e6f:7264:92f6:52ff:fe2b:5f44;
allow 2a03:2267:4e6f:7264:92f6:52ff:fe2e:4c5c;
allow 2a03:2267:4e6f:7264:92f6:52ff:fee1:5df6;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe24:312c;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe36:09ac;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe4c:058c;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe4c:2988;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:7aac;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:cc82;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe6c:3bf6;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe72:71fe;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fea8:bcd6;
allow 2a03:2267:4e6f:7264:9ade:d0ff:feaa:917e;
allow 2a03:2267:4e6f:7264:9ade:d0ff:feb4:1e18;
allow 2a03:2267:4e6f:7264:9ade:d0ff:feb9:3e4c;
allow 2a03:2267:4e6f:7264:9ade:d0ff:feb9:9b42;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fec1:d858;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fed5:c83c;
allow 2a03:2267:4e6f:7264:9ade:d0ff:fefe:beb2;
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe05:6960;
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe12:dbfa;
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe77:faea;
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe7a:422c;
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe7a:4298;
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe9f:16e4;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fea9:db7e;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fea9:f368;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fea9:f60e;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fea9:f8fc;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fead:9824;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fead:a310;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fead:a42c;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fec7:b058;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:8252;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:82c0;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:82ee;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:8496;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:9c80;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed1:a8e4;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1b50;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1b9e;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed7:1f42;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:6d78;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:7d66;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:7ff6;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:80fe;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8106;
allow 2a03:2267:4e6f:7264:c24a:00ff:fefa:33fc;
allow 2a03:2267:4e6f:7264:c66e:1fff:fe41:730f;
allow 2a03:2267:4e6f:7264:c66e:1fff:fe53:ef9c;
allow 2a03:2267:4e6f:7264:c66e:1fff:fe7a:9c96;
allow 2a03:2267:4e6f:7264:c66e:1fff:fe82:1342;
allow 2a03:2267:4e6f:7264:c66e:1fff:fea1:1d72;
allow 2a03:2267:4e6f:7264:c66e:1fff:feb3:b2c2;
allow 2a03:2267:4e6f:7264:c66e:1fff:feb8:fe3c;
allow 2a03:2267:4e6f:7264:c66e:1fff:feb9:1450;
allow 2a03:2267:4e6f:7264:c66e:1fff:feb9:14e4;
allow 2a03:2267:4e6f:7264:c66e:1fff:fec5:4fc0;
allow 2a03:2267:4e6f:7264:c66e:1fff:fec5:5004;
allow 2a03:2267:4e6f:7264:c66e:1fff:fec9:c3b4;
allow 2a03:2267:4e6f:7264:c66e:1fff:fec9:d620;
allow 2a03:2267:4e6f:7264:c66e:1fff:fede:bae0;
allow 2a03:2267:4e6f:7264:c66e:1fff:fee7:ad06;
allow 2a03:2267:4e6f:7264:c66e:1fff:fee8:677a;
allow 2a03:2267:4e6f:7264:c66e:1fff:fef4:8ff8;
allow 2a03:2267:4e6f:7264:c66e:1fff:fefe:41ee;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe33:3e4e;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe59:08da;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe6e:6922;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe79:dd48;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe96:ed9c;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe9d:b0fa;
allow 2a03:2267:4e6f:7264:c6e9:84ff:febe:2374;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fed1:a47a;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fef7:b568;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fef7:c0d0;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fef7:d734;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fefd:2118;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fefd:29c6;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fefd:32aa;
allow 2a03:2267:4e6f:7264:cad3:a3ff:fe5c:c88c;
allow 2a03:2267:4e6f:7264:da5d:4cff:fec4:0dbe;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe4c:0b1e;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe69:5d62;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe6b:dbc7;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe78:358e;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fea1:cea8;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fea1:e6bc;
allow 2a03:2267:4e6f:7264:ea94:f6ff:feaa:f608;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fecd:519c;
allow 2a03:2267:4e6f:7264:eade:27ff:fe65:c5f8;
allow 2a03:2267:4e6f:7264:eade:27ff:fefd:2cce;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe20:6d7e;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe2b:aaa8;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe34:bb96;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:0b44;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:3052;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:30da;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:3128;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:3148;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe3f:ce32;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe3f:ce84;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe50:1ff8;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe56:b5ca;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe57:e77e;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe57:efd6;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:2748;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:32c8;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:32f6;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:3314;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:5790;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6c:53ea;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6c:6538;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe73:72ac;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:67b2;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:67f2;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:7276;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73d0;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:7f14;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe7c:ce70;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:b84a;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:ba9c;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:bb14;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:bf8c;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:bf9c;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:bfda;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c1ee;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c278;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c77e;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c8a2;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:ce44;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ca32;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ca46;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:caa0;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cac0;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc02;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc26;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc3a;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc50;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc58;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc66;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccd0;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cdf4;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:d5c4;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:dd9c;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ddfa;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e262;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e27e;
allow 2a03:2267:4e6f:7264:ee08:6bff:feaa:9e34;
allow 2a03:2267:4e6f:7264:ee08:6bff:feab:0a5c;
allow 2a03:2267:4e6f:7264:ee08:6bff:feab:4c6a;
allow 2a03:2267:4e6f:7264:ee08:6bff:fec5:02d0;
allow 2a03:2267:4e6f:7264:ee08:6bff:feec:2bfc;
allow 2a03:2267:4e6f:7264:ee08:6bff:fef1:5908;
allow 2a03:2267:4e6f:7264:ee08:6bff:fef1:5b4e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:93e0;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:a8e0;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:ab42;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:ae9e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:0420;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:2a6a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ae70;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:b0a2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:c848;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:cf7c;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:cfb8;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d0aa;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d18a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e07e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e186;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e1fe;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e6aa;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e6ca;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e778;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e7b6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e82c;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e9d6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:eb9e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:eca2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:edd6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:eedc;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:eee4;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef32;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef52;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:feb2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ff22;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ff6e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ffd8;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:132a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:468a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:4e24;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:5468;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:55d2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:803e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:9c9c;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:d136;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:d6e6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:db66;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:dfb6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e3ba;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e466;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e4d2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:8fae;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:8ffc;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:b27a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe53:154a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe53:336e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe70:a164;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:a406;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:a880;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:aa9a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:be64;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bfa6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:cf32;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d008;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d58e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:e004;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:e090;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe8f:04ac;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:112c;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9a82;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fef2:6f6c;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fefe:d00e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:feff:4462;
allow 2a03:2267:4e6f:7264:f6f2:6dff:feff:6396;
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe3e:8432;
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe3e:93d2;
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe5a:6dae;
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe87:29bc;
allow 2a03:2267:4e6f:7264:fad1:11ff:fe94:773a;
allow 2a03:2267:4e6f:7264:fced:beff:feef:ff02;


#17.05.2017

allow 2a03:2267:4e6f:7264:20c:29ff:fe8e:f1b2;
allow 2a03:2267:4e6f:7264:223:cdff:fe20:bcd2;
allow 2a03:2267:4e6f:7264:618:d6ff:fe27:c9d7;
allow 2a03:2267:4e6f:7264:126f:3fff:fe25:d3f8;
allow 2a03:2267:4e6f:7264:12fe:edff:fe2b:9f3c;
allow 2a03:2267:4e6f:7264:12fe:edff:fe3b:a4f6;
allow 2a03:2267:4e6f:7264:12fe:edff:fe56:a5e4;
allow 2a03:2267:4e6f:7264:12fe:edff:fe84:6d66;
allow 2a03:2267:4e6f:7264:12fe:edff:fef1:4f1a;
allow 2a03:2267:4e6f:7264:16cc:20ff:fe09:2d29;
allow 2a03:2267:4e6f:7264:16cc:20ff:fe31:7266;
allow 2a03:2267:4e6f:7264:16cc:20ff:fe5a:5f14;
allow 2a03:2267:4e6f:7264:16cc:20ff:fe63:83ba;
allow 2a03:2267:4e6f:7264:16cc:20ff:fe7b:83bf;
allow 2a03:2267:4e6f:7264:16cc:20ff:fe9a:a412;
allow 2a03:2267:4e6f:7264:16cc:20ff:fec0:2dc0;
allow 2a03:2267:4e6f:7264:16cc:20ff:fec2:5172;
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:1efc;
allow 2a03:2267:4e6f:7264:16cc:20ff:fee3:dd96;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe27:63c0;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe59:b58c;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe74:2f20;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feb6:ebc4;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:91c4;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed9:857a;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:d288;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e040;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e4c8;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e4f8;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e4fa;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e54c;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e578;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feea:915c;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:5e18;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:5e2a;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:6030;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:6b7e;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:7a40;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:7a96;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:7c0a;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fef5:10b0;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fef5:150e;
allow 2a03:2267:4e6f:7264:1ad6:c7ff:fe45:1766;
allow 2a03:2267:4e6f:7264:1ad6:c7ff:fe51:82f4;
allow 2a03:2267:4e6f:7264:1ad6:c7ff:fe51:b374;
allow 2a03:2267:4e6f:7264:1ad6:c7ff:fee5:d6ea;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:8466;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:8ffa;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:908a;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:90a4;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:9618;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:97e0;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:aa36;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:aa82;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe22:d0b2;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe23:0436;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe3e:82a0;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe3e:8430;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe56:6aa0;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe6f:3380;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe7e:2b5c;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe81:6360;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe85:f026;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe86:3fbe;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe87:80bc;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:460c;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:4dd0;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:4dd4;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe88:6028;
allow 2a03:2267:4e6f:7264:32b5:c2ff:feb3:74a6;
allow 2a03:2267:4e6f:7264:32b5:c2ff:feb3:7634;
allow 2a03:2267:4e6f:7264:32b5:c2ff:feb3:76dc;
allow 2a03:2267:4e6f:7264:32b5:c2ff:feb8:7239;
allow 2a03:2267:4e6f:7264:32b5:c2ff:febd:0158;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fec2:6002;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fec6:bb2e;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fee3:1eb0;
allow 2a03:2267:4e6f:7264:32b5:c2ff:feed:3890;
allow 2a03:2267:4e6f:7264:52c7:bfff:fe36:b5aa;
allow 2a03:2267:4e6f:7264:56e6:fcff:fef1:254a;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe2a:610a;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe2a:67ba;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe36:6d74;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe36:78d2;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe36:83c8;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe53:0692;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe57:4fba;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe57:5108;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe57:5210;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe57:5852;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe5f:ed4e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe60:3a1e;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe67:e454;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe76:26b0;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe81:c34a;
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:de94;
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:9938;
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:9c24;
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:a064;
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:b78c;
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:d098;


	root /opt/www;
	autoindex on;
	deny all;




	index index.html index.htm;

	# Make site accessible from http://localhost/
	server_name localhost;

	

	# Only for nginx-naxsi used with nginx-naxsi-ui : process denied requests
	#location /RequestDenied {
	#	proxy_pass http://127.0.0.1:8080;    
	#}

	#error_page 404 /404.html;

	# redirect server error pages to the static page /50x.html
	#
	#error_page 500 502 503 504 /50x.html;
	#location = /50x.html {
	#	root /usr/share/nginx/www;
	#}

	# pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
	#
	#location ~ \.php$ {
	#	fastcgi_split_path_info ^(.+\.php)(/.+)$;
	#	# NOTE: You should have "cgi.fix_pathinfo = 0;" in php.ini
	#
	#	# With php5-cgi alone:
	#	fastcgi_pass 127.0.0.1:9000;
	#	# With php5-fpm:
	#	fastcgi_pass unix:/var/run/php5-fpm.sock;
	#	fastcgi_index index.php;
	#	include fastcgi_params;
	#}

	# deny access to .htaccess files, if Apache's document root
	# concurs with nginx's one
	#
	#location ~ /\.ht {
	#	deny all;
	#}
}


# another virtual host using mix of IP-, name-, and port-based configuration
#
#server {
#	listen 8000;
#	listen somename:8080;
#	server_name somename alias another.alias;
#	root html;
#	index index.html index.htm;
#
#	location / {
#		try_files $uri $uri/ =404;
#	}
#}


# HTTPS server
#
#server {
#	listen 443;
#	server_name localhost;
#
#	root html;
#	index index.html index.htm;
#
#	ssl on;
#	ssl_certificate cert.pem;
#	ssl_certificate_key cert.key;
#
#	ssl_session_timeout 5m;
#
#	ssl_protocols SSLv3 TLSv1;
#	ssl_ciphers ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv3:+EXP;
#	ssl_prefer_server_ciphers on;
#
#	location / {
#		try_files $uri $uri/ =404;
#	}
#}
