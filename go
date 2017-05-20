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

#20.5.2017

#Dorfstrasse27_TL-WR1043ND_Test
allow 2a03:2267:4e6f:7264:8616:f9ff:fe5b:8846;
#Dorfstr_27
allow 2a03:2267:4e6f:7264:6670:2ff:fe67:899c;

#FF-Silberstedt-HS25
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9a4a;


#Angeln-Satrup-0025
allow 2a03:2267:4e6f:7264:92f6:52ff:fef2:6f7c;


#Christian_Sylt-Futro
allow 2a03:2267:4e6f:7264:219:99ff:fe7a:704e;

  # Norderstedt-Harksheide01
  allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e75c;


  #1. MK19_02_ffnord
  allow 2a03:2267:4e6f:7264:c66e:1fff:feaf:4dc0;
  #2. NDS-FFNH-224-MK19
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ca56;
  #3. NDS-FFNH-100-MK19
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e2fc;
  #zuletzt dann noch 4. MK19_FFNH_Offloader
  allow 2a03:2267:4e6f:7264:2e0:4cff:fe4d:de21;    
    
  #19.05.2017
  

  allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:a2a4;
  allow 2a03:2267:4e6f:7264:32b5:c2ff:fe6e:f6dc;
  allow 2a03:2267:4e6f:7264:62e3:27ff:fe36:6ea4;
  allow 2a03:2267:4e6f:7264:62e3:27ff:fe60:1fb0;
  allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:8640;
  allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:acd8;
  allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae8a;
  allow 2a03:2267:4e6f:7264:62e3:27ff:fece:fffa;
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:5dc6;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:aa78;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fb14;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fe9b:fc48;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:979e;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:9948;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:9a72;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:abc6;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:ae38;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:b4b2;
  allow 2a03:2267:4e6f:7264:8616:f9ff:fee8:9192;
  allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:82ae;
  allow 2a03:2267:4e6f:7264:a62b:b0ff:fed2:1aa6;
  allow 2a03:2267:4e6f:7264:ea94:f6ff:fe51:3c7c;
  allow 2a03:2267:4e6f:7264:eade:27ff:fe65:8a91;
  allow 2a03:2267:4e6f:7264:ee08:6bff:fe3f:ca82;
  allow 2a03:2267:4e6f:7264:ee08:6bff:fe6c:5ba4;
  allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:670e;
  allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6864;
  allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6894;
  allow 2a03:2267:4e6f:7264:ee08:6bff:feab:0aec;
  allow 2a03:2267:4e6f:7264:ee08:6bff:feab:65aa;
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:b0a2;
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ee5e;
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef32;
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:fed4;
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:53a4;
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe5f:14d2;
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

  #NDS-FFNH-0101-Bispingen-EK - done
  #NDS-FFNH-0509-Bispingen-CPE1
  allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:8514;

  #S77-Suelfeld-Marktplatz-01
  allow 2a03:2267:4e6f:7264:ea94:f6ff:fef2:ed38;

  #ferienhof_moeller_4
  allow 2a03:2267:4e6f:7264:eade:27ff:fe37:e090;

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

  ####### Unterhalb keine Änderungen durchführen! ###############
  
  root /opt/www;
  autoindex on;
  deny all;
  index index.html index.htm;
  server_name localhost;
}
