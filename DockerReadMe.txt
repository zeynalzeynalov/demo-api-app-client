//building docker image
docker build . -t taxdooclientwithgo

//running docker image
docker run -p 8080:8080 taxdooclientwithgo
