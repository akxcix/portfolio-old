# Portfolio
Source code for my portfolio
[website](https://iamadarshk.com).

## Running
Most convinient way is to use [docker](https://docker.com):
- Command Line Arguments to go program passed as Environment Variables to
  docker compose:
  - | ARGV  | ENV | Default | Description |
    | :---: | :---: | :---: | :---: |
    | -addr | ADDR | 8080 | address (port) for the server to listen on |
    | -tls  | TLS | false | specifies whether to use TLS server or not |
    | -     | TLS_DIR | ./tls | location of directory containing cert.pem and key.pem |
    ##### Note: `-` in `ARGV` column means that the Environment Variable is used only for `docker-compose`
- All the Command Line Arguments are passed as Environment Variables using 
`docker-compose`
 (could be passed using a separate `.env` file)
- ```bash
  docker-compose up --build
  ```

## License
The code is released under the MIT License. See `LICENSE` for more details.

## Acknowledgements
- [Hover.css](http://ianlunn.github.io/Hover/) for UI element effects.
- [Font Awesome](https://fontawesome.com/) for Icons.

