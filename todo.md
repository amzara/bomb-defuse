// set up http server 
// set up bomb service logic to generate a random pin
// set up bomb service endpoint to start bomb, take in the random pin, and store timer and random memory
// set up bomb service logic to explode
// set up bomb service endpoint to allow other service to communicate
// set up defuser service to communicate with the bomb service 
// set up mutex to prevent bomb defusers from spamming bomb defuse, only 1 attempt per sec
// auto explode bomb if more than 3 tries
// at app start, set up bomb and user can only interact with defuser service


//small mini game i made for fun on a sunday evening to revise goroutines, channel and mutex 