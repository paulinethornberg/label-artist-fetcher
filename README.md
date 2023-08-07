# label-artist-fetcher
fetching labels and their respective artists 

## API Documentation
This API contains one endpoint: to fetch labels and their respective artist that are played at a channel at a specific time period. 

### GET /labels?channel={channel}&from={fromTimestamp}&to={toTimestamp}
Returns a map of labels and their respective artists, that have been played at the radio channel of choice {channel} during the time period specified {from} and {to}. 

+ Limitations:
   + the API is limited to fetching only 100 items. 

+ Parameters
    + channel: `p3` (required, string) - choose from `p1`/`p2`/`p3`
    + from: `from` (required, UNIX timestamp)
    + to: `to` (required, UNIX timestamp)

+ Request

    + Body
        + map of label names and artists belonging to that label. 
       ```json //TODO: UPDATE
       [
            
       ]
       ```

+ Response 200
    + Headers

            Content-Type:application/json
    + Body
        + map of label names and artists belonging to that label.
       ```json //TODO: UPDATE
       [
          
       ]
       ```
      
+ Response 400
  Faulty request, faulty input variables
      
+ Response 500
  Failed to retrieve information.


