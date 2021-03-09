# Kummer JWT GO MOD
This is the main exported class.
It has only 2 public method.

## Required
**env**: `SECRET_TOKEN_AUTH`

### How to implement

#### CreateToken

```go
func SingUp(ctx *fiber.Ctx) error {
	// ...
  
	t, err := jwt.CreateToken(
		models.Token{
			UserID:    userID,
			CompanyID: "",
			Admin:     false,
		})

	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
  
  // ...

```

#### MetadataToken
```go
func GetUser(ctx *fiber.Ctx) error {

	fastHTTPRequest := ctx.Request()

	claims, err := jwt.MetadataToken(fastHTTPRequest)

	data := claims["Data"].(map[string]interface{}) // add this in JWT

	token := models.DataToken{
		UserID:    data["user_id"].(string),
		Admin:     data["admin"].(bool),
		CompanyID: data["company_id"].(string),
		Exp:       claims["exp"].(string),
	}

	if err != nil {
		return ctx.SendStatus(fiber.StatusNetworkAuthenticationRequired)
	}
  //...
}
```
