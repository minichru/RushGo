# RushGo - HTTP Wrapper For Go

RushGo is a simple, yet effective HTTP client wrapper library designed to streamline HTTP request handling in Go applications. While initially developed because of pain points I had with the standard library, it encapsulates commonly needed functionalities, making it a potentially useful tool for other Go developers seeking a straightforward approach to HTTP communications.

### Note on Maintenance and Contributions
- This library is actively maintained as a public project.
- If you find RushGo useful or have suggestions for improvements, feel free to open an issue in the repository. Community input is always welcome and can inspire further development or public maintenance.

### Features
- RushGo offers a variety of features to manage HTTP requests with ease:
  - Simplified methods for GET and POST requests (Will add moe methods soon).
  - Convenient functions to set custom headers, cookies, and user agents.
  - Support for HTTP/2.
  - Customizable timeout settings for requests.
  - Utility functions to streamline common tasks such as downloading files, and generating random user-agents.


### To-Do List / Future Directions
- [ ] Implement additional HTTP methods (PUT, DELETE, PATCH, etc.).
- [ ] Extend support for parsing different data formats (e.g., YAML, TOML).
- [ ] Integrate advanced error handling and logging mechanisms.
- [ ] Add support for HTTP/3.
- [ ] Add support for proxies.
- [ ] Add support for asynchronous requests.
- [ ] Incorporate OAuth and other authentication protocols.
- [ ] Enhance user-agent randomization features.
- [ ] Integrate TLS support.
- [ ] Create comprehensive documentation and usage examples.
- [ ] Optimize performance for large-scale applications.

### Future Directions
- Keep an eye on the repository for any updates or enhancements.
- Wehter this is useful or not isn't something I thought about, It was made for my own personal use which I then have released to others. It is VERY incomplete, but that will change in the future.

### Acknowledgments
- RushGo is a personal project developed to address specific needs in Go-based HTTP handling. It is inspired by the simplicity and utility of existing HTTP client libraries but tailored to personal preferences and use cases.
