server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

  location / {
    try_files $uri $uri/ /index.html;
  }
#### Trage deine Knoten hier ein ######
## 21.05.2017
#blaue-elise
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe4c:b1e;

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

#Elbmarsch_RS3
allow 2a03:2267:4e6f:7264:ee08:6bff:fe6f:6c08;

#Itzstedt
allow 2a03:2267:4e6f:7264:c24a:ff:fe2c:a26c;

#FB-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3d6e;
#FB-3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:388c;

#behard
allow	2a03:2267:4e6f:7264:a62b:b0ff:fed7:1eb2;

#FF_SL_Husumer_Str_001
allow 2a03:2267:4e6f:7264:eade:27ff:fefd:2cce;

#FF_SL_Holm_Johanniskloster_A
allow 2a03:2267:4e6f:7264:ee08:6bff:fe56:cfa4;
#FF_SL_Holm_Johanniskloster_B
allow 2a03:2267:4e6f:7264:c24a:ff:fe0b:a2d2;

#ffn-Werft-Eberhardt-2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe4f:fe48;
#arnis-netzperfekt2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c942;

#FF-Ferienwohnung-Wiese-Nickel-03
allow 2a03:2267:4e6f:7264:9ade:d0ff:fea8:c0c6;
#FF-Ferienwohnung-Wiese-Nickel-02
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe2f:e632;

#www_schlei-charter-arnis_de_1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe58:fe4c;
#www_schlei-charter-arnis_de_2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe2b:b498;
#www_schlei-charter-arnis_de_03
allow 2a03:2267:4e6f:7264:9ade:d0ff:fec2:152e;

#FF-DIT-Kappeln3
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:15be;

#ffn-Schleihalle-3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee1:e032;
#ffn-Schleihalle-4
allow 2a03:2267:4e6f:7264:62e3:27ff:feed:e31c;
#ffn-Schleihalle-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fed5:edfe;

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


#FF-DIT-Schwimmbad3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:988e;
#FF-DIT-Schwimmbad1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4c:7639;

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
