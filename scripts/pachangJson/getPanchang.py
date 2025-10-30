import requests

base_url = "http://localhost:4200/api/v1/panchang/getMonth?birthData.tob.d=1&birthData.tob.m=10&birthData.tob.y=2025&birthData.pob.city=Chennai&birthData.pob.state=TN&birthData.pob.country=IN&birthData.pob.lat=13.8784&birthData.pob.long=80.27847"

for year in range(2025,2029):
    for month in range(1, 13):
        url = f"http://localhost:4200/api/v1/panchang/getMonth?birthData.tob.d=1&birthData.tob.m={month}&birthData.tob.y={year}&birthData.pob.city=Chennai&birthData.pob.state=TN&birthData.pob.country=IN&birthData.pob.lat=13.8784&birthData.pob.long=80.27847"
        #url = f"{base_url}{i}"
        response = requests.get(url)
        response.raise_for_status()  # To raise error if the request failed
        json_data = response.json()
        filename = f"panchang/data_{year}_{month}.json"
        with open(filename, 'w') as f:
            import json
            json.dump(json_data, f, indent=4)
        print(f"Saved JSON from {url} to {filename}")
