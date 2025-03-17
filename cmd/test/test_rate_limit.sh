for i in {1..110}; do 
  response=$(curl -s -w "%{http_code}" -X POST http://localhost:8080/api/urls \
    -H "Content-Type: application/json" \
    -d '{"long_url":"https://github.com/bharabhi01"}' -o /dev/null)
  echo "Request $i: HTTP Status $response"
  if [ "$response" -eq 429 ]; then
    echo "Rate limit triggered after $i requests!"
    break
  fi
done