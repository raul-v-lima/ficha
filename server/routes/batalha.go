package routes

import (
	"context"
	"fmt"
	"net/http"
	"server/helpers"
	"server/server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var batalhaCollection *mongo.Collection = OpenCollection(Client, "batalhas")
var environmentCollection *mongo.Collection = OpenCollection(Client, "environment")

func CalcIniciative() {}
func CalcDamage()     {}
func AddEnvironment(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var environment models.Environment
	if err := c.BindJSON(&environment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(environment)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	environment.ID = primitive.NewObjectID()
	result, insertErr := environmentCollection.InsertOne(ctx, environment)
	if insertErr != nil {
		msg := fmt.Sprintf("environment wasn't created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)

}
func AddCharacter(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var character models.Character

	if err := c.BindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(character)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	character.ID = primitive.NewObjectID()
	result, insertErr := batalhaCollection.InsertOne(ctx, character)

	if insertErr != nil {
		msg := fmt.Sprintf("character não foi criado")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)

}
func GetEnvironment(c *gin.Context) {
	environmentName := c.Params.ByName("name")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var environment []bson.M

	cursor, err := environmentCollection.Find(ctx, bson.M{"name": environmentName})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &environment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(environment)

	c.JSON(http.StatusOK, environment)
}
func GetEnvironments(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var environment []bson.M

	cursor, err := environmentCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &environment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	//fmt.Println(environment)

	c.JSON(http.StatusOK, environment)
}
func CalcFa(c *gin.Context) int32 {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var magic bool = true

	name := c.Params.ByName("name")
	//forca, _ := strconv.ParseInt(c.Params.ByName("Vigor"), 10, 32)
	vigor := helpers.ConvertStringToInt32(c.Params.ByName("vigor"))
	dexterity := helpers.ConvertStringToInt32(c.Params.ByName("dexterity"))
	empiricism := helpers.ConvertStringToInt32(c.Params.ByName("empiricism"))

	//intVar, err := strconv.Atoi(strVar)
	var fa int32 = vigor + dexterity
	if magic {
		fa = empiricism + dexterity
	}

	// limpeza de chamdas de recursos
	defer cancel()
	fmt.Println(fa, name)
	//c.JSON(http.StatusOK,)
	return fa
}

func UpdateCharacter(c *gin.Context) {

	characterID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(characterID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	type Character struct {
		Name       *string `json:"name"`
		Vigor      *int32  `json:"vigor"`
		Empiricism *int32  `json:"resistency"`
		Dexterity  *int32  `json:"Dexterity"`
		Mana       *int32  `json:"mana"`
		Xp         *int32  `json:"xp"`
		Level      *int32  `json:"nivel"`
	}

	var character Character
	if err := c.BindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := batalhaCollection.UpdateOne(ctx, bson.M{"_id": docID}, bson.D{
		{Key: "$set", Value: bson.D{{Key: "name", Value: character.Name}}},
		{Key: "$set", Value: bson.D{{Key: "dexterity", Value: character.Dexterity}}},
		{Key: "$set", Value: bson.D{{Key: "empiricism", Value: character.Empiricism}}},
		{Key: "$set", Value: bson.D{{Key: "level", Value: character.Level}}},
		{Key: "$set", Value: bson.D{{Key: "vigor", Value: character.Vigor}}},
		{Key: "$set", Value: bson.D{{Key: "mana", Value: character.Mana}}},
		{Key: "$set", Value: bson.D{{Key: "xp", Value: character.Xp}}},
	},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.ModifiedCount)

}
func DeleteCharacter(c *gin.Context) {

	character := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(character)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := batalhaCollection.DeleteOne(ctx, bson.M{"_id": docId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)

}
func Skirmish()          {}
func RelatorioSkirmish() {}
func GetCharacters(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var characters []bson.M

	cursor, err := batalhaCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &characters); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	//fmt.Println(characters)

	c.JSON(http.StatusOK, characters)
}

func Getcharacter(c *gin.Context) {

	nomecharacter := c.Params.ByName("nome")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var character []bson.M

	cursor, err := batalhaCollection.Find(ctx, bson.M{"nome": nomecharacter})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &character); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()

	fmt.Println(character)

	c.JSON(http.StatusOK, character)
}
