# Real-time Leaderboard

This is a real-time leaderboard system built in Go using Redis and PostgreSQL. The application allows users to register, submit scores, and view leaderboards in real-time, with scores being stored in Redis sorted sets.

## Features

- User Authentication (using JWT)
- Submit Scores for various games/activities
- Real-time Leaderboard Updates
- View Global Leaderboard
- View User's Rank on Leaderboard
- Automatic Database Migrations on Startup

## Technologies Used

- Backend:  Golang
- Database: PostgreSQL and Redis (Sorted Set)
- Real-time: Redis
- go-redis package for Redis connection
- go-sqlx  package for PostgreSQL

## This is Roadmap.sh Project solution in Golang Roadmap
   ```bash
       https://roadmap.sh/projects/realtime-leaderboard-system
   ```


## Setup

1. Clone the repository:
    ```bash
    git clone  github.com/rupak26/Real-time-Leaderboard
    ```
2. Install dependencies:
    ```bash
    cd Real-time Leaderboard
    
    Create a .env file in the root of the project with the following content:
        VERSION = 1.0.0 
        SERVICE_NAME = REAL-TIME-LADERBORD 
        HTTP_PORT = 8080

        DB_HOST=localhost
        DB_PORT=5432
        DB_USER=postgres
        DB_PASSWORD=12345678
        DB_NAME=Real_Time_LaderBoard
        ENABLE_SSL_MODE = false



        # DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${ENABLE_SSL_MODE}


        JWT_SECRET="a4bc166839d6ff3c83eb9d1cbb0ddda2a65f74eeea9a07413357200dfad1d808d708e7d5"


        REDIS_ADDR=localhost:6379
        REDIS_PASSWORD=""
        REDIS_DB=0
    ```

4. Install Dependencies 
   ```
      go mod download
   ```
3. Start the development server:
    ```bash
    go run main.go
    ```

## API Endpoints

1. Register User 
   POST users/

   Request:
   ```
      {
        "username" : "jhon" , 
        "email" : "jhon@gmail.com" ,
        "password" : "1234"
      }
   ```
2. Login User
   POST users/login/

   Request:
   ```
      {      
        "email": "jhon@gmail.com",
        "password": "1234"
      }
   ```
   Response (Success):
   ```
       {
         "status": 201,
         "message": "Access Token",
         "data": "your_jwt_token_here"
       }
   ```
3. Submit Score
   POST /submit-score

   Request:
   ```
      {
        "game_id" : "cricket__12" ,
        "score"   :  1600
      }
   ``` 

4. Get Global Leaderboard 
   GET /leaderboard?limit=10
   
   Response(Success):
   ```
     {
        "datalist": [
            {
                "user_name": "Rupak",
                "score": 1600,
                "rank": 1
            }
        ],
        "Pagination": {
            "limit": 5
        }
    }
   ```

5. Get User Ranking
   GET /user-ranking/{userId}
   ```
    {
        "user_name": "jhon",
        "score": 1600,
        "rank": 1
    }  
   ```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.