# TDEE API
This is a simple API that calculates the Total Daily Energy Expenditure (TDEE) of a person based on their age, weight, height

## Installation
1. Clone the repository
2. Install the dependencies
    ```bash
    go mod download
    ```
3. Run the server
    ```bash
    go run main.go
    ```

## Docker
1. Build the image
    ```bash
    docker build -t tdee-api .
    ```
2. Run the container
    ```bash
    docker run -p 3000:3000 tdee-api
    ```

## Usage
Send a GET request to the `/tdee/daily` endpoint with the following URL parameters:
```http request
GET http://localhost:3000/api/v1/tdee/daily
    ?age=25
    &weight=
    &height=
    &sex=male
    &activity_level=
```
Weight is in pounds, height is in inches. The activity level can be one of the following:
- sedentary
  - little or no exercise
- light_active
  - light exercise/sports 1-3 days/week
- moderate_active
  - moderate exercise/sports 3-5 days/week
- active
  - hard exercise/sports 6-7 days a week
- very_active
  - very hard exercise/sports & physical job or 2x training

## Example
```http request
GET http://localhost:3000/api/v1/tdee/daily
    ?age=25
    &weight=180
    &height=72
    &sex=male
    &activity_level=active
```
The response will be:

```json
{
  "tdee": 3200,
  "deficits": {
    "deficit_250": 2950,
    "deficit_500": 2700,
    "deficit_750": 2450,
    "deficit_1000": 2200
  }
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)


