# AlagouAPI

## Running with Docker
- Build an image as **development** mode
  - `docker build -t alagou-api --build-arg app_env=development .`
- Build an image as production mode
  - `docker build -t alagou-api .`
- Run the generated image
  - `docker run alagou-api`
- Test it with _Status_ endpoint
  - `http://localhost:8080/status`
