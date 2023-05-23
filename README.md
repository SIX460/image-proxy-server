# Image-Proxy-Server

![banner](banner.png)

`Image-Proxy-Server` is a Go-based server that proxies image requests. Given an image URL, it fetches the image, transforms the URL into an MD5 hash for tracking, stores the image (either in-memory or on disk), and serves it back to the requester. If the image has already been requested, the server will serve the image from memory or disk, depending on the configuration settings.

This server can handle requests from a variety of sources, including but not limited to `LuisaViaRoma`, and is designed to handle infinite pending/redirect requests efficiently. For example you can let your Discord monitor have always the images loaded :P

## Installation

To install and run the Image-Proxy-Server, you'll need to have Go installed on your system. If you haven't installed Go yet, you can download it from the [official website](https://golang.org/dl/).

Once you have Go installed, clone this repository to your local machine:

```bash
git clone https://github.com/glizzykingdreko/image-proxy-server.git
cd image-proxy-server
```

## Usage

To start the server with default settings, run:

```bash
go run main.go
```
By default, the server will start on port 3777 and will store images on disk in the ./img directory.

To proxy an image request, simply send a GET request to the running server:
```bash
http://localhost:3777/proxy?url=https://images.lvrcdn.com/BigRetina77I/3FL/006_597aaf87-6887-49d6-892d-39bded9f3693.JPG
```
The format is as follows:
```bash
http://localhost:3777/proxy?url=<image_url>
```

### Command Line Flags

- `-in-memory`: If set to `true`, the server will store images in memory instead of writing them to disk.
  Example: 
  ```bash
  go run main.go -in-memory=true
  ```
- `-port`: Allows you to specify the port number the server will listen on. The default port is `3777`.
  
  Example: 
  ```bash
  go run main.go -port=8080
  ```

## Contact

If you have any questions, suggestions, or just want to say hi, you can reach out to me on the following platforms:

- Twitter: [@glizzykingdreko](https://twitter.com/glizzykingdreko)
- Medium: [@glizzykingdreko](https://medium.com/@glizzykingdreko)
- Github: [@glizzykingdreko](https://github.com/glizzykingdreko)
- Mail: [send mail](mailto:glizzykingdreko@protonmail.com)

I look forward to hearing from you!

## License
This project is licensed under the MIT License. See the [LICENSE](/LICENSE) file for details.