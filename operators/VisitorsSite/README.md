# Visitors site
The Visitors Site tracks information about each request to its home page. 
Each time the page is refreshed, an entry is stored with details about the client, 
backend server, and timestamp. The home page displays a list of the most recent visits.
![Visitors site](images/VisitorsSite.png)

## run
- kubectl apply -f database.yaml
- kubectl applay -f backend.yaml
- kubectl apply -f frontend.yaml
- http://localhost:30686/


