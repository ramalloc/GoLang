package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ramalloc/go-jwt/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// JWT uses hashes, So it gives token for all the data which provided to it
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UserType  string
	UserId    string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var secretKey string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(
	email string,
	firstName string,
	lastName string,
	userType string,
	userId string) (signedToken string, signedRefreshToken string, err error) {

	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserType:  userType,
		UserId:    userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(), //24 hours

		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(), //168 hours
		},
	}
	token, errToken:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if errToken != nil {
		log.Panic(errToken)
		return
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Making a object to store updated tokens
	var updateObj primitive.D

	// appending refresh token in upddatedObj
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", UpdatedAt})

	upsert := true
	filter := bson.M{"user_id": userId}
	updateOption := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&updateOption,
	)

	if err != nil {
		log.Panic(err)
		return
	}
	return
}


func ValidateToken(clientToken string) (claims *SignedDetails, msg string) {
	log.Println("Validating Token")
	token, err := jwt.ParseWithClaims(
		clientToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	log.Println("Token Parsed")
	fmt.Printf("Token: %v\n", token)
	if err != nil {
		msg = err.Error()
		return 
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = fmt.Sprintf("The Token in invalid !")
		msg = err.Error()
		return
	}


	// Checking for the token is expired or not 
	if claims.ExpiresAt < time.Now().Local().Unix(){
		msg = fmt.Sprintf("Token is Expired !")
		msg = err.Error()
		return
	}
	return claims, msg
}
