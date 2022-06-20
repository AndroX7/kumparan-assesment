package middleware

import (
	"net/http"

	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stoewer/go-strcase"
)

func (m *Middleware) ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Only run if there are some errors to handle
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Find out what type of error it is

				switch e.Type {
				case gin.ErrorTypePublic:
					// Only output public errors if nothing has been written yet
					if !c.Writer.Written() {
						// check if it is part of custom error
						if err, ok := e.Err.(errors.CustomError); ok {
							// print the underlying error and return the specified message to user
							c.JSON(err.HTTPCode, gin.H{"errors": err.Message})
						} else {
							c.JSON(c.Writer.Status(), gin.H{"errors": e.Error()})
						}

					}
				case gin.ErrorTypeBind:
					errs, ok := e.Err.(validator.ValidationErrors)
					if ok {
						list := make(map[string]string)
						for _, err := range errs {
							list[strcase.SnakeCase(err.Field())] = validationErrorToText(err)
						}

						// Make sure we maintain the preset response status
						status := http.StatusUnprocessableEntity
						if c.Writer.Status() != http.StatusOK {
							status = c.Writer.Status()
						}
						c.JSON(status, gin.H{"errors": list})
					} else {
						c.JSON(422, gin.H{"errors": "please make sure to provide the correct data type or format"})
					}

				default:
					// Log all other errors
					//rollbar.RequestError(rollbar.ERR, c.Request, e.Err)
				}

			}
			// If there was no public or bind error, display default 500 message
			if !c.Writer.Written() {
				c.JSON(http.StatusInternalServerError, gin.H{"errors": "something went wrong"})
			}
		}
	}
}
