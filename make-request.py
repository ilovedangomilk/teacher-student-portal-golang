import requests

url = 'http://localhost:8080/register'
data = {'studentId': 1} # Replace 1 with the actual StudentID

response = requests.post(url, json=data)

print(response.text)
