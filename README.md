# ASCII-ART-WEB
## Description

ASCII-ART-WEB is a web application written in Go that allows users to generate ASCII art from text through a graphical web interface. Users can enter text, select one of the available banners (`standard`, `shadow`, or `thinkertoy`), and view the generated ASCII art directly in their browser.
The application uses HTTP handlers and HTML templates to provide an interactive user experience.
Author

* Samuel Ojedo
Usage
Requirements

* Go installed on your machine.
How to Run

1. Clone the repository:

```bash
git clone <repository-url>
cd ascii-art-web

```

1. Run the server:

```bash
go run .

```

1. Open your browser and visit:

```text
http://localhost:8080

```

1. Enter text, choose a banner, and click the submit button to generate ASCII art.
Implementation Details
Algorithm

1. The server starts and listens on port `8080`.
2. A `GET /` request serves the main HTML page containing:
   * A text input field.
   * Banner selection options.
   * A submit button.
3. When the user submits the form, a `POST /ascii-art` request is sent to the server.
4. The server:
   * Retrieves the text and selected banner from the form.
   * Validates the request.
   * Reads the corresponding banner file.
   * Converts the input text into ASCII art.
5. The generated ASCII art is rendered and displayed on the webpage.
6. Appropriate HTTP status codes are returned:
   * `200 OK` for successful requests.
   * `400 Bad Request` for invalid input or request method.
   * `404 Not Found` when templates or banner files are missing.
   * `500 Internal Server Error` for unexpected server errors.
Project Structure

```text
.
├── main.go
├── handlers.go
├── templates/
│   └── index.html
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── README.md

```
