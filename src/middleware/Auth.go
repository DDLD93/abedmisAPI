package middleware

import (
	"errors"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type Header struct {
	API_KEY string `header:"api-key"`
}
type Keys struct {
	PSP       string
	Test_Keys string
	Live_Keys string
}

var PSP = map[string]Keys{
	"upperLink": {
		PSP:       "upperLink",
		Test_Keys: "test_hbbIVvbjhkBKJBkbKKJBjkbkbkjbKBJKbjhvjhVJKHVjkhvkjhvb",
		Live_Keys: "live_gjhGjkgJKGjklGJKgUIGUIfgvyBGVftrDSeaweSdrytfuiPOIoHG",
	},
	"visualICT": {
		PSP:       "visualICT",
		Test_Keys: "test_OUYYGtFRTwserwserxCVgjvbKljIOhuGYdfrtDRWSErdrtDYFG",
		Live_Keys: "live_RRTGghhjJKKLMmbBvccwqQwEEfvbbbhbHJppOObBCcCddDf",
	},
	"unifiedPayment": {
		PSP:       "unifiedPayment",
		Test_Keys: "live_BBuUuiKmMNjVGHFcxdrSrstEqWAserDYfJHopkmnBNnVcFGXd",
		Live_Keys: "test_jjkkuurrrryVVUbiNMKmppNnCCxSAQWeRtYYhnjNJnOJojJojJQ",
	},
}

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()
		if len(c.GetHeader("api-key")) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "Require auth key",
				"Payload": nil,
				"Error":   errors.New("auth key not found"),
			})
			return
		}
		for _, v := range PSP {
			if c.GetHeader("api-key") == v.Live_Keys {
				c.Request.Header.Add("mode", "live")
				c.Request.Header.Add("PSP", v.PSP)
				c.Next()
				latency := time.Since(t)
				c.Writer.WriteString(latency.String())
				return

			}
			if c.GetHeader("api-key") == v.Test_Keys {
				c.Request.Header.Add("mode", "test")
				c.Request.Header.Add("PSP", v.PSP)
				c.Next()
				latency := time.Since(t)
				c.Writer.WriteString(latency.String())				
				return
			}

		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Invalid api key",
			"Payload": nil,
			"Error":   errors.New("invalid api key"),
		})
	}
}
