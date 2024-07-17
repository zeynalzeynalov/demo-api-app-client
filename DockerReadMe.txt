//building docker image
docker build . -t demoapiclientwithgo

//running docker image
docker run -p 8080:8080 demoapiclientwithgo
