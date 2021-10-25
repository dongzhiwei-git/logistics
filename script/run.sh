docker build -t logistics .
docker run -v /nginx/logistics /logistics -p 8000:8000 logistics