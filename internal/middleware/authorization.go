package middleware

import (
	"context"
	"net/http"
	"crypto/hmac"
	"crypto/sha256"
	"strings"
	"encoding/json"
	"github.com/rupak26/Real-time-Leaderboard/utils"
)

func (m *Middlewares) Authorization(next http.Handler) http.Handler {
	 return http.HandlerFunc(func (w http.ResponseWriter , r *http.Request) {
		// parse jwt 
		// parse header and payload or claim 
		// hmac-sha256 algorithm -> hash( header , payload , secret_key )
		// parse signature part from jwt 
		// if signature and hash is same => forward create products 
		// otherwise 401 status code with Unauthorize
		
	
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w , "Unauthorized" , http.StatusUnauthorized) 
			return 
		}
		
		headerArry := strings.Split(header , " ") 

		if len(headerArry) != 2 {
			http.Error(w , "Unauthorized" , http.StatusUnauthorized) 
			return 
		}
		
		accessToken := headerArry[1] 

		tokenParts := strings.Split(accessToken , ".") 
		
		if len(tokenParts) != 3 {
			http.Error(w , "Unauthorized" , http.StatusUnauthorized) 
			return 
		}

	// fmt.Println("-------------------tokenParts---------------" , tokenParts)

		jwtHeader := tokenParts[0] 
		jwtPayload := tokenParts[1] 
		signature := tokenParts[2]
	
	// Need to chage this	
		cnf := m.cnf
		key := cnf.SecretKey 

		message := jwtHeader + "." + jwtPayload 
		byteKey := []byte(key)
		byteMessage := []byte(message) 
		
		h := hmac.New(sha256.New , byteKey)
		h.Write(byteMessage)

		hash := h.Sum(nil)
		signatureB64 := utils.Base64UrlEncode(hash) 

		if signature != signatureB64 {
			http.Error(w , "Unauthorized" , http.StatusUnauthorized) 
			return 	
		}
        
		payloadBytes, err := utils.Base64UrlDecode(jwtPayload)
		if err != nil {
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		var claims map[string]interface{}
		if err := json.Unmarshal(payloadBytes, &claims); err != nil {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
        
        userIdFloat , okk := claims["user_id"].(float64) 

		if !okk {
			http.Error(w, "Invalid or missing user_id in token", http.StatusUnauthorized)
			return
		}
        userId := int(userIdFloat)
		userName, ok := claims["user_name"]
		if !ok {
			http.Error(w, "Invalid or missing user_name in token", http.StatusUnauthorized)
			return
		}

		

		// ✅ Put user_id in context
		ctx := context.WithValue(r.Context(), "user_id", userId) 
		ctx = context.WithValue(ctx, "user_name", userName)

		r = r.WithContext(ctx)

		next.ServeHTTP(w , r)
	})
}