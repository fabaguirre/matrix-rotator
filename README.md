# Matrix Rotator API 🚀

## Description

API that receives a matrix, performs QR decomposition, rotates the matrix, and forwards the result to [another API](https://github.com/fabaguirre/matrix-stats) for statistical analysis. 🧩

## Installation

1. Clone the repository:

   ```bash
   git clone git@github.com:fabaguirre/matrix-rotator.git
   ```

2. Navigate to the project directory:

   ```bash
   cd matrix-rotator
   ```

3. Copy and configure the environment variables:

   ```bash
   cp .env.example .env
   ```

   Edit the `.env` file to fill in the necessary environment variables. The `PORT` variable defaults to an empty string, so it must be set in the `.env` file for the application to function properly. 📑

4. Build the Go application:

   ```bash
   go build -o bin/app ./cmd
   ```

## Usage

1. Start the application:

   ```bash
   ./bin/app
   ```

   By default, the application will be available at `http://localhost:4000`. 🚀

2. Make requests to the API:

   - **POST /api/matrix/rotate**

     Request body:

     ```json
     {
       "matrix": [[1, 2], [3, 4]]
     }
     ```

     Response:

     ```json
     {
       "qrDecomposition": {
         "Q": [[...], [...]],
         "R": [[...], [...]]
       },
       "rotatedMatrix": [[...], [...]],
       "stats": {...}
     }
     ```

## Environment Variables

- **PORT:** The port where the server will listen. Default is `4000`. 🎯
- **API_URL:** The URL of the [API](https://github.com/fabaguirre/matrix-stats/) for statistical analysis.

## Docker

1. Build the Docker image:

   ```bash
   docker build -t matrix-rotator .
   ```

2. Run the Docker container:

   ```bash               
   docker run -d --name matrix-rotator-fiber --network my-network -p 4000:4000 -e PORT=4000 -e API_URL=http://matrix-stats-express:3000/api/ matrix-rotator
   ```