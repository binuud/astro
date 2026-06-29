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
./astrobin -b1.1.2026 -ut05:15:00 -geopos80.27,13.08,0 -n60 -s15 -p2 -fTlbr
```