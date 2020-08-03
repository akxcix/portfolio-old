# Portfolio
Source code for my portfolio
[website](https://iamadarshk.com).

## Running
Using [docker](https://docker.com):
```bash
docker build -t portfolio-server .

# PORT is the enviorment variable containg number of the desired port
docker run -v ~/tls:/build/tls -p ${PORT}:8080 portfolio-server
```

## License
The code is released under the MIT License. See `LICENSE` for more details.

## Acknowledgements
I have used:
- [Hover.css](http://ianlunn.github.io/Hover/) for UI element effects.
- Icons from [Font Awesome](https://fontawesome.com/).
