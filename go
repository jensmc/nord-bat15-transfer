server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

  location / {
    try_files $uri $uri/ /index.html;
  }

    location /batman15/ {
        alias /opt/www/firmware/;
        autoindex on;
    }

  

#### Trage deine Knoten hier ein ######

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

#FF-BU-BA140
allow	2a03:2267:4e6f:7264:f6f2:6dff:feef:4662;
#FF-BU-Am_Kamin_2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e568;

#FF-BU-ABH7
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1b70;

#NDS-FFNH-0188-Hoeckel-AG02
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:684c;
#NDS-FFNH-0189-Hoeckel-AG03
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6592;

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
#allow 2a03:2267:4e6f:7264:8616:f9ff:fe5b:834e:
#FF-BU-Riedels-Hotel_Burg_02
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:9f5a;
#FF-BU-Kasten-NST2
#allow 2a03:2267:4e6f:7264:ee08:6bff:fe35:3062;
#FF-BU-GB8
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:c71c;
#FF-BU-Kasten_BU2
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e43e;
#FF-BU-MiZ-BU2
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:adf6;
#FF-BU-Aussen_BA1
#allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:2102;
#FF-BU-BA1
#allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:aed0;
#FF-BU-Gemeinde_Ost
#allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:22da;
#FF-BU-Info-Point
#allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:31ac;
#FF-BU-Amt-Burg
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:eda2:

#ff-nf-gar-vermittlung-hoffmann
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:82ee;

#FF-DIT-Schwimmbad3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:988e;
#FF-DIT-Schwimmbad1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4c:7639;

#FF-BU-OFF_03
allow 2a03:2267:4e6f:7264:219:99ff:fe67:4dce;

#FF-BU-OFF_05
allow 2a03:2267:4e6f:7264:219:99ff:fe50:404d;

#NDS-FFNH-128-NeuWu-LeA
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed6:fe28;
#NDS-FFNH-132-NeuWu-LeA
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8134;
#NDS-FFNH-127-NeuWu-LeA
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:80f6;
#NDS-FFNH-138-NeuWu-LeA
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:634e;
#NDS-FFNH-133-NeuWu-LeA
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:7b8c;

#FF-BU-BU37_1
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e58e;

#FF-BU-AM8
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae1a;
#FF-BU-evGem_GB6
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae92;

#Angeln-0006
allow 	2a03:2267:4e6f:7264:c66e:1fff:fea1:234a;
#Angeln-0007
allow 2a03:2267:4e6f:7264:c66e:1fff:fea1:2bf0;
#Angeln-0016
allow 2a03:2267:4e6f:7264:c66e:1fff:fe3b:59a0;
#Angeln-0009
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe6f:38c8;


#Tarp_CPE210_0003
allow 2a03:2267:4e6f:7264:62e3:27ff:fe53:126e;
#Tarp_MobiFEG_04
#allow 2a03:2267:4e6f:7264:c66e:1fff:fea1:2ab2;
#Tarp_Futro_03
#allow 2a03:2267:4e6f:7264:218:71ff:fe4d:3785;

#Kirschhof2
allow 2a03:2267:4e6f:7264:6670:2ff:fe9d:ee8a;
#Kirschhof3
#allow 2a03:2267:4e6f:7264:62e3:27ff:feed:960c;
#Kirschhof
#allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:21c;

#3. Freifunk_Meldorf18
allow 2a03:2267:4e6f:7264:c2a0:bbff:fef0:f146;

#Tarp_Italia_02
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:f7f6;
#Tarp_Italia_03
allow	2a03:2267:4e6f:7264:62e3:27ff:feb7:dd22;

#Tarp_MobiFEG_04
allow	2a03:2267:4e6f:7264:c66e:1fff:fea1:2ab2;

#Christian_Sylt-Futro
allow 2a03:2267:4e6f:7264:219:99ff:fe7a:704e;

# Norderstedt-Harksheide01 
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e75c;


allow 2a03:2267:4e6f:7264:219:99ff:fe7a:704e;
allow 2a03:2267:4e6f:7264:16cc:20ff:fe5e:bc80;
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:53c0;
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe23:0436;
allow 2a03:2267:4e6f:7264:62e3:27ff:feb7:9cbc;
allow 2a03:2267:4e6f:7264:62e3:27ff:febe:0068;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8574;
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8640;
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:fffa;
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:6c2a;
allow 2a03:2267:4e6f:7264:62e3:27ff:feee:3090;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe7b:8a80;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:aa78;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fc48;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:9948;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:9a72;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:a7a0;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:ae38;
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:b4b2;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feaa:0dca;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed7:1aaa;
allow 2a03:2267:4e6f:7264:c24a:00ff:fe49:7622;
allow 2a03:2267:4e6f:7264:c24a:00ff:fee5:83b8;
allow 2a03:2267:4e6f:7264:c66e:1fff:fede:ba5a;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fefc:f80e;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:606e;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:68ca;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:7b5a;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c214;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cbe0;
allow 2a03:2267:4e6f:7264:ee08:6bff:fef1:5908;
allow 2a03:2267:4e6f:7264:ee08:6bff:fef1:5b4e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:b0a2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:e7b6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ecb0;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef24;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef32;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:5476;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e4a2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bfa6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fef2:50bc;


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

  #HammeForumOfficeCPE
  allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:c0ba;

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #NDS-FFNH-0204-Bispingen-VGH
  allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73e6;

  #NDS-FFNH-0509-Bispingen-CPE1
  allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:8514;

  #S77-Suelfeld-Marktplatz-01
  allow 2a03:2267:4e6f:7264:ea94:f6ff:fef2:ed38;

  #ferienhof_moeller_4
  allow 2a03:2267:4e6f:7264:eade:27ff:fe37:e090;

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #FF-DIT-Meerwald2
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fec2:490;
  #freifunk-nazar-kebabhaus
  #allow 2a03:2267:4e6f:7264:62e3:27ff:fee6:fec0;
  #FF-BU-Info-Point_Marne_DS2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:3a50;

  #FF-HB-161
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:43a8;

  #FressflashRuftDasKaenguru
  allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d73e;


  #FF-BU-JKS7
  allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:35b8;

  #FF-BU-HAS35
  allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:6b56;

  #FF-BU-KWKDEICH35-2
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:e4a2;

  ####### Unterhalb keine Änderungen durchführen! ###############
  
  root /opt/www;
  autoindex on;
  deny all;
  index index.html index.htm;
  server_name localhost;
}
