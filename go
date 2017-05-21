
server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

  location / {
    try_files $uri $uri/ /index.html;
  }
#### Trage deine Knoten hier ein ######
## 21.05.2017
#FF-HB-Asch-Huett-15
#weiter mit https://mesh.nord.freifunk.net/#!v:m;n:f4f26d3fd1cc
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d138; 

#Meldorfer_Buecherstube
#weiter mit https://mesh.nord.freifunk.net/#!v:m;n:c8d3a35cc166
allow 2a03:2267:4e6f:7264:f6f2:6dff:fef6:66fe;

#Schuhhaus_Carstensen
allow 2a03:2267:4e6f:7264:62e3:27ff:feed:db02;

#Eiscafe-Boethern
allow 2a03:2267:4e6f:7264:c66e:1fff:fed6:6f7c;
#VHS_Meldorf_1
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe3e:8e34;
#FF_Meldorf_DerKulturonkel
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:aa9c;

#FF-DIT-Stadtbuecherei2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d3d8;
#FF-DIT-Stadtbuecherei1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d542;
#VHS_Meldorf_2
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe0e:8fc0;

#FF-BU-St_Annen-Kirche
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:49c4;

#FF-HB-029
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:6f8c;

#FF-HB-Asch-ses-iwi
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4110;
#FF-HB-Asch-ses-iwi-0
allow 2a03:2267:4e6f:7264:6666:b3ff:fe2e:fa5;
#FF-HB-Asch-Schu-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:dfec;

#MTV-von-1860-ev-Heide_Gaststube
allow 2a03:2267:4e6f:7264:ee08:6bff:feec:273e;

#FF-BU-WH-Heide_Haus_2-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9f04;

#FF-BU-Woe-Ring11a
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e470;

#FF-WesthofWoehrden-2
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e296;

#FF-WesthofWoehrden-1
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e498;

#FF-BU-Woe-Chstr2
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e494;
#FF-BU-Woe-Chstr8
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9fc6;

#FF-BU-Woe-CNS6
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:ef32;

#NDS-Hue-Rathaus-0292
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc00;
#NDS-Hue-Rathaus-0291
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:e11c;

#Am-Exer-01
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe6e:6922;
#SeasideAppartment02
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe6e:3ed2;

#Eckernfoerde02
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bf12;
#Eckernfoerde01
allow 2a03:2267:4e6f:7264:c66e:1fff:fee8:eec4;

#Doerpstedt-FFW-Sport
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe33:3e4e;

#FF-BU-GB8
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:c71c;

#FF-BU-223
allow 2a03:2267:4e6f:7264:a62b:b0ff:feaa:dca;

#FFNH-Deutsches-Haus-02
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed9:7bb6;
#FFNH-Deutsches-Haus-01
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fed9:79ce;

#FFNH-Origene-Praxis
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:ae9e;

#Freifunk_Geestlandhalle
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:9190;

#Elbmarsch_OL3
allow 2a03:2267:4e6f:7264:219:99ff:fe5b:45cc;
#Elbmarsch_RS1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6f:68a4;
#Elbmarsch_RS2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:8cac;
#Elbmarsch_RS3
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6f:6c08;

#FB-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3d6e;
#FB-3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:388c;

#behard
allow	2a03:2267:4e6f:7264:a62b:b0ff:fed7:1eb2;

#FF_SL_Husumer_Str_001
allow 2a03:2267:4e6f:7264:eade:27ff:fefd:2cce;

#FF-DIT-Kappeln1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:9710;
#FF-DIT-Kappeln2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:15c4;
#FF-DIT-Kappeln3
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:15be;

#Napapijri
allow 2a03:2267:4e6f:7264:ea94:f6ff:fea1:d60e;
#Tobacco
allow 2a03:2267:4e6f:7264:62e3:27ff:fecf:2276;


#RBZ10
allow 2a03:2267:4e6f:7264:6670:2ff:fed3:9fc;

#FF-BU-WH-Heide_Haus_1
allow 2a03:2267:4e6f:7264:f6f2:6dff:feef:468c;
#FF-BU-WH-Heide_Haus_2-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:9f04;

allow 2a03:2267:4e6f:7264:219:99ff:fe7a:932f;
allow 2a03:2267:4e6f:7264:16cc:20ff:fec0:177e;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee5:e55a;
allow 2a03:2267:4e6f:7264:1aa6:f7ff:feeb:6b7e;
allow 2a03:2267:4e6f:7264:32b5:c2ff:feed:537a;
allow 2a03:2267:4e6f:7264:5054:00ff:fee9:59f9;
allow 2a03:2267:4e6f:7264:62e3:27ff:fece:d30c;
allow 2a03:2267:4e6f:7264:6670:02ff:feba:03d4;
allow 2a03:2267:4e6f:7264:8616:f9ff:fe2a:03b4;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:82ee;
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:9190;
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:6d78;
allow 2a03:2267:4e6f:7264:c24a:00ff:fe2c:a26c;
allow 2a03:2267:4e6f:7264:c6e9:84ff:fe33:3e4e;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe4c:0b1e;
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e75c;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe61:32d0;
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c278;
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:c97a;
allow 2a03:2267:4e6f:7264:ee08:6bff:feab:254a;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:ae9e;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:9e26;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:c71c;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe49:def2;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bfa6;
allow 2a03:2267:4e6f:7264:f6f2:6dff:fef2:50bc;

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

#FF-BU-THW_BA35_1
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:4b60;

#HollenbeckFF_0x06
allow 2a03:2267:4e6f:7264:a2f3:c1ff:feac:26c8;
#HollenbeckFF_0x02
allow 2a03:2267:4e6f:7264:12fe:edff:feb7:5dcc;
#HollenbeckFF_0x04
allow 2a03:2267:4e6f:7264:6666:b3ff:feaf:dfca;

#FF-BU-223
allow 2a03:2267:4e6f:7264:a62b:b0ff:feaa:dca;

#FF-BU-LS30
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:ae0a;

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

  #freifunk-nazar-kebabhaus
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee6:fec0;

  #FF-HB-161
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:43a8;

  #FressflashRuftDasKaenguru
  allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d73e;

  ####### Unterhalb keine Änderungen durchführen! ###############
  
  root /opt/www;
  autoindex on;
  deny all;
  index index.html index.htm;
  server_name localhost;
}
