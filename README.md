# astro

## ephimeris server

Get shell access to server.
Binary is  ./astroBin/astrobin

To see table of all planets with house information
```
./astrobin -b4.2.1978 -ut05:15:00 -geopos80.27,13.08,0 -n1 -s1 -sid1 -fPl -pd -eswe -house
```

To see table of mercury with declination and longitude
```
./astrobin -b1.1.2026 -ut05:15:00 -geopos80.27,13.08,0 -n60 -s15 -p2 -fTlbDr
```

To see the table of sun with declination and longitude
```
./astrobin -b1.1.2026 -ut05:15:00 -geopos80.27,13.08,0 -n60 -s15 -p0 -fTlbrD
```

## panchang

Panchang server is part of combined astro binary

For UI panchang, we create static json files using
```
cd ./scripts/panchangJson
python getPanchang.py
```

```
curl -X GET http://localhost:9080/v1/panchang/getMonth?birthData.tob.d=1&birthData.tob.m=10&birthData.tob.y=2025&birthData.pob.city=Chennai&birthData.pob.state=TN&birthData.pob.country=IN&birthData.pob.lat=13.8784&birthData.pob.long=80.27847
```

Check ephimeris using docker
```
docker run --rm -it --entrypoint /astroBin/astrobin dronasys-com/astro-ephimeris -sid1 -b14.7.2026 -ut06:45:00 -geopos80.699997,27.080000,0 ;5D-n1 -s1 -pd -eswe -house -fPlbrds
```