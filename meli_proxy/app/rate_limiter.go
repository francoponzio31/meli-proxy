package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RateLimiter(endpoint string) gin.HandlerFunc {
	return func(c *gin.Context) {

		config := LoadConfig()
		MaxRequests := int64(config.RateLimiter)
		TimeWindow := time.Duration(config.TimeWindow) * time.Second

		clientIP := c.ClientIP()
		currentTime := time.Now().Unix()
		key := fmt.Sprintf("ratelimit:%s:%s", clientIP, endpoint)

		// Delete requests outside the time window
		RedisClient.ZRemRangeByScore(RedisCtx, key, "0", fmt.Sprint(currentTime-int64(TimeWindow)))

		requestCount, err := RedisClient.ZCard(RedisCtx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		if requestCount >= MaxRequests {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded. Try again later."})
			c.Abort()
			return
		}

		// Register the new request with the current timestamp
		_, err = RedisClient.ZAdd(RedisCtx, key, &redis.Z{Score: float64(currentTime), Member: currentTime}).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering request in Redis"})
			c.Abort()
			return
		}

		RedisClient.Expire(RedisCtx, key, TimeWindow)

		c.Next()
	}
}
