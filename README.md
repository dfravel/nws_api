### National Weather Service API Proof of Concept

API written in Go that queries the National Weather Service endpoint. 
This was written as part of a coding challenge by CrossnoKaye for a 
full-stack developer role, but I'm planning to continue updating it 
since the NWS API has a ton of information available within it. 

**To run this locally:**

1. clone the project from GitHub (no uploaded yet)
2. signup for a geocoding account at [ZipCodes.com](https://zipcodes.com)
3. update config.yaml (Geocode --> AUTHTOKEN) with the authToken you receive from zipcodes.com
4. run `go mod tidy` at the command line to make sure you have all of the dependencies installed 
   1. note: I'm honestly not sure if this is a required step, but it couldn't hurt
5. run `./dev` or `go run main.go` from the command line
6. the API should start on port 9999


**Areas for Improvement**
1. tests - the CK team suggested this challenge would take 4-5 hours. It didn't and I cut a few corners to fit this in with my regular job and responsibilities
2. better error handling - right now I"m returning a simple `error` object
3. if I had more time, I'd play with **Goa**, but based on the suggested time constraint, I went with what I know
4. I'm sure there are a dozen other ways to make this better. Hopefully I'll have the time to do that. 


