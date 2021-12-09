docker build -t logistics .
docker run -di -v /app/logistics:/logistics/logistics -p 8000:8000 logistic