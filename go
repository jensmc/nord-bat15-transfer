server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

  location / {
    try_files $uri $uri/ /index.html;
  }
#### Trage deine Knoten hier ein ######

##20.05.2017
#ff-bu-bahn31_3
allow 2a03:2267:4e6f:7264:32b5:c2ff:feb3:75c6;
#FF-BU-NE44B
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:3658;
#FF-BU-THW_BA35_1
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4b60;

#FF-BU-VA24
allow 2a03:2267:4e6f:7264:8616:f9ff:fe2a:3dc;

#HBJensen_Sylt_1
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:dbe6;

#FF-BU-BA140
allow 2a03:2267:4e6f:7264:f6f2:6dff:feef:4662;
#FF-BU-Am_Kamin_2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e568;

#FF-BU-ABH7
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1b70;

#NDS-FFNH-0187-Hoeckel-AG01
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:682c;
#NDS-FFNH-0188-Hoeckel-AG02
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:684c;
#NDS-FFNH-0189-Hoeckel-AG03
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6592;

#HollenbeckFF_0x07
allow 2a03:2267:4e6f:7264:12fe:edff:fe89:f3ea;
#HollenbeckFF_0x08
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe9f:16e4;

#HollenbeckFF_0x06
allow 2a03:2267:4e6f:7264:a2f3:c1ff:feac:26c8;
#HollenbeckFF_0x02
allow 2a03:2267:4e6f:7264:12fe:edff:feb7:5dcc;
#HollenbeckFF_0x04
allow 2a03:2267:4e6f:7264:6666:b3ff:feaf:dfca;

#FF-BU-223
allow 2a03:2267:4e6f:7264:a62b:b0ff:feaa:dca;
#FF-BU-OFF_08
allow 2a03:2267:4e6f:7264:219:99ff:fe67:4dda;
#FF-BU-LS30
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae0a;

#FF-BU-AM9
allow 2a03:2267:4e6f:7264:62e3:27ff:fe57:51c2;
#FF-BU-Riedels-Hotel_Burg_01
allow 2a03:2267:4e6f:7264:8616:f9ff:fe5b:834e:

  ####### Unterhalb keine Änderungen durchführen! ###############
  
  root /opt/www;
  autoindex on;
  deny all;
  index index.html index.htm;
  server_name localhost;
}
