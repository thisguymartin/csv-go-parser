# csv-go-parser


Simple go script that converts csv file into a json document.

CSV Input: 

```
id,first_name,last_name,email,avatar,ip_address
1,Pauline,Hirth,phirth0@nationalgeographic.com,https://robohash.org/doloremdignissimosoccaecati.png?size=50x50&set=set1,199.109.72.212
2,Gilly,Challicombe,gchallicombe1@altervista.org,https://robohash.org/voluptatumomnisaliquid.png?size=50x50&set=set1,14.139.200.228
3,Laverna,Ratledge,lratledge2@com.com,https://robohash.org/quiavoluptatemplaceat.png?size=50x50&set=set1,22.254.66.241
4,Clary,Forbes,cforbes3@wordpress.com,https://robohash.org/temporequidolor.png?size=50x50&set=set1,234.201.73.224
5,Alia,Bilyard,abilyard4@boston.com,https://robohash.org/eumetneque.png?size=50x50&set=set1,172.242.193.201
6,Giacobo,Hulle,ghulle5@etsy.com,https://robohash.org/autemodioporro.png?size=50x50&set=set1,177.34.216.2

```

JSON output:

```json
[
  {
    "last_name": "last_name",
    "email": "email",
    "avatar": "avatar",
    "ip_address": "ip_address"
  },
  {
    "last_name": "Hirth",
    "email": "phirth0@nationalgeographic.com",
    "avatar": "https://robohash.org/doloremdignissimosoccaecati.png?size=50x50&set=set1",
    "ip_address": "199.109.72.212"
  },
  {
    "last_name": "Challicombe",
    "email": "gchallicombe1@altervista.org",
    "avatar": "https://robohash.org/voluptatumomnisaliquid.png?size=50x50&set=set1",
    "ip_address": "14.139.200.228"
  },
  {
    "last_name": "Ratledge",
    "email": "lratledge2@com.com",
    "avatar": "https://robohash.org/quiavoluptatemplaceat.png?size=50x50&set=set1",
    "ip_address": "22.254.66.241"
  },
  {
    "last_name": "Forbes",
    "email": "cforbes3@wordpress.com",
    "avatar": "https://robohash.org/temporequidolor.png?size=50x50&set=set1",
    "ip_address": "234.201.73.224"
  },
  {
    "last_name": "Bilyard",
    "email": "abilyard4@boston.com",
    "avatar": "https://robohash.org/eumetneque.png?size=50x50&set=set1",
    "ip_address": "172.242.193.201"
  },
  {
    "last_name": "Hulle",
    "email": "ghulle5@etsy.com",
    "avatar": "https://robohash.org/autemodioporro.png?size=50x50&set=set1",
    "ip_address": "177.34.216.2"
  }
]

```
