// sending the request to orchestrator server for creating user space folder
data := url.Values{}
data.Set("username", "myuser")

// 2. Target URL
apiUrl := "http://localhost:3003/user_creation"

// 3. Create request with encoded body
req, _ := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))

// 4. Set mandatory header
req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 5. Send
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
	fmt.Println("Error:", err)
	return
}
defer resp.Body.Close()

fmt.Println("Status:", resp.Status)