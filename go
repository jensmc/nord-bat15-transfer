server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

  location / {
    try_files $uri $uri/ /index.html;
  }
#### Trage deine Knoten hier ein ######
## 21.05.2017

#FF-BU-BA29_2
allow 2a03:2267:4e6f:7264:62e3:27ff:fe57:5dd0;

#FF-BU-HHSTR83
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:3630;

#KreissparkasseRitterhude
##weiter bei https://mesh.nord.freifunk.net/#!v:m;n:98ded065c6c6
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:ccfc;

#FF-BU-WH-Heide_Haus_1
allow 2a03:2267:4e6f:7264:f6f2:6dff:feef:468c;
#FF-BU-WH-Heide_Haus_2-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9f04;
#FF-BU-WH-Heide_LVJSt9_5
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:66e8;
#FF-BU-WH-Heide_LVJSt9_4
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c390;
#FF-BU-WH-Heide_Haus_3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:f04;


##20.05.2017

#HBJensen_Sylt_1
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:dbe6;
#HBJensen_Sylt_4
#allow 2a03:2267:4e6f:7264:62e3:27ff:febd:dc36;
#HBJensen_Sylt_3
#allow 2a03:2267:4e6f:7264:62e3:27ff:febe:75a;
#HBJensen_Sylt_2
#allow 2a03:2267:4e6f:7264:62e3:27ff:fec5:d53e;
#HBJensen_Sylt_7
#allow 2a03:2267:4e6f:7264:92f6:52ff:feb5:2c8a;
#Etoile-V
#allow 2a03:2267:4e6f:7264:46d9:e7ff:fedb:fb24;
#DIAVOLO
#allow 2a03:2267:4e6f:7264:16cc:20ff:febb:d1ae;
#Napapijri
#allow 2a03:2267:4e6f:7264:ea94:f6ff:fea1:d60e;
#Tobacco
#allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:2276;
#Freifunk-SYLT
#allow 2a03:2267:4e6f:7264:16cc:20ff:fee1:5654;

#ff-bu-bahn31_3
allow 2a03:2267:4e6f:7264:32b5:c2ff:feb3:75c6;
#FF-BU-NE44B
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:3658;
#FF-BU-THW_BA35_1
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4b60;

#FF-BU-VA24
allow 2a03:2267:4e6f:7264:8616:f9ff:fe2a:3dc;

#FF-BU-BA140
allow 2a03:2267:4e6f:7264:f6f2:6dff:feef:4662;
#FF-BU-Am_Kamin_2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e568;

#FF-BU-ABH7
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1b70;

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

#FF-DIT-Schwimmbad3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:988e;
#FF-DIT-Schwimmbad1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4c:7639;

#FF-BU-BU37_1
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e58e;

#FF-BU-evGem_GB6
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae92;


#Kirschhof2
allow 2a03:2267:4e6f:7264:6670:2ff:fe9d:ee8a;
#Kirschhof3
allow 2a03:2267:4e6f:7264:62e3:27ff:feed:960c;
#Kirschhof
allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:21c;

#Tarp_Italia_Futro
allow 2a03:2267:4e6f:7264:218:71ff:fe4d:3764;
#Tarp_Italia_Buero
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6c:725e;
#Tarp_Italia_02
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:f7f6;
#Tarp_Italia_03
allow	2a03:2267:4e6f:7264:62e3:27ff:feb7:dd22;

#Christian_Sylt-Futro
allow 2a03:2267:4e6f:7264:219:99ff:fe7a:704e;

# Norderstedt-Harksheide01 
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e75c;

  #18052017
  #FF_Windbergen_Wodansberg_3
  allow 2a03:2267:4e6f:7264:c66e:1fff:fe2d:a208;

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

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #NDS-FFNH-0509-Bispingen-CPE1
  allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:8514;

  #ferienhof_moeller_4
  allow 2a03:2267:4e6f:7264:eade:27ff:fe37:e090;

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #FF-DIT-Meerwald2
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fec2:490;
  #freifunk-nazar-kebabhaus
  #allow 2a03:2267:4e6f:7264:62e3:27ff:fee6:fec0;

  #FF-HB-161
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:43a8;

  #FressflashRuftDasKaenguru
  allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d73e;

  #FF-BU-JKS7
  allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:35b8;


  ####### Unterhalb keine Änderungen durchführen! ###############
  
  root /opt/www;
  autoindex on;
  deny all;
  index index.html index.htm;
  server_name localhost;
}
