package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ramalloc/go-jwt/database"
	"github.com/ramalloc/go-jwt/helpers"
	"github.com/ramalloc/go-jwt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytePass)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	// We are using bcrypt library to compare hashed password,
	// while giving arguments to bcrypt.CompareHashAndPassword, first argument should be hashed password
	// and second argument should be the password provided by the user
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email or password is incorrect")
		check = false
	}

	return check, msg
}

func Singnup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			// return
		}

		// Checking for the email in the database already exists or not
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
			return
		}

		password := HashPassword(*&user.Password)
		user.Password = password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this email or phone number already exists"})
			return
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		// Saving the user's ID in Hexadecimal format in user_id
		user.UserId = user.ID.Hex()

		// Now we want the token for the details we will send to the user, which is done by helper.
		token, refreshToken, jwtErr := helpers.GenerateAllTokens(*&user.Email, *&user.FirstName, *&user.LastName, *&user.UserType, *&user.UserId)
		if jwtErr != nil {
			log.Panic(jwtErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while signup"})
			return
		}
		user.Token = token
		user.RefreshToken = refreshToken

		// Insert all data in database
		resultInsertionCount, InsertError := userCollection.InsertOne(ctx, user)
		if InsertError != nil {
			msg := fmt.Sprintf("Uer item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionCount)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		// Validating Icoming Data structure
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"email": user.Email}
		err := userCollection.FindOne(ctx, filter).Decode(&foundUser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		if foundUser.Email == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not Found"})
			return
		}

		passIsValid, message := VerifyPassword(*&user.Password, *&foundUser.Password)
		defer cancel()

		if !passIsValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": message})
			return
		}

		token, refreshTOken, _ := helpers.GenerateAllTokens(*&foundUser.Email, *&foundUser.FirstName, *&foundUser.LastName, *&foundUser.UserType, *&foundUser.UserId)
		// Update the token and refresh token in the database
		helpers.UpdateAllTokens(token, refreshTOken, *&foundUser.UserId)

		userCollection.FindOne(ctx, bson.M{"user_id": foundUser.UserId}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Loggedin Successfully", "user": foundUser})
	}
}

// Only Admin can access users data regular user can't access this
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		log.Println("User ID:", userId)
		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		filter := bson.M{"user_id": userId}
		err := userCollection.FindOne(ctx, filter).Decode(&user)
		defer cancel()

		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Println("User not found:", err)
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				log.Println("Error fetching user:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			return
		}
		user.UserType = ""
		log.Println("Fetched User:", user)
		c.JSON(http.StatusOK, user)
	}
}
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		pageSize, pageSizeErr := strconv.Atoi(c.Query("page_size"))
		if pageSizeErr != nil || pageSize < 1 {
			pageSize = 10
		}

		page, pageErr := strconv.Atoi(c.Query("page"))
		if pageErr != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * pageSize

		matchStage := bson.D{{"$match", bson.D{}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}},
		}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, pageSize}}}},
			}},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while listing user data items"})
			return
		}

		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}
		defer cancel()

		c.JSON(http.StatusOK, gin.H{"data": allUsers[0]})
	}
}
