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

#NDS-FFNH-138-NeuWu-LeA
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:634e;
#NDS-FFNH-133-NeuWu-LeA
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:7b8c;

#FF-BU-BU37_2
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e5a0;

#FF-BU-AM8
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae1a;
#FF-BU-evGem_GB6
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae92;

#FF-BU-Aussen-2
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1b0e;
#FF-BU-312
allow	2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e55a;

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


#20.5.2017
#Kirschhof2
allow 2a03:2267:4e6f:7264:6670:2ff:fe9d:ee8a;
#Kirschhof3
#allow 2a03:2267:4e6f:7264:62e3:27ff:feed:960c;
#Kirschhof
#allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:21c;

#2. Altes_Rathaus_Garding
#allow 	2a03:2267:4e6f:7264:f6f2:6dff:fe85:cb7a;
#3. Freifunk_Meldorf18
#allow 2a03:2267:4e6f:7264:c2a0:bbff:fef0:f146;


#Angeln-0011
allow 2a03:2267:4e6f:7264:32b5:c2ff:feed:32c0;

#Eggebek_0001
allow 2a03:2267:4e6f:7264:c66e:1fff:fee8:eafc;

#Eggebek_0002
allow 2a03:2267:4e6f:7264:12fe:edff:fec4:68f8;

#Tarp_Tornschauer_Str_17
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe90:9828;

#Tarp_Italia_01
allow 2a03:2267:4e6f:7264:eade:27ff:fe55:a0d4;
#Tarp_Italia_04
allow 2a03:2267:4e6f:7264:c66e:1fff:fe41:e97a;

#FF-Silberstedt-MW8
allow 2a03:2267:4e6f:7264:c24a:ff:fe49:7622;
#FF-Silberstedt-HS25
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9a4a;


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
