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

#FF_Windbergen_Wodansberg_3
allow 2a03:2267:4e6f:7264:c66e:1fff:fe2d:a208;

#FF-CC-IZ-AP03
allow 2a03:2267:4e6f:7264:ee08:6bff:fe2d:6ee0;

#AWO_Leck_Obergeschoss
allow 2a03:2267:4e6f:7264:c66e:1fff:fea1:2b82;


#Pension-Galerie_Vierhoefen_03
allow 2a03:2267:4e6f:7264:c66e:1fff:fe63:4724;

#DieMuehle
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed1:e178;

#NDS-FFNH-120-NeuWu-Rathaus-SO
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8178;

#FF-BU-HDHB41
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4bac;

#Faehrhaus_Hodorf2
allow 2a03:2267:4e6f:7264:6a72:51ff:fe36:c0c;

#FF-Hanstedt-12
allow 2a03:2267:4e6f:7264:62e3:27ff:feed:86fc;

#HammeForumOfficeCPE
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:c0ba;

#NDS-ReiseAtelier
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:df62;

#NDS-HotelzurPost
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccb4;

#NDS-FFNH-0145-Bispingen-AKR2
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

#NDS-FFNH-0204-Bispingen-VGH
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73e6;

#NDS-FFNH-0101-Bispingen-EK
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf2;
#NDS-FFNH-0509-Bispingen-CPE1
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:8514;

#ffrow-nds-huelsefrucht-2-i
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:adc6;

#GalerieHotelHaar-Worpswede-1823
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fb14;

#S77-Suelfeld-Marktplatz-01
allow 2a03:2267:4e6f:7264:ea94:f6ff:fef2:ed38;

#FreifunkKaltenkirchen1
allow 2a03:2267:4e6f:7264:1ad6:c7ff:fee8:ae7f;
#FreifunkKaltenkirchen3
allow 2a03:2267:4e6f:7264:1ad6:c7ff:fee8:ac8d;

#TheRing
allow 2a03:2267:4e6f:7264:6670:2ff:feae:60de;
#ferienhof_moeller_4
allow 2a03:2267:4e6f:7264:eade:27ff:fe37:e090;

## SY-Rathaus-Futro beachten !!!

#SY-Rathaus-Maybachstrasse
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe6a:93ba;
#SY-GemeindeSylt05
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:9f52;
#SY-GemeindeSylt06
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:9a54;
#SY-GemeindeSyltIT
allow 2a03:2267:4e6f:7264:62e3:27ff:fe93:a9de;

#NDS-FFNH-0145-Bispingen-AKR2
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;


#FunPoint-01
allow 2a03:2267:4e6f:7264:ee08:6bff:fe3f:ce32;
#mediamor-01
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:5272;
#Captain-Ahabs-01
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:44e4;


#FF-DIT-Meerwald2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fec2:490;
#freifunk-nazar-kebabhaus
#allow 2a03:2267:4e6f:7264:62e3:27ff:fee6:fec0;
#FF-BU-Info-Point_Marne_DS2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:3a50;

#FFNH-VVUndeloh-003
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe5f:14d2;

#FF-HB-161
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:43a8;

#Hotel-Neptuns-Ankerplatz-03
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:49da;

#FressflashRuftDasKaenguru
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d73e;

#FeWo-Demo
allow 2a03:2267:4e6f:7264:62e3:27ff:fe60:1fb0;

#tarptown2
allow 2a03:2267:4e6f:7264:16cc:20ff:fe70:423a;

#FF-BU-JKS7
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:35b8;

#FF-BU-WA124
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:acd8;

#FF-BU-wa214b
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1aa6;

#FF-BU-GM60
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ee5e;

#FF-BU-STA28
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:c756;

#FF-BU-SoeBe-Haupt
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:fed4;

#FF-BU-HAS35
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:6b56;

#FF-BU-STMA-Dostr-19
allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:10b2;

#VHSBRB04
allow 2a03:2267:4e6f:7264:ee08:6bff:fe3f:ca82;

#VHSBRB01
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:bed4;

#FF-BU-KWKDEICH35-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e4a2;



#FF-BU-BBST4_2
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe7a:4bc6;


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
