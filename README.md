# label-artist-fetcher
This is an API that fetches what labels and artists are being represented at three different Sveriges Radio channels: `p1`, `p2` and `p3`. 

To run this API, please make sure you have Docker installed. Run the API by running: `docker-compose up --build` in the root. This should start the API on port :3000. 

## API Documentation
This API contains one endpoint: to fetch labels and their respective artist that are played at a specific radio channel during a specific time period at Sveriges Radio. 

## GET /labels?channel={channel}&from={fromTimestamp}&to={toTimestamp}
Returns an array of label names and their respective artists, that have been played at the radio channel of choice {channel} during the time period specified {from} and {to}. 

+ Limitations:
   + the API is limited to fetching only 100 items. 

+ Parameters
    + channel: `p3` (required, string) - choose from `p1`/`p2`/`p3`
    + from: `from` (required, UNIX timestamp)
    + to: `to` (required, UNIX timestamp)

+ Request
    + see Parameters above

+ Response 200
    + Headers

            Content-Type:application/json
+ Body
    + map of label names and artists belonging to that label.
   ```json 
     [
        {
          "label_name":"Ador",
          "artists":[{
            "name":"Newjeans"}
          ]},
        {
          "label_name":"Atlantic",
          "artists":[{
            "name":"Jay-Z, Alicia Keys"}
          ]}
     ]
   ```
      
+ Response 400
  Faulty request, faulty input variables
      
+ Response 500
  Failed to retrieve information.

+ Example query: 
  `http://localhost:3000/labels?channel=p3&to=1691413764&from=1691320112`
