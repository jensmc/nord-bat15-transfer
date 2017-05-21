
server {
  listen   80; ## listen for ipv4; this line is default and implied
  listen   [::]:80 default_server ipv6only=on; ## listen for ipv6

  location / {
    try_files $uri $uri/ /index.html;
  }
#### Trage deine Knoten hier ein ######
## 21.05.2017
#FF-Hanstedt-20
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:bfa6;

#NDS-FFNH-0532-Erikastrasse02
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fee1:79a8;

#FF-HB-Borg-Ferienhaus_Britta
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d194;
#FF-HB-Borg-Ferienhaus_Tina
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:64b2;

#FF-HB-Borg-Trae-3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d1a6;
#FF-HB-Borg-Trae-1a
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d102;

#NDS-FFNH-242-Hanstedt-009
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c75a;

#NDS-FFNH-243-Hanstedt-010
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:c152;
#FF-Hanstedt-22
allow 2a03:2267:4e6f:7264:a62b:b0ff:fea2:3384;

#drestedt_asyl_jupiter
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:b16e;
#drestedt_asyl_saturn
allow 2a03:2267:4e6f:7264:62e3:27ff:febd:9be2;
#mushreq
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe9f:1bce;
#drestedt_asyl_mars
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:9a9a;

#NDS-FFNH-131-NeuWu-Turek
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8010;
#NDS-FFNH-1401-NeuWu-Uhu
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe65:cac8;
#NDS-FFNH-139-NeuWu-Plover
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:cc8a;
#NDS-FFNH-212-NeuWu-Sterna
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:679a;
##weiter mit https://mesh.nord.freifunk.net/#!v:m;n:0019999ae5b3

#NDS-FFNH-213-Moisburg-Rathaus
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:6884;
#NDS-FFNH-211-Moisburg-Rathaus
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:678c;

#Bodderbrod_mit_Kees
allow 2a03:2267:4e6f:7264:c6e9:84ff:fefd:db0;
#Leuchte_des_Nordens
allow 2a03:2267:4e6f:7264:6a72:51ff:fe34:afc2;

#Wiep4
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:4018;
#Wiep5
allow 2a03:2267:4e6f:7264:f6f2:6dff:fee8:262c;
#Wiep2
allow 2a03:2267:4e6f:7264:62e3:27ff:fec6:fc5a;
#hechthausen-001
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe6e:c514;
#hechthausen-002
allow 2a03:2267:4e6f:7264:9ade:d0ff:fe92:4de8;

#Privilegierte_Stadtapotheke
#allow 2a03:2267:4e6f:7264:62e3:27ff:fece:ff4c;
#Ristorante_Mama_Leone
#allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:dcfe;
#freifunk_Meld_Suederm
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe40:5406;
#Provinzial_Nord_Meldorf
#allow 2a03:2267:4e6f:7264:62e3:27ff:feb8:49e;
#MED-Rathausplatz-FF-03
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:dfea;
#MED-Rathausplatz-FF-02
#allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:d1d4;
#MED-Rathausplatz-FF-04
#allow a03:2267:4e6f:7264:f6f2:6dff:fe85:e0ea;

#HV-Stelle-1043-v2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3b62;
#HV-Stelle-1043
allow 2a03:2267:4e6f:7264:8616:f9ff:fe66:5f92;
#HV-Stelle
allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:8caa;

#ff-nf-spo-fewo-carstens-1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73c8;

#InuYasha
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe69:4277;
#Kikyo
allow 2a03:2267:4e6f:7264:6a72:51ff:fe06:5e95;
#Kaede
allow 2a03:2267:4e6f:7264:c66e:1fff:feb9:13f8;
#Kiara
allow 2a03:2267:4e6f:7264:eade:27ff:fe65:cae4;
#Kagome
allow 2a03:2267:4e6f:7264:32b5:c2ff:fe6e:9cba;

#FF-NMS-CB-01
allow 2a03:2267:4e6f:7264:12fe:edff:fe92:eb24;
#FF-NMS-CB-04
allow 2a03:2267:4e6f:7264:ee08:6bff:fe56:d9a0;
#FF-NMS-NSM-01
allow 2a03:2267:4e6f:7264:6a72:51ff:fe16:6376;
#FF-NMS-AB
allow 2a03:2267:4e6f:7264:12fe:edff:fe7a:38fa;

#tangogolf
allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:d83e;
#tangogolf_2
allow 2a03:2267:4e6f:7264:a62b:b0ff:fec7:cdb8;

#FF-Stolpe-DGH-0
allow 2a03:2267:4e6f:7264:219:99ff:fe68:dece;
#FF-Stolpe-DGH-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe22:840c;
#FF-Stolpe-DGH-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe22:8092;

#eae-badsegeberg-01
allow 2a03:2267:4e6f:7264:46d9:e7ff:fe74:5a5b;
#eae-segeberg-ap2
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:a998;

#KelliNet-009
allow 2a03:2267:4e6f:7264:32b5:c2ff:fed5:c96c;
#KelliNet-005
allow 2a03:2267:4e6f:7264:16cc:20ff:fe2b:7568;

#KelliNet-002
allow 2a03:2267:4e6f:7264:c66e:1fff:feb3:967e;
#KelliNet-001
allow 2a03:2267:4e6f:7264:6a72:51ff:fe50:a341;
#KelliNet-006
allow 2a03:2267:4e6f:7264:1aa6:f7ff:fe3d:f346;
#KelliNet-008
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe3f:6394;

#MBD_Freifunk_Ehndorf
allow 2a03:2267:4e6f:7264:a62b:b0ff:feca:8aa4;
#MBD_Freifunk_Ehndorf2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:a3de;

#Ferienhof-Struve-11-c
allow 2a03:2267:4e6f:7264:ea94:f6ff:fe67:e9c6;
#Ferienhof-Struve-14
allow 2a03:2267:4e6f:7264:92f6:52ff:fec4:b47f;
#Ferienhof-Struve-11-d
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:21aa;
#Ferienhof-Struve-13a
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe3e:8dd6;
#Ferienhof-Struve-11-d
allow 2a03:2267:4e6f:7264:16cc:20ff:fecd:21aa;
#Ferienhof-Struve-11b
allow 2a03:2267:4e6f:7264:92f6:52ff:fec4:a9f6;
#Ferienhof-Struve-16
allow 2a03:2267:4e6f:7264:fa1a:67ff:fe3e:8bd4;
#Ferienhof-Struve-15
allow 2a03:2267:4e6f:7264:92f6:52ff:fec4:b2c6;
#Ferienhof-Struve-11a
allow 2a03:2267:4e6f:7264:6666:b3ff:fe60:6f23;

#FF-DIT-Albersdorf-Freizeitbad2
allow 2a03:2267:4e6f:7264:6a72:51ff:fe44:da41;
#FF-DIT-Albersdorf-Freizeitbad1
allow 2a03:2267:4e6f:7264:6a72:51ff:fe4e:28ba;

#FNF_Corcuscant
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3e:7146;
#FNF_Alderaan
allow 2a03:2267:4e6f:7264:46d9:e7ff:feb7:e715;

#schiri012
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:dd34;
#freifunk_hagen_sk_09
#allow 2a03:2267:4e6f:7264:62e3:27ff:fe60:6536;
#freifunk_hagen_sk_06
#allow 2a03:2267:4e6f:7264:62e3:27ff:febd:be20;
#freifunk_hagen_sk_04
#allow 2a03:2267:4e6f:7264:16cc:20ff:feb1:f38;

#FF-Elpersb-HausDithmarschen
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe85:acee;
#FF-Elpersb-HausDithmarschen-02
allow 2a03:2267:4e6f:7264:c66e:1fff:fea7:52f2;
#FF-Elpersb-HausDithmarschen-01
allow 2a03:2267:4e6f:7264:32b5:c2ff:fef0:987a;

#ff-nf-spo-fewo-carstens-1
allow 2a03:2267:4e6f:7264:ee08:6bff:fe78:73c8;
#ff-nf-spo-fewo-carstens-2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe8a:bb9c;
#FF-RD_BI-26B
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d1d0;

#FF-HB-Asch-Huett-10
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe3f:d1cc; 

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

#MTV-von-1860-ev-Heide
allow 2a03:2267:4e6f:7264:8616:f9ff:fec8:a5c4;
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

#FB-2
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3c1e;
#FB-1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:3d6e;
#FB-3
allow 2a03:2267:4e6f:7264:f6f2:6dff:fedc:388c;

#behard
allow 2a03:2267:4e6f:7264:a62b:b0ff:fed7:1eb2;

#FF_SL_Husumer_Str_001
allow 2a03:2267:4e6f:7264:eade:27ff:fefd:2cce;

#FF-DIT-Kappeln1
allow 2a03:2267:4e6f:7264:f6f2:6dff:fe52:9710;
#FF-DIT-Kappeln2
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:15c4;
#FF-DIT-Kappeln3
allow 2a03:2267:4e6f:7264:ee08:6bff:fe69:15be;

#Etoile-V
allow 2a03:2267:4e6f:7264:46d9:e7ff:fedb:fb24;
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
  #FF_Windbergen_Wodansberg
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe54:c26e;
  #FF_Windbergen_Wodansberg_2
  allow 2a03:2267:4e6f:7264:f6f2:6dff:fe54:c28e;
  #FF_Windbergen_Wodansberg_3
  allow 2a03:2267:4e6f:7264:c66e:1fff:fe2d:a208;

  #Pension-Galerie_Vierhoefen_03
  allow 2a03:2267:4e6f:7264:c66e:1fff:fe63:4724;

  #DieMuehle
  allow 2a03:2267:4e6f:7264:a62b:b0ff:fed1:e178;

#NDS-FFNH-123-NeuWu-Courage
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:80ae;
#NDS-FFNH-125-NeuWu-Rathaus-NW
allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:814e;
  #NDS-FFNH-120-NeuWu-Rathaus-SO
  allow 2a03:2267:4e6f:7264:a62b:b0ff:fefe:8178;

#Faehrhaus_Hodorf1
allow 2a03:2267:4e6f:7264:16cc:20ff:fec0:177e;
  #Faehrhaus_Hodorf2
  allow 2a03:2267:4e6f:7264:6a72:51ff:fe36:c0c;

  #FF-Hanstedt-12
  allow 2a03:2267:4e6f:7264:62e3:27ff:feed:86fc;

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #NDS-FFNH-0509-Bispingen-CPE1
  allow 2a03:2267:4e6f:7264:1aa6:f7ff:feca:8514;

#ferienhof_moeller
allow 2a03:2267:4e6f:7264:62e3:27ff:fec7:77c;
#ferienhof_moeller_2
allow 2a03:2267:4e6f:7264:62e3:27ff:fe9a:e3ee;

  #ferienhof_moeller_4
  allow 2a03:2267:4e6f:7264:eade:27ff:fe37:e090;

  #NDS-FFNH-0145-Bispingen-AKR2
  allow 2a03:2267:4e6f:7264:ee08:6bff:fea4:ccf4;

  #freifunk-nazar-kebabhaus
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee6:fec0;

  #FF-HB-161
  allow 2a03:2267:4e6f:7264:62e3:27ff:fee7:43a8;

#Norbi
allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d6ba;
  #FressflashRuftDasKaenguru
  allow 2a03:2267:4e6f:7264:a2f3:c1ff:fe48:d73e;

  ####### Unterhalb keine Änderungen durchführen! ###############
  
  root /opt/www;
  autoindex on;
  deny all;
  index index.html index.htm;
  server_name localhost;
}
