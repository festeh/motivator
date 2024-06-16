# Motivator

Motivator is a Go application that checks the status of the "Distract Me Not" extension in Google Chrome. 

## Features

- Checks if the "Distract Me Not" extension is enabled in Google Chrome.
- Outputs the status of the extension in JSON format.

## Usage

To run the application, use the following command:

```bash
go run main.go
```

This will output the status of the "Distract Me Not" extension. If the extension is enabled, the output will be `{"DmnOn":true}`. If the extension is disabled or not found, the output will be `{"DmnOn":false}`.

## License

[MIT](https://choosealicense.com/licenses/mit/)
